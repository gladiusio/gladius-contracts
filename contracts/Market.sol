pragma solidity ^0.4.19;

import "./Pool.sol";
import "./AbstractBalance.sol";

/// Maps functions of Gladius Token to the base Token contract
contract Token {
  mapping (address => uint256) public balanceOf;
  function getBal(address owner) public returns(uint256 bal);
  function transferFrom(address _from, address _to, uint256 _value) public returns (bool success);
}

contract Market is AbstractBalance {

  mapping(address => Pool[]) public marketPools;        // All Pools in marketplace, ignores non-market Pools
  mapping(address => Pool[]) public ownedPools;         // All Pools created off this contract, regardless of listing in marketplace
  mapping(address => Pool) public clientToPool;         // Client that pays the Pool for service
  mapping(address => address) public poolToOwner;       // Owner of a Pool
  mapping(address => uint32) tokensPaid;                // Account balance of the clients
  mapping(address => Payout[]) public payouts;

  address[] public allPoolsList;                        // Array of all pools
  address[] public marketPoolsList;                     // Array of all pools on the marketplace

  address public owner;                                 // Owner of the market
  uint256 maxPayout;                                    // Max amount a pool can withdraw daily
  uint256 joinCost;                                     // Cost to join marketplace

  Token gladiusToken;

  struct Payout {
    address pool;
    address user;
    uint256 amount;
    uint timestamp;
  }

  /**
   * Marketplace constructor
   *
   * Logic
   * @param _owner Owners address
   * @param _gladiusToken GladiusToken address
   * @param _joinCost Cost to join the marketplace
   * @param _maxPayout Maxium payout to pool owners in a given day
   */
  function Market(address _owner, address _gladiusToken, uint256 _joinCost, uint256 _maxPayout) public {
    owner = _owner;
    gladiusToken = Token(_gladiusToken);
    joinCost = _joinCost;
    maxPayout = _maxPayout;
  }

  /**
   * Create a new pool. This is FREE but DOES NOT add pool to the marketplace
   *
   * Instantiate a new Pool and set the sender as the owner
   * @param _publicKey Public RSA key used to encrypt traffic
   * @return address Address to the new Pool
   */
   function createPool(string _publicKey) public returns(address) {
     Pool newPool = new Pool(_publicKey, msg.sender);
     ownedPools[msg.sender].push(newPool);
     allPoolsList.push(newPool);
     poolToOwner[newPool] = msg.sender;

     return newPool;
   }

  /**
   * List out pools owned by given address
   *
   * Returns a list of pools
   * @param _ownerAddress Description
   * @return Pool[] List of Pools owned by _ownerAddress
   */
  function getOwnedPools(address _ownerAddress) public view returns (Pool[]) {
    return ownedPools[_ownerAddress];
  }

  /**
   * Returns an array of all pools that a website(client) uses
   *
   * @param _client Address of client to query pools
   * @return poolAddress Address to a list of pools the client is a part of
   */
  function getClientPool(address _client) public view returns(Pool) {
    return clientToPool[_client];
  }

  /**
   * Records an amount of work done in a pool for a node by a client
   *
   * @param _pool Pool address to record work to
   * @param _node Node that completed the work
   * @param _client Client that will be billed for the nodes work
   * @param _amount Amount of work completed
   */
  function logWorkFrom(address _pool, address _node, address _client, uint _amount) public {
    Pool pool = Pool(_pool);

    require(pool.logWorkFrom(_node, _client, _amount));
    require(work(_amount));
  }

  /**
   * Returns a list of receipts of payouts for a given user
   *
   * @param _user User to retrieve payouts from
   * @return Payout[] array of payouts
   */
  function getPayouts(address _user) public view returns(Payout[]) {
    return payouts[_user];
  }

  /**
   * Determines if a node can withdraw money
   *
   * Checks the daily maximum and the timeframe of last payout
   * @param _pool Pool address to record work to
   * @param _node Node that completed the work
   * @param _amount Amount of work completed
   * @return bool true if node can withdraw
   */
  function canPayNode(address _pool, address _node, uint256 _amount) private view returns (bool) {
    Payout[] storage poolPayouts = payouts[_node];

    // Arbitrary number for now
    uint256 dailyMaximum = 5;

    if (_amount > dailyMaximum) {
      return false;
    }

    if (poolPayouts.length == 0) {
      return true;
    } else {
      uint256 recentPayouts = 0;
      for(uint i = poolPayouts.length - 1; i>=0; i++) {
        Payout storage payout = poolPayouts[i];
        if (now - payout.timestamp > 86400) {
          break;
        } else if (_pool == payout.pool) {
          recentPayouts += payout.amount;
        }
      }

      if (recentPayouts < dailyMaximum && _amount < Pool(_pool).getOwedBalanceFor(_node)) {
        return true;
      } else {
        return false;
      }
    }
  }

  /**
   * Lets a node withdraw money owed to them if canPayNode() passes
   *
   * Actual payout function that transfers money
   * @param _pool Pool address to record work to
   * @param _node Node that completed the work
   * @param _amount Amount of work completed
   * @return bool true if withdraw was successful
   */
  function payout(address _pool, address _node, uint256 _amount) public returns (bool) {
    if (!canPayNode(_pool, _node, _amount)) {
      return false;
    }

    Pool(_pool).payout(_node, _amount);
    pay(_amount);

    payouts[_node].push(Payout(_pool,  _node, _amount, now));

    return true;
  }

  /**
   * Adds the pool address to the marketplace, and charges the owner of the pool the join cost
   *
   * Require that the pool's owner is the sender
   * Checks account balance and transfers money from the sender to the pool's address for the join cost
   * Creates a pool instance from the address and pushes the pool to the market place pools
   *
   * @param _poolAddress Param Description
   */
  function joinMarketplace(address _poolAddress) public returns(bool) {
    require(poolToOwner[_poolAddress] == msg.sender); //caller owns pool and pool exists

    gladiusToken.transferFrom(msg.sender, address(this), joinCost); //charge the caller to add to marketplace (right now just burn)

    Pool p = Pool(_poolAddress);
    marketPools[msg.sender].push(p); //add pool to the marketplace
    marketPoolsList.push(p);

    return true;
  }

  /**
   * Initial allocation of funds from Client to Market
   *
   * Moves money from a client into the MarketPlace
   * @param _pool address
   * @param _client address
   * @param _amount uint
   * @return bool TODO run through checks to ensure transfer logic passes
   */
  function allocateClientFundsTo(address _pool, address _client, uint _amount) public returns (bool) {
    // TODO - Allocate pool funds first to ensure marketplace should allocate
    allocateFunds(_amount);

    Pool pool = Pool(_pool);
    // Allocate pool balance
    return pool.allocateFundsFrom(_client, _amount);
  }

  function getAllPools() public view returns (address[]) {
      return allPoolsList;
    }
}

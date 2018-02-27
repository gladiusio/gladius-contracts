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
    mapping(address => Withdrawal[]) public withdrawals;

    address [] public allPoolsList;                       // Array of all pools
    address [] public marketPoolsList;                    // Array of all pools on the marketplace

    address public owner;                                 // Owner of the market
    uint256 maxPayout;                                    // Max amount a pool can withdraw daily
    uint256 joinCost;                                     // Cost to join marketplace

    Token gladiusToken;

    struct Withdrawal {
      address pool;
      address user;
      uint256 amount;
      uint timestamp;
    }

    /**
     * Marketplace constructor
     *
     * Logic
     * @param _owner Owner's address
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
     * @param publicKey Public RSA key used to encrypt traffic
     * @return address Address to the new Pool
     */
    function createPool(string publicKey) public returns(address) {
        Pool newPool = new Pool(publicKey, msg.sender);
        ownedPools[msg.sender].push(newPool);
        allPoolsList.push(newPool);

        return newPool;
    }

    function getOwnedPools(address ownerAddress) public view returns (Pool[]) {
      return ownedPools[ownerAddress];
    }

    /**
     * Returns an array of all pools that a website(client) uses
     *
     * @param client Address of client to query pools
     * @return poolAddress Address to a list of pools the client is a part of
     */
    function getClientPool(address client) public view returns(Pool) {
        return clientToPool[client];
    }

    function getWithdrawals(address _user) public view returns(Withdrawal[]) {
        return withdrawals[_user];
    }

    function canWithdraw(address _pool, address _user, uint256 _amount) public view returns (bool) {
        Withdrawal[] storage userWithdrawals = withdrawals[_user];

        // Arbitrary number for now
        uint256 dailyMaximum = 5;

        if (_amount > dailyMaximum) {
          return false;
        }

        if (userWithdrawals.length == 0) {
          return true;
        } else {
          uint256 withdrawn = 0;
          for(uint i = userWithdrawals.length - 1; i>=0; i++){
              Withdrawal storage withdrawal = userWithdrawals[i];
              if (now - withdrawal.timestamp > 86400) {
                break;
              } else if (_pool == withdrawal.pool) {
                withdrawn += withdrawal.amount;
              }
          }

          if (withdrawn < dailyMaximum && _amount < Pool(_pool).getWithdrawable()) {
            return true;
          } else {
            return false;
          }
        }
    }

    function withdraw(address _pool, address _user, uint256 _amount) public returns (bool) {
        if (!canWithdraw(_pool, _user, _amount)) {
          return false;
        }

        // Do withdraw in AbstractBalance
        Pool(_pool).withdrawFunds(_amount, _user);
        this.withdrawFunds(_amount);

        withdrawals[_user].push(Withdrawal(_pool,  _user, _amount, now));

        return true;
    }

    /**
     * Adds the pool address to the marketplace, and charges the owner of the pool the join cost
     *
     * Require that the pool's owner is the sender
     * Checks account balance and transfers money from the sender to the pool's address for the join cost
     * Creates a pool instance from the address and pushes the pool to the market place pools
     *
     * @param poolAddress Param Description
     */
    function joinMarketplace(address poolAddress) public returns(bool) {
        require(poolToOwner[poolAddress] == msg.sender); //caller owns pool and pool exists

        gladiusToken.transferFrom(msg.sender, address(this), joinCost); //charge the caller to add to marketplace (right now just burn)

        Pool p = Pool(poolAddress);
        marketPools[msg.sender].push(p); //add pool to the marketplace
        marketPoolsList.push(newPool);

        return true;
    }

    function allocateClientFundsTo(address poolAddress, address userAddress, uint256 allocationAmount) public returns (bool) {
        allocateFunds(allocationAmount);

        Pool pool = Pool(poolAddress);
        // Allocate pool balance
        return pool.allocateClientFundsFrom(userAddress, allocationAmount);
    }
}

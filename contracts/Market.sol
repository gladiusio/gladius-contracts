pragma solidity ^0.4.18;

import "./Pool.sol";

contract Owned {
    function owned() public { owner = msg.sender; }
    address owner;

    // This contract only defines a modifier but does not use
    // it: it will be used in derived contracts.
    // The function body is inserted where the special symbol
    // `_;` in the definition of a modifier appears.
    // This means that if the owner calls this function, the
    // function is executed and otherwise, an exception is
    // thrown.
    modifier onlyOwner {
        require(msg.sender == owner);
        _;
    }
}

/// Maps functions of Gladius Token to the base Token contract
contract Token {
    mapping (address => uint256) public balanceOf;
    function getBal(address owner) public returns(uint256 bal);
    function transferFrom(address _from, address _to, uint256 _value) public returns (bool success);
}

contract Market {

    mapping(address => Pool[]) public marketPools;        // All Pools in marketplace, ignores non-market Pools
    mapping(address => Pool[]) public ownedPools;         // All Pools created off this contract, regardless of listing in marketplace
    mapping(address => Pool) public clientToPool;         // Client that pays the Pool for service
    mapping(address => address) public poolToOwner;       // Owner of a Pool
    mapping(address => uint32) tokensPaid;                // Account balance of the clients

    address public owner;                                 // Owner of the market
    uint256 maxPayout;                                    // Max amount a pool can withdraw daily
    uint256 joinCost;                                     // Cost to join marketplace

    // Ratios

    // Work Ratio           - Amount set aside to work balance

    // ----------------------

    // Balances
    
    uint256 public balanceTotal;                                 // Total balance        - Complete Balance
    uint256 public balanceAvailable;                             // Available Balance    - Allocatable Balance
    uint256 public balanceTransactionCosts;                      // Market Fee Balance   - Balance to hold transaction fees
    uint256 public balanceWorkable;                              // Work Balance         - Allocated Balance
    uint256 public balanceCompleted;                             // Work Completed       - Work completed balance
    uint256 public balancePayable;                               // Payout Balance       - Balance able to withdraw

    uint256 public balanceWithdrawable;                          // Payout Balance       - Balance able to withdraw
    Token gladiusToken;

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

        return newPool;
    }

    function getOwnedPools(address ownerAddress) public returns (Pool[]) {
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

        return true;
    }

    function allocateClientFundsTo(address poolAddress, address userAddress, uint32 allocationAmount) public returns (bool) {
        uint256 _balanceTotal = allocationAmount;
        uint256 _balanceAvailable = (3 * allocationAmount) / 5;
        uint256 _balanceWorkable = allocationAmount - _balanceAvailable;
        uint256 _balanceWithdrawable = 0;
        uint256 _balanceCompleted = 0;
        uint256 _balancePayable = 0;

        // Allocate market funds
        balanceTotal += _balanceTotal;
        balanceAvailable += _balanceAvailable;
        balanceWorkable += _balanceWorkable;
        balanceWithdrawable += _balanceWithdrawable;
        balanceCompleted += _balanceCompleted;
        balancePayable += _balancePayable;

        Pool pool = Pool(poolAddress);
        // Allocate pool balance
        return pool.allocateClientFundsFrom(userAddress, allocationAmount);
    }
}

pragma solidity ^0.4.18;

import "./Pool.sol";

contract owned {
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

contract Token {
    mapping (address => uint256) public balanceOf;
    function getBal(address owner) public returns(uint256 bal);
    function transferFrom(address _from, address _to, uint256 _value) public returns (bool success);
}

contract Market {

    mapping(address => Pool[]) public marketPools; // all pools in market
    mapping(address => Pool[]) public ownedPools; // all pools created
    mapping(address => Pool) public clientToPool; // pool that website uses
    mapping(address => address) public poolToOwner;
    mapping(address => uint32) tokensPaid; // account balance of the clients

    address owner;
    uint256 maxPayout; //max amount a pool can withdraw daily
    uint256 joinCost; //cost to join marketplace

    Token gladiusToken;

    /**
     * Marketplace constructor
     *
     * Logic
     * @param _gladiusToken GladiusToken address
     * @param _joinCost Cost to join the marketplace
     * @param _maxPayout Maxium payout to pool owners in a given day
     */
    function Market(address _gladiusToken, uint256 _joinCost, uint256 _maxPayout) public {
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

    /**
     * Returns an array of all pools that a website(client) uses
     *
     * Logic
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

    /* Pay pool certain amount. Manual as of now, automated in the future */
    /* Only the owner of a pool can withdraw money from its pool */
    /* Gladius must be able to withdraw money from listing a pool on the marketplace */
    function withdraw(uint256 amount) public returns(bool){
        if( msg.sender == owner)
            return true;
        else
            return false;
    }
}

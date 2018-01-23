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

    uint256 maxPayout; //max amount a pool can withdraw daily
    uint256 joinCost; //cost to join marketplace

    Token gladiusToken;

    /* Marketplace contrusctor, add the token address and the initial cost per pool */
    function Market(address tokenAddress, uint256 cost, uint256 maxPay) public {
        gladiusToken = Token(tokenAddress);
        joinCost = cost;
        maxPayout = maxPay;
    }

    /* Create a new pool. This is FREE but DOES NOT add pool to the marketplace */
    function createPool(string publicKey) public returns(address) {
        Pool p = new Pool(publicKey, msg.sender);
        ownedPools[msg.sender].push(p);

        return p;
    }

    /* Returns an array of all of the pools that belong to owner */
    function getOwnedPools(address owner) public view returns(Pool[]) {
        return ownedPools[owner];
    }
    /* Returns an array of all pools that a website(client) uses */
    function getClientPool(address client) public view returns(Pool) {
        return clientPool[client];
    }

    /* Add pool to the marketplace and charge the owner */
    function joinMarketplace(address poolAddress) public returns(bool) {
        require(poolToOwner[poolAddress] == msg.sender); //caller owns pool and pool exists

        gladiusToken.transferFrom(msg.sender, address(this), joinCost); //charge the caller to add to marketplace (right now just burn)

        Pool p = Pool(poolAddress);
        marketPools[msg.sender].push(p); //add pool to the marketplace

        return true;
    }

    /* Pay pool certain amount. Manual as of now, automated in the future */
    function withdraw(uint256 amount) public returns(bool){
        if( msg.sender == owner)
            return true;
        else
            return false;
    }
}

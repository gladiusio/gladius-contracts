pragma solidity ^0.4.18;

import "./Pool.sol";

contract Token {
    mapping (address => uint256) public balanceOf;
    function getBal(address owner) public returns(uint256 bal);
    function transferFrom(address _from, address _to, uint256 _value) public returns (bool success);
}

contract Market {

    mapping(address => Pool[]) marketPools; // all pools in market
    mapping(address => Pool[]) ownedPools; // all pools created
    mapping(address => Pool) clientPool; // pool that website uses
    mapping(address => address) poolToOwner;
    mapping(address => uint32) tokensPayed; // The number of tokens a person has payed to the market.

    Token token;

    uint256 joinCost; //cost to create a pool

    /* Marketplace contrusctor, add the token address and the initial cost per pool */
    function Market(address tokenAddress, uint256 cost) public {
        token = Token(tokenAddress);
        joinCost = cost;
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

        token.transferFrom(msg.sender, address(this), joinCost); //charge the caller to add to marketplace (right now just burn)

        Pool p = Pool(poolAddress);
        marketPools[msg.sender].push(p); //add pool to the marketplace

        return true;
    }

    /* Pay pool certain amount. Manual as of now, automated in the future */
    function payPool(address poolAddress, uint amount) public returns(bool){
        token.transferFrom(address(this), poolAddress, amount);

        return true;
    }




}

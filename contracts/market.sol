pragma solidity ^0.4.18;

import "./Pool.sol";

contract Token {
    function getBal(address owner) public returns(uint256 bal);
    mapping (address => uint256) public balanceOf;
}

contract Market {

    mapping(address => Pool[]) marketPools; // all pools in market
    mapping(address => Pool[]) ownedPools; // all pools created
    mapping(address => Pool) clientPool; // pool that website uses
    mapping(address => address) poolToOwner;
    mapping(address => uint32) tokensPayed; // The number of tokens a person has payed to the market.

    address tokenAddress; // address of the gladius token contract

    function Market(address token) public {
        tokenAddress = token;
    }

    function createPool(string _name) public returns(address) {
        Pool p = new Pool("password", msg.sender);
        ownedPools[msg.sender].push(p);

        return p;
    }

    function getOwnedPools(address owner) public view returns(Pool[]) {
        return ownedPools[owner];
    }

    function getClientPool(address client) public view returns(Pool) {
        return clientPool[client];
    }

    function joinMarketplace(address poolAddress) public returns(bool) {
        require (poolToOwner[poolAddress] == msg.sender); //caller owns pool

        Token t = Token(tokenAddress);
        require(t.balanceOf(poolAddress) > 0); //caller has sufficient funds to do this

        marketPools[msg.sender].push(poolAddress);

        return true;
    }

    /*function payPool(address poolAddress, uint amount) public returns(bool){

    }*/




}

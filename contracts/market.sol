pragma solidity ^0.4.0;

import "./Pool.sol";


contract Market {

    // mapping(address => Pool[]) marketPools; // all pools in market
    mapping(address => Pool[]) ownedPools; // all pools created
    mapping(address => Pool) clientPool; // pool that website uses
    mapping(address => address) poolToOwner;

    mapping(address => uint32) tokensPayed; // The number of tokens a person has payed to the market.

    function createPool(string _name) public returns(address) {
        Pool p = new Pool("password", msg.sender);
        ownedPools[msg.sender].push(p);

        return p;
    }

    function getOwnedPools(address owner) public returns(Pool[]) {
        return ownedPools[owner];
    }

    function getClientPool(address client) public returns(Pool) {
        return clientPool[client];
    }

    function joinMarketplace(address poolAddress) public returns(bool) {
        require (poolToOwner[poolAddress] == msg.sender);
        // TODO: make sure they have GLA to do this
        // marketPools[msg.sender].push(poolAddress);

        return true;
    }

    /*function payPool(address poolAddress, uint amount) public returns(bool){

    }*/




}

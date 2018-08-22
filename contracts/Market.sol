pragma solidity ^0.4.19;

import "./Pool.sol";

contract Market {
  mapping(address => Pool[]) public marketPoolOwners;   // Owners of pools in the marketplace
  mapping(address => Pool[]) public poolOwners;         // Owners of all the pools
  mapping(address => address) public poolToOwner;       // Pool to Owner

  address[] public allPoolsList;                        // Array of all pools
  address[] public marketPoolsList;                     // Array of all pools in the marketplace

  address public owner;                                 // Owner of the market

  string public data;

  /**
   * Marketplace constructor
   *
   * @param _owner Owners address
   */
  constructor(address _owner) public {
    owner = _owner;
  }

  /**
   * Create a new pool. This is FREE but DOES NOT add pool to the marketplace
   *
   * Instantiate a new Pool and set the sender as the owner
   * @return address Address to the new Pool
   */
   function createPool() public returns(address) {
     Pool newPool = new Pool(msg.sender);
     poolOwners[msg.sender].push(newPool);
     allPoolsList.push(newPool);
     poolToOwner[newPool] = msg.sender;

     return newPool;
   }

   function getAllPools() public view returns (address[]) {
     return allPoolsList;
   }

   function getMarketPools() public view returns (address[]) {
     return marketPoolsList;
   }

   function getOwner() public view returns (address) {
     return owner;
   }

   function getData() public view returns (string) {
     return data;
   }

  /**
   * List out pools owned by given address
   *
   * @param _ownerAddress Owner of pools
   * @return Pool[] List of Pools owned by _ownerAddress
   */
  function getOwnerAllPools(address _ownerAddress) public view returns (Pool[]) {
    return poolOwners[_ownerAddress];
  }

  /**
   * List out pools owned by given address in the marketplace
   *
   * @param _ownerAddress Owner of pools
   * @return Pool[] List of Pools owned by _ownerAddress that are also in the marketplace
   */
  function getOwnerMarketPools(address _ownerAddress) public view returns (Pool[]) {
    return marketPoolOwners[_ownerAddress];
  }

  /**
  * Get the owner of a pool
  *
  * @param _pool Pool address
  * @return Owner of _pool
  */
  function getPoolOwner(address _pool) public view returns (address) {
    return poolToOwner[_pool];
  }

  /**
  * Set pool data
  *
  * @param _data Pool data
  */
  function setData(string _data) public {
    require(msg.sender == owner);
    data = _data;
  }

  /**
   * Adds the pool address to the marketplace
   *
   * @param _poolAddress Pool address to add
   */
  function addToMarketplace(address _poolAddress) public returns(bool) {
    require(msg.sender == owner); // only marketplace owner can do this

    Pool p = Pool(_poolAddress);

    marketPoolOwners[p.getOwner()].push(p);
    marketPoolsList.push(p.getOwner());

    return true;
  }


}

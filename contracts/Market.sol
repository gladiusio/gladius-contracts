pragma solidity >=0.4.21 <0.6.0;

import "./Pool.sol";

contract Market {
    mapping(address => Pool[]) public marketPoolOwners;   // Owners of pools in the marketplace

    address[] public marketPoolsList;                     // Array of all pools in the marketplace

    address public owner;                                 // Owner of the market
    PoolFactory poolFactory;

    string public data;

    string private enforceOwner = enforceOwner;

    /**
    * Marketplace constructor
    *
    * @param _owner Owners address
    */
    constructor(address _owner, address _poolFactory) public {
        owner = _owner;
        poolFactory = PoolFactory(_poolFactory);
    }

    /**
    * Create a new pool. This is FREE but DOES NOT add pool to the marketplace
    *
    * Instantiate a new Pool and set the sender as the owner
    * @return address Address to the new Pool
    */
    function createPool() public returns(address) {
        address poolAddress = poolFactory.createPool(msg.sender);
        Pool newPool = Pool(poolAddress);

        return address(newPool);
    }

    function getAllPools() public view returns (address[] memory) {
        return poolFactory.getAllPools();
    }

    function getMarketPools() public view returns (address[] memory) {
        return marketPoolsList;
    }

    function getOwner() public view returns (address) {
        return owner;
    }

    function getData() public view returns (string memory) {
        return data;
    }

    function getPoolFactory() public view returns (address) {
        return address(poolFactory);
    }

    /**
    * List out pools owned by given address
    *
    * @param _ownerAddress Owner of pools
    * @return Pool[] List of Pools owned by _ownerAddress
    */
    function getOwnerAllPools(address _ownerAddress) public view returns (Pool[] memory) {
        return poolFactory.getOwnedPool(_ownerAddress);
    }

    /**
    * List out pools owned by given address in the marketplace
    *
    * @param _ownerAddress Owner of pools
    * @return Pool[] List of Pools owned by _ownerAddress that are also in the marketplace
    */
    function getOwnerMarketPools(address _ownerAddress) public view returns (Pool[] memory) {
        return marketPoolOwners[_ownerAddress];
    }

    /**
    * Get the owner of a pool
    *
    * @param _pool Pool address
    * @return Owner of _pool
    */
    function getPoolOwner(address _pool) public view returns (address) {
        return poolFactory.getPoolOwner(_pool);
    }

    /**
    * Change the owner
    *
    * @param _owner New owner
    */
    function changeOwner(address _owner) public {
        require(msg.sender == owner, enforceOwner);
        owner = _owner;
    }

    /**
    * Set pool data
    *
    * @param _data Pool data
    */
    function setData(string memory _data) public {
        require(msg.sender == owner, enforceOwner);
        data = _data;
    }

    /**
    * Change pool factory
    *
    * @param _factory Pool address
    */
    function changePoolFactory(address _factory) public {
        require(msg.sender == owner, enforceOwner);
        poolFactory = PoolFactory(_factory);
    }

    /**
    * Adds the pool address to the marketplace
    *
    * @param _poolAddress Pool address to add
    */
    function addToMarketplace(address _poolAddress) public returns(bool) {
        require(msg.sender == owner, enforceOwner); // only marketplace owner can do this

        Pool p = Pool(_poolAddress);

        marketPoolOwners[p.getOwner()].push(p);
        marketPoolsList.push(p.getOwner());

        return true;
    }
}

pragma solidity ^0.4.19;

contract PoolFactoryTest {

    mapping(address => PoolTest[]) public ownerToPools;       // Owner to their pool(s)
    mapping(address => address) public poolToOwner;       // Pool to their Owner
    address[] public allPools;                            // All pools

    /**
    * Create a new pool.
    *
    * Instantiate a new Pool and set the sender as the owner
    * @return address Address to the new pool
    */
    function createPool(address _owner) public returns(address) {
        PoolTest pool = new PoolTest(_owner);

        ownerToPools[_owner].push(pool);
        poolToOwner[address(pool)] = _owner;
        allPools.push(address(pool));

        return address(pool);
    }

    function ownerToPools(address _owner) public view returns(PoolTest[]) {
        return ownerToPools[_owner];
    }

    function poolToOwner(address _pool) public view returns(address) {
        return poolToOwner[_pool];
    }

    function getAllPools() public view returns(address[]) {
        return allPools;
    }
}

contract PoolTest {

    address[] public masterNodes;
    address public owner;

    string public seedNode;
    string public data;
    string public url;


    /**
    * Create new Pool and assign owner
    *
    * Data is assigned owner and uses the owner's public key
    * @param _owner Address of the owner
    */
    constructor(address _owner) public {
        owner = _owner;
    }

    function getOwner() public view returns(address) {
        return owner;
    }

    function getMasterNodes() public view returns(address[]) {
        return masterNodes;
    }

    function getSeedNode() public view returns(string) {
        return seedNode;
    }

    function getData() public view returns(string) {
        return data;
    }

    function getUrl() public view returns(string) {
        return url;
    }

    /**
    * Change the owner
    *
    * @param _owner New owner
    */
    function changeOwner(address _owner) public {
        require(msg.sender == owner, "Must be owner of contract");
        owner = _owner;
    }


    /**
    * Set the pool data
    *
    * @param _data new pool data
    */
    function setData(string _data) public {
        require(msg.sender == owner, "Must be owner of contract");
        data = _data;
    }

    /**
    * Set the url data
    *
    * @param _url new url
    */
    function setUrl(string _url) public {
        require(msg.sender == owner, "Must be owner of contract");
        url = _url;
    }

    /**
    * Set the seed node
    *
    * @param _node New seed node
    */
    function setSeedNode(string _node) public {
        require(msg.sender == owner, "Must be owner of contract");
        seedNode = _node;
    }

    /**
    * Add a master node to the list of master nodes
    *
    * @param _node Node to add to the list
    */
    function addMasterNode(address _node) public {
        require(msg.sender == owner, "Must be owner of contract");
        masterNodes.push(_node);
    }
}
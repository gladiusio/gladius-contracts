pragma solidity >=0.4.21 <0.6.0;

contract PoolFactory {
    mapping(address => Pool[]) public ownerToPools;       // Owner to their pool(s)
    mapping(address => address) public poolToOwner;       // Pool to their Owner
    address[] public allPools;                            // All pools

    /**
    * Create a new pool.
    *
    * Instantiate a new Pool and set the sender as the owner
    * @return address Address to the new pool
    */
    function createPool(address _owner) public returns(address) {
        Pool pool = new Pool(_owner);

        ownerToPools[_owner].push(pool);
        poolToOwner[address(pool)] = _owner;
        allPools.push(address(pool));

        return address(pool);
    }

    function getOwnedPool(address _owner) public view returns(Pool[] memory) {
        return ownerToPools[_owner];
    }

    function getPoolOwner(address _pool) public view returns(address) {
        return poolToOwner[_pool];
    }

    function getAllPools() public view returns(address[] memory) {
        return allPools;
    }
}

contract Pool {
    mapping(address => bool) public masterNodes;
    address[] public masterNodesList;

    mapping(address => bool) public approvedNodes;
    address[] public approvedNodesList;

    mapping(address => bool) public infrastructureNodes;
    address[] public infrastructureNodesList;

    address public owner;

    string public seedNode;
    string public data;
    string public poolDomain;
    string public cdnDomain;
    string public certificateBundle;


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

    function isMasterNode(address _node) public view returns(bool) {
        return masterNodes[_node];
    }

    function isApprovedNode(address _node) public view returns(bool) {
        return approvedNodes[_node];
    }

    function isInfrastructureNode(address _node) public view returns(bool) {
        return infrastructureNodes[_node];
    }

    function getAllApprovedNodes() public view returns(address[] memory) {
        return approvedNodesList;
    }

    function getAllMasterNodes() public view returns(address[] memory) {
        return masterNodesList;
    }

    function getAllInfrastructureNodes() public view returns(address[] memory) {
        return infrastructureNodesList;
    }

    function getSeedNode() public view returns(string memory) {
        return seedNode;
    }

    function getData() public view returns(string memory) {
        return data;
    }

    function getPoolDomain() public view returns(string memory) {
        return poolDomain;
    }

    function getCertificateBundle() public view returns(string memory) {
        return certificateBundle;
    }

    function getCDNDomain() public view returns(string memory) {
        return cdnDomain;
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
    function setData(string memory _data) public {
        require(msg.sender == owner, "Must be owner of contract");
        data = _data;
    }

    /**
    * Set the certiciate bundle link
    *
    * @param _url new cdn certificate link
    */
    function setCertificateBundle(string memory _url) public {
        require(msg.sender == owner, "Must be owner of contract");
        certificateBundle = _url;
    }

    /**
    * Set the poolDomain data
    *
    * @param _url new url
    */
    function setPoolDomain(string memory _url) public {
        require(msg.sender == owner, "Must be owner of contract");
        poolDomain = _url;
    }

    /**
    * Set the cdnDomain data
    *
    * @param _url new url
    */
    function setCDNDomain(string memory _url) public {
        require(msg.sender == owner, "Must be owner of contract");
        cdnDomain = _url;
    }

    /**
    * Set the seed node
    *
    * @param _node New seed node
    */
    function setSeedNode(string memory _node) public {
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
        masterNodes[_node] = true;
        masterNodesList.push(_node);
    }

    /**
    * Add an infrastructure node to the list of master nodes
    *
    * @param _node Node to add to the list
    */
    function addInfrastructureNode(address _node) public {
        require(msg.sender == owner, "Must be owner of contract");
        infrastructureNodes[_node] = true;
        infrastructureNodesList.push(_node);
    }


    /**
    * Add a node to the list of approved nodes
    *
    * @param _node Node to add to the list
    */
    function addApprovedNode(address _node) public {
        require(msg.sender == owner, "Must be owner of contract");
        approvedNodes[_node] = true;
        approvedNodesList.push(_node);
    }

    /**
    * Remove a master node to the list of master nodes
    *
    * @param _node Node to remove from the list
    */
    function removeMasterNode(address _node) public {
        require(msg.sender == owner, "Must be owner of contract");
        if (masterNodes[_node]) {
          masterNodes[_node] = false;
        }
    }

        /**
    * Remove an infrastructure node to the list of master nodes
    *
    * @param _node Node to remove from the list
    */
    function removeInfrastructureNode(address _node) public {
        require(msg.sender == owner, "Must be owner of contract");
        if (infrastructureNodes[_node]) {
          infrastructureNodes[_node] = false;
        }
    }

    /**
    * Remove a node to the list of approved nodes
    *
    * @param _node Node to remove from the list
    */
    function removeApprovedNode(address _node) public {
        require(msg.sender == owner, "Must be owner of contract");
        if (approvedNodes[_node]) {
          approvedNodes[_node] = false;
        }
    }
}

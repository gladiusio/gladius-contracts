pragma solidity ^0.4.19;

contract Pool {

  address public owner;
  address[] public masterNodes;

  string public seedNode;
  string public data;


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

  /**
   * Change the owner
   *
   * @param _owner New owner
   */
  function changeOwner(address _owner) public {
    require(msg.sender == owner);
    owner = _owner;
  }


  /**
   * Set the pool data
   *
   * @param _data new pool data
   */
  function setData(string _data) public {
    require(msg.sender == owner);
    data = _data;
  }

  /**
   * Set the seed node
   *
   * @param _node New seed node
   */
  function setSeedNode(string _node) public {
    require(msg.sender == owner);
    seedNode = _node;
  }

  /**
   * Add a master node to the list of master nodes
   *
   * @param _node Node to add to the list
   */
  function addMasterNode(address _node) public {
    require(msg.sender == owner);
    masterNodes.push(_node);
  }
}

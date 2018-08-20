pragma solidity ^0.4.19;

//ALL DATA VARIABLES ARE SUBJECT TO CHANGE (STC)

contract Pool {

  address public owner;
  address[] public nodes;
  address[] public clients;
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

  function getNodes() public view returns(address[]) {
    return nodes;
  }

  function getClients() public view returns(address[]) {
    return clients;
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

  /**
   * Add a node to the list of master nodes
   *
   * @param _node Node to add to the list
   */
  function addMasterNode(address _node) public {
    require(msg.sender == owner);
    masterNodes.push(_node);
  }

  /**
   * Add a node to the list of nodes
   *
   * @param _node Node to add to the list
   */
  function addNode(address _node) public {
    require(msg.sender == owner);
    nodes.push(_node);
  }

  /**
   * Add a client to the list of clients
   *
   * @param _node Node to add to the list
   */
  function addClient() external {
    require(msg.sender == owner);
    clients.push(client);
  }
}

pragma solidity ^0.4.19;

import "./Node.sol";
import "./Client.sol";
import "./AbstractBalance.sol";
//ALL DATA VARIABLES ARE SUBJECT TO CHANGE (STC)

contract Pool is AbstractBalance {
  string public publicKey;                           //a public RSA key to encrypt against
  string public data;

  address private owner;                             //msg.sender = marketplace; therefore we need to pass in an owner manually
  bytes32[] nameServers;

  mapping (address => Client) private clients;       //client requesting pool for protection/cdn
  mapping (address => Node) private nodes;         //node information (proposal)

  address[] private client_list;                     //list of client proposals
  address[] private node_list;                       //list of node proposals

  mapping (address => Balance) userBalance;          //maps a client's address to a balance struct

  /**
   * Create new Pool and assign owner
   *
   * Data is assigned owner and uses the owner's public key
   * @param _publicKey Owner's public RSA key to encrypt against
   * @param _owner Address of the owner
   */
  function Pool(string _publicKey, address _owner) public {
    publicKey = _publicKey;
    owner = _owner;
  }

  function setData(string _data) external {
    data = _data;
  }

  /**
   * Makes and returns the Balance struct for a given user
   *
   * @param _user User address to query userBalance array
   * @return (uint,uint,uint,uint)
   */
  function getBalanceStructFor(address _user) public view returns (uint,uint,uint,uint) {
    return (userBalance[_user].owed, userBalance[_user].total, userBalance[_user].completed, userBalance[_user].paid);
  }

  /**
   * Returns the total from a given users balance
   *
   * @param _user User address to query userBalance[_user].total
   * @return (uint)
   */
  function getTotalBalanceFor(address _user) public view returns (uint) {
    return userBalance[_user].total;
  }

  /**
   * Returns the owed balance from a given users balance
   *
   * @param _user User address to query userBalance[_user].owed
   * @return (uint)
   */
  function getOwedBalanceFor(address _user) public view returns (uint) {
    userBalance[_user].owed;
  }

  /**
   * Allocates a clients funds into the Pools balance and the users balance
   *
   * @param _client address
   * @param _amount uint
   * @return bool TODO run through checks to ensure transfer logic passes
   */
  function allocateFundsFrom(address _client, uint _amount) public returns (bool) {
    allocateFunds(_amount);

    Balance storage _userBalance = userBalance[_client];

    userBalance[_client] = Balance({
      owed : _userBalance.owed,
      total : _userBalance.total + _amount,
      completed : _userBalance.completed,
      paid : _userBalance.paid
    });

    return true;
  }

  /**
   * Records an amount of work done by a node to the balance
   *
   * @param _node Node that completed the work
   * @param _client Client that will be billed for the nodes work
   * @param _amount Amount of work completed
   */
  function logWorkFrom(address _node, address _client, uint _amount) public returns (bool) {
    Balance storage _nodeBalance = userBalance[_node];
    Balance storage _clientBalance = userBalance[_client];

    if (work(_amount) && _amount > _clientBalance.total) {
      revert();
    }

    // Node Balance
    userBalance[_node] = Balance({
      owed : _nodeBalance.owed + _amount,
      total : _nodeBalance.total,
      completed : _nodeBalance.completed + _amount,
      paid : _nodeBalance.paid
    });

    // Client Balance
    userBalance[_client] = Balance({
      owed : _clientBalance.owed,
      total : _clientBalance.total - _amount,
      completed : _clientBalance.completed + _amount,
      paid : _clientBalance.paid
    });

    return true;
  }

  /**
   * Recalculates the balances of the pool and user after a payout
   *
   * @param _node Node that completed the work
   * @param _amount Amount of work completed
   * @return bool true if withdraw was successful
   */
  function payout(address _node, uint _amount) public returns (bool) {
    Balance storage _nodeBalance = userBalance[_node];

    if (pay(_amount) && _amount > _nodeBalance.owed) {
      revert();
    }

    userBalance[_node] = Balance({
      owed : _nodeBalance.owed - _amount,
      total : _nodeBalance.total,
      completed : _nodeBalance.completed,
      paid : _nodeBalance.paid + _amount
    });

    return true;
  }

  function getNodeList() public view returns(address[]) {
    return node_list;
  }

  function getClientList() public view returns(address[]) {
    return client_list;
  }

  /**
   * Add a node to the list of nodes only if it's applied
   * called by the node contract durring applyToPool
   *
   */
  function addNode() external {
    Node _newNode = Node(msg.sender);
    require(_newNode.getStatus(address(this)) == 2); //make sure the node has applied to this pool

    node_list.push(msg.sender);
    nodes[msg.sender] = _newNode;
  }

  /**
   * Add a client to the list of clients only if it's applied
   * called by the client contract durring applyToPool
   *
   */
  function addClient() external {
    Client _newClient = Client(msg.sender);
    require(_newClient.getStatus(address(this)) == 2); //make sure the client has applied to this pool

    client_list.push(msg.sender);
    clients[msg.sender] = _newClient;
  }

  /**
   * Accept a node
   *
   * @param _node address of the applying node
   */
  function acceptNode(address _node) public {
    require(msg.sender == owner);
    Node _applicant = Node(_node);
    require(_applicant.getStatus(address(this)) == 2);
    _applicant.setStatus(1);
  }

  /**
   * Accept a client
   *
   * @param _client address of the applying node
   */
  function acceptClient(address _client) public {
    require(msg.sender == owner);
    Client _applicant = Client(_client);
    require(_applicant.getStatus(address(this)) == 2);
    _applicant.setStatus(1);
  }

  /**
   * Remove a member
   * Must be the owner
   * @param _node address to be removed
   */
  function rejectNode(address _node) public {
    require(msg.sender == owner);
    Node _applicant = Node(_node);
    require(_applicant.getStatus(address(this)) == 2);
    _applicant.setStatus(0);
  }

  /**
  * Remove a member
  * Must be the owner
  * @param _client address to be removed
  */
  function rejectClient(address _client) public {
    require(msg.sender == owner);
    Client _applicant = Client(_client);
    require(_applicant.getStatus(address(this)) == 2);
    _applicant.setStatus(0);
  }
}

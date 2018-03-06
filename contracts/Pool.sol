pragma solidity ^0.4.19;

import "./Node.sol";
import "./Client.sol";
import "./AbstractBalance.sol";
//ALL DATA VARIABLES ARE SUBJECT TO CHANGE (STC)

contract Pool is AbstractBalance {
  string public publicKey;                           //a public RSA key to encrypt against

  address private owner;                                     //msg.sender = marketplace; therefore we need to pass in an owner manually
  Client client;                                     //client/website that is employing the pool (1 per pool for Beta)
  bytes32[] nameServers;

  mapping (address => Client) private clients;       //client requesting pool for protection/cdn
  mapping (address => Node) private nodes;           //node information (proposal)

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

  function getNodeList() public returns(address[]) {
    return node_list;
  }

  function getClientList() public returns(address[]) {
    return node_list;
  }

  function getNode(address _node) public returns(Node) {
    return nodes[_node];
  }

  function getClient(address _client) public returns(Client) {
    return clients[_client];
  }

  /** STC
   * Make a proposal to join this pool (from a node)
   *
   * @param _publicKey for encryption
   * @param _data information about this node
   */
  function applyNode(address _node) public {
    require(msg.sender == _node);
    Node _newNode = Node(_node);

    node_list.push(_node);
    nodes[_node] = _newNode;

    _newNode.addPool();
  }

  /** STC
   * Client calls this to apply to this pool
   *
   * @param _publicKey RSA public key
   * @param _data Application or any information that is being sent from the client to the pool
   */
  function applyClient(string _publicKey, string _data) public {
    clients[msg.sender] = Client({
      publicKey : _publicKey,
      name : _data,
      email: "client@gladius.io",
      bio: "hello world",
      ip_address: "1.1.1.1",
      location: "usa",
      wallet_address: msg.sender,
      status: 2,
      exists: true
    });
    client_list.push(msg.sender);
  }

  /**
   * Accept a node
   *
   * @param _node address of the applying node
   */
  function acceptNode(address _node) public {
    require(msg.sender == owner);
    require(nodes[_node].exists);
    nodes[_node].status = 1;
  }

  /**
   * Remove a member
   * Must be the owner
   * @param _node address to be removed
   */
  function rejectNode(address _node) public {
    require(msg.sender == owner);
    require(nodes[_node].exists);
    nodes[_node].status = 0;
  }

  /**
   * Sets the client for this pool (1 client per pool in Beta)
   *
   * @param _clientAddress clientAddress
   */
  function acceptClient(address _clientAddress) public {
    require(msg.sender == owner);
    require(clients[_clientAddress].exists);
    clients[_clientAddress].status = 1;
    client = clients[_clientAddress];
  }

  /**
  * Remove a member
  * Must be the owner
  * @param _client address to be removed
  */
  function rejectClient(address _client) public {
    require(msg.sender == owner);
    require(clients[_client].exists);
    clients[_client].status = 0;
    //$client doesnt change but since it's status is 0 it should still work
  }
}

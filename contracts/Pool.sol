pragma solidity ^0.4.19;

import "./AbstractBalance.sol";
//ALL DATA VARIABLES ARE SUBJECT TO CHANGE (STC)

contract Pool is AbstractBalance {
  string publicKey;                                 //a public RSA key to encrypt against
  bytes32[] nameServers;
  address owner;                                    //msg.sender = marketplace; therefore we need to pass in an owner manually
  mapping (address => string) dataForNode;          //data for the pool to send its nodes (node_address => data) STC
  string dataForClient;                             //data for the pool to send its nodes (node_address => data) STC

  Client client;                                    //client/website that is employing the pool (1 per pool for Beta)

  mapping (address => Client) clients;              //client requesting pool for protection/cdn
  mapping (address => Node) nodes;                  //node information (proposal)

  address[] public client_list;                    //list of client proposals
  address[] public node_list;                      //list of node proposals

  // Struct to store node data
  struct Node {
    string publicKey;

    //these will be encrypted in the future and are STC
    string name;
    string email;
    string bio;
    string ip_address;
    string location;
    address wallet_address;
    int status; // 0 = rejected, 1 = approved, 2 = pending
    bool exists;
  }

  // Struct to store client data
  struct Client {
    string publicKey;

    //these will be encrypted in the future and STC
    string name;
    string email;
    string bio;
    string ip_address;
    string location;
    address wallet_address;
    int status; // 0 = rejected, 1 = approved, 2 = pending
    bool exists;
  }

  mapping (address => Balance) userBalance;         //maps a client's address to a balance struct

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

  /** STC
   * Make a proposal to join this pool (from a node)
   *
   * @param _publicKey for encryption
   * @param _data information about this node
   */
  function applyNode(string _publicKey, string _data) public {
    nodes[msg.sender] = Node({
      publicKey : _publicKey,
      name : _data,
      email: "node@gladius.io",
      bio: "hello world",
      ip_address: "1.1.1.1",
      location: "usa",
      wallet_address: msg.sender,
      status: 2,
      exists: true
    });
    node_list.push(msg.sender);
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

  //STC
  function getNodeData(address _node) constant public returns (string){
    return nodes[_node].name;
  }

  //STC
  function getClientData(address _client) constant public returns (string){
    return clients[_client].name;
  }

  //STC
  function getPoolDataForNode(address _node) constant public returns (string){
    return dataForNode[_node];
  }

  //STC
  function getPoolDataForClient(address _node) constant public returns (string){
    return dataForNode[_node];
  }

  function getNodes() constant public returns (address[]) {
    return node_list;
  }

  function getClients() constant public returns (address[]) {
    return client_list;
  }

  /** STC
   * Update the pool data for a node
   * Must be pool owner to execute
   * @param _node node to update
   * @param _newData data to set for this node
   */
  function updateDataForNode(address _node, string _newData) public {
    require(msg.sender == owner);
    dataForNode[_node] = _newData;
  }

  /**
   * Update the data inside of a node
   * Must be the node to execute
   * @param _node node to update
   * @param _newData data to set for this node
   */
  function updateNodeData(address _node, string _newData) public {
    require(msg.sender == _node);
    nodes[_node].name = _newData;
  }

  /** STC
   * Set the client data variable
   *
   * @param _client client
   * @param _newData newData
   */
  function updateClientData(address _client, string _newData) public {
    require(msg.sender == _client);
    clients[_client].name = _newData;
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
   * @param _node address to be removed
   */
  function rejectNode(address _node) public {
    require(msg.sender == owner);
    require(nodes[_node].exists);
    nodes[_node].status = 0;
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

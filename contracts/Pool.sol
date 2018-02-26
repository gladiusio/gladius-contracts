pragma solidity ^0.4.19;

import "./AbstractBalance.sol";

contract Pool is AbstractBalance {

    // Data about a pool
    struct PoolData {
        mapping (address => string) encryptedPoolData;
        mapping (address => Node) members;

        bytes32[] nameServers;
        string publicKey;
        address[] membersList;
        address owner;
    }

    // Struct to store node data
    struct Node {
        string encryptedData;
        string publicKey;
    }

    struct Client {
        string encryptedData;
        string publicKey;
    }

    PoolData data;
    // Later versions will support multiple clients
    Client client;
    // Encrypted data for the client
    string clientData = "";

    mapping (address => Node) proposals;
    mapping (address => Client) clientRequests;
    // Maps a client's address to a balance struct
    mapping (address => Balance) userBalance;

    address[] private proposedAddresses;


    /**
     * Create new Pool and assign owner
     *
     * Data is assigned owner and uses the owners public key
     * @param publicKey Owners public RSA key to encrypt against
     * @param owner Address of the owner
     */
    function Pool(string publicKey, address owner) public {
        data.owner = owner;
        data.publicKey = publicKey;
    }

    /**
     * Sets the nameServers for the Pool
     *
     * Ensure that the sender is the data owner before allowing changes
     * @param nameServers Name Servers for the Pool
     */
    function setNameServers(bytes32[] nameServers) public {
        require(msg.sender == data.owner);
        data.nameServers = nameServers;
    }

    /**
     * Return the public key for a node member of a node
     *
     * @param node Address of a node
     * @return publicKey RSA public key for the given node
     */
    function getMemberPublicKey(address node) public constant returns (string) {
        return data.members[node].publicKey;
    }

    /**
     * Stores the client encrypted data and the public key
     *
     * @param encryptedData Encrypted data payload
     * @param publicKey RSA public key
     */
    function clientRequest(string encryptedData, string publicKey) public {
        clientRequests[msg.sender] = Client({
            encryptedData : encryptedData,
            publicKey : publicKey
        });
    }

    function getBalanceStructFor(address _user) public view returns (uint, uint, uint, uint) {
      return (userBalance[_user].owed, userBalance[_user].total, userBalance[_user].completed, userBalance[_user].paid);
    }

    function getTotalBalanceFor(address _client) public view returns (uint) {
      return userBalance[_client].total;
    }

    function getOwedBalanceFor(address _node) public view returns (uint) {
      userBalance[_node].owed;
    }

    function allocateFundsFrom(address _client, uint _amount) public returns (bool) {
        allocateFunds(_amount);

        Balance storage _userBalance = userBalance[_client];

        userBalance[_client] = Balance({
          owed : _userBalance.owed,
          total : _userBalance.total + _amount,
          completed : _userBalance.completed,
          paid : _userBalance.paid
        });

        if (userBalance[_client].total != _userBalance.total + _amount) { revert(); }

        return true;
    }

    function authorizeClient(address clientAddress) public {
        require(msg.sender == data.owner);
        client = clientRequests[clientAddress];
    }

    // Set the encrypted data for the client to decrypt
    function setClientData(string clientDataIn) public {
        require(msg.sender == data.owner);
        clientData = clientDataIn;
    }

    function getProposalPublicKey(address node) constant public returns (string) {
        return proposals[node].publicKey;
    }

    function getMemberData(address node) constant public returns (string) {
        return data.members[node].encryptedData;
    }

    function getProposalData(address node) constant public returns (string) {
        return proposals[node].encryptedData;
    }

    function getPoolDataForNode(address node) constant public returns (string) {
        return data.encryptedPoolData[node];
    }

    /// Update the pool data assosiated with the Node node with data newData
    /// Must be owner to execute
    function updateEncryptedPoolData(address node, string newData) public {
        require(msg.sender == data.owner);
        data.encryptedPoolData[node] = newData;
    }

    /// Update the encrypted member data assosiated with the Node node
    /// Must be that member to execute
    function updateEncryptedMemberData(address node, string newData) public {
        require(msg.sender == node);
        data.members[node].encryptedData = newData;
    }

    // Get the data about a proposed node
    function getProposedNodeData(address node) constant public returns (string) {
        return proposals[node].encryptedData;
    }

    /// Make proposal
    function proposeNode(address node, string publicKey, string eData) public {
        proposals[node] = (Node({
          encryptedData : eData,
          publicKey : publicKey
        }));
        proposedAddresses.push(node);
    }

    /// Return a list of addresses of the proposals
    function getProposals() constant public returns (address[]) {
        return proposedAddresses;
    }

    /// Update a member's key. Owner or that member can do it.
    function updateMemberPublicKey(address node, string newKey) public {
        require(msg.sender == node || msg.sender == data.owner);
        data.members[node].publicKey = newKey;
    }

    /// Accept a node to the pool (can only be done by the owner)
    function acceptNode(address node) public {
        require(msg.sender == data.owner);
        data.members[node] = (proposals[node]);
    }

    // Remove a node from the pool
    function removeNode(address node) public {
        require(msg.sender == data.owner);
        delete data.members[node];
    }
}

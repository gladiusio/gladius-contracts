pragma solidity ^0.4.19;

import "./Client.sol";

contract ClientFactory {

  mapping(address=>address) ownerToClient;
  uint256 totalClients = 0;

  /**
   * Create a new node.
   *
   * Instantiate a new Client and set the sender as the owner
   * @param _data data for the client
   * @return address Address to the new Client
   */
   function createClient(string _data) public returns(address) {
     Client client = new Client(_data, msg.sender);
     ownerToClient[msg.sender] = address(client);
     totalClients += 1;

     return address(client);
   }

   function getClientAddress() public returns(address) {
     return ownerToClient[msg.sender];
   }

   function getClientCount() public returns(uint256){
     return totalClients;
   }
}
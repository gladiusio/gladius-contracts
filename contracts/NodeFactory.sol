pragma solidity ^0.4.19;

import "./Node.sol";

contract NodeFactory {

  mapping(address=>address) ownerToNode;
  uint256 totalNodes = 0;

  /**
   * Create a new node.
   *
   * Instantiate a new Node and set the sender as the owner
   * @param _data data for the node
   * @return address Address to the new Node
   */
   function createNode(string _data) public returns(address){
     Node node = new Node(_data, msg.sender);
     ownerToNode[msg.sender] = address(node);
     totalNodes += 1;

     return address(node);
   }

   function getNodeAddress() public returns(address){
     return ownerToNode[msg.sender];
   }

   function getNodeCount() public returns(uint256){
     return totalNodes;
   }

}

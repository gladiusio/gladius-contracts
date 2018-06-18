pragma solidity ^0.4.19;

import "./Pool.sol";

contract Node {
  address owner;
  address [] poolList;

  mapping (address=>int) status; // 0 = not available, 1 = approved, 2 = rejected, 3 = pending

  function Node(address _owner) public {
    owner = _owner;
  }

  function setStatus(int _status) external {
    status[msg.sender] = _status;
  }

  function getStatus(address _pool) public view returns(int) {
    require(msg.sender == _pool || msg.sender == owner);
    return status[_pool];
  }

  function getPoolList() public view returns(address[]) {
    return poolList;
  }

  function getOwner() public view returns(address) {
    return owner;
  }

  /**
  * Apply to be a node of a pool
  *
  * @param _pool address of pool you are applying to
  * @param _data hello
  */
 function applyToPool(address _pool, string _data) public {
   require(msg.sender == owner);
   require(status[_pool] == 0);

   Pool p = Pool(_pool);

   status[_pool] = 3;
   poolList.push(_pool);

   p.addNode();
 }
}

pragma solidity ^0.4.19;

import "./Pool.sol";

contract Node {
  string data; //encrypted data with the NODE'S public key
  address owner;
  mapping (address=>int) status; // 0 = rejected, 1 = approved, 2 = pending
  mapping (address=>string) poolData; //encrypted data with the POOL'S public key
  address [] poolList;

  function Node(string _data, address _owner) public {
    data = _data;
    owner = _owner;
  }

  function setStatus(int _status) external {
    status[msg.sender] = _status;
  }

  function getStatus(address _pool) public view returns(int) {
    require(msg.sender == _pool || msg.sender == owner);
    return status[_pool];
  }

  function getData() public view returns(string) {
    return data;
  }

  function getPoolData(address _pool) public view returns(string) {
    return poolData[_pool];
  }

  function getPoolList() public view returns(address[]) {
    return poolList;
  }

  /**
   * change node's core data
   *
   * @param _data new data
   */
  function changeData(string _data) public {
    require(msg.sender == owner);
    data = _data;
  }

  /**
   * change data provided to pool
   *
   * @param _data new data
   */
  function changePoolData(address _pool, string _data) public {
    require(msg.sender == owner);
    poolData[_pool] = _data;
  }

  /**
   * Apply to be a node of a pool
   *
   * @param _pool address of pool you are applying to
   * @param _data hello
   */
  function applyToPool(address _pool, string _data) public {
    require(msg.sender == owner);

    Pool p = Pool(_pool);

    status[_pool] = 2;
    poolData[_pool] = _data;
    poolList.push(_pool);

    p.addNode();
  }
}

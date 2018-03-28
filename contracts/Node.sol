pragma solidity ^0.4.19;

import "./Pool.sol";

contract Node {
  string public data; //encrypted data with the NODE'S public key
  string public publicData;

  address owner;
  address [] poolList;

  mapping (address=>int) status; // 0 = not available, 1 = approved, 2 = rejected, 3 = pending
  mapping (address=>string) poolData; //encrypted data with the POOL'S public key

  function Node(address _owner) public {
    owner = _owner;
    data = '';
    publicData = '';
  }

  function setData(string _data) external {
    require(msg.sender == owner);
    data = _data;
  }

  function setPublicData(string _publicData) external {
    publicData = _publicData;
  }

  function setStatus(int _status) external {
    status[msg.sender] = _status;
  }

  function getStatus(address _pool) public view returns(int) {
    require(msg.sender == _pool || msg.sender == owner);
    return status[_pool];
  }

  function getPoolData(address _pool) public view returns(string) {
    return poolData[_pool];
  }

  function getPoolList() public view returns(address[]) {
    return poolList;
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
    require(status[_pool] == 0);

    Pool p = Pool(_pool);

    status[_pool] = 3;
    poolData[_pool] = _data;
    poolList.push(_pool);

    p.addNode();
  }
}

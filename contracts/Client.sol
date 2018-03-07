pragma solidity ^0.4.19;

import "./Pool.sol";

contract Client {
  string data; //encrypted data with the CLIENT'S public key
  address owner;
  mapping (address=>int) status; // 0 = rejected, 1 = approved, 2 = pending
  mapping (address=>string) poolData;
  address [] poolList; //encrypted data with the POOL'S public key

  function Client(string _data) public{
    data = _data;
    owner = msg.sender;
  }

  function setStatus(int _status) external {
    status[msg.sender] = _status;
  }

  function getStatus(address _pool) public view returns(int){
    require(msg.sender == _pool || msg.sender == owner);
    return status[_pool];
  }

  /**
   * change client's core data
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
    poolData[address(this)] = _data;
    poolList.push(address(this));

    p.addClient();
  }

}

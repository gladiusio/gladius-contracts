pragma solidity ^0.4.19;

contract Balance {

  address public owner;
  address public pool;

  mapping(address => uint256) public balances;

  constructor(address _owner, address _pool) public {
    owner = _owner;
    pool = _pool;
  }

  function addBalance(address node, uint256 amount) external {
    require(msg.sender == pool);
    balances[node] = amount;
  }


}
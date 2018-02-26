pragma solidity ^0.4.19;

contract AbstractBalance {
  struct Balance {
    uint owed;
    uint total;
    uint completed;
    uint paid;
  }

  Balance public balance;

  function allocateFunds(uint _amount) public {
      // Allocate market funds
      balance.total += _amount;
  }

  function pay(uint _amount) public returns (bool) {
      if (_amount > balance.owed) { return false; }

      balance.owed -= _amount;
      balance.paid += _amount;

      return true;
  }

  function work(uint _amount) public returns (bool) {
      if (_amount > balance.total) { return false; }

      balance.total -= _amount;
      balance.owed += _amount;

      return true;
  }
}

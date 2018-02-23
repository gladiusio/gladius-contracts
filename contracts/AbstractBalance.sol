pragma solidity ^0.4.19;

contract AbstractBalance {
  struct Balance {
    int owed;
    int total;
    int completed;
    int paid;
  }

  Balance public balance;

  function allocateFunds(int _amount) public {
      // Allocate market funds
      balance.total += _amount;
  }

  function getBalance() public view returns (int) {
      return balance.total;
  }

  function pay(int _amount) public returns (bool) {
      if (_amount > owed) { return false; }

      owed -= _amount;
      paid += _amount;

      return true;
  }

  function work(int _amount) public returns (bool) {
      if (_amount > total) { return false; }
      
      total -= _amount;
      owed += _amount;

      return true;
  }
}

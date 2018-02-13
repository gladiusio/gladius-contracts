pragma solidity ^0.4.19;

contract AbstractBalance {
  struct Balance {
    uint256 total;                                 // Total balance        - Complete Balance
    uint256 available;                             // Available Balance    - Allocatable Balance
    uint256 transactionCosts;                      // Market Fee Balance   - Balance to hold transaction fees
    uint256 workable;                              // Work Balance         - Allocated Balance
    uint256 completed;                             // Work Completed       - Work completed balance
    uint256 withdrawable;                          // Payout Balance       - Balance able to withdraw
  }

  Balance public balance;

  function allocateFunds(uint256 _amount) public returns (bool) {
      uint256 _total = _amount;
      uint256 _available = (2 * _amount) / 10;
      uint256 _withdrawable = (2 * _amount) / 10;
      uint256 _transactionCosts = 0; // TODO transaction cost calculation
      uint256 _workable = _amount - _available - _withdrawable;
      uint256 _completed = 0;

      // Allocate market funds
      balance.total += _total;
      balance.available += _available;
      balance.transactionCosts += _transactionCosts;
      balance.workable += _workable;
      balance.completed += _completed;
      balance.withdrawable += _withdrawable;

      // TODO add validation logic to ensure success
      return true;
  }

  function getBalance() public view returns (uint256,uint256,uint256,uint256,uint256,uint256) {
      return (balance.total, balance.available, balance.transactionCosts, balance.workable, balance.completed, balance.withdrawable);
  }

  function getWithdrawable() public view returns (uint256) {
    return balance.withdrawable;
  }

  function withdrawFunds(uint256 _amount) public {
    balance.total -= _amount;
    balance.withdrawable -= _amount;
  }
}

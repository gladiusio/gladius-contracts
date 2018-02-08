pragma solidity ^0.4.19;

contract AbstractBalance {
  struct Balance {
    uint256 total;                                 // Total balance        - Complete Balance
    uint256 available;                             // Available Balance    - Allocatable Balance
    uint256 transactionCosts;                      // Market Fee Balance   - Balance to hold transaction fees
    uint256 workable;                              // Work Balance         - Allocated Balance
    uint256 completed;                             // Work Completed       - Work completed balance
    uint256 transferrable;                         // Payout Balance       - Balance able to withdraw
    uint256 withdrawable;                          // Payout Balance       - Balance able to withdraw
  }
}

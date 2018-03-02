pragma solidity ^0.4.19;

contract AbstractBalance {
  struct Balance {
    uint owed;
    uint total;
    uint completed;
    uint paid;
  }

  Balance public balance;

  /**
   * Allocate funds by incrementing the total
   *
   * Used to increase the total amount of funds available for the Marketplace or Pool 
   * @param _amount The amount of funds to allocate
   * @return bool Stub for now, but will return true / false when actual funds are used
   */
  function allocateFunds(uint _amount) public returns (bool) {
    // Allocate market funds
    balance.total += _amount;
    return true;
  }

  /**
   * Payout
   *
   * Payout and rebalance the accounts for the Marketplace or Pool
   * @param _amount The amount of funds to allocate
   * @return bool Stub for now, but will return true / false when actual funds are used
   */
  function pay(uint _amount) public returns (bool) {
    if (_amount > balance.owed) { return false; }

    balance.owed -= _amount;
    balance.paid += _amount;

    return true;
  }

  /**
   * Work
   *
   * Allocates work amount for work that has been calculated for a node
   * @param _amount The amount of funds to allocate
   * @return bool Stub for now, but will return true / false when actual funds are used
   */
  function work(uint _amount) public returns (bool) {
    if (_amount > balance.total) { return false; }

    balance.total -= _amount;
    balance.owed += _amount;
    balance.completed += _amount;

    return true;
  }
}

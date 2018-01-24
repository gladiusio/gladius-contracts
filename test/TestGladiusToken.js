// Specifically request an abstraction for MetaCoin
let PoolContract = artifacts.require('Pool');
let MarketContract = artifacts.require('Market');
let GladiusToken = artifacts.require('GladiusToken');

contract('GladiusToken', function(accounts) {
  // Accounts
  let creator = accounts[0]
  let user = accounts[1]

  let initialSupply = 10000 * (10**8)

  function getInstance() {
    return GladiusToken.deployed().then(function(instance) {
      return instance
    })
  }

  // Test creation of tokens
  it('Deploy Tokens to Creator\'s Account', function() {
    // Create initial supply of tokens
    getInstance().then(function(instance) {
      return instance.balanceOf.call(creator)
    }).then(function(balance) {
      assert.equal(balance.valueOf(), initialSupply, 'Initial tokens were not allocated to the creator\'s account')
    })  
  })

  // Test balance transfers
  it('Simulate transfer from creator to user', function() {
    let self
    
    // Balance ladger
    let creatorStartBalance
    let userStartBalance
    let creatorEndBalance
    let userEndBalance

    let transferAmount = 10
    
    // Grab instance
    getInstance().then(function(instance) {
      self = instance
      return self.balanceOf.call(creator) 
    }).then(function(creatorBalance) {
      creatorStartBalance = creatorBalance.toNumber()     // Save creator's starting balance
      return self.balanceOf.call(user)
    }).then(function(userBalance) {
      userStartBalance = userBalance.toNumber()           // Save user's starting balance
      return self.transfer(user, transferAmount)          // Initiate Transfer
    }).then(function() {
      return self.balanceOf.call(creator) 
    }).then(function(creatorBalance) {
      creatorEndBalance = creatorBalance.toNumber()       // Save creator's ending balance
      return self.balanceOf.call(user)
    }).then(function(userBalance) {
      userEndBalance = userBalance.toNumber()             // Save creator's ending balance

      // Compare transfers to ledger math
      assert.equal(creatorEndBalance, creatorStartBalance - transferAmount, 'Amount wasn\'t correctly taken from the sender')
      assert.equal(userEndBalance, userStartBalance + transferAmount, 'Amount wasn\'t correctly sent to the receiver')
    })
  })

  it('Simulate burning tokens', function() {
    let self
    let burnAmount = 50
    let startBalance
    
    getInstance().then(function(instance) {
      self = instance
      return self.totalSupply.call()
    }).then(function(total) {
      startBalance = total.toNumber()
      return self.burn(burnAmount)
    }).then(function() {
      return self.totalSupply.call()
    }).then(function(endTotal) {
      console.log(startBalance)
      console.log(burnAmount)
      let endingSupply = endTotal.valueOf()
      console.log(endingSupply)
      assert.equal(endingSupply, startBalance - burnAmount, 'Tokens did not successfully burn')
    })
  })
})

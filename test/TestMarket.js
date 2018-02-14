let Pool = artifacts.require('Pool')
let Market = artifacts.require('Market')
let GladiusToken = artifacts.require('GladiusToken')

contract('Market', function(accounts) {
  // Accounts
  let owner = accounts[0]
  let user = accounts[1]

  // Test creation of marketplace
  describe('Test Market Contract', function() {
    it('Check Owner', async function() {
      let market = await Market.deployed()
      let marketOwner = await market.owner.call()

      assert.equal(marketOwner, owner, 'Market\'s owner is not the creator of the market')
    })

    it('Pool creation and addition', async function() {
      let market = await Market.deployed()

      let initialMarketPools = await market.getOwnedPools.call(user)

      market.createPool('FAKE_KEY', {from: user})
      market.createPool('FAKE_KEY', {from: owner})
      market.createPool('FAKE_KEY', {from: user})

      let endingMarketPools = await market.getOwnedPools.call(user)

      // Assert that 2 of the 3 pools created are owned by the user
      assert.equal(initialMarketPools.length + 2, endingMarketPools.length, 'Pools added to the marketplace do not equal the market pools')
    })

    it('Pool token allocation', async function () {
      let market = await Market.deployed()

      let amount = 10

      let userOwnedPools = await market.getOwnedPools.call(user)
      let userOwnedPool = Pool.at(userOwnedPools[0])

      await market.allocateClientFundsTo(userOwnedPools[0], user, amount, { from: user })

      let ownerOwnedPools = await market.getOwnedPools.call(owner)
      let ownerOwnedPool = Pool.at(userOwnedPools[0])

      await market.allocateClientFundsTo(ownerOwnedPools[0], owner, amount, { from: owner })

      let balance = await market.balance.call()

      // Keeping console.logs for debugging
      console.log("\n\nAllocated 2 Clients' Funds to Marketplace...")

      console.log("\n\n----- Market Balance ------")
      console.log("Total: " + balance[0].toNumber())
      console.log("Available: " + balance[1].toNumber())
      console.log("TransactionCosts: " + balance[2].toNumber())
      console.log("Workable: " + balance[3].toNumber())
      console.log("Completed: " + balance[4].toNumber())
      console.log("Withdrawable: " + balance[5].toNumber())

      let userBalance = await userOwnedPool.getuserBalance(user)

      console.log("\n\n------ Client Balance ------")
      console.log("Total: " + userBalance[0].toNumber())
      console.log("Available: " + userBalance[1].toNumber())
      console.log("TransactionCosts: " + userBalance[2].toNumber())
      console.log("Workable: " + userBalance[3].toNumber())
      console.log("Completed: " + userBalance[4].toNumber())
      console.log("Withdrawable: " + userBalance[5].toNumber())

      let poolBalance = await userOwnedPool.getBalance()

      console.log("\n\n------ Pool Balance ------")
      console.log("Total: " + poolBalance[0].toNumber())
      console.log("Available: " + poolBalance[1].toNumber())
      console.log("TransactionCosts: " + poolBalance[2].toNumber())
      console.log("Workable: " + poolBalance[3].toNumber())
      console.log("Completed: " + poolBalance[4].toNumber())
      console.log("Withdrawable: " + poolBalance[5].toNumber())

      assert.equal(balance[0].toNumber(), amount * 2, 'Total market balance is not twice the allocated test amount')
      assert.equal(userBalance[0].toNumber(), amount, 'Total client balance is not the allocated test amount')
      assert.equal(userBalance[1].toNumber(), (amount * 0.2), 'Total client balance is not the allocated test amount')
    })

    it('Market withdrawal', async function () {
      let market = await Market.deployed()

      let amount = 2
      let overageAmount = 8

      let userOwnedPools = await market.getOwnedPools.call(user)
      let userOwnedPool = Pool.at(userOwnedPools[0])

      let balance = await market.balance.call()

      // Keeping console.logs for debugging
      console.log("\n\nChecking Market Balance")

      console.log("\n\n----- Market Balance ------")
      console.log("Total: " + balance[0].toNumber())
      console.log("Available: " + balance[1].toNumber())
      console.log("TransactionCosts: " + balance[2].toNumber())
      console.log("Workable: " + balance[3].toNumber())
      console.log("Completed: " + balance[4].toNumber())
      console.log("Withdrawable: " + balance[5].toNumber())

      let userBalance = await userOwnedPool.getuserBalance(user)
      console.log("\n\n----- User Balance ------")
      console.log("Total: " + userBalance[0].toNumber())
      console.log("Available: " + userBalance[1].toNumber())
      console.log("TransactionCosts: " + userBalance[2].toNumber())
      console.log("Workable: " + userBalance[3].toNumber())
      console.log("Completed: " + userBalance[4].toNumber())
      console.log("Withdrawable: " + userBalance[5].toNumber())

      await market.withdraw(userOwnedPools[0], user, overageAmount)

      let balanceAfter = await market.balance.call()

      assert.equal(balanceAfter[5].toNumber(), balance[5].toNumber(), 'Money was withdrawn outside of the daily limit')

      let canWithdraw = await market.withdraw(userOwnedPools[0], user, amount)
      let balanceFinal = await market.balance.call()

      console.log("\n\n----- Market Ending Balance ------")
      console.log("Total: " + balanceFinal[0].toNumber())
      console.log("Available: " + balanceFinal[1].toNumber())
      console.log("TransactionCosts: " + balanceFinal[2].toNumber())
      console.log("Workable: " + balanceFinal[3].toNumber())
      console.log("Completed: " + balanceFinal[4].toNumber())
      console.log("Withdrawable: " + balanceFinal[5].toNumber())

      assert.equal(balanceFinal[5].toNumber(), balance[5].toNumber() - amount, 'Money did not properly withdraw')

      let userEndingBalance = await userOwnedPool.getuserBalance(user)
      console.log("\n\n----- User Ending Balance ------")
      console.log("Total: " + userEndingBalance[0].toNumber())
      console.log("Available: " + userEndingBalance[1].toNumber())
      console.log("TransactionCosts: " + userEndingBalance[2].toNumber())
      console.log("Workable: " + userEndingBalance[3].toNumber())
      console.log("Completed: " + userEndingBalance[4].toNumber())
      console.log("Withdrawable: " + userEndingBalance[5].toNumber())
    })
  })
})

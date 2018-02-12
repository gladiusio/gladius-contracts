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
      console.log("Transferrable: " + balance[5].toNumber())
      console.log("Withdrawable: " + balance[6].toNumber())

      let clientBalance = await userOwnedPool.getClientBalance(user)

      console.log("\n\n------ Client Balance ------")
      console.log("Total: " + clientBalance[0].toNumber())
      console.log("Available: " + clientBalance[1].toNumber())
      console.log("TransactionCosts: " + clientBalance[2].toNumber())
      console.log("Workable: " + clientBalance[3].toNumber())
      console.log("Completed: " + clientBalance[4].toNumber())
      console.log("Transferrable: " + clientBalance[5].toNumber())
      console.log("Withdrawable: " + clientBalance[6].toNumber())

      let poolBalance = await userOwnedPool.getBalance()

      console.log("\n\n------ Pool Balance ------")
      console.log("Total: " + poolBalance[0].toNumber())
      console.log("Available: " + poolBalance[1].toNumber())
      console.log("TransactionCosts: " + poolBalance[2].toNumber())
      console.log("Workable: " + poolBalance[3].toNumber())
      console.log("Completed: " + poolBalance[4].toNumber())
      console.log("Transferrable: " + poolBalance[5].toNumber())
      console.log("Withdrawable: " + poolBalance[6].toNumber())

      assert.equal(balance[0].toNumber(), amount * 2, 'Total market balance is not twice the allocated test amount')
      assert.equal(clientBalance[0].toNumber(), amount, 'Total client balance is not the allocated test amount')
      assert.equal(clientBalance[1].toNumber(), (amount * 0.6), 'Total client balance is not the allocated test amount')
    })
  })
})

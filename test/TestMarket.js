let Pool = artifacts.require('Pool')
let Market = artifacts.require('Market')
let GladiusToken = artifacts.require('GladiusToken')

function printBalance(balance) {
  console.log("Owed: " + balance[0].toNumber())
  console.log("Total: " + balance[1].toNumber())
  console.log("Completed: " + balance[2].toNumber())
  console.log("Paid: " + balance[3].toNumber())
}

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

    // it('Pool token allocation', async function () {
    //   let market = await Market.deployed()
    //
    //   let amount = 10
    //
    //   let userOwnedPools = await market.getOwnedPools.call(user)
    //   let userOwnedPool = Pool.at(userOwnedPools[0])
    //
    //   await market.allocateClientFundsTo(userOwnedPools[0], user, amount, { from: user, gas: 999999 })
    //
    //   let ownerOwnedPools = await market.getOwnedPools.call(owner)
    //   let ownerOwnedPool = Pool.at(userOwnedPools[0])
    //
    //   await market.allocateClientFundsTo(ownerOwnedPools[0], owner, amount, { from: owner, gas: 999999 })
    //
    //   // let balance = await market.balance.call()
    //
    //   // Keeping console.logs for debugging
    //   console.log("\n\nAllocated 2 Clients' Funds to Marketplace...")
    //
    //   console.log("\n\n----- Market Balance ------")
    //   // printBalance(balance)
    //
    //   // let userBalance = await userOwnedPool.getBalanceStructFor(user)
    //   // console.log(userBalance)
    //
    //   // console.log(userOwnedPool.balance)
    //   // printBalance(userBalance)
    //
    //   let poolBalance = await userOwnedPool.balance.call()
    //
    //   console.log("\n\n------ Pool Balance ------")
    //   printBalance(poolBalance)
    //
    //   // assert.equal(balance[1].toNumber(), amount * 2, 'Total market balance is not twice the allocated test amount')
    //   // assert.equal(userBalance[1].toNumber(), amount, 'Total client balance is not the allocated test amount')
    //
    //   assert.eqaul(poolBalance.total,10, '')
    // })
  })
})

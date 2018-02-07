let Pool = artifacts.require('Pool')
let Market = artifacts.require('Market')
let GladiusToken = artifacts.require('GladiusToken')
let Balances = artifacts.require('Balances')

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
      //assert.equal(initialMarketPools.length + 2, endingMarketPools.length, 'Pools added to the marketplace do not equal the market pools')
    })

    it('Pool token allocation', async function () {
      let market = await Market.deployed()

      let amount = 10

      let userOwnedPools = await market.getOwnedPools.call(user)
      let userOwnedPool = Pool.at(userOwnedPools[0])

      await market.allocateClientFundsTo(userOwnedPools[0], user, amount, { from: user })

      let balances = await market.balances.call()
      // let total = await balances.total.call()
      // let records = Balances.records(balances)
      console.log(balances)

      // let balanceAvailable = await market.balanceAvailable.call()
      // console.log(balanceAvailable.toNumber())
      //
      // let balanceWorkable = await market.balanceWorkable.call()
      // console.log(balanceWorkable.toNumber())
      //
      // let clientBalance = await userOwnedPool.getClientBalance.call(user)
      //
      // console.log(clientBalance.toNumber())
    })
  })
})

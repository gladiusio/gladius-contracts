let Pool = artifacts.require('Pool')
let Market = artifacts.require('Market')
let GladiusToken = artifacts.require('GladiusToken')
let NodeFactory = artifacts.require('NodeFactory')
let ClientFactory = artifacts.require('ClientFactory')
let Node = artifacts.require('Node')
let Client = artifacts.require('Client')

contract('Market', function(accounts) {
  // Accounts
  let owner = accounts[0]
  let user = accounts[1]

  // Test creation of marketplace
  describe('Test Market Contract', function() {
    it('Check Owner', async function() { // jshint ignore:line
      let market = await Market.deployed()
      let marketOwner = await market.owner.call()

      assert.equal(marketOwner, owner, 'Market\'s owner is not the creator of the market')
    }) // jshint ignore:line

    it('Pool creation and addition', async function() { // jshint ignore:line
      let market = await Market.deployed()

      let initialMarketPools = await market.getOwnedPools(user)

      await market.createPool('FAKE_KEY', {from: user})
      await market.createPool('FAKE_KEY', {from: owner})
      await market.createPool('FAKE_KEY', {from: user})

      let endingMarketPools = await market.getOwnedPools(user)

      // Assert that 2 of the 3 pools created are owned by the user
      assert.equal(initialMarketPools.length + 2, endingMarketPools.length, 'Pools added to the marketplace do not equal the market pools')
    }) // jshint ignore:line
  })
})
let Pool = artifacts.require('Pool')
let Market = artifacts.require('Market')
let GladiusToken = artifacts.require('GladiusToken')

contract('Market', function(accounts) {
  // Accounts
  let owner = accounts[0]
  let user = accounts[1]

  // Test creation of marketplace
  describe('Market Contract', function() {
    it('Check Market Owner', async function() {
      let market = await Market.deployed()
      let marketOwner = await market.owner.call()
      
      assert.equal(marketOwner, owner, 'Market\'s owner is not the creator of the market')
    })
  })
})

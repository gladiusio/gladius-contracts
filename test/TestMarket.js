let Pool = artifacts.require('Pool')
let Market = artifacts.require('Market')
let GladiusToken = artifacts.require('GladiusToken')

contract('Market', function(accounts) {
  // Accounts
  let marketOwner = accounts[0]
  let user = accounts[1]

  // Test creation of marketplace
  describe('Market Contract', function() {
    it('Deploy Tokens to Creator\'s Account', async function() {
      //let gladius = await GladiusToken.deployed()
      let market = await Market.deployed()
      console.log(market)
    })
  })
})

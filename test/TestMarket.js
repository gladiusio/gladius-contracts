let PoolFactory1 = artifacts.require('PoolFactory')
let PoolFactory2 = artifacts.require('PoolFactoryTest')
let Market = artifacts.require('Market')

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

    it('Pool creation', async function() { // jshint ignore:line
      let market = await Market.deployed()

      let initialMarketPools = await market.getOwnerAllPools(user)

      await market.createPool({from: user})
      await market.createPool({from: owner})
      await market.createPool({from: user})

      let endingMarketPools = await market.getOwnerAllPools(user)

      // Assert that 2 of the 3 pools created are owned by the user
      assert.equal(initialMarketPools.length + 2, endingMarketPools.length, 'Pools added to the marketplace do not equal the market pools')
    }) // jshint ignore:line

    it('Check Pool ownership', async function() { // jshint ignore:line
      let market = await Market.deployed()
      let userPools = await market.getOwnerAllPools(user)
      let ownerPools = await market.getOwnerAllPools(owner)

      await market.addToMarketplace(userPools[0], {from:owner})

      let userMarketPools = await market.getOwnerMarketPools(user)

      assert.equal(ownerPools.length, 1, "Owner missing pools")
      assert.equal(userPools.length, 2, "User missing pools")
      assert.equal(userMarketPools.length, 1, "User missing market pools")

    }) // jshint ignore:line

    it('Switch the pool factory', async function() { // jshint ignore:line
      let market = await Market.deployed()
      let pools1 = await PoolFactory2.deployed()

      await market.changePoolFactory(pools1.address, {from: owner})

      let pools2 = await market.getPoolFactory()

      assert.equal(pools1.address, pools2, "Market did not switch pool factories")
    }) // jshint ignore:line

    it('Get/Set data', async function() {
      let market = await Market.deployed()

      await market.setData("data goes here", {from:owner})

      try {
        await market.setData("data doesnt go here", {from:accounts[1]})
      } catch (error) {

      }

      let data = await market.getData()

      assert.equal(data, "data goes here", "Data changed or not set")
    })

  })
})
let Pool = artifacts.require('Pool')
let Market = artifacts.require('Market')
let GladiusToken = artifacts.require('GladiusToken')

contract('Pool', async function(accounts) {
  // Accounts
  let owner = accounts[0]

  describe('Test Pool Contract', function() {
    it('Check owner', async function() {
      let market = await Market.deployed()
      await market.createPool({from: owner})

      let allPools = await market.getAllPools.call();
      let poolOwner = await market.getPoolOwner.call(allPools[0]);

      assert.equal(poolOwner, owner, 'Pool\'s owner is not the creator of the market')
    })

    it('Get/Set masternodes', async function() {
      let market = await Market.deployed()
      let allPools = await market.getAllPools.call();
      let pool = await Pool.at(allPools[0])

      await pool.addMasterNode(accounts[1], {from:owner})

      try{
        await pool.addMasterNode(accounts[2], {from:accounts[1]})
      } catch (error) {

      }

      let masternodes = await pool.getMasterNodes()

      assert.equal(masternodes.length, 1, 'Masternodes not added to list')
    })

    it('Get/Set seed node', async function() {
      let market = await Market.deployed()
      let allPools = await market.getAllPools.call();
      let pool = await Pool.at(allPools[0])

      await pool.setSeedNode("1.1.1", {from:owner})

      try {
        await pool.setSeedNode("2.2.2", {from:accounts[1]})
      } catch (error) {

      }

      let seedNode = await pool.getSeedNode()

      assert.equal(seedNode, "1.1.1", 'Seed node not added to list')
    })

    it('Get/Set data', async function() {
      let market = await Market.deployed()
      let allPools = await market.getAllPools.call();
      let pool = await Pool.at(allPools[0])

      await pool.setData("data goes here", {from:owner})

      try {
        await pool.setData("data doesnt go here", {from:accounts[1]})
      } catch (error) {

      }

      let data = await pool.getData()

      assert.equal(data, "data goes here","Data changed or not set")
    })

  })
})

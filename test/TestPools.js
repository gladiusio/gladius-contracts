let Pool = artifacts.require('Pool')
let Market = artifacts.require('Market')
let GladiusToken = artifacts.require('GladiusToken')

contract('Market', function(accounts) {
  // Accounts
  let owner = accounts[0]
  let node = accounts[1]
  let client = accounts[2]

  // Test creation of marketplace
  describe('Test Pool Contract', function() {
    it('Check Owner', async function() {
      let market = await Market.deployed()

      await market.createPool("TEST_KEY1", {from: owner})

      let plist = await market.getAllPools.call();
      let poolOwner = await market.poolToOwner(plist[0]);

      assert.equal(poolOwner, owner, 'Pool\'s owner is not the creator of the market')
    })

    it('Apply as Node', async function() {
      let market = await Market.deployed()

      await market.createPool("TEST_KEY2", {from: owner})

      let plist = await market.getAllPools.call()
      console.log(plist);

      let pool = Pool.at(plist[0])
      console.log("my pool:" + plist[0]);
      console.log("======================");

      // await pool.applyNode("public_key", "celo-node", {from: node})

      let key = await pool.publicKey.call()
      console.log(key)
      console.log("======================");
      // console.log("===RE ROO===");
      // console.log(pool.nodes.call('asd'));

      // assert.equal(allPoolsList, owner, 'Node applicant is added to list')
    })

  })
})

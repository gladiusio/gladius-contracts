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

      let initialMarketPools = await market.getOwnedPools(user)

      market.createPool('FAKE_KEY', {from: user})
      market.createPool('FAKE_KEY', {from: owner})
      market.createPool('FAKE_KEY', {from: user})

      let endingMarketPools = await market.getOwnedPools(user)

      // Assert that 2 of the 3 pools created are owned by the user
      assert.equal(initialMarketPools.length + 2, endingMarketPools.length, 'Pools added to the marketplace do not equal the market pools')
    })

    it('Allocate Client funds to a Pool', async function() {
      let market = await Market.deployed()

      let amount = 10

      let userOwnedPools = await market.getOwnedPools(user)
      let userOwnedPool = Pool.at(userOwnedPools[0])

      let ownerOwnedPools = await market.getOwnedPools(owner)
      let ownerOwnedPool = Pool.at(ownerOwnedPools[0])

      await market.allocateClientFundsTo(userOwnedPools[0], user, amount, { from: user })
      await market.allocateClientFundsTo(ownerOwnedPools[0], owner, amount, { from: owner })
      
      let userBalance = await userOwnedPool.getBalanceStructFor(user)
      console.log('\n\nUser Balance Sheet')
      printBalance(userBalance)
      
      let ownerBalance = await ownerOwnedPool.getBalanceStructFor(owner)
      console.log('\n\nOwner Balance Sheet')
      printBalance(ownerBalance)

      assert.equal(userBalance[1], amount, 'Amount allocated to user\'s balance does not equal ' + amount)
      assert.equal(ownerBalance[1], amount, 'Amount allocated to owner\'s balance does not equal ' + amount)

      let marketBalance = await market.balance.call()
      console.log('\n\nMarket Balance Sheet')
      printBalance(marketBalance)
      assert.equal(marketBalance[1], 2 * amount, 'Amount allocated to market\'s balance does not equal ' + (2 * amount))
    })

    it('Payout Node from Pool\'s payout escrow account', async function() {
      let market = await Market.deployed()

      let amount = 1

      let userOwnedPools = await market.getOwnedPools(user)
      let userOwnedPool = Pool.at(userOwnedPools[0])

      await userOwnedPool.applyNode('fake_key', 'name')

      let node = await userOwnedPool.node_list.call(0)
      console.log(node)
      await market.logWorkFrom.call(userOwnedPools[0], node, user, amount)

      let userBalance = await userOwnedPool.getBalanceStructFor(user)
      console.log('\n\nClient Balance Sheet')
      printBalance(userBalance)
      
      let nodeBalance = await userOwnedPool.getBalanceStructFor(node)
      console.log('\n\nNode Balance Sheet')
      printBalance(nodeBalance)
    })
  })
})

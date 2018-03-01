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
      let ownerBalance = await ownerOwnedPool.getBalanceStructFor(owner)
      let marketBalance = await market.balance.call()
      
      assert.equal(userBalance[1], amount, 'Amount allocated to user\'s balance does not equal ' + amount)
      assert.equal(ownerBalance[1], amount, 'Amount allocated to owner\'s balance does not equal ' + amount)
      assert.equal(marketBalance[1], 2 * amount, 'Amount allocated to market\'s balance does not equal ' + (2 * amount))
    })

    it('Payout Node from Pool\'s payout escrow account', async function() {
      let market = await Market.deployed()

      let amount = 1

      let userOwnedPools = await market.getOwnedPools(user)
      let userOwnedPool = Pool.at(userOwnedPools[0])

      await userOwnedPool.applyNode('fake_key', 'name')

      let node = await userOwnedPool.node_list.call(0)

      let startingPoolBalance = await userOwnedPool.balance.call()
      let startingUserBalance = await userOwnedPool.getBalanceStructFor(user)
      let startingNodeBalance = await userOwnedPool.getBalanceStructFor(node)
      let startingMarketBalance = await market.balance.call()
      
      await market.logWorkFrom(userOwnedPools[0], node, user, amount)

      let poolBalance = await userOwnedPool.balance.call()
      let userBalance = await userOwnedPool.getBalanceStructFor(user)
      let nodeBalance = await userOwnedPool.getBalanceStructFor(node)
      let marketBalance = await market.balance.call()

      // Check Market balance
      assert.equal(marketBalance[0].toNumber(), amount, 'Market: Amount owed (expected:' + amount + ') does not equal ' + amount + ' worked')
      assert.equal(marketBalance[1].toNumber(), startingMarketBalance[1].toNumber() - amount, 'Market: Total amount did not decrease by ' + amount)
      assert.equal(marketBalance[2].toNumber(), amount, 'Market: Amount completed (expected:' + amount + ') does not equal ' + amount + ' completed')
      // Check Pool Balance
      assert.equal(poolBalance[0].toNumber(), amount, 'Pool: Amount owed (expected:' + amount + ') does not equal ' + amount + ' worked')
      assert.equal(poolBalance[1].toNumber(), startingPoolBalance[1].toNumber() - amount, 'Pool: Total amount did not decrease by ' + amount)
      assert.equal(poolBalance[2].toNumber(), amount, 'Pool: Amount completed (expected:' + amount + ') does not equal ' + amount + ' completed')
      // Check Node Balance
      assert.equal(nodeBalance[0].toNumber(), amount, 'Node: Amount owed (expected:' + amount + ') does not equal ' + amount + ' worked')
      assert.equal(nodeBalance[1].toNumber(), startingNodeBalance[1], 'Node: Total amount changed unexpectedly')
      assert.equal(nodeBalance[2].toNumber(), amount, 'Node: Amount completed (expected:' + amount + ') does not equal ' + amount + ' completed')
      // Check Client Balance
      assert.equal(userBalance[0].toNumber(), startingUserBalance[0].toNumber(), 'Client: Amount owed  changed unexpectedly')
      assert.equal(userBalance[1].toNumber(), startingUserBalance[1].toNumber() - amount, 'Client: Total amount did not decrease by ' + amount)
      assert.equal(userBalance[2].toNumber(), amount, 'Client: Amount completed (expected:' + amount + ') does not equal ' + amount + ' completed')
    })
  })
})

function printBalance(balance) {
  console.log("Owed: " + balance[0].toNumber())
  console.log("Total: " + balance[1].toNumber())
  console.log("Completed: " + balance[2].toNumber())
  console.log("Paid: " + balance[3].toNumber())
}


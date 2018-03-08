let Pool = artifacts.require('Pool')
let Market = artifacts.require('Market')
let GladiusToken = artifacts.require('GladiusToken')
let Node = artifacts.require('Node')
let Client = artifacts.require('Client')

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

    it('Add work to a node\'s  balance', async function() {
      let market = await Market.deployed()
      let node = await Node.deployed()

      let amount = 1

      let userOwnedPools = await market.getOwnedPools(user)
      let userOwnedPool = Pool.at(userOwnedPools[0])

      await node.applyToPool.sendTransaction(userOwnedPools[0], 'name', {from: user})
      //
      // console.log(node.address);

      let startingPoolBalance = await userOwnedPool.balance.call()
      let startingUserBalance = await userOwnedPool.getBalanceStructFor(user)
      let startingNodeBalance = await userOwnedPool.getBalanceStructFor(node.address)
      let startingMarketBalance = await market.balance.call()

      await market.logWorkFrom(userOwnedPools[0], node.address, user, amount)

      let poolBalance = await userOwnedPool.balance.call()
      let userBalance = await userOwnedPool.getBalanceStructFor(user)
      let nodeBalance = await userOwnedPool.getBalanceStructFor(node.address)
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

    it('Pay out a node from owed balance', async function() {
      let market = await Market.deployed()

      let amount = 1

      let userOwnedPools = await market.getOwnedPools(user)
      let userOwnedPool = Pool.at(userOwnedPools[0])
      let nodeList = await userOwnedPool.getNodeList()
      let node = nodeList[0]

      let startingPoolBalance = await userOwnedPool.balance.call()
      let startingUserBalance = await userOwnedPool.getBalanceStructFor(user)
      let startingNodeBalance = await userOwnedPool.getBalanceStructFor(node)
      let startingMarketBalance = await market.balance.call()

      await market.payout(userOwnedPools[0], node, amount)

      let poolBalance = await userOwnedPool.balance.call()
      let userBalance = await userOwnedPool.getBalanceStructFor(user)
      let nodeBalance = await userOwnedPool.getBalanceStructFor(node)
      let marketBalance = await market.balance.call()

      // Check Market balance
      assert.equal(marketBalance[0].toNumber(), startingMarketBalance[0].toNumber() - amount, 'Market:  (expected:' + amount + ') does not equal ' + amount + ' owed')
      assert.equal(marketBalance[1].toNumber(), startingMarketBalance[1].toNumber(), 'Market: Total amount changed unexpectedly')
      assert.equal(marketBalance[2].toNumber(), startingMarketBalance[2].toNumber(), 'Market: Completed (expected:' + amount + ') does not equal ' + amount + ' completed')
      assert.equal(marketBalance[3].toNumber(), amount, 'Market: Paid (expected:' + amount + ') does not equal ' + amount + ' paid')

      // Check Pool Balance
      assert.equal(poolBalance[0].toNumber(), startingPoolBalance[0].toNumber() - amount, 'Pool:  (expected:' + amount + ') does not equal ' + amount + ' owed')
      assert.equal(poolBalance[1].toNumber(), startingPoolBalance[1].toNumber(), 'Pool: Total amount changed unexpectedly')
      assert.equal(poolBalance[2].toNumber(), startingPoolBalance[2].toNumber(), 'Pool: Completed (expected:' + amount + ') does not equal ' + amount + ' completed')
      assert.equal(poolBalance[3].toNumber(), amount, 'Pool: Paid (expected:' + amount + ') does not equal ' + amount + ' paid')
      // Check Node Balance
      assert.equal(nodeBalance[0].toNumber(), startingNodeBalance[0].toNumber() - amount, 'Node:  (expected:' + amount + ') does not equal ' + amount + ' owed')
      assert.equal(nodeBalance[1].toNumber(), startingNodeBalance[1].toNumber(), 'Node: Total amount changed unexpectedly')
      assert.equal(nodeBalance[2].toNumber(), startingNodeBalance[2].toNumber(), 'Node: Completed (expected:' + amount + ') does not equal ' + amount + ' completed')
      assert.equal(nodeBalance[3].toNumber(), amount, 'Node: Paid (expected:' + amount + ') does not equal ' + amount + ' paid')
      // Check Client Balance
      assert.equal(userBalance[0].toNumber(), startingUserBalance[0].toNumber(), 'Node: Owed amount changed unexpectedly')
      assert.equal(userBalance[1].toNumber(), startingUserBalance[1].toNumber(), 'Node: Total amount changed unexpectedly')
      assert.equal(userBalance[2].toNumber(), startingUserBalance[2].toNumber(), 'Node: Completed amount changed unexpectedly')
      assert.equal(userBalance[3].toNumber(), startingUserBalance[3].toNumber(), 'Node: Paid amount changed unexpectedly')
    })
  })
})

function printBalance(balance) {
  console.log("Owed: " + balance[0].toNumber())
  console.log("Total: " + balance[1].toNumber())
  console.log("Completed: " + balance[2].toNumber())
  console.log("Paid: " + balance[3].toNumber())
}

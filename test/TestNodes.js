let Pool = artifacts.require('Pool')
let Market = artifacts.require('Market')
let GladiusToken = artifacts.require('GladiusToken')
let NodeFactory = artifacts.require('NodeFactory')
let ClientFactory = artifacts.require('ClientFactory')
let Node = artifacts.require('Node')
let Client = artifacts.require('Client')

contract('Node', function(accounts) {
  // Accounts
  let owner = accounts[0]
  let nodeAddress4 = accounts[4]
  let clientAddress3 = accounts[7]

  describe('Test Node Contract', function() {

    it('Create a node', async function() {
      let market = await Market.deployed()
      let nFactory = await NodeFactory.deployed()

      await nFactory.createNode.sendTransaction("node4_data", {from:nodeAddress4})

      let _nAddress = await nFactory.getNodeAddress.call({from:nodeAddress4});

      let node = await Node.at(_nAddress);
      let count = await nFactory.getNodeCount.call();

      assert.equal(count, 1, "Node made successfully");

    })


  })
})

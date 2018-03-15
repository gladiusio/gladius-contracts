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
  let nodeAddress = accounts[1]
  let clientAddress = accounts[2]

  describe('Test Node Contract', function() {

    it('Create a node', async function() {
      let market = await Market.deployed()
      let nFactory = await NodeFactory.deployed()

      await nFactory.createNode.sendTransaction({from:nodeAddress})

      let _nAddress = await nFactory.getNodeAddress.call({from:nodeAddress});

      let node = await Node.at(_nAddress);
      let count = await nFactory.getNodeCount.call();

      assert.equal(count, 1, "Node made successfully");

    })


  })
})

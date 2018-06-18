let Pool = artifacts.require('Pool')
let Market = artifacts.require('Market')
let GladiusToken = artifacts.require('GladiusToken')
let NodeFactory = artifacts.require('NodeFactory')
let ClientFactory = artifacts.require('ClientFactory')
let Node = artifacts.require('Node')
let Client = artifacts.require('Client')

contract('Pool', async function(accounts) {
  // Accounts
  let owner = accounts[0]
  let nodeAddress1 = accounts[1]
  let nodeAddress2 = accounts[2]
  let nodeAddress3 = accounts[3]
  let clientAddress1 = accounts[5]
  let clientAddress2 = accounts[6]

  let market = await Market.deployed()
  let nFactory = await NodeFactory.deployed()
  let cFactory = await ClientFactory.deployed()


  describe('Test Pool Contract', function() {

    it('Check owner', async function() {
      await market.createPool("TEST_KEY1", {from: owner})

      let plist = await market.getAllPools.call();
      let poolOwner = await market.poolToOwner(plist[0]);

      assert.equal(poolOwner, owner, 'Pool\'s owner is not the creator of the market')
    })

    it('Node and client added to list', async function() {
      await nFactory.createNode({from:nodeAddress1})
      let _node1 = await nFactory.getNodeAddress.call({from:nodeAddress1})
      let node1 = await Node.at(_node1);

      await nFactory.createNode({from:nodeAddress2})
      let _node2 = await nFactory.getNodeAddress.call({from:nodeAddress2})
      let node2 = await Node.at(_node2);

      await cFactory.createClient({from:clientAddress1})
      let _client1 = await cFactory.getClientAddress.call({from:clientAddress1})
      let client1 = await Client.at(_client1);

      await cFactory.createClient({from:clientAddress2})
      let _client2 = await cFactory.getClientAddress.call({from:clientAddress2})
      let client2 = await Client.at(_client2);

      await market.createPool("TEST_KEY2", {from: owner})
      await market.createPool("TEST_KEY3", {from: owner})

      let plist = await market.getAllPools.call()
      let pool = Pool.at(plist[0])

      await node1.applyToPool.sendTransaction(plist[0], "celo-node1", {from: nodeAddress1})
      await node2.applyToPool.sendTransaction(plist[0], "celo-node2", {from: nodeAddress2})
      await client1.applyToPool.sendTransaction(plist[0], "celo-client1", {from: clientAddress1})
      await client2.applyToPool.sendTransaction(plist[0], "celo-client2", {from: clientAddress2})

      let nodeList = await pool.getNodeList.call()
      let clientList = await pool.getClientList.call()

      assert.equal(nodeList[0], node1.address, 'Node applicant is added to list')
      assert.equal(clientList[1], client2.address, 'Client applicant is added to list')
    })

    it('Get node and client lists', async function() {
      let plist = await market.getAllPools.call()
      let pool = Pool.at(plist[0])

      // 1 client and 1 node applied for pool in the previous tests

      let nodeList   = await pool.getNodeList.call()
      let clientList = await pool.getClientList.call()

      assert.equal(nodeList.length, 2, 'Nodes being added to list')
      assert.equal(clientList.length, 2, 'Clients being added to list')
    })

    it('Accept node and clients', async function() {
      let plist = await market.getAllPools.call()
      let pool = Pool.at(plist[0])
      let nodeList   = await pool.getNodeList.call()
      let clientList = await pool.getClientList.call()

      await pool.acceptNode.sendTransaction(nodeList[0], {from:owner})
      await pool.acceptClient.sendTransaction(clientList[0], {from:owner})

      let _node = await nFactory.getNodeAddress.call({from:nodeAddress1})
      let node = await Node.at(_node)

      let _client = await cFactory.getClientAddress.call({from:clientAddress1})
      let client = await Client.at(_client);

      let nodeStatus = await node.getStatus.call(plist[0], {from:nodeAddress1})
      let clientStatus = await client.getStatus.call(plist[0], {from:clientAddress1})

      assert.equal(nodeStatus.toNumber(), 1, 'Nodes being added to count')
      assert.equal(clientStatus.toNumber(), 1, 'Clients being added to count')
    })

    it('Reject node and clients', async function() {
      let plist = await market.getAllPools.call()
      let pool = Pool.at(plist[0])
      let nodeList   = await pool.getNodeList.call()
      let clientList = await pool.getClientList.call()

      await pool.rejectNode.sendTransaction(nodeList[1], {from:owner})
      await pool.rejectClient.sendTransaction(clientList[1], {from:owner})

      let _node = await nFactory.getNodeAddress.call({from:nodeAddress2})
      let node = await Node.at(_node)

      let _client = await cFactory.getClientAddress.call({from:clientAddress2})
      let client = await Client.at(_client);

      let nodeStatus = await node.getStatus.call(plist[0], {from:nodeAddress2})
      let clientStatus = await client.getStatus.call(plist[0], {from:clientAddress2})

      assert.equal(nodeStatus.toNumber(), 2, 'Nodes being added to count')
      assert.equal(clientStatus.toNumber(), 2, 'Clients being added to count')
    })

    it('Records all Node owners', async function() {
      let plist = await market.getAllPools.call()
      let pool = Pool.at(plist[0])
      let nodeList = await pool.getNodeList.call()
      let nodeOwnersList = await pool.getNodeOwnerList.call()
      assert.equal(nodeList.length, nodeOwnersList.length, 'Nodes being added to owners list')
      assert.equal(nodeOwnersList[0], accounts[1], 'Nodes being added to owners list')
      assert.equal(nodeOwnersList[1], accounts[2], 'Nodes being added to owners list')
    })

    it('Records all Client owners', async function() {
      let plist = await market.getAllPools.call()
      let pool = Pool.at(plist[0])
      let clientList = await pool.getClientList.call()
      let clientOwnersList = await pool.getClientOwnerList.call()
      assert.equal(clientList.length, clientOwnersList.length, 'Clients being added to owners list')
      assert.equal(clientOwnersList[0], accounts[5], 'Clients being added to owners list')
      assert.equal(clientOwnersList[1], accounts[6], 'Clients being added to owners list')
    })
  })
})

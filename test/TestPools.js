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
  let nodeAddress = accounts[1]
  let clientAddress = accounts[2]

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
      await nFactory.createNode.sendTransaction("node_data", {from:nodeAddress})
      let _node = await nFactory.getNodeAddress.call({from:nodeAddress})
      let node = await Node.at(_node);

      await cFactory.createClient.sendTransaction("client_data", {from:clientAddress})
      let _client = await cFactory.getClientAddress.call({from:clientAddress})
      let client = await Client.at(_client);

      await market.createPool("TEST_KEY2", {from: owner})
      await market.createPool("TEST_KEY3", {from: owner})

      let plist = await market.getAllPools.call()
      let pool = Pool.at(plist[0])

      await node.applyToPool.sendTransaction(plist[0], "celo-node1", {from: nodeAddress})
      await client.applyToPool.sendTransaction(plist[0], "celo-client1", {from: clientAddress})

      let data = await node.getData.call(plist[0])

      let nodeList = await pool.getNodeList.call()
      let clientList = await pool.getClientList.call()

      assert.equal(nodeList[0], node.address, 'Node applicant is added to list')
      assert.equal(clientList[0], client.address, 'Client applicant is added to list')
    })

    it('Get node and client information', async function() {
      let plist = await market.getAllPools.call()
      let pool = Pool.at(plist[0])

      await nFactory.createNode.sendTransaction("node_data", {from:nodeAddress})
      let _node = await nFactory.getNodeAddress.call({from:nodeAddress})
      let node = await Node.at(_node);

      await cFactory.createClient.sendTransaction("client_data", {from:clientAddress})
      let _client = await cFactory.getClientAddress.call({from:clientAddress})
      let client = await Client.at(_client);

      await node.applyToPool.sendTransaction(plist[0], "celo-node1", {from: nodeAddress})
      await client.applyToPool.sendTransaction(plist[0], "celo-client1", {from: clientAddress})

      let nodeData = await node.getData.call(plist[0])
      let clientData = await client.getData.call(plist[0])

      assert.equal(nodeData, "celo-node1", 'Able to get node data')
      assert.equal(clientData, "celo-client1", 'Able to get client data')
    })

    // it('Get node and client lists', async function() {
    //   let plist = await market.getAllPools.call()
    //   let pool = Pool.at(plist[0])
    //
    //   await nFactory.createNode.sendTransaction("node_data", {from:nodeAddress})
    //   let _node = await nFactory.getNodeAddress.call({from:nodeAddress})
    //   let node = await Node.at(_node);
    //
    //   await cFactory.createClient.sendTransaction("client_data", {from:clientAddress})
    //   let _client = await cFactory.getClientAddress.call({from:clientAddress})
    //   let client = await Client.at(_client);
    //
    //   await node.applyToPool.sendTransaction(plist[0], "celo-node2", {from: nodeAddress})
    //   await node.applyToPool.sendTransaction(plist[0], "celo-node3", {from: nodeAddress})
    //   await client.applyToPool.sendTransaction(plist[0], "celo-client2", {from: clientAddress})
    //   await client.applyToPool.sendTransaction(plist[0], "celo-client3", {from: clientAddress})
    //
    //   let nodeList   = await pool.getNodeList.call()
    //   let clientList = await pool.getClientList.call()
    //
    //   assert.equal(nodeList.length, 3, 'Nodes being added to count')
    //   assert.equal(clientList.length, 3, 'Clients being added to count')
    // })

    // it('Accept node and clients', async function() {
    //   let market = await Market.deployed()
    //   let plist = await market.getAllPools.call()
    //   let pool = Pool.at(plist[0])
    //   let node = await Node.deployed()
    //   let client = await Client.deployed()
    //
    //   let clientList = await pool.getClientList.call()
    //   let nodeList   = await pool.getNodeList.call()
    //
    //   console.log(clientList);
    //   console.log("========");
    //   console.log(nodeList);
    //
    //
    //   assert.equal(resultClient, "nate-client1", 'Clients being added to count')
    //   assert.equal(resultNode, "nate-node1", 'Nodes being added to count')
    // })

  })
})

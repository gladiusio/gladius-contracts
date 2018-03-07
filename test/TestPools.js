let Pool = artifacts.require('Pool')
let Market = artifacts.require('Market')
let GladiusToken = artifacts.require('GladiusToken')
let Node = artifacts.require('Node')
let Client = artifacts.require('Client')

contract('Pool', function(accounts) {
  // Accounts
  let owner = accounts[0]
  let nodeAddress = accounts[1]
  let clientAddress = accounts[2]

  // Test creation of marketplace
  describe('Test Pool Contract', function() {
    it('Check owner', async function() {
      let market = await Market.deployed()

      await market.createPool("TEST_KEY1", {from: owner})

      let plist = await market.getAllPools.call();
      let poolOwner = await market.poolToOwner(plist[0]);

      assert.equal(poolOwner, owner, 'Pool\'s owner is not the creator of the market')
    })

    // Nodes apply and added to the list
    it('Node and client added to list', async function() {
      let market = await Market.deployed()
      let node = await Node.deployed()
      let client = await Client.deployed()

      await market.createPool("TEST_KEY2", {from: owner})
      await market.createPool("TEST_KEY3", {from: owner})

      let plist = await market.getAllPools.call()
      let pool = Pool.at(plist[0])


      await node.applyToPool.sendTransaction(plist[0], "celo-node1", {from: nodeAddress})
      await client.applyToPool.sendTransaction(plist[0], "celo-client1", {from: clientAddress})

      let nodeList = await pool.getNodeList.call()
      let clientList = await pool.getClientList.call()

      assert.equal(nodeList[0], node.address, 'Node applicant is added to list')
      assert.equal(clientList[0], client.address, 'Client applicant is added to list')
    })
    //
    // // Get node and client information (right now just name)
    // it('Get node and client information', async function() {
    //   let market = await Market.deployed()
    //   let plist = await market.getAllPools.call()
    //   let pool = Pool.at(plist[0])
    //
    //   let clientList = await pool.getClientList.call()
    //   let nodeList   = await pool.getNodeList.call()
    //
    //   let clientData = await pool.getClientData.call(clientList[0])
    //   let nodeData = await pool.getNodeData.call(nodeList[0])
    //
    //   assert.equal(clientData, "celo-client1", 'Able to get client data')
    //   assert.equal(nodeData, "celo-node1", 'Able to get node data')
    // })
    //
    // // Get node and client lists
    // it('Get node and client lists', async function() {
    //   let market = await Market.deployed()
    //   let plist = await market.getAllPools.call()
    //   let pool = Pool.at(plist[0])
    //
    //   await pool.applyClient("public_key", "celo-client2", {from: client})
    //   await pool.applyClient("public_key", "celo-client3", {from: client})
    //   await pool.applyNode("public_key", "celo-node2", {from: node})
    //   await pool.applyNode("public_key", "celo-node3", {from: node})
    //
    //   let clientList = await pool.getClientList.call()
    //   let nodeList   = await pool.getNodeList.call()
    //
    //   assert.equal(clientList.length, 3, 'Clients being added to count')
    //   assert.equal(nodeList.length, 3, 'Nodes being added to count')
    // })
    //
    // WIP
    //
    // it('Update node and client data', async function() {
    //   let market = await Market.deployed()
    //   let plist = await market.getAllPools.call()
    //   let pool = Pool.at(plist[0])
    //
    //   let clientList = await pool.getClientList.call()
    //   let nodeList   = await pool.getNodeList.call()
    //
    //   console.log("==================");
    //   console.log(clientList[0]);
    //   console.log("==================");
    //
    //   await pool.updateClientData.sendTransaction(clientList[0], "nate-client1", {from:client})
    //   await pool.updateNodeData.sendTransaction(nodeList[0], "nate-node1", {from:node})
    //
    //    let resultClient = await pool.getClient.call(clientList[0])
    //    console.log("==================");
    //    console.log(resultClient);
    //    console.log("==================");
    //   // let resultNode = await pool.getNode.call(nodeList[0])
    //
    //   assert.equal(resultClient, "nate-client1", 'Clients being added to count')
    //   assert.equal(resultNode, "nate-node1", 'Nodes being added to count')
    // })

  })
})

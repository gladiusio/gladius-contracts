let NodeFactory = artifacts.require('NodeFactory')
let ClientFactory = artifacts.require('ClientFactory')
let Node = artifacts.require('Node')
let Client = artifacts.require('Client')

module.exports = function(deployer, network, accounts) {
  deployer.deploy(NodeFactory, {from: accounts[1]});
  deployer.deploy(ClientFactory, {from: accounts[2]});
}

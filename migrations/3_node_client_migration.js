let Node = artifacts.require('Node')
let Client = artifacts.require('Client')

module.exports = function(deployer, network, accounts) {
  deployer.deploy(Node, "node_data", {from: accounts[1]});
  deployer.deploy(Client, "client_data", {from: accounts[2]});
}

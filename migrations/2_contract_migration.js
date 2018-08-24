let PoolContract = artifacts.require('Pool')
let MarketContract = artifacts.require('Market')

module.exports = function(deployer, network, accounts) {
  deployer.deploy(MarketContract, accounts[0])
  deployer.deploy(PoolContract, accounts[0])
}

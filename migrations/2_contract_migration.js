let Pool = artifacts.require('Pool')
let PoolFactory = artifacts.require('PoolFactory')
let Market = artifacts.require('Market')

module.exports = function(deployer, network, accounts) {
  deployer.deploy(PoolFactory, accounts[0]);
  deployer.deploy(PoolFactory, accounts[0]).then(function() {
    return deployer.deploy(Market, accounts[0], PoolFactory.address)
  })
}
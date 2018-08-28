let PoolFactory1 = artifacts.require('PoolFactory')
let PoolFactory2 = artifacts.require('PoolFactoryTest')
let Market = artifacts.require('Market')

module.exports = function(deployer, network, accounts) {
  // I only need this to test switching out contracts
  if (network == "travisci") {
    deployer.deploy(PoolFactory2, accounts[0]);
  }

  // here's the "real" deployment
  deployer.deploy(PoolFactory1, accounts[0]).then(function() {
    return deployer.deploy(Market, accounts[0], PoolFactory1.address)
  })
}
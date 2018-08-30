let Balance = artifacts.require('Balance')
let Pool = artifacts.require('Pool')

module.exports = function(deployer, network, accounts) {
  // I only need this to test switching out contracts
  if (network == "truffle") {
    deployer.deploy(Balance, accounts[0]);
  }

  // here's the "real" deployment
  deployer.deploy(PoolFactory1, accounts[0]).then(function() {
    return deployer.deploy(Market, accounts[0], PoolFactory1.address)
  })
}

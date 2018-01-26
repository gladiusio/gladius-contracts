let PoolContract = artifacts.require('Pool')
let MarketContract = artifacts.require('Market')
let GladiusToken = artifacts.require('GladiusToken')

module.exports = function(deployer) {
  deployer.deploy(GladiusToken, 10000, 'Gladius Token', 'GLA').then( ()=> {
    return deployer.deploy(MarketContract, GladiusToken.address, 10, 10)
  });
}

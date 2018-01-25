let PoolContract = artifacts.require('Pool')
let MarketContract = artifacts.require('Market')
let Token = artifacts.require('GladiusToken')

module.exports = async function(deployer) {
  await deployer.deploy(Token, 10000, 'Gladius Token', 'GLA')
  await deployer.link(Token, MarketContract)
  await deployer.deploy(PoolContract)
  await deployer.deploy(MarketContract, Token.address, 10, 5)
}

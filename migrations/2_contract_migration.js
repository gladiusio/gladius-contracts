var PoolContract = artifacts.require("./Pool.sol");
var MarketContract = artifacts.require("./Market.sol");
var Token = artifacts.require("./GladiusToken.sol");


module.exports = function(deployer) {
    deployer.deploy(PoolContract);
    deployer.deploy(MarketContract);
    deployer.deploy(Token, 10000, "Gladius Token", "GLA");
};

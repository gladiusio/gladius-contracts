var PoolContract = artifacts.require("./Pool.sol");
var MarketContract = artifacts.require("./Market.sol");
var Token = artifacts.require("./GladiusToken.sol");


module.exports = function(deployer) {
    deployer.deploy(Token, 10000, "Gladius Token", "GLA").then(function() {
        deployer.deploy(PoolContract);
        deployer.deploy(MarketContract,Token.address);
    });
};

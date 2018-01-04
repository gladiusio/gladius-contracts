var Contract = artifacts.require("./market.sol");

module.exports = function(deployer) {
  deployer.deploy(Migrations);
};

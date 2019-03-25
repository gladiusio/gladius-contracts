const PoolFactory = artifacts.require("PoolFactory");
const Pool = artifacts.require("Pool");
const Market = artifacts.require("Market");

module.exports = function(deployer, network, accounts) {
	deployer.deploy(PoolFactory).then(function() {
		deployer.link(PoolFactory, Market)
 		return deployer.deploy(Market, accounts[0], PoolFactory.address);
 	});
};


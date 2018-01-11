// Specifically request an abstraction for MetaCoin
let PoolContract = artifacts.require("Pool");
let MarketContract = artifacts.require("Market");
let GladiusToken = artifacts.require("GladiusToken");



contract('GladiusToken', function(accounts) {
    it("Creator has all the tokens", function() {
        return GladiusToken.deployed().then(function(instance) {
            return instance.balanceOf.call(accounts[0]);
        }).then(function(balance) {
            assert.equal(balance.valueOf(), 10000 * (10**8), "Initial tokens were not allocated to the creator's account");
        });
    });

    it("Tokens send correctly", function() {
        var meta;

        // Get initial balances of first and second account.
        var account_one = accounts[0];
        var account_two = accounts[1];

        var account_one_starting_balance;
        var account_two_starting_balance;
        var account_one_ending_balance;
        var account_two_ending_balance;

        var amount = 10;

        return GladiusToken.deployed().then(function(instance) {
          meta = instance;
          return meta.balanceOf.call(account_one);
        }).then(function(balance) {
          account_one_starting_balance = balance.toNumber();
          return meta.balanceOf.call(account_two);
        }).then(function(balance) {
          account_two_starting_balance = balance.toNumber();
          return meta.transfer(account_two, amount);
        }).then(function() {
          return meta.balanceOf.call(account_one);
        }).then(function(balance) {
          account_one_ending_balance = balance.toNumber();
          return meta.balanceOf.call(account_two);
        }).then(function(balance) {
          account_two_ending_balance = balance.toNumber();

          assert.equal(account_one_ending_balance, account_one_starting_balance - amount, "Amount wasn't correctly taken from the sender");
          assert.equal(account_two_ending_balance, account_two_starting_balance + amount, "Amount wasn't correctly sent to the receiver");
        });
    });

    it("Able to create new tokens", function() {
        var meta;

        // Get initial balances of first and second account.
        var account_one = accounts[0];

        var account_one_starting_balance;
        var account_one_ending_balance;

        var amount = 10;

        return GladiusToken.deployed().then(function(instance) {
          meta = instance;
          return meta.balanceOf.call(account_one);
        }).then(function(balance) {
          account_one_starting_balance = balance.toNumber();
          return meta.createToken(amount);
        }).then(function() {
          return meta.balanceOf.call(account_one);
        }).then(function(balance) {
          account_one_ending_balance = balance.toNumber();
          assert.equal(account_one_ending_balance, account_one_starting_balance + amount, "Amount created was not given to caller");
        });
    });

    it("Able to get balance through the smart contract", function() {
        return GladiusToken.deployed().then(function(instance) {
            return instance.balanceOf.call(accounts[0]);
        }).then(function(bal) {
            assert.equal(bal,10000*10**8, "SHIT");
        })
    })

    it("Add pool to Marketplace", function() {
        // TODO
    })
});

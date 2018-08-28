let HDWalletProvider = require('truffle-hdwallet-provider')
let auth = require('./auth.json')

let infura_apikey = auth.infuraKey
let mnemonic = auth.mnemonic

module.exports = {
  solc: {
    optimizer: {
      enabled: true,
      runs: 200
    }
  },
  networks: {
      mainnet: {
        provider: function() {return new HDWalletProvider(mnemonic.mainnet, "https://mainnet.infura.io/"+infura_apikey)},
        gasPrice:10000000000,
        network_id: 1
      },
      truffle: {
        host: "localhost",
        port: 9545,
        network_id: "*",
      },
      ganache: {
        host: "localhost",
        port: 8545,
        network_id: "*"
      },
      ropsten: {
        provider: function() {return new HDWalletProvider(mnemonic.ropsten, "https://ropsten.infura.io/"+infura_apikey)},
        network_id: 3,
        gas: 4700000
      }
  }
};

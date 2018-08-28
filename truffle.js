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
      development: {
          host: "localhost",
          port: 7545,
          network_id: "*" // Match any network id
      },
      truffle: {
        host: "localhost",
        port: 9545,
        network_id: "*", // Match any network id
      },
      travisci: {
          host: "localhost",
          port: 8545,
          network_id: "*"
      },
      ropsten: {
        provider: new HDWalletProvider(mnemonic, "https://ropsten.infura.io/"+infura_apikey),
        network_id: 3,
        gas: 4700000
      }
  }
};

let HDWalletProvider = require('truffle-hdwallet-provider')
let auth = require('./auth.json')

let infura_apikey = auth.infuraKey
let mnemonic = auth.mnemonic

module.exports = {
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
        },
        rinkeby: {
            host: "localhost", // Connect to geth on the specified
            port: 8545,
            from: "0x0085f8e72391Ce4BB5ce47541C846d059399fA6c", // default address to use for any transaction Truffle makes during migrations
            network_id: 4,
            gas: 4612388 // Gas limit used for deploys
        }
    }
};

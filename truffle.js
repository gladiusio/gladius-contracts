let HDWalletProvider = require('truffle-hdwallet-provider')
let auth, infura_apikey, mnemonic
try {
    auth = require('./auth.json')
    infura_apikey = auth.infuraKey
    mnemonic = auth.mnemonic
} catch(e){
    console.warn("Could not load auth.json file, continuing without those credentials")
}


module.exports = {
    networks: {
        development: {
            host: "localhost",
            port: 7545,
            network_id: "*" // Match any network id
        },
        docker: {
            host: "ganache",
            port: 6545,
            network_id: "*"
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
          provider: () => auth ? new HDWalletProvider(mnemonic, "https://ropsten.infura.io/"+infura_apikey): null,
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

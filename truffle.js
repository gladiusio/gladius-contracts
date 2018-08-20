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
        }
    }
};

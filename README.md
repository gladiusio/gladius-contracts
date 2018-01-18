# Gladius Marketplace Backend

# Market.sol
Main marketplace contract:
- Acts as a factory to create new pools
- Adds pools to the marketplace list
- Holds payments (security deposit) from pool owners?

# Pool.sol
Main pool contract:
- Pays nodes in the pool
- Sets rules of the pool
- Keeps track of data from the nodes
- Sends data to nodes

# GladiusToken.sol
Fake ERC-20 Token

## How to deploy to a local blockchain
- Run local blockchain (we like Ganache)
- Make sure to have the right port in `truffle.js`
- Compile: `truffle compile`
- Migrate(deploy): `truffle migrate --reset`

# Gladius Marketplace Backend

[![Build Status](https://travis-ci.com/gladiusio/gladius-contracts.svg?branch=master)](https://travis-ci.com/gladiusio/gladius-contracts)

## Installation & Testing

* `npm install`
* `npm run test`


## Platform Requirements
* Node.js, `>=7.6.0`

## How to deploy to a local blockchain
- `npm run truffleDevelop` (Terminal 1)
- `npm run truffleDeploy` (Terminal 2)

## How to deploy to Main Ethereum Network
- `npm run truffleDeployMain`


## Code Overview

### Market.sol (Market Contract)
The `Marketplace` contract will hold all information regarding the Gladius Marketplace. This includes which pools are added to the marketplace and who the owners are. In addition the `Marketplace` has the ability to create and track all pools via the Pool Factory contract.

### Pool.sol (PoolFactory and Pool Contracts)
The `PoolFactory` contract creates and tracks `Pool` contracts. `Pool` contracts contain information regarding the P2P network, application server, etc... . We use a factory pattern to create `Pool`'s because it allows us to be more flexible with deploying contacts.


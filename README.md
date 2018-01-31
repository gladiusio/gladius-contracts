# Gladius Marketplace Backend

[![Build Status](https://travis-ci.org/gladiusio/gladius-marketplace-backend.svg?branch=master)](https://travis-ci.org/gladiusio/gladius-marketplace-backend)

## Installation & Testing

* Run local blockchain (we like Ganache)
* Make sure to have the right port in `truffle.js`
* Compile: `truffle compile`
* Migrate (deploy): `truffle migrate --reset`
* Test: `truffle test`

## Contribution Guide

### Platform Requirements

* Node.js, `>=7.6.0`
* Truffle
* Ganache, (Optional)
* Default port in `truffle.js` is `7545` for development network

### General

* Issue Requests
    * Please search for existing issues before creating
    * Search Solidity, Truffle, and Ganache issues as well
    * Provide system info
        * System, Ubuntu, macOS, Windows 10, etc.
        * Node version, `node -v`
    * Console output
    * Behavior
        * Expected
        * Actual
    * Steps to reproduce
* Pull Requests
    * Pull Requests should be tied to a specific and single issue
    * PR's will have to pass the automated test suite before manual review
    * Keep comments and opinions constructive for all parties
* Code Style
    * We will be updating our code to fully conform to the [Solidity Code Style Guide](http://solidity.readthedocs.io/en/develop/style-guide.html)
    * Any Pull Requests should also conform to this guide
    * Add doc blocks where necessary to comment functionality
    ```
    /**
     * Title
     *
     * Description (Optional)
     * @param paramName Description of parameter
     * @return
     */
    ```

## Code Overview

More in-depth descriptions coming soon

### Market.sol

Main marketplace contract

### Pool.sol

Main pool contract

### GladiusToken.sol

Fake ERC-20 Token

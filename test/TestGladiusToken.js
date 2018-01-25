// Specifically request an abstraction for MetaCoin
let Pool = artifacts.require('Pool')
let Market = artifacts.require('Market')
let GladiusToken = artifacts.require('GladiusToken')

contract('GladiusToken', function(accounts) {
  // Accounts
  let creator = accounts[0]
  let user = accounts[1]

  let initialSupply = 10000 * (10**8)

  describe('Gladius Token Contract', function() {
    // Test creation of tokens
    it('Deploy Tokens to Creator\'s Account', async function() {
      let gladius = await GladiusToken.deployed()
      let balance = await gladius.balanceOf.call(accounts[0])
      assert.equal(balance.valueOf(), initialSupply, 'Initial tokens were not allocated to the creator\'s account')
    })

    // Test balance transfers
    it('Simulate transfer from creator to user', async function() {
      let transferAmount = 10

      // Grab instance
      let gladius = await GladiusToken.deployed()
      let creatorStartBalance = await gladius.balanceOf.call(creator)
      let userStartBalance = await gladius.balanceOf.call(user)

      await gladius.transfer(user, transferAmount)

      let creatorEndBalance = await gladius.balanceOf.call(creator)
      let userEndBalance = await gladius.balanceOf.call(user)

      assert.equal(
        creatorEndBalance.toNumber(),
        creatorStartBalance.toNumber() - transferAmount,
        'Amount wasn\'t correctly taken from the sender'
      )
      assert.equal(
        userEndBalance.toNumber(),
        userStartBalance.toNumber() + transferAmount,
        'Amount wasn\'t correctly sent to the receiver'
      )
    })

    it('Simulate burning tokens', async function() {
      let burnAmount = 50.0

      let gladius = await GladiusToken.deployed()
      let totalInitialSupply = await gladius.totalSupply.call()

      await gladius.burn(burnAmount)

      let totalEndingSupply = await gladius.totalSupply.call()

      assert.equal(
        totalEndingSupply.toNumber(),
        totalInitialSupply.toNumber() - burnAmount,
        'Tokens did not successfully burn'
      )
    })

    it('Approval of spending from other accounts', async function() {
      let approvalAmount = 25.0

      let gladius = await GladiusToken.deployed()

      let userStartBalance = await gladius.balanceOf.call(user)

      await gladius.approve(user, approvalAmount, { from: creator })

      await gladius.transferFrom(creator, user, approvalAmount, { from: user })

      let creatorBalance = await gladius.balanceOf.call(creator)
      let userBalance = await gladius.balanceOf.call(user)

      assert.equal(
        userBalance.toNumber(),
        userStartBalance.toNumber() + approvalAmount,
        'Tokens were not approved for transfer'
      )
    })
  })
})

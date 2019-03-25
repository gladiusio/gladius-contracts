package commands

import (
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gladiusio/gladius-contracts/go/cli/backend"
	ishell "gopkg.in/abiosoft/ishell.v2"
)

func accountPicker(bm *backend.Manager, c *ishell.Context) accounts.Account {
	accounts := bm.GetWallet().Accounts()
	accountStrings := make([]string, len(accounts))

	for i, a := range accounts {
		accountStrings[i] = a.Address.String()
	}

	accountChoice := c.MultiChoice(accountStrings, "Pick an account to use")

	return accounts[accountChoice]
}

func marketPicker(bm *backend.Manager, c *ishell.Context) common.Address {
	if addr, set := bm.GetMarketPlaceAddress(); set {
		return addr
	}

	c.Print("Enter the marketplace address you want to connect to: ")
	marketAddress := common.HexToAddress(c.ReadLine())
	bm.SetMarketPlaceAddress(marketAddress)
	c.Printf("Caching %s as default market (you can change this with `market connect`)\n", marketAddress.String())

	return marketAddress

}

package commands

import (
	"errors"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gladiusio/gladius-contracts/go/cli/backend"
	"github.com/gladiusio/gladius-contracts/go/contracts"
	ishell "gopkg.in/abiosoft/ishell.v2"
)

// MarketCreate walks a user through creating a market
func MarketCreate(bm *backend.Manager) func(c *ishell.Context) {
	return func(c *ishell.Context) {
		if bm.GetClient() == nil {
			c.Err(errors.New("no backend configured, run `connect`"))
			return
		}

		c.Print("This operation will cost Ether, are you sure you want to deploy the market? (y/n): ")
		answer := c.ReadLine()
		if !(strings.EqualFold(answer, "y") || strings.EqualFold(answer, "yes")) {
			return
		}

		account := accountPicker(bm, c)
		c.Println("Picked: " + account.Address.String())

		key, err := bm.GetWallet().PrivateKey(account)
		if err != nil {
			c.Err(fmt.Errorf("could not get private key: %v", err))
			return
		}
		auth := bind.NewKeyedTransactor(key)

		pfAddress, _, _, err := contracts.DeployPoolFactory(auth, bm.GetClient())
		if err != nil {
			c.Err(fmt.Errorf("could not deploy pool factory contract: %v", err))
			return
		}

		marketAddress, _, _, err := contracts.DeployMarket(auth, bm.GetClient(), account.Address, pfAddress)
		if err != nil {
			c.Err(fmt.Errorf("could not deploy market contract: %v", err))
			return
		}

		c.Printf("Deployed contracts:\n - Marketplace: %s\n - Pool Factory: %s\n", marketAddress.String(), pfAddress.String())
		bm.SetMarketPlaceAddress(marketAddress)
	}
}

// MarketAddPool walks a user through adding a pool to the marketplace
func MarketAddPool(bm *backend.Manager) func(c *ishell.Context) {
	return func(c *ishell.Context) {
		if bm.GetClient() == nil {
			c.Err(errors.New("no backend configured, run `connect`"))
			return
		}

		c.Print("This operation will cost Ether, are you sure you want to add a pool to the market? (y/n): ")
		answer := c.ReadLine()
		if !(strings.EqualFold(answer, "y") || strings.EqualFold(answer, "yes")) {
			return
		}

		account := accountPicker(bm, c)
		c.Println("Picked: " + account.Address.String())

		key, err := bm.GetWallet().PrivateKey(account)
		if err != nil {
			c.Err(fmt.Errorf("could not get private key: %v", err))
			return
		}
		auth := bind.NewKeyedTransactor(key)

		marketAddress := marketPicker(bm, c)

		market, err := contracts.NewMarket(marketAddress, bm.GetClient())
		if err != nil {
			c.Err(fmt.Errorf("could not bind to marketplace: %v", err))
			return
		}

		owner, err := market.GetOwner(&bind.CallOpts{})
		if err != nil {
			c.Err(fmt.Errorf("could not get owner of contract: %v", err))
			return
		}

		if owner != account.Address {
			c.Err(fmt.Errorf("you do not own that contract: %v", err))
			return
		}

		c.Print("Enter a pool address to add: ")
		poolAddress := common.HexToAddress(c.ReadLine())

		_, err = contracts.NewPool(poolAddress, bm.GetClient())
		if err != nil {
			c.Println("Error: That does not look like a pool contract")
			return
		}

		market.AddToMarketplace(&bind.TransactOpts{From: account.Address, Signer: auth.Signer, GasLimit: 4000000}, poolAddress)
	}
}

// MarketConnect walks a user through setting up a new marketplace default
func MarketConnect(bm *backend.Manager) func(c *ishell.Context) {
	return func(c *ishell.Context) {
		if bm.GetClient() == nil {
			c.Err(errors.New("no backend configured, run `connect`"))
			return
		}

		c.Print("Enter market address: ")
		marketAddress := common.HexToAddress(c.ReadLine())

		_, err := contracts.NewMarket(marketAddress, bm.GetClient())
		if err != nil {
			c.Println("Error: that does not look like a marketplace")
			return
		}

		bm.SetMarketPlaceAddress(marketAddress)
	}
}

// MarketList lists all added pools in the marketplace
func MarketList(bm *backend.Manager) func(c *ishell.Context) {
	return func(c *ishell.Context) {
		if bm.GetClient() == nil {
			c.Err(errors.New("no backend configured, run `connect`"))
			return
		}

		marketAddress := marketPicker(bm, c)

		market, err := contracts.NewMarket(marketAddress, bm.GetClient())
		if err != nil {
			c.Println("Error: that does not look like a marketplace")
			return
		}

		pools, err := market.GetMarketPools(&bind.CallOpts{})
		if err != nil {
			c.Err(fmt.Errorf("could not get owned pools from marketplace: %v", err))
			return
		}
		c.Println("Marketplace pools:")
		for _, p := range pools {
			c.Println(" - " + p.String())
		}
	}
}

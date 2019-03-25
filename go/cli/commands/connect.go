package commands

import (
	"strings"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gladiusio/gladius-contracts/go/cli/backend"

	ishell "gopkg.in/abiosoft/ishell.v2"

	"context"
	"fmt"
)

// Connect returns an ishell command that updates the given backend
func Connect(bm *backend.Manager) func(c *ishell.Context) {
	return func(c *ishell.Context) {
		choice := c.MultiChoice([]string{
			"Infura",
			"Ganache",
			"Other",
		}, "Pick a node type:")

		switch choice {
		case 0:
			c.Print("Infura Key: ")
			key := c.ReadLine()

			networkChoice := c.MultiChoice([]string{
				"Mainnet",
				"Ropsten",
				"Rinkeby",
				"Kovan",
			}, "Network to connect to:")

			network := ""
			switch networkChoice {
			case 0:
				network = "mainnet"
			case 1:
				network = "ropsten"
			case 2:
				network = "rinkeby"
			case 3:
				network = "kovan"
			}

			url := fmt.Sprintf("https://%s.infura.io/v3/%s", network, key)

			c.Printf("Using %s to connect...\n", url)
			var err error
			newBackend, err := ethclient.Dial(url)
			if err != nil {
				c.Err(err)
				return
			}

			// Check if the API key provided works
			_, err = newBackend.NetworkID(context.Background())
			if err != nil {
				c.Err(err)
				return
			}

			// Pass the backend out
			bm.SetClient(newBackend)

			c.Println("Connected!")
		case 1:
			networkChoice := c.MultiChoice([]string{
				"7545",
				"8545",
				"9545",
				"Custom",
			}, "Port to use:")

			port := "7545"
			switch networkChoice {
			case 1:
				port = "8545"
			case 2:
				port = "9545"
			case 3:
				c.Print("Port: ")
				port = c.ReadLine()
			}

			url := fmt.Sprintf("http://localhost:%s", port)

			c.Printf("Connecting to Ganache on %s...\n", url)
			var err error
			newBackend, err := ethclient.Dial(url)
			if err != nil {
				c.Err(err)
				return
			}

			// Check if the connection provided works
			_, err = newBackend.NetworkID(context.Background())
			if err != nil {
				c.Err(err)
				return
			}

			// Pass the backend out
			bm.SetClient(newBackend)
			c.Println("Connected!")
		case 2:
			c.Print("Connection URL: ")
			url := c.ReadLine()
			c.Printf("Connecting to service on %s...\n", url)
			var err error
			newBackend, err := ethclient.Dial(url)
			if err != nil {
				c.Err(err)
				return
			}

			// Check if the connection provided works
			_, err = newBackend.NetworkID(context.Background())
			if err != nil {
				c.Err(err)
				return
			}

			// Pass the backend out
			bm.SetClient(newBackend)
			c.Println("Connected!")
		}

		c.Print("Enter (or paste) the mnemonic for the accounts you want to use: ")
		mnemonic := c.ReadPassword()
		err := bm.SetupAccounts(mnemonic)

		if err != nil {
			c.Err(err)
			return
		}

		accountStrings := make([]string, len(bm.GetWallet().Accounts()))

		for i, a := range bm.GetWallet().Accounts() {
			accountStrings[i] = a.Address.String()
		}

		c.Printf("Loaded accounts: [%s]\n", strings.Join(accountStrings, ","))
	}
}

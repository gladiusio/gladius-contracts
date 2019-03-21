package main

import (
	"github.com/ethereum/go-ethereum/ethclient"
	ishell "gopkg.in/abiosoft/ishell.v2"

	"context"
	"fmt"
)

// Connect returns an ishell command that updates the given backend
func Connect(c *ishell.Context) {
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
		backend, err = ethclient.Dial(url)
		if err != nil {
			c.Err(err)
			return
		}

		// Check if the API key provided works
		_, err = backend.NetworkID(context.Background())
		if err != nil {
			c.Err(err)
			return
		}

		c.Println("Connected!")
	case 1:
	}

}

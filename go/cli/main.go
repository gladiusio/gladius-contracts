package main

import (
	"github.com/ethereum/go-ethereum/ethclient"
	ishell "gopkg.in/abiosoft/ishell.v2"
)

// Global backend
var backend *ethclient.Client

func main() {

	shell := ishell.New()

	shell.Println("Gladius Contract Manager")
	shell.Println("Start with `connect`, type `help` for more information")

	shell.AddCmd(&ishell.Cmd{
		Name: "connect",
		Help: "Connect to Ethereum RPC",
		Func: Connect,
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "pool",
		Help: "Connect to Ethereum RPC",
		Func: func(c *ishell.Context) {

		},
	})

	// run shell
	shell.Run()
}

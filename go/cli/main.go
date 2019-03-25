package main

import (
	"fmt"

	"github.com/gladiusio/gladius-contracts/go/cli/backend"
	"github.com/gladiusio/gladius-contracts/go/cli/commands"
	ishell "gopkg.in/abiosoft/ishell.v2"
)

func main() {
	shell := ishell.New()
	shell.Println("Gladius Contract Manager")
	shell.Println("Start with `connect`, type `help` for more information")

	bm := &backend.Manager{}

	connectCommand := &ishell.Cmd{
		Name: "connect",
		Help: "Connect to Ethereum RPC",
		Func: commands.Connect(bm),
	}

	poolCommand := &ishell.Cmd{
		Name: "pool",
		Help: "Allows creation and interaction of a pool contract",
		Func: func(c *ishell.Context) { fmt.Println(c.Cmd.HelpText()) },
	}

	poolCreateCommand := &ishell.Cmd{
		Name: "create",
		Help: "Creates a pool",
		Func: commands.PoolCreate(bm),
	}

	poolListCommand := &ishell.Cmd{
		Name: "list",
		Help: "Lists all pools you own",
		Func: commands.PoolList(bm),
	}

	poolInfoCommand := &ishell.Cmd{
		Name: "info",
		Help: "Get's information about a specific pool",
		Func: commands.PoolInformation(bm),
	}

	poolUpdateCommand := &ishell.Cmd{
		Name: "update",
		Help: "Update fields in a pool",
		Func: commands.PoolUpdate(bm),
	}

	poolCommand.AddCmd(poolCreateCommand)
	poolCommand.AddCmd(poolListCommand)
	poolCommand.AddCmd(poolInfoCommand)
	poolCommand.AddCmd(poolUpdateCommand)

	shell.AddCmd(connectCommand)
	shell.AddCmd(poolCommand)

	// run shell
	shell.Run()
}

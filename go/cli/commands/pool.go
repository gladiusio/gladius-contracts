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

// PoolCreate walks a user through creating a pool
func PoolCreate(bm *backend.Manager) func(c *ishell.Context) {
	return func(c *ishell.Context) {
		if bm.GetClient() == nil {
			c.Err(errors.New("no backend configured, run `connect`"))
			return
		}

		account := accountPicker(bm, c)
		c.Println("Picked: " + account.Address.String())

		marketAddress := marketPicker(bm, c)
		market, err := contracts.NewMarket(marketAddress, bm.GetClient())
		if err != nil {
			c.Err(fmt.Errorf("could not bind to marketplace: %v", err))
			return
		}

		c.Print("This operation will cost Ether, are you sure you want to deploy the pool? (y/n): ")
		answer := c.ReadLine()
		if strings.EqualFold(answer, "y") || strings.EqualFold(answer, "yes") {
			key, err := bm.GetWallet().PrivateKey(account)
			if err != nil {
				c.Err(fmt.Errorf("could not get private key: %v", err))
				return
			}
			auth := bind.NewKeyedTransactor(key)
			_, err = market.CreatePool(&bind.TransactOpts{From: account.Address, Signer: auth.Signer, GasLimit: 4000000})
			if err != nil {
				c.Err(fmt.Errorf("could not deploy pool: %v", err))
				return
			}

			pools, err := market.GetOwnerAllPools(&bind.CallOpts{From: account.Address, Pending: true}, account.Address)
			if err != nil {
				c.Err(fmt.Errorf("could not get owned pools from marketplace: %v", err))
				return
			}

			poolAddrs := make([]string, len(pools))
			for i, p := range pools {
				poolAddrs[i] = p.String()
			}
			c.Printf("Your owned pools: %v\n", poolAddrs)

		} else {
			c.Println("Not creating pool")
		}
	}
}

// PoolList lists all pools owned by the user
func PoolList(bm *backend.Manager) func(c *ishell.Context) {
	return func(c *ishell.Context) {
		if bm.GetClient() == nil {
			c.Err(errors.New("no backend configured, run `connect`"))
			return
		}

		account := accountPicker(bm, c)
		c.Println("Picked: " + account.Address.String())

		marketAddress := marketPicker(bm, c)
		market, err := contracts.NewMarket(marketAddress, bm.GetClient())
		if err != nil {
			c.Err(fmt.Errorf("could not bind to marketplace: %v", err))
			return
		}

		pools, err := market.GetOwnerAllPools(&bind.CallOpts{From: account.Address}, account.Address)
		if err != nil {
			c.Err(fmt.Errorf("could not get owned pools from marketplace: %v", err))
			return
		}
		c.Println("Your pools:")
		for _, p := range pools {
			c.Println(" - " + p.String())
		}
	}
}

// PoolUpdate asks the user which data they want to update and updates it
func PoolUpdate(bm *backend.Manager) func(c *ishell.Context) {
	return func(c *ishell.Context) {
		if bm.GetClient() == nil {
			c.Err(errors.New("no backend configured, run `connect`"))
			return
		}

		c.Print("This operation will cost Ether, are you sure you want to update pool information? (y/n): ")
		answer := c.ReadLine()
		if !(strings.EqualFold(answer, "y") || strings.EqualFold(answer, "yes")) {
			return
		}

		account := accountPicker(bm, c)
		c.Println("Picked: " + account.Address.String())

		marketAddress := marketPicker(bm, c)
		market, err := contracts.NewMarket(marketAddress, bm.GetClient())
		if err != nil {
			c.Err(fmt.Errorf("could not bind to marketplace: %v", err))
			return
		}

		pools, err := market.GetOwnerAllPools(&bind.CallOpts{From: account.Address}, account.Address)
		if err != nil {
			c.Err(fmt.Errorf("could not get pools from marketplace: %v", err))
			return
		}

		poolAddrs := make([]string, len(pools))
		for i, p := range pools {
			poolAddrs[i] = p.String()
		}

		poolChoice := c.MultiChoice(poolAddrs, "Which pool do you want to update?")

		pool, err := contracts.NewPool(pools[poolChoice], bm.GetClient())
		if err != nil {
			c.Err(fmt.Errorf("could not bind to pool contract: %v", err))
			return
		}

		key, err := bm.GetWallet().PrivateKey(account)
		if err != nil {
			c.Err(fmt.Errorf("could not get private key: %v", err))
			return
		}
		auth := bind.NewKeyedTransactor(key)
		opts := &bind.TransactOpts{From: account.Address, Signer: auth.Signer, GasLimit: 4000000}

		toUpdate := c.MultiChoice([]string{
			"Approve Node",
			"Update Seed Node Address",
			"Update Pool Domain",
			"Update CDN Domain",
			"Update Certificate Bundle Link",
			"Add Masternode",
			"Add Infrastructure Node",
			"Remove Node",
			"Remove Masternode",
			"Remove Infrastructure Node",
		}, "What do you want to do?")

		switch toUpdate {
		case 0:
			c.Print("Enter node: ")
			node := common.HexToAddress(c.ReadLine())
			tx, err := pool.AddApprovedNode(opts, node)
			if err != nil {
				c.Err(err)
				return
			}
			c.Printf("Success! TX Hash: %s\n", tx.Hash().String())
		case 1:
			c.Print("Enter new IP and port (ex. 1.1.1.1:7946): ")
			addr := c.ReadLine()
			tx, err := pool.SetSeedNode(opts, addr)
			if err != nil {
				c.Err(err)
				return
			}
			c.Printf("Success! TX Hash: %s\n", tx.Hash().String())
		case 2:
			c.Print("Enter new link: ")
			addr := c.ReadLine()
			tx, err := pool.SetPoolDomain(opts, addr)
			if err != nil {
				c.Err(err)
				return
			}
			c.Printf("Success! TX Hash: %s\n", tx.Hash().String())
		case 3:
			c.Print("Enter new link: ")
			addr := c.ReadLine()
			tx, err := pool.SetCDNDomain(opts, addr)
			if err != nil {
				c.Err(err)
				return
			}
			c.Printf("Success! TX Hash: %s\n", tx.Hash().String())
		case 4:
			c.Print("Enter new link: ")
			addr := c.ReadLine()
			tx, err := pool.SetCertificateBundle(opts, addr)
			if err != nil {
				c.Err(err)
				return
			}
			c.Printf("Success! TX Hash: %s\n", tx.Hash().String())
		case 5:
			c.Print("Enter Masternode address: ")
			node := common.HexToAddress(c.ReadLine())
			tx, err := pool.AddMasterNode(opts, node)
			if err != nil {
				c.Err(err)
				return
			}
			c.Printf("Success! TX Hash: %s\n", tx.Hash().String())
		case 6:
			c.Print("Enter Infrastructure Node address: ")
			node := common.HexToAddress(c.ReadLine())
			tx, err := pool.AddInfrastructureNode(opts, node)
			if err != nil {
				c.Err(err)
				return
			}
			c.Printf("Success! TX Hash: %s\n", tx.Hash().String())
		case 7:
			c.Print("Node address to remove: ")
			node := common.HexToAddress(c.ReadLine())
			tx, err := pool.RemoveApprovedNode(opts, node)
			if err != nil {
				c.Err(err)
				return
			}
			c.Printf("Success! TX Hash: %s\n", tx.Hash().String())
		case 8:
			c.Print("Masternode address to remove: ")
			node := common.HexToAddress(c.ReadLine())
			tx, err := pool.RemoveMasterNode(opts, node)
			if err != nil {
				c.Err(err)
				return
			}
			c.Printf("Success! TX Hash: %s\n", tx.Hash().String())
		case 9:
			c.Print("Infrastructure Node address to remove: ")
			node := common.HexToAddress(c.ReadLine())
			tx, err := pool.RemoveInfrastructureNode(opts, node)
			if err != nil {
				c.Err(err)
				return
			}
			c.Printf("Success! TX Hash: %s\n", tx.Hash().String())
		default:
			c.Println("Not yet implemented...")
		}
	}
}

// PoolInformation lists information about a specific pool
func PoolInformation(bm *backend.Manager) func(c *ishell.Context) {
	return func(c *ishell.Context) {
		if bm.GetClient() == nil {
			c.Err(errors.New("no backend configured, run `connect`"))
			return
		}

		account := accountPicker(bm, c)
		marketAddress := marketPicker(bm, c)
		market, err := contracts.NewMarket(marketAddress, bm.GetClient())
		if err != nil {
			c.Err(fmt.Errorf("could not bind to marketplace: %v", err))
			return
		}

		pools, err := market.GetAllPools(&bind.CallOpts{From: account.Address})
		if err != nil {
			c.Err(fmt.Errorf("could not get pools from marketplace: %v", err))
			return
		}

		poolAddrs := make([]string, len(pools))
		for i, p := range pools {
			poolAddrs[i] = p.String()
		}

		poolChoice := c.MultiChoice(poolAddrs, "Which pool do you want to see information on?")

		pool, err := contracts.NewPool(pools[poolChoice], bm.GetClient())
		if err != nil {
			c.Err(fmt.Errorf("could not bind to pool contract: %v", err))
			return
		}

		opts := &bind.CallOpts{}
		c.Printf("Pool information for %s:\n", pools[poolChoice].String())
		owner, err := pool.GetOwner(opts)
		if err != nil {
			c.Err(fmt.Errorf("could not get owner: %v", err))
			return
		}
		c.Printf("Owner: %s\n", owner.String())

		cdndomain, err := pool.GetCDNDomain(opts)
		if err != nil {
			c.Err(fmt.Errorf("could not get cdn domain: %v", err))
			return
		}
		c.Printf("CDN Domain: %s\n", cdndomain)

		certLink, err := pool.GetCertificateBundle(opts)
		if err != nil {
			c.Err(fmt.Errorf("could not get certificate link: %v", err))
			return
		}
		c.Printf("Certificate bundle link: %s\n", certLink)

		poolDomain, err := pool.GetPoolDomain(opts)
		if err != nil {
			c.Err(fmt.Errorf("could not get pool domain: %v", err))
			return
		}
		c.Printf("Pool domain: %s\n", poolDomain)

		seedNodeAddress, err := pool.GetSeedNode(opts)
		if err != nil {
			c.Err(fmt.Errorf("could not get seed node address: %v", err))
			return
		}
		c.Printf("Seed Node Address: %s\n", seedNodeAddress)

		nodes, err := pool.GetAllApprovedNodes(opts)
		if err != nil {
			c.Err(fmt.Errorf("could not get approved nodes: %v", err))
			return
		}
		nodeStrings := make([]string, len(nodes))
		for i, n := range nodes {
			nodeStrings[i] = n.String()
		}
		c.Printf("Approved Nodes: [%s]\n", strings.Join(nodeStrings, ","))

		mNodes, err := pool.GetAllMasterNodes(opts)
		if err != nil {
			c.Err(fmt.Errorf("could not get master nodes: %v", err))
			return
		}
		mNodeStrings := make([]string, len(mNodes))
		for i, n := range mNodes {
			mNodeStrings[i] = n.String()
		}
		c.Printf("Master Nodes: [%s]\n", strings.Join(mNodeStrings, ","))
	}
}

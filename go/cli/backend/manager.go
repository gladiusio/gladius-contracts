package backend

import (
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
)

// Manager manages ethereum clients and keys
type Manager struct {
	client     *ethclient.Client
	wallet     *hdwallet.Wallet
	marketAddr common.Address
	marketSet  bool
}

// SetClient sets the ethereum client
func (bm *Manager) SetClient(c *ethclient.Client) {
	bm.client = c
}

// SetMarketPlaceAddress sets the current marketplace address
func (bm *Manager) SetMarketPlaceAddress(addr common.Address) {
	bm.marketAddr = addr
	bm.marketSet = true
}

// GetMarketPlaceAddress gets the current marketplace address
func (bm *Manager) GetMarketPlaceAddress() (common.Address, bool) {
	return bm.marketAddr, bm.marketSet
}

// GetClient gets the current ethereum client
func (bm *Manager) GetClient() *ethclient.Client {
	return bm.client
}

// SetupAccounts takes a mnemonic and builds accounts from it
func (bm *Manager) SetupAccounts(mnemonic string) error {
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		return err
	}

	for i := 0; i < 10; i++ {
		path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/" + strconv.Itoa(i))
		_, err := wallet.Derive(path, true)
		if err != nil {
			return err
		}
	}

	bm.wallet = wallet

	return nil
}

// GetWallet returns the wallet that's stored
func (bm *Manager) GetWallet() *hdwallet.Wallet {
	return bm.wallet
}

package wallet

import (
	"github.com/wealdtech/go-ed25519hd"
	e2wallet "github.com/wealdtech/go-eth2-wallet"
	keystorev4 "github.com/wealdtech/go-eth2-wallet-encryptor-keystorev4"
	types "github.com/wealdtech/go-eth2-wallet-types/v2"
)

func CreateHDWalletFromMnemonic(walletName, password, mnemonic string, store types.Store) types.Wallet {
	seed, err := ed25519hd.SeedFromMnemonic(mnemonic, password)
	if err != nil {
		panic(err)
	}
	opts := []e2wallet.Option{
		e2wallet.WithType("hd"),
		e2wallet.WithPassphrase([]byte(password)),
		e2wallet.WithStore(store),
		e2wallet.WithEncryptor(keystorev4.New()),
		e2wallet.WithSeed(seed),
	}
	wallet, err := e2wallet.CreateWallet(walletName, opts...)
	if err != nil {
		panic(err)
	}
	return wallet
}

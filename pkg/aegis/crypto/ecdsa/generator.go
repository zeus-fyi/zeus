package ecdsa

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/tyler-smith/go-bip32"
	"github.com/wealdtech/go-ed25519hd"
	aegis_random "github.com/zeus-fyi/zeus/pkg/aegis/crypto/random"
)

func GenerateAddresses(mnemonic string, count int) error {
	mnemonic, err := aegis_random.GenerateMnemonic()
	if err != nil {
		return err
	}
	seed, err := ed25519hd.SeedFromMnemonic(mnemonic, "password")
	if err != nil {
		return err
	}
	masterKey, err := bip32.NewMasterKey(seed)

	// Use BIP44: m / purpose' / coin_type' / account' / change / address_index
	// Ethereum path: m/44'/60'/0'/0/0

	for i := 0; i <= count; i++ {
		child, cerr := masterKey.NewChildKey(uint32(i))
		if cerr != nil {
			return err
		}
		privateKeyECDSA := crypto.ToECDSAUnsafe(child.Key)
		address := crypto.PubkeyToAddress(privateKeyECDSA.PublicKey)
		leadingZeroesCount := countLeadingZeroes(address.Hex())
		fmt.Println("Leading zeroes: ", leadingZeroesCount)
		fmt.Println("Ethereum Address: ", address.Hex())
		fmt.Println("Private Key: ", hexutil.Encode(crypto.FromECDSA(privateKeyECDSA)))
	}
	return nil
}

func countLeadingZeroes(address string) int {
	address = strings.TrimPrefix(address, "0x")
	leadingZeros := 0

	for _, char := range address {
		if char == '0' {
			leadingZeros++
		} else {
			break
		}
	}
	return leadingZeros
}

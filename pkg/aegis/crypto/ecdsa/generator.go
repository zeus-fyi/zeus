package ecdsa

import (
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/tyler-smith/go-bip32"
	"github.com/wealdtech/go-ed25519hd"
	aegis_random "github.com/zeus-fyi/zeus/pkg/aegis/crypto/random"
)

type AddressGenerator struct {
	Mnemonic string
	AddressDetails
}

type AddressDetails struct {
	PathIndex          int
	Address            string
	LeadingZeroesCount int
}

func GenerateAddresses(mnemonic string, count, numWorkers int) (AddressGenerator, error) {
	var wg sync.WaitGroup
	wg.Add(numWorkers)
	resultCh := make(chan AddressDetails)
	ag := AddressGenerator{
		Mnemonic: mnemonic,
		AddressDetails: AddressDetails{
			PathIndex:          0,
			Address:            "",
			LeadingZeroesCount: 0,
		},
	}
	mnemonic, err := aegis_random.GenerateMnemonic()
	if err != nil {
		return ag, err
	}
	seed, err := ed25519hd.SeedFromMnemonic(mnemonic, "password")
	if err != nil {
		return ag, err
	}
	masterKey, err := bip32.NewMasterKey(seed)

	// Use BIP44: m / purpose' / coin_type' / account' / change / address_index
	// Ethereum path: m/44'/60'/0'/0/0
	for i := 0; i < numWorkers; i++ {
		go func(workerId, start, end int) {
			defer wg.Done()
			for j := start; j < end; j++ {
				child, _ := masterKey.NewChildKey(uint32(j))
				privateKeyECDSA := crypto.ToECDSAUnsafe(child.Key)
				address := crypto.PubkeyToAddress(privateKeyECDSA.PublicKey)
				leadingZeroesCount := countLeadingZeroes(address.Hex())
				resultCh <- AddressDetails{j, address.Hex(), leadingZeroesCount}
			}
		}(i, i*(count/numWorkers), (i+1)*(count/numWorkers))
	}

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	for res := range resultCh {
		if res.LeadingZeroesCount > ag.LeadingZeroesCount {
			ag.Mnemonic = mnemonic
			ag.PathIndex = res.PathIndex
			ag.Address = res.Address
			ag.LeadingZeroesCount = res.LeadingZeroesCount
		}
	}

	//for i := 0; i <= count; i++ {
	//	child, cerr := masterKey.NewChildKey(uint32(i))
	//	if cerr != nil {
	//		return ag, err
	//	}
	//	privateKeyECDSA := crypto.ToECDSAUnsafe(child.Key)
	//	address := crypto.PubkeyToAddress(privateKeyECDSA.PublicKey)
	//	leadingZeroesCount := countLeadingZeroes(address.Hex())
	//	if leadingZeroesCount > ag.LeadingZeroesCount {
	//		ag.Mnemonic = mnemonic
	//		ag.PathIndex = i
	//		ag.Address = address.Hex()
	//		ag.LeadingZeroesCount = leadingZeroesCount
	//	}
	//	//fmt.Println("Leading zeroes: ", leadingZeroesCount)
	//	//fmt.Println("Ethereum Address: ", address.Hex())
	//	//fmt.Println("Private Key: ", hexutil.Encode(crypto.FromECDSA(privateKeyECDSA)))
	//}

	return ag, err
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

package zeus_ecdsa

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

func GenerateZeroPrefixAddresses(mnemonic, pw string, count, numWorkers int) (AddressGenerator, error) {
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	maxMutex := &sync.Mutex{}
	maxAddressDetails := AddressDetails{
		PathIndex:          0,
		Address:            "",
		LeadingZeroesCount: 0,
	}

	mnemonic, err := aegis_random.GenerateMnemonic()
	if err != nil {
		return AddressGenerator{}, err
	}
	seed, err := ed25519hd.SeedFromMnemonic(mnemonic, pw)
	if err != nil {
		return AddressGenerator{}, err
	}
	masterKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		return AddressGenerator{}, err
	}

	for i := 0; i < numWorkers; i++ {
		go func(workerId, start, end int) {
			defer wg.Done()
			for j := start; j < end; j++ {
				child, _ := masterKey.NewChildKey(uint32(j))
				privateKeyECDSA := crypto.ToECDSAUnsafe(child.Key)
				address := crypto.PubkeyToAddress(privateKeyECDSA.PublicKey)
				leadingZeroesCount := countLeadingZeroes(address.Hex())

				// Compare and replace if necessary
				maxMutex.Lock()
				if leadingZeroesCount > maxAddressDetails.LeadingZeroesCount {
					maxAddressDetails = AddressDetails{j, address.Hex(), leadingZeroesCount}
				}
				maxMutex.Unlock()
			}
		}(i, i*(count/numWorkers), (i+1)*(count/numWorkers))
	}

	wg.Wait()
	return AddressGenerator{
		Mnemonic:       mnemonic,
		AddressDetails: maxAddressDetails,
	}, nil
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

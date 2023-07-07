package ecdsa

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
	aegis_random "github.com/zeus-fyi/zeus/pkg/aegis/crypto/random"
)

type EcdsaTestSuite struct {
	suite.Suite
}

// m/44'/60'/0'/0/0
func (s *EcdsaTestSuite) TestEthWalletGeneration() {
	mnemonic, err := aegis_random.GenerateMnemonic()
	s.Require().Nil(err)

	ag, err := GenerateAddresses(mnemonic, 10000)
	s.Require().Nil(err)

	fmt.Println("Mnemonic: ", ag.Mnemonic)
	fmt.Println("Path Index: ", ag.PathIndex)
	fmt.Println("Address: ", ag.Address)
	fmt.Println("Leading Zeroes Count: ", ag.LeadingZeroesCount)
}

func TestEcdsaTestSuite(t *testing.T) {
	suite.Run(t, new(EcdsaTestSuite))
}

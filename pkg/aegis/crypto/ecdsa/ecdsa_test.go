package ecdsa

import (
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

	err = GenerateAddresses(mnemonic, 10)
	s.Require().Nil(err)
}

func TestEcdsaTestSuite(t *testing.T) {
	suite.Run(t, new(EcdsaTestSuite))
}

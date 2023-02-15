package ethereum_automation_cookbook

import (
	age_encryption "github.com/zeus-fyi/zeus/pkg/crypto/age"
	"testing"

	"github.com/stretchr/testify/suite"
	ethereum_cookbook_test_suite "github.com/zeus-fyi/zeus/cookbooks/ethereum/test"
)

type EthereumAutomationCookbookTestSuite struct {
	ethereum_cookbook_test_suite.EthereumCookbookTestSuite
}

// TestDecryptThenEncryptBatchWithAgeEncryption will decrypt your keystores only into memory, so they're never
// written to disk, then encrypt them with age encryption and write the encrypted values to disk

func (t *EthereumAutomationCookbookTestSuite) TestDecryptThenEncryptBatchWithAgeEncryption() {
	enc := age_encryption.NewAge(t.Tc.AgePrivKey, t.Tc.AgePubKey)
	keystoresPath := KeystorePath
	keystoresPath.DirOut = "./ethereum/automation/validator_keys/tmp"
	keystoresPath.DirIn = "./ethereum/automation/validator_keys/keystores"
	err := GenerateAgeKeystores(keystoresPath, enc, t.Tc.HDWalletPassword)
	t.Require().Nil(err)
}

func TestEthereumAutomationCookbookTestSuite(t *testing.T) {
	suite.Run(t, new(EthereumAutomationCookbookTestSuite))
}

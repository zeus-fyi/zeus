package ethereum_automation_cookbook

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
	ethereum_cookbook_test_suite "github.com/zeus-fyi/zeus/cookbooks/ethereum/test"
	age_encryption "github.com/zeus-fyi/zeus/pkg/aegis/crypto/age"
	signing_automation_ethereum "github.com/zeus-fyi/zeus/pkg/artemis/web3/signing_automation/ethereum"
)

type EthereumAutomationCookbookTestSuite struct {
	ethereum_cookbook_test_suite.EthereumCookbookTestSuite
}

// TestDecryptThenEncryptBatchWithAgeEncryption will decrypt your keystores only into memory, so they're never
// written to disk, then encrypt them with age encryption and write the encrypted values to disk

func (t *EthereumAutomationCookbookTestSuite) TestDecryptThenEncryptBatchWithAgeEncryption() {
	keystoresPath := KeystorePath
	keystoresPath.DirOut = "./ethereum/automation/validator_keys/keystores"
	keystoresPath.DirIn = "./ethereum/automation/validator_keys/tmp"

	agePubKey := t.Tc.AgePubKey
	agePrivKey := t.Tc.AgePrivKey
	hdWalletPassword := t.Tc.HDWalletPassword
	network := "ephemery"

	if agePubKey == "" || agePrivKey == "" {
		agePubKey, agePrivKey = age_encryption.GenerateNewKeyPair()
		fmt.Println("no credentials provided, generating new age keypair")
		fmt.Println("agePubKey: ", agePubKey)
		fmt.Println("agePrivKey: ", agePrivKey)
	}

	if hdWalletPassword == "" {
		hdWalletPassword = "password"
		fmt.Println("no hd wallet password provided, using default password: password")
	}

	offset := 0
	numKeys := 10

	vdg := signing_automation_ethereum.ValidatorDepositGenerationParams{
		Fp:                   keystoresPath,
		Mnemonic:             t.Tc.LocalMnemonic24Words,
		Pw:                   t.Tc.HDWalletPassword,
		ValidatorIndexOffset: offset,
		NumValidators:        numKeys,
		Network:              network,
	}

	enc := age_encryption.NewAge(agePrivKey, agePubKey)
	err := GenerateValidatorDepositsAndCreateAgeEncryptedKeystores(ctx, t.Web3SignerClientTestClient, vdg, enc, hdWalletPassword)
	t.Require().Nil(err)
}

func TestEthereumAutomationCookbookTestSuite(t *testing.T) {
	suite.Run(t, new(EthereumAutomationCookbookTestSuite))
}

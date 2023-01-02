package signing_automation_ethereum

import (
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	strings_filter "github.com/zeus-fyi/zeus/pkg/utils/strings"
	"github.com/zeus-fyi/zeus/test/configs"
)

var depositDataPath = filepaths.Path{
	PackageName: "",
	DirIn:       "./mocks/validator_keys",
	DirOut:      "../mocks/validator_keys",
	FnOut:       "",
	Env:         "",
	FilterFiles: &strings_filter.FilterOpts{},
}

/*
EIP-2334 defines derivation path indices for withdrawal and validator keys.
For a given index i the keys will be at the following paths:

withdrawal key: m/12381/3600/i/0
validator key: m/12381/3600/i/0/0
*/

// TestEphemeralDepositsFromMnemonicInEth2KeystoreFormat is useful for generating deposits for testing using
// the config file to provide your mnemonic

func (t *Web3SignerClientTestSuite) TestEphemeralDepositsFromMnemonicInEth2KeystoreFormat() {
	configs.ForceDirToConfigLocation()
	offset := 0
	numKeys := 3

	vdg := ValidatorDepositGenerationParams{
		Fp:                   depositDataPath,
		Mnemonic:             t.TestMnemonic,
		Pw:                   t.TestHDWalletPassword,
		ValidatorIndexOffset: offset,
		NumValidators:        numKeys,
	}
	err := vdg.GenerateAndEncryptValidatorKeysFromSeedAndPath(ctx, "ephemery")
	t.Require().Nil(err)
}

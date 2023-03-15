package signing_automation_ethereum

import (
	"encoding/base64"

	age_encryption "github.com/zeus-fyi/zeus/pkg/crypto/age"
	"github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/memfs"
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
		Network:              "ephemery",
	}
	err := vdg.GenerateAndEncryptValidatorKeysFromSeedAndPath(ctx)
	t.Require().Nil(err)

}

func (t *Web3SignerClientTestSuite) TestAgeEncryptedKeystoresGen() {
	t.Tc = configs.InitLocalTestConfigs()
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
	inMemFs := memfs.NewMemFs()
	enc := age_encryption.NewAge(t.Tc.AgePrivKey, t.Tc.AgePubKey)
	b, err := vdg.GenerateAgeEncryptedValidatorKeysInMemZipFile(ctx, inMemFs, enc)
	t.Require().Nil(err)
	t.Require().NotNil(b)
	base64EncodedData := base64.StdEncoding.EncodeToString(b)
	vdg.Fp.DirOut = "/Users/alex/go/Olympus/Zeus/pkg/artemis/signing_automation/ethereum/"
	vdg.Fp.FnOut = "keystores.zip"
	dst := make([]byte, base64.StdEncoding.DecodedLen(len(base64EncodedData)))
	_, err = base64.StdEncoding.Decode(dst, []byte(base64EncodedData))
	err = vdg.Fp.WriteToFileOutPath(dst)
	t.Require().Nil(err)
}

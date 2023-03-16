package signing_automation_ethereum

import (
	"encoding/json"

	aws_aegis_auth "github.com/zeus-fyi/zeus/pkg/aegis/aws/auth"
	aegis_aws_secretmanager "github.com/zeus-fyi/zeus/pkg/aegis/aws/secretmanager"
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

	t.Tc = configs.InitLocalTestConfigs()

	region := "us-west-1"
	a := aws_aegis_auth.AuthAWS{
		AccessKey: t.Tc.AccessKeyAWS,
		SecretKey: t.Tc.SecretKeyAWS,
		Region:    region,
	}
	sm, err := aegis_aws_secretmanager.InitSecretsManager(ctx, a)
	t.Require().Nil(err)
	t.Require().NotNil(sm)

	secretInfo := aegis_aws_secretmanager.SecretInfo{
		Region: region,
		Name:   "ageEncryptionKeyEphemery",
	}
	b, err := sm.GetSecretBinary(ctx, secretInfo)
	t.Require().Nil(err)

	m := make(map[string]any)
	err = json.Unmarshal(b, &m)
	t.Require().Nil(err)
	var enc age_encryption.Age
	for pubkey, privkey := range m {
		enc = age_encryption.NewAge(privkey.(string), pubkey)
	}
	offset := 0
	numKeys := 10

	vdg := ValidatorDepositGenerationParams{
		Fp:                   depositDataPath,
		Mnemonic:             t.TestMnemonic,
		Pw:                   t.TestHDWalletPassword,
		ValidatorIndexOffset: offset,
		NumValidators:        numKeys,
	}
	inMemFs := memfs.NewMemFs()
	zipBytes, err := vdg.GenerateAgeEncryptedValidatorKeysInMemZipFile(ctx, inMemFs, enc)
	t.Require().Nil(err)
	t.Require().NotEmpty(zipBytes)

	//p := filepaths.Path{}
	//p.DirOut = "./"
	//p.FnOut = "keystores.zip"
	//err = p.WriteToFileOutPath(zipBytes.Bytes())
	//t.Require().Nil(err)
}

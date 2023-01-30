package signing_automation_ethereum

import (
	"encoding/json"

	"github.com/zeus-fyi/zeus/cookbooks"
	bls_signer "github.com/zeus-fyi/zeus/pkg/crypto/bls"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
)

/*
	encryptor := keystorev4.New()
	input := make(map[string]interface{})
	err := json.Unmarshal([]byte(test.input), &input)
*/

func (t *Web3SignerClientTestSuite) TestKeystoreDecrypt() {
	ksPath := filepaths.Path{
		PackageName: "",
		DirIn:       "./ethereum/automation/validator_keys/ephemery",
		DirOut:      "./ethereum/automation/validator_keys/ephemery",
		FnIn:        "keystore-ephemery-m_12381_3600_0_0_0.json",
		FnOut:       "",
		Env:         "",
		FilterFiles: nil,
	}
	cookbooks.ChangeToCookbookDir()

	f := ksPath.ReadFileInPath()

	input := make(map[string]interface{})
	err := json.Unmarshal(f, &input)
	t.Require().Nil(err)
	out, err := DecryptKeystoreCipher(ctx, input, t.TestHDWalletPassword)
	t.Require().Nil(err)

	t.Assert().NotEmpty(out)

	key := bls_signer.NewEthSignerBLSFromExistingKeyBytes(out)

	pubKeyExp := "8a7addbf2857a72736205d861169c643545283a74a1ccb71c95dd2c9652acb89de226ca26d60248c4ef9591d7e010288"
	t.Assert().Equal(pubKeyExp, key.PublicKeyString())

	acc, err := DecryptKeystoreCipherIntoEthSignerBLS(ctx, input, t.TestHDWalletPassword)
	t.Require().Nil(err)

	t.Assert().Equal(pubKeyExp, acc.PublicKeyString())
}

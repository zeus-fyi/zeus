package hercules_client

import (
	"encoding/json"

	"github.com/zeus-fyi/hercules/api/v1/common/aegis"
	"github.com/zeus-fyi/zeus/cookbooks"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
)

func (t *HerculesClientTestSuite) TestImportKeystores() {
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

	ksImport := aegis.Keystore{
		KeystoreJSON: input,
		Password:     t.Tc.HDWalletPassword,
	}
	rr := aegis.ImportValidatorsRequest{ImportKeystores: []aegis.Keystore{ksImport}}
	err = t.HerculesTestClient.ImportKeystores(ctx, rr)
	t.Assert().Nil(err)
}

package signing_automation_ethereum

import (
	"context"

	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	strings_filter "github.com/zeus-fyi/zeus/pkg/utils/strings"
	test_base "github.com/zeus-fyi/zeus/test"
)

func (t *Web3SignerClientTestSuite) TestKeystoreParse() {
	keystorePath := filepaths.Path{
		PackageName: "",
		DirIn:       "./mocks/validator_keys",
		DirOut:      "",
		FnIn:        "deposit_data-1671500394.json",
		FnOut:       "",
		Env:         "",
		FilterFiles: strings_filter.FilterOpts{},
	}
	// points working dir to inside /test
	test_base.ForceDirToTestDirLocation()
	k, err := ParseKeystoreJSON(context.Background(), keystorePath)
	t.Require().Nil(err)
	t.Assert().NotEmpty(k)
}

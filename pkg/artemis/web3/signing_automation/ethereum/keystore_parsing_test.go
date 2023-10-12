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
		FnIn:        "",
		FnOut:       "",
		Env:         "",
		FilterFiles: &strings_filter.FilterOpts{StartsWith: "deposit_data-ephemery"},
	}
	// points working dir to inside /test
	test_base.ForceDirToTestDirLocation()
	k, err := ParseValidatorDepositSliceJSON(context.Background(), keystorePath)
	t.Require().Nil(err)
	t.Assert().NotEmpty(k)
}

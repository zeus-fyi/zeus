package signing_automation_ethereum

import (
	"fmt"
	"time"

	bls_signer "github.com/zeus-fyi/zeus/pkg/crypto/bls"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	strings_filter "github.com/zeus-fyi/zeus/pkg/utils/strings"
)

var depositDataPath = filepaths.Path{
	PackageName: "",
	DirIn:       "./mocks/validator_keys",
	DirOut:      "../mocks/validator_keys",
	FnOut:       "",
	Env:         "",
	FilterFiles: strings_filter.FilterOpts{},
}

func (t *Web3SignerClientTestSuite) TestEphemeralDepositGenerator() {
	s := bls_signer.NewSignerBLS()
	wd, err := ValidateAndReturnEcdsaPubkeyBytes(t.TestAccount1.PublicKey())
	t.Require().Nil(err)
	dd, err := GenerateEphemeralDepositData(s, wd)
	t.Require().Nil(err)
	t.Assert().NotEmpty(dd)

	depositDataPath.FnOut = fmt.Sprintf("deposit_data-ephemeral-%d.json", time.Now().Unix())
	dd.PrintJSON(depositDataPath)
}

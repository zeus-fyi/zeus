package signing_automation_ethereum

import (
	"context"

	artemis_client "github.com/zeus-fyi/zeus/pkg/artemis/client"
	artemis_req_types "github.com/zeus-fyi/zeus/pkg/artemis/client/req_types"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	strings_filter "github.com/zeus-fyi/zeus/pkg/utils/strings"
	test_base "github.com/zeus-fyi/zeus/test"
)

var ctx = context.Background()

func (t *Web3SignerClientTestSuite) TestSignedValidatorDepositTxPayload() {
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
	ks, err := ParseKeystoreJSON(ctx, keystorePath)
	t.Require().Nil(err)
	t.Require().NotEmpty(ks)

	acc := t.TestAccount1.Account

	st := acc.Address().String()
	t.Assert().NotNil(st)

	t.Web3SignerClientTestClient.Dial()
	defer t.Web3SignerClientTestClient.Client.Close()
	bal, err := t.Web3SignerClientTestClient.GetBalance(ctx, acc.Address().String(), nil)
	t.Require().Nil(err)
	t.Require().NotEmpty(bal)

	for i, k := range ks {
		params := k.ValidatorDepositParams
		signedTx, signErr := t.Web3SignerClientTestClient.SignValidatorDeposit(ctx, params)
		t.Require().Nil(signErr)
		t.Require().NotNil(signedTx)

		if i == 0 {
			payload := artemis_req_types.SignedTxPayload{Transaction: *signedTx}
			resp, aerr := t.ArtemisTestClient.SendSignedTx(ctx, &payload, artemis_client.ArtemisEthereumEphemeral)
			t.Require().Nil(aerr)
			t.Require().NotNil(resp)
		}
	}
}

func (t *Web3SignerClientTestSuite) TestValidatorABI() {
	ForceDirToEthSigningDirLocation()
	f, err := ABIOpenFile(validatorAbiFileLocation)
	t.Require().Nil(err)
	t.Require().NotEmpty(f)

	depositExists := false
	for _, mn := range f.Methods {
		if mn.Name == validatorDepositMethodName {
			depositExists = true
		}
	}
	t.Require().True(depositExists)
}

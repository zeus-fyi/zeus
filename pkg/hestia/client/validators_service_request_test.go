package hestia_client

import (
	"github.com/zeus-fyi/zeus/cookbooks"
	ethereum_automation_cookbook "github.com/zeus-fyi/zeus/cookbooks/ethereum/automation"
	signing_automation_ethereum "github.com/zeus-fyi/zeus/pkg/artemis/signing_automation/ethereum"
	hestia_req_types "github.com/zeus-fyi/zeus/pkg/hestia/client/req_types"
)

func (t *HestiaClientTestSuite) TestValidatorServiceRequest() {
	hs := hestia_req_types.CreateValidatorServiceRequest{}
	cookbooks.ChangeToCookbookDir()
	dp, err := signing_automation_ethereum.ParseValidatorDepositSliceJSON(ctx, ethereum_automation_cookbook.KeystorePath)
	t.Require().Nil(err)
	t.Assert().NotEmpty(dp)
	sr := hestia_req_types.ServiceRequestWrapper{
		GroupName:         "test",
		ProtocolNetworkID: hestia_req_types.EthereumEphemeryProtocolNetworkID,
		FeeRecipient:      t.TestAccount1.PublicKey(),
		Enabled:           true,
	}
	hs.CreateValidatorServiceRequest(dp, sr)
	t.Assert().NotEmpty(hs)
	resp, err := t.HestiaTestClient.ValidatorsServiceRequest(ctx, hs)
	t.Require().Nil(err)
	t.Assert().NotEmpty(resp)
}

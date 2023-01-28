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
		Enabled:           true,
	}
	keyOne := hestia_req_types.ValidatorServiceOrgGroup{
		Pubkey:       "0x8a7addbf2857a72736205d861169c643545283a74a1ccb71c95dd2c9652acb89de226ca26d60248c4ef9591d7e010288",
		FeeRecipient: "0xF7Ab1d834Cd0A33691e9A750bD720cb6436cA1B9",
	}
	keyTwo := hestia_req_types.ValidatorServiceOrgGroup{
		Pubkey:       "0x9a7addbf2857a72736205d861169c643545283a74a1ccb71c95dd2c9652acb89de226ca26d60248c4ef9591d7e010288",
		FeeRecipient: "0xF7Ab1d834Cd0A33691e9A750bD720cb6436cA1B9",
	}
	pubkeys := hestia_req_types.ValidatorServiceOrgGroupSlice{keyOne, keyTwo}

	hs.CreateValidatorServiceRequest(pubkeys, sr)
	t.Assert().NotEmpty(hs)
	resp, err := t.HestiaTestClient.ValidatorsServiceRequest(ctx, hs)
	t.Require().Nil(err)
	t.Assert().NotEmpty(resp)
}

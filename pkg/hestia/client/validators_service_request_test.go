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
		GroupName:         "testGroup",
		ProtocolNetworkID: hestia_req_types.EthereumEphemeryProtocolNetworkID,
		Enabled:           true,
		ServiceAuth: hestia_req_types.ServiceAuthConfig{
			AuthLamdbaAWS: &hestia_req_types.AuthLamdbaAWS{
				ServiceURL:   t.Tc.ServerlessSignerFuncBLS,
				SecretName:   t.Tc.ServerlessSignerFuncSecretName,
				AccessKey:    t.Tc.AwsAccessKeyLambdaInvoke,
				AccessSecret: t.Tc.AwsSecretKeyLambdaInvoke,
			}},
	}

	pubkeys := hestia_req_types.ValidatorServiceOrgGroupSlice{}
	for _, validatorDepositInfo := range dp {

		pubkeys = append(pubkeys, hestia_req_types.ValidatorServiceOrgGroup{
			Pubkey:       validatorDepositInfo.Pubkey,
			FeeRecipient: "0xF7Ab1d834Cd0A33691e9A750bD720cb6436cA1B9",
		})
		if len(pubkeys) == 100 {
			break
		}
	}

	hs.CreateValidatorServiceRequest(pubkeys, sr)
	t.Assert().NotEmpty(hs)
	resp, err := t.HestiaTestClient.ValidatorsServiceRequest(ctx, hs)
	t.Require().Nil(err)
	t.Assert().NotEmpty(resp)
}

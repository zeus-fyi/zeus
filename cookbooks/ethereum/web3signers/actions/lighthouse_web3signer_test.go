package ethereum_web3signer_actions

import (
	"context"
	"fmt"

	validator_cookbooks "github.com/zeus-fyi/zeus/cookbooks/ethereum/validators"
	signing_automation_ethereum "github.com/zeus-fyi/zeus/pkg/artemis/signing_automation/ethereum"
	strings_filter "github.com/zeus-fyi/zeus/pkg/utils/strings"
)

func (t *EthereumWeb3SignerCookbookTestSuite) TestLighthouseImportWeb3SignerAPI() {
	fl := strings_filter.FilterOpts{StartsWith: "deposit"}
	keystorePath.FilterFiles = &fl
	k, err := signing_automation_ethereum.ParseValidatorDepositSliceJSON(context.Background(), keystorePath)
	t.Require().Nil(err)
	lh := LighthouseWeb3SignerRequests{}
	lh.Web3SignerURL = "http://zeus-web3signer:9000"
	lh.Enable = true
	lh.FeeAddr = t.TestAccount1.PublicKey()
	lh.ReadDepositParamsAndExtractToEnableKeysOnWeb3Signer(ctx, k)
	kns := validator_cookbooks.ValidatorCloudCtxNs
	w3 := Web3SignerActionsClient{t.ZeusTestClient}
	resp, err := w3.EnableWeb3SignerLighthouse(ctx, kns, lh.Slice)
	t.Require().Nil(err)
	t.Assert().NotEmpty(resp)
	fmt.Println(resp)
}

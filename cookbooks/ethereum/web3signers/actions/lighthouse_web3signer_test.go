package ethereum_web3signer_actions

import (
	"context"
	"fmt"
	"time"

	validator_cookbooks "github.com/zeus-fyi/zeus/cookbooks/ethereum/validators"
	strings_filter "github.com/zeus-fyi/zeus/pkg/utils/strings"
)

func (t *EthereumWeb3SignerCookbookTestSuite) TestLighthouseHerculesAuthRoute() {
	kns := validator_cookbooks.ValidatorCloudCtxNs
	w3 := Web3SignerActionsClient{t.ZeusTestClient}
	resp, err := w3.GetLighthouseAuth(ctx, kns)
	t.Require().Nil(err)
	t.Assert().NotEmpty(resp)

	token := resp.GetAnyValue()
	fmt.Println(token)
}

func (t *EthereumWeb3SignerCookbookTestSuite) TestLighthouseImportWeb3SignerAPI() {
	w3 := Web3SignerActionsClient{t.ZeusTestClient}
	kns := validator_cookbooks.ValidatorCloudCtxNs

	resp, err := w3.GetLighthouseAuth(ctx, kns)
	t.Require().Nil(err)
	t.Assert().NotEmpty(resp)

	token := resp.GetAnyValue()
	fl := strings_filter.FilterOpts{StartsWith: "deposit"}
	KeystorePath.FilterFiles = &fl
	k, err := signing_automation_ethereum.ParseValidatorDepositSliceJSON(context.Background(), KeystorePath)
	t.Require().Nil(err)
	lh := LighthouseWeb3SignerRequests{}
	lh.Web3SignerURL = "http://zeus-web3signer:9000"
	lh.Enabled = true
	lh.FeeAddr = t.TestAccount1.PublicKey()
	lh.ReadDepositParamsAndExtractToEnableKeysOnWeb3Signer(ctx, k)

	resp, err = w3.EnableWeb3SignerLighthouse(ctx, kns, lh.Slice, string(token))
	t.Require().Nil(err)
	t.Assert().NotEmpty(resp)
	fmt.Println(resp)
	tmp := LighthouseWeb3SignerRequests{
		Web3SignerURL: "http://zeus-web3signer:9000",
		Enabled:       true,
		FeeAddr:       t.TestAccount1.PublicKey(),
		Slice:         []LighthouseWeb3SignerRequest{},
	}

	for _, v := range lh.Slice {
		tmp.Slice = []LighthouseWeb3SignerRequest{v}
		resp, err = w3.EnableWeb3SignerLighthouse(ctx, kns, tmp.Slice, string(token))
		t.Require().Nil(err)
		t.Assert().NotEmpty(resp)
		fmt.Println(resp)
		time.Sleep(1 * time.Second)
	}

	resp, err = w3.EnableWeb3SignerLighthouse(ctx, kns, tmp.Slice, string(token))
	t.Require().Nil(err)
	t.Assert().NotEmpty(resp)
	fmt.Println(resp)
}

func (t *EthereumWeb3SignerCookbookTestSuite) TestRemoteKeystoresLH() {
	w3 := Web3SignerActionsClient{t.ZeusTestClient}
	kns := validator_cookbooks.ValidatorCloudCtxNs
	resp, err := w3.GetLighthouseRemoteKeystores(ctx, kns)
	t.Require().Nil(err)
	t.Assert().NotEmpty(resp)
	clientR := resp.GetAnyValue()
	fmt.Println(string(clientR))
}

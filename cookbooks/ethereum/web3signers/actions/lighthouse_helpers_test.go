package ethereum_web3signer_actions

import (
	"context"

	signing_automation_ethereum "github.com/zeus-fyi/zeus/pkg/artemis/signing_automation/ethereum"
	strings_filter "github.com/zeus-fyi/zeus/pkg/utils/strings"
)

func (t *EthereumWeb3SignerCookbookTestSuite) TestWriteYamlConfig() {
	k := Web3SignerKeystores{}
	k.ReadKeystoreDirAndAppendPw(ctx, KeystorePath, t.Tc.HDWalletPassword)
	t.Assert().NotEmpty(k.Keystores)

	fl := strings_filter.FilterOpts{StartsWith: "deposit"}
	KeystorePath.FilterFiles = &fl
	dp, err := signing_automation_ethereum.ParseValidatorDepositSliceJSON(context.Background(), KeystorePath)
	t.Require().Nil(err)
	lh := LighthouseWeb3SignerRequests{}
	lh.Web3SignerURL = "http://zeus-web3signer:9000"
	lh.Enable = true
	lh.FeeAddr = t.TestAccount1.PublicKey()
	lh.ReadDepositParamsAndExtractToEnableKeysOnWeb3Signer(ctx, dp)

	od := "./ethereum/automation/lighthouse_config_yaml"

	KeystorePath.DirOut = od
	err = lh.WriteYamlConfig(KeystorePath)
	t.Require().Nil(err)
}

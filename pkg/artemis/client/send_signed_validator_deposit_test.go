package artemis_client

import (
	artemis_req_types "github.com/zeus-fyi/zeus/pkg/artemis/client/req_types"
	signing_automation_ethereum "github.com/zeus-fyi/zeus/pkg/artemis/signing_automation/ethereum"
)

// TestSendSignedValidatorDepositTxPayload uses ephemery network for this test
func (t *ArtemisClientTestSuite) TestSendSignedValidatorDepositTxPayload() {
	vdg := signing_automation_ethereum.ValidatorDepositGenerationParams{
		Mnemonic:             t.Tc.LocalMnemonic24Words,
		Pw:                   t.Tc.HDWalletPassword,
		ValidatorIndexOffset: 0,
		NumValidators:        1,
	}
	dp, err := t.Web3SignerClientTestClient.GenerateEphemeryDepositDataWithDefaultWd(ctx, vdg)
	t.Require().Nil(err)

	t.Require().Len(dp, 1)
	t.Require().Nil(err)
	signedTx, err := t.Web3SignerClientTestClient.SignValidatorDepositTxToBroadcast(ctx, dp[0])
	t.Require().Nil(err)
	t.Require().NotNil(signedTx)
	payload := artemis_req_types.SignedTxPayload{Transaction: *signedTx}
	resp, err := t.ArtemisTestClient.SendSignedTx(ctx, &payload, ArtemisEthereumEphemeral)
	t.Assert().Nil(err)
	t.Assert().NotNil(resp)
}

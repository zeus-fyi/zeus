package artemis_client

import (
	artemis_req_types "github.com/zeus-fyi/zeus/pkg/artemis/client/req_types"
	signing_automation_ethereum "github.com/zeus-fyi/zeus/pkg/artemis/signing_automation/ethereum"
)

// TestSendSignedValidatorDepositTxPayload uses ephemery network for this test
func (t *ArtemisClientTestSuite) TestSendSignedValidatorDepositTxPayload() {
	wc, err := signing_automation_ethereum.ValidateAndReturnEcdsaPubkeyBytes(t.TestAccount1.PublicKey())
	t.Require().Nil(err)
	dd, err := signing_automation_ethereum.GenerateEphemeralDepositData(t.TestBLSAccount, wc)
	t.Require().Nil(err)
	signedTx, err := t.Web3SignerClientTestClient.SignValidatorDepositTxToBroadcast(ctx, dd)
	t.Require().Nil(err)
	t.Require().NotNil(signedTx)
	payload := artemis_req_types.SignedTxPayload{Transaction: *signedTx}
	resp, err := t.ArtemisTestClient.SendSignedTx(ctx, &payload, ArtemisEthereumEphemeral)
	t.Assert().Nil(err)
	t.Assert().NotNil(resp)
}

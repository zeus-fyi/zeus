package signing_automation_ethereum

import (
	"context"
)

var ctx = context.Background()

func (t *Web3SignerClientTestSuite) TestSignedValidatorDepositTxPayload() {
	params := ValidatorDepositParams{
		Pubkey:                "",
		WithdrawalCredentials: "",
		Signature:             "",
		DepositDataRoot:       "",
	}
	signedTx, err := t.Web3SignerClientTestClient.SignValidatorDeposit(ctx, params)
	t.Require().Nil(err)
	t.Require().NotNil(signedTx)
}

package signing_automation_ethereum

import (
	"github.com/zeus-fyi/gochain/web3/accounts"
	web3_actions "github.com/zeus-fyi/gochain/web3/client"
)

func (t *Web3SignerClientTestSuite) TestEtherSend() {
	t.Web3SignerClientTestClient.Dial()
	defer t.Web3SignerClientTestClient.Close()
	sendEthTx := web3_actions.SendEtherPayload{
		TransferArgs: web3_actions.TransferArgs{
			Amount:    Finney,
			ToAddress: accounts.Address(t.TestAccount2.Address()),
		},
		GasPriceLimits: web3_actions.GasPriceLimits{},
	}
	rx, err := t.Web3SignerClientTestClient.Send(ctx, sendEthTx)
	t.Require().Nil(err)
	t.Require().NotNil(rx)
}

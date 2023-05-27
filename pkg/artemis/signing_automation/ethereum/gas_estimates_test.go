package signing_automation_ethereum

import (
	"fmt"

	"github.com/zeus-fyi/gochain/v4/common"
	web3_actions "github.com/zeus-fyi/gochain/web3/client"
)

func (t *Web3SignerClientTestSuite) TestValidatorDepositPayloadGasEstimate() {
	t.Web3SignerClientTestClient.Dial()
	defer t.Web3SignerClientTestClient.Close()

	est, err := t.Web3SignerClientTestClient.Web3Actions.C.SuggestGasPrice(ctx)
	t.Require().Nil(err)
	t.Require().NotNil(est)
	fmt.Println(est.Uint64())

	thirtyTwoEthInGwei := uint64(32000000000)
	t.Require().Equal(ValidatorDeposit32EthInGweiUnits.Uint64(), thirtyTwoEthInGwei)
}

func (t *Web3SignerClientTestSuite) TestValidatorDeposit() {
	t.Web3SignerClientTestClient.Dial()
	defer t.Web3SignerClientTestClient.Close()
	vdg := ValidatorDepositGenerationParams{
		Fp:                   depositDataPath,
		Mnemonic:             t.Tc.LocalMnemonic24Words,
		Pw:                   t.Tc.HDWalletPassword,
		ValidatorIndexOffset: 0,
		NumValidators:        1,
	}
	dd, err := t.Web3SignerClientTestClient.GenerateEphemeryDepositDataWithDefaultWd(ctx, vdg)
	t.Require().Nil(err)

	t.Require().Len(dd, 1)
	broadcast, err := t.Web3SignerClientTestClient.SignValidatorDepositTxToBroadcast(ctx, dd[0])
	t.Require().Nil(err)
	t.Require().NotNil(broadcast)

	rx, err := t.Web3SignerClientTestClient.SubmitSignedTxAndReturnTxData(ctx, broadcast)
	t.Require().Nil(err)
	t.Require().NotNil(rx)
}

func (t *Web3SignerClientTestSuite) TestSendEtherGasEstimates() {
	t.Web3SignerClientTestClient.Dial()
	defer t.Web3SignerClientTestClient.Close()

	est, err := t.Web3SignerClientTestClient.C.SuggestGasPrice(ctx)
	t.Require().Nil(err)
	t.Require().NotNil(est)

	fmt.Println(est.Uint64())
	est, err = t.Web3SignerClientTestClient.C.SuggestGasPrice(ctx)
	t.Require().Nil(err)
	t.Require().NotNil(est)
	fmt.Println(est.Uint64())
}

func (t *Web3SignerClientTestSuite) TestSendEther() {
	t.Web3SignerClientTestClient.Dial()
	defer t.Web3SignerClientTestClient.Close()
	sendEthTx := web3_actions.SendEtherPayload{
		TransferArgs: web3_actions.TransferArgs{
			Amount:    Finney,
			ToAddress: common.Address(t.TestAccount2.Address()),
		},
		GasPriceLimits: web3_actions.GasPriceLimits{},
	}
	est, err := t.Web3SignerClientTestClient.Web3Actions.C.SuggestGasPrice(ctx)
	t.Require().Nil(err)
	sendEthTx.GasPrice = est
	rx, err := t.Web3SignerClientTestClient.Send(ctx, sendEthTx)
	t.Require().Nil(err)
	t.Require().NotNil(rx)
	fmt.Println(rx)
}

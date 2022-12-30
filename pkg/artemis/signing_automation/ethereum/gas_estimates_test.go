package signing_automation_ethereum

import (
	"fmt"

	"github.com/zeus-fyi/gochain/web3/web3_actions"
)

func (t *Web3SignerClientTestSuite) TestValidatorDepositPayloadGasEstimate() {
	t.Web3SignerClientTestClient.Dial()
	defer t.Web3SignerClientTestClient.Close()
	wc, err := ValidateAndReturnEcdsaPubkeyBytes(t.TestAccount1.PublicKey())
	t.Require().Nil(err)
	dd, err := GenerateEphemeralDepositData(t.TestBLSAccount, wc)
	t.Require().Nil(err)
	payload, err := getValidatorDepositPayload(ctx, dd)
	t.Require().Nil(err)

	from := t.TestAccount1.Address()
	txPayload, err := extractCallMsgFromSendContractTxPayload(&from, payload)
	t.Require().Nil(err)
	est, err := t.Web3SignerClientTestClient.GetGasPriceEstimateForTx(ctx, txPayload)
	t.Require().Nil(err)
	t.Require().NotNil(est)
	fmt.Println(est.Uint64())

	thirtyTwoEthInGwei := uint64(32000000000)
	t.Require().Equal(ValidatorDeposit32EthInGweiUnits.Uint64(), thirtyTwoEthInGwei)
}

func (t *Web3SignerClientTestSuite) TestValidatorDeposit() {
	t.Web3SignerClientTestClient.Dial()
	defer t.Web3SignerClientTestClient.Close()
	wc, err := ValidateAndReturnEcdsaPubkeyBytes(t.TestAccount1.PublicKey())
	t.Require().Nil(err)
	dd, err := GenerateEphemeralDepositData(t.TestBLSAccount, wc)
	t.Require().Nil(err)

	broadcast, err := t.Web3SignerClientTestClient.SignValidatorDepositTxToBroadcast(ctx, dd)
	t.Require().Nil(err)
	t.Require().NotNil(broadcast)

	rx, err := t.Web3SignerClientTestClient.SubmitSignedTxAndReturnTxData(ctx, broadcast)
	t.Require().Nil(err)
	t.Require().NotNil(rx)
}

func (t *Web3SignerClientTestSuite) TestSendEtherGasEstimates() {
	sendEthTx := web3_actions.SendEtherPayload{
		TransferArgs: web3_actions.TransferArgs{
			Amount:    Finney,
			ToAddress: t.TestAccount2.Address(),
		},
		GasPriceLimits: web3_actions.GasPriceLimits{},
	}
	from := t.TestAccount1.Address()
	msg := extractCallMsgFromSendEtherPayload(&from, sendEthTx)
	t.Web3SignerClientTestClient.Dial()
	defer t.Web3SignerClientTestClient.Close()

	est, err := t.Web3SignerClientTestClient.GetGasPrice(ctx)
	t.Require().Nil(err)
	t.Require().NotNil(est)

	fmt.Println(est.Uint64())
	est, err = t.Web3SignerClientTestClient.GetGasPriceEstimateForTx(ctx, msg)
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
			ToAddress: t.TestAccount2.Address(),
		},
		GasPriceLimits: web3_actions.GasPriceLimits{},
	}
	from := t.TestAccount1.Address()
	msg := extractCallMsgFromSendEtherPayload(&from, sendEthTx)
	est, err := t.Web3SignerClientTestClient.GetGasPriceEstimateForTx(ctx, msg)
	t.Require().Nil(err)
	sendEthTx.GasPrice = est

	rx, err := t.Web3SignerClientTestClient.Send(ctx, sendEthTx)
	t.Require().Nil(err)
	t.Require().NotNil(rx)

	fmt.Println(rx)

}

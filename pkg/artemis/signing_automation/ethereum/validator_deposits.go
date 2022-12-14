package signing_automation_ethereum

import (
	"context"
	"encoding/hex"
	"strings"

	"github.com/gochain/gochain/v4/core/types"
	"github.com/rs/zerolog/log"
	"github.com/zeus-fyi/gochain/web3/web3_actions"
)

const (
	validatorDepositMethodName   = "deposit"
	validatorAbiFileLocation     = "smart_contracts/eth_deposit_contract.json"
	EphemeralDepositContractAddr = "0x4242424242424242424242424242424242424242"
	EphemeralBeacon              = "https://eth.ephemeral.zeus.fyi"
	BeaconGenesisPath            = "/eth/v1/beacon/genesis"
	BeaconForkPath               = "/eth/v1/beacon/states/head/fork"
)

func (w *Web3SignerClient) SignValidatorDepositTxToBroadcast(ctx context.Context, depositParams *DepositDataParams) (*types.Transaction, error) {
	w.Dial()
	defer w.Close()
	params, err := getValidatorDepositPayload(ctx, depositParams)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("Web3SignerClient: SignValidatorDepositTxToBroadcast")
		return nil, err
	}
	from := w.Address()
	txPayload, err := extractCallMsgFromSendContractTxPayload(ctx, &from, params)
	if err != nil {
		panic(err)
	}
	est, err := w.GetGasPriceEstimateForTx(ctx, txPayload)
	if err != nil {
		panic(err)
	}
	params.GasPrice = est
	params.GasLimit = 100000
	signedTx, err := w.GetSignedTxToCallFunctionWithArgs(ctx, &params)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("Web3SignerClient: SignValidatorDepositTxToBroadcast")
		return nil, err
	}
	return signedTx, err
}

func getValidatorDepositPayload(ctx context.Context, depositParams *DepositDataParams) (web3_actions.SendContractTxPayload, error) {
	ForceDirToEthSigningDirLocation()
	abiFile, err := ABIOpenFile(ctx, validatorAbiFileLocation)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("Web3SignerClient: SignValidatorDeposit: ABIOpenFile")
		return web3_actions.SendContractTxPayload{}, err
	}

	pubkey, err := hex.DecodeString(strings.TrimPrefix(depositParams.PublicKey.String(), "0x"))
	if err != nil {
		panic(err)
	}
	sig, err := hex.DecodeString(strings.TrimPrefix(depositParams.Signature.String(), "0x"))
	if err != nil {
		panic(err)
	}

	params := web3_actions.SendContractTxPayload{
		SmartContractAddr: EphemeralDepositContractAddr,
		ContractABI:       abiFile,
		MethodName:        validatorDepositMethodName,
		SendEtherPayload: web3_actions.SendEtherPayload{
			TransferArgs: web3_actions.TransferArgs{
				Amount: ValidatorDeposit32Eth,
			},
			GasPriceLimits: web3_actions.GasPriceLimits{},
		},
		Params: []interface{}{pubkey, depositParams.WithdrawalCredentials, sig, depositParams.DepositDataRoot},
	}
	return params, err
}

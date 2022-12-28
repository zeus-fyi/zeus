package signing_automation_ethereum

import (
	"context"

	"github.com/gochain/gochain/v4/common"
	"github.com/gochain/gochain/v4/core/types"
	"github.com/rs/zerolog/log"
	"github.com/zeus-fyi/gochain/web3/web3_actions"
)

const (
	validatorDepositMethodName   = "deposit"
	validatorAbiFileLocation     = "smart_contracts/eth_deposit_contract.json"
	EphemeralDepositContractAddr = "0x4242424242424242424242424242424242424242"
)

func (w *Web3SignerClient) SignValidatorDeposit(ctx context.Context, depositParams ValidatorDepositParams) (*types.Transaction, error) {

	abiFile, err := ABIOpenFile(validatorAbiFileLocation)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("Web3SignerClient: SignValidatorDeposit: ABIOpenFile")
		return nil, err
	}
	params := web3_actions.SendContractTxPayload{
		SmartContractAddr: EphemeralDepositContractAddr,
		ContractABI:       abiFile,
		MethodName:        validatorDepositMethodName,
		SendEtherPayload: web3_actions.SendEtherPayload{
			TransferArgs: web3_actions.TransferArgs{
				Amount:    ValidatorDeposit32Eth,
				ToAddress: common.Address{},
			},
			GasPriceLimits: web3_actions.GasPriceLimits{},
		},
		Params: []interface{}{"0x" + depositParams.Pubkey, "0x" + depositParams.WithdrawalCredentials, "0x" + depositParams.Signature, "0x" + depositParams.DepositDataRoot},
	}
	signedTx, err := w.GetSignedTxToCallFunctionWithArgs(ctx, &params)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("Web3SignerClient: SignValidatorDeposit")
	}
	return signedTx, err
}

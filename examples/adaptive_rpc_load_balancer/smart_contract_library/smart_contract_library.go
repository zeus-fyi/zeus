package smart_contract_library

import (
	"context"

	web3_actions "github.com/zeus-fyi/zeus/pkg/artemis/web3/client"
	signing_automation_ethereum "github.com/zeus-fyi/zeus/pkg/artemis/web3/signing_automation/ethereum"
)

func MustLoadERC20AbiPayload(ctx context.Context) web3_actions.SendContractTxPayload {
	abiFile := signing_automation_ethereum.MustReadAbiString(ctx, Erc20Abi)
	params := web3_actions.SendContractTxPayload{
		SendEtherPayload: web3_actions.SendEtherPayload{},
		ContractABI:      abiFile,
		Params:           []interface{}{},
	}
	return params
}

package signing_automation_ethereum

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/zeus-fyi/gochain/v4/common"
	web3_actions "github.com/zeus-fyi/gochain/web3/client"
	web3_types "github.com/zeus-fyi/gochain/web3/types"
)

func extractCallMsgFromSendContractTxPayload(ctx context.Context, from *common.Address, sendContractTxPayload web3_actions.SendContractTxPayload) (web3_types.CallMsg, error) {
	msg := extractCallMsgFromSendEtherPayload(from, sendContractTxPayload.SendEtherPayload)
	if sendContractTxPayload.ContractABI != nil {
		b, err := sendContractTxPayload.ContractABI.Pack(sendContractTxPayload.MethodName,
			sendContractTxPayload.Params...)
		if err != nil {
			log.Ctx(ctx).Err(err)
			return web3_types.CallMsg{}, err
		}
		msg.Data = b
	}
	return msg, nil
}

func extractCallMsgFromSendEtherPayload(from *common.Address, payload web3_actions.SendEtherPayload) web3_types.CallMsg {
	var msg web3_types.CallMsg
	msg.From = from
	msg.To = &payload.ToAddress
	msg.Gas = payload.GasLimit
	msg.GasPrice = payload.GasPrice
	msg.Value = payload.Amount
	msg.Data = nil
	return msg
}

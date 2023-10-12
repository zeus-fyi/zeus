package web3_actions

import (
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func GetSender(tx *types.Transaction) (common.Address, error) {
	if tx == nil {
		return common.Address{}, errors.New("tx is nil")
	}
	if tx.ChainId() == nil {
		return common.Address{}, errors.New("tx.ChainId() is nil")
	}
	sender := types.LatestSignerForChainID(tx.ChainId())
	from, err := sender.Sender(tx)
	if err != nil {
		return common.Address{}, err
	}
	return from, nil
}

package web3_actions

import (
	"context"
	"fmt"

	common2 "github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog/log"
	"github.com/zeus-fyi/gochain/v4/common"
	"github.com/zeus-fyi/gochain/v4/crypto"
)

func (w *Web3Actions) GetNonce(ctx context.Context) (uint64, error) {
	w.Dial()
	defer w.C.Close()
	publicKeyECDSA := w.EcdsaPublicKey()
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	return w.GetNonceByPublicAddress(ctx, fromAddress)
}

func (w *Web3Actions) GetNonceByPublicAddress(ctx context.Context, fromAddress common.Address) (uint64, error) {
	w.Dial()
	defer w.C.Close()
	nonce, err := w.C.PendingNonceAt(ctx, common2.Address(fromAddress))
	if err != nil {
		log.Ctx(ctx).Err(err).Interface("fromAddress", fromAddress).Msg("Web3Actions: GetNonceByPublicAddress")
		return 0, fmt.Errorf("cannot get nonce: %v", err)
	}
	return nonce, err
}

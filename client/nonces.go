package web3_actions

import (
	"context"
	"fmt"

	common2 "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/rs/zerolog/log"
	"github.com/zeus-fyi/gochain/web3/accounts"
)

func (w *Web3Actions) GetNonce(ctx context.Context) (uint64, error) {
	w.Dial()
	defer w.C.Close()
	publicKeyECDSA := w.EcdsaPublicKey()
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	return w.GetNonceByPublicAddress(ctx, accounts.Address(fromAddress))
}

func (w *Web3Actions) GetNonceByPublicAddress(ctx context.Context, fromAddress accounts.Address) (uint64, error) {
	w.Dial()
	defer w.C.Close()
	nonce, err := w.C.PendingNonceAt(ctx, common2.Address(fromAddress))
	if err != nil {
		log.Ctx(ctx).Err(err).Interface("fromAddress", fromAddress).Msg("Web3Actions: GetNonceByPublicAddress")
		return 0, fmt.Errorf("cannot get nonce: %v", err)
	}
	return nonce, err
}

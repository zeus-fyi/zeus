package web3_actions

import (
	"context"
	"math/big"
	"net/url"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	zlog "github.com/rs/zerolog/log"
	"github.com/zeus-fyi/gochain/web3/accounts"
)

type Web3Actions struct {
	C *ethclient.Client
	*accounts.Account
	Headers map[string]string
	NodeURL string
	RelayTo string
	Network string
}

func (w *Web3Actions) Dial() {
	if w.Headers == nil {
		w.Headers = make(map[string]string)
	}
	nodeUrl := w.NodeURL
	if len(w.RelayTo) > 0 {
		relayUrlVal, rerr := url.ParseRequestURI(w.RelayTo)
		if rerr == nil {
			w.Headers["Proxy-Relay-To"] = nodeUrl
			nodeUrl = relayUrlVal.String()
		}
	}
	ctx := context.Background()
	cli, err := ethclient.DialContext(ctx, nodeUrl)
	if err != nil {
		panic(err)
	}
	w.C = cli
	for k, h := range w.Headers {
		w.C.Client().SetHeader(k, h)
	}
}

func (w *Web3Actions) Close() {
	w.C.Close()
}

func NewWeb3ActionsClient(nodeUrl string) Web3Actions {
	return Web3Actions{
		NodeURL: nodeUrl,
	}
}

func NewWeb3ActionsClientWithRelay(nodeUrl, relayUrl string, accounts *accounts.Account) Web3Actions {
	return Web3Actions{
		NodeURL: nodeUrl,
		RelayTo: relayUrl,
		Account: accounts,
	}
}

func NewWeb3ActionsClientWithAccount(nodeUrl string, account *accounts.Account) Web3Actions {
	return Web3Actions{
		NodeURL: nodeUrl,
		Account: account,
	}
}

func (w *Web3Actions) MineBlock(ctx context.Context, blocksToMine hexutil.Big) error {
	err := w.C.Client().CallContext(ctx, nil, "hardhat_mine", blocksToMine.String())
	return err
}

func (w *Web3Actions) GetStorageAt(ctx context.Context, addr, slot string) (hexutil.Bytes, error) {
	var result hexutil.Bytes
	err := w.C.Client().CallContext(ctx, &result, "eth_getStorageAt", addr, slot)
	return result, err
}

func (w *Web3Actions) SetStorageAt(ctx context.Context, addr, slot, value string) error {
	err := w.C.Client().CallContext(ctx, nil, "hardhat_setStorageAt", addr, slot, value)
	return err
}

func (w *Web3Actions) GetEVMSnapshot(ctx context.Context) (*big.Int, error) {
	var result hexutil.Big
	err := w.C.Client().CallContext(ctx, &result, "evm_snapshot")
	return (*big.Int)(&result), err
}

func (w *Web3Actions) ResetNetwork(ctx context.Context, rpcUrl string, blockNumber int) error {
	if rpcUrl != "" && blockNumber != 0 {
		args := toForkingArg(rpcUrl, blockNumber)
		return w.C.Client().CallContext(ctx, nil, "hardhat_reset", args)
	}
	return w.C.Client().CallContext(ctx, nil, "hardhat_reset")
}

func (w *Web3Actions) ImpersonateAccount(ctx context.Context, address string) error {
	var result any
	err := w.C.Client().CallContext(ctx, &result, "hardhat_impersonateAccount", accounts.HexToAddress(address))
	return err
}

func (w *Web3Actions) StopImpersonatingAccount(ctx context.Context, address string) error {
	err := w.C.Client().CallContext(ctx, nil, "hardhat_stopImpersonatingAccount", accounts.HexToAddress(address))
	return err
}

func (w *Web3Actions) SetNonce(ctx context.Context, address string, nonce hexutil.Big) error {
	err := w.C.Client().CallContext(ctx, nil, "hardhat_setNonce", accounts.HexToAddress(address), nonce.String())
	return err
}

func (w *Web3Actions) SetCode(ctx context.Context, address string, bytes string) error {
	err := w.C.Client().CallContext(ctx, nil, "hardhat_setCode", accounts.HexToAddress(address), bytes)
	return err
}

func (w *Web3Actions) SetBalance(ctx context.Context, address string, balance hexutil.Big) error {
	err := w.C.Client().CallContext(ctx, nil, "hardhat_setBalance", accounts.HexToAddress(address), balance)
	if err != nil {
		zlog.Err(err).Msg("HardHatSetBalance error")
		return err
	}
	return err
}

func (w *Web3Actions) SendRawTransaction(ctx context.Context, tx *types.Transaction) error {
	data, err := rlp.EncodeToBytes(tx)
	if err != nil {
		return err
	}
	err = w.C.Client().CallContext(ctx, nil, "eth_sendRawTransaction", hexutil.Encode(data))
	return err
}

func (w *Web3Actions) GetNumber(ctx context.Context, address string, blockNumber *big.Int) (*big.Int, error) {
	var result hexutil.Big
	err := w.C.Client().CallContext(ctx, &result, "eth_getBalance", accounts.HexToAddress(address), toBlockNumArg(blockNumber))
	return (*big.Int)(&result), err
}

func (w *Web3Actions) GetBalance(ctx context.Context, address string, blockNumber *big.Int) (*big.Int, error) {
	var result hexutil.Big
	err := w.C.Client().CallContext(ctx, &result, "eth_getBalance", accounts.HexToAddress(address), toBlockNumArg(blockNumber))
	return (*big.Int)(&result), err
}

func (w *Web3Actions) GetCode(ctx context.Context, address string, blockNumber *big.Int) ([]byte, error) {
	var result hexutil.Bytes
	err := w.C.Client().CallContext(ctx, &result, "eth_getCode", accounts.HexToAddress(address), toBlockNumArg(blockNumber))
	if err != nil {
		zlog.Err(err).Msg("GetCode: CallContext")
		return result, err
	}
	return result, err
}

func (w *Web3Actions) GetTxPoolContent(ctx context.Context) (map[string]map[string]map[string]*types.Transaction, error) {
	var txPool map[string]map[string]map[string]*types.Transaction
	if err := w.C.Client().CallContext(ctx, &txPool, "txpool_content"); err != nil {
		zlog.Err(err).Msg("GetTxPoolContent: CallContext")
		return nil, err
	}
	return txPool, nil
}

func toBlockNumArg(number *big.Int) string {
	if number == nil {
		return "latest"
	}
	return hexutil.EncodeBig(number)
}

func toForkingArg(jsonRpcURL string, blockNumber int) interface{} {
	arg := map[string]map[string]any{
		"forking": {
			"jsonRpcUrl":  jsonRpcURL,
			"blockNumber": blockNumber,
		},
	}
	return arg
}

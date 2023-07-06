package web3_actions

import (
	"context"
	"math/big"
	"net/url"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/google/uuid"
	zlog "github.com/rs/zerolog/log"
	"github.com/zeus-fyi/gochain/web3/accounts"
)

const (
	defaultProxyUrl    = "https://iris.zeus.fyi/v1/"
	proxyHeader        = "Proxy-Relay-To"
	SessionLockHeader  = "Session-Lock-ID"
	DurableExecutionID = "Durable-Execution-ID"
	EndSessionLockID   = "End-Session-Lock-ID"
)

type Web3Actions struct {
	C *ethclient.Client
	*accounts.Account
	Headers          map[string]string
	NodeURL          string
	RelayProxyUrl    string
	Network          string
	IsAnvilNode      bool
	DurableExecution bool
}

func (w *Web3Actions) Dial() {
	if w.Headers == nil {
		w.Headers = make(map[string]string)
	}
	nodeUrl := w.NodeURL
	if len(w.RelayProxyUrl) > 0 {
		proxyRelayUrlVal, rerr := url.ParseRequestURI(w.RelayProxyUrl)
		if rerr == nil {
			// the node becomes the destination through the proxy now
			w.Headers[proxyHeader] = nodeUrl
			nodeUrl = proxyRelayUrlVal.String()
		}
	}
	if w.DurableExecution {
		w.Headers[DurableExecutionID] = uuid.New().String()
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

func (w *Web3Actions) AddEndSessionLockHeader(sessionID string) {
	if w.Headers == nil {
		w.Headers = make(map[string]string)
	}
	w.Headers[EndSessionLockID] = sessionID
}

func (w *Web3Actions) AddEndSessionLockToHeaderIfExisting() {
	if w.Headers == nil {
		w.Headers = make(map[string]string)
	}
	if sessionID, ok := w.Headers[SessionLockHeader]; ok {
		w.Headers[EndSessionLockID] = sessionID
	} else {
		zlog.Warn().Msg("no session lock header found")
	}
}

func (w *Web3Actions) EndHardHatSessionReset(ctx context.Context, nodeURL string, blockNum int) {
	w.AddEndSessionLockToHeaderIfExisting()
	err := w.ResetNetwork(ctx, nodeURL, blockNum)
	if err != nil {
		zlog.Warn().Err(err).Msg("error resetting hardhat session")
		return
	}
}

func (w *Web3Actions) AddDurableExecutionIDHeader(reqID string) {
	if w.Headers == nil {
		w.Headers = make(map[string]string)
	}
	w.Headers[DurableExecutionID] = reqID
}

func (w *Web3Actions) AddSessionLockHeader(sessionID string) {
	if w.Headers == nil {
		w.Headers = make(map[string]string)
	}
	w.Headers[SessionLockHeader] = sessionID
}

func (w *Web3Actions) AddBearerToken(token string) {
	if w.Headers == nil {
		w.Headers = make(map[string]string)
	}
	w.Headers["Authorization"] = "Bearer " + token
}

func (w *Web3Actions) Close() {
	w.C.Close()
}

func NewWeb3ActionsClient(nodeUrl string) Web3Actions {
	return Web3Actions{
		NodeURL: nodeUrl,
	}
}

func NewWeb3ActionsClientWithDefaultRelayProxy(nodeUrl string, accounts *accounts.Account) Web3Actions {
	return Web3Actions{
		NodeURL:       nodeUrl,
		RelayProxyUrl: defaultProxyUrl,
		Account:       accounts,
	}
}

func NewWeb3ActionsClientWithRelayProxy(relayProxyUrl, nodeUrl string, accounts *accounts.Account) Web3Actions {
	return Web3Actions{
		NodeURL:       nodeUrl,
		RelayProxyUrl: relayProxyUrl,
		Account:       accounts,
	}
}

func NewWeb3ActionsClientWithAccount(nodeUrl string, account *accounts.Account) Web3Actions {
	return Web3Actions{
		NodeURL: nodeUrl,
		Account: account,
	}
}

type RpcMessage struct {
	Method string        `json:"method"`
	Id     int           `json:"id"`
	Result any           `json:"result,omitempty"`
	Params []interface{} `json:"params"`
}

func replacePrefix(input string, prefix string, replacement string) string {
	if strings.HasPrefix(input, prefix) {
		return replacement + input[len(prefix):]
	}
	return input
}

func (w *Web3Actions) swapToAnvil(method string) string {
	if w.IsAnvilNode {
		return replacePrefix(method, "hardhat_", "anvil_")
	}
	return method
}

func (w *Web3Actions) MineBlock(ctx context.Context, blocksToMine hexutil.Big) error {
	err := w.C.Client().CallContext(ctx, nil, w.swapToAnvil("hardhat_mine"), blocksToMine.String())
	return err
}

func (w *Web3Actions) GetStorageAt(ctx context.Context, addr, slot string) (hexutil.Bytes, error) {
	var result hexutil.Bytes
	err := w.C.Client().CallContext(ctx, &result, "eth_getStorageAt", addr, slot)
	return result, err
}

func (w *Web3Actions) SetStorageAt(ctx context.Context, addr, slot, value string) error {
	err := w.C.Client().CallContext(ctx, nil, w.swapToAnvil("hardhat_setStorageAt"), addr, slot, value)
	return err
}

func (w *Web3Actions) GetEVMSnapshot(ctx context.Context) (*big.Int, error) {
	var result hexutil.Big
	err := w.C.Client().CallContext(ctx, &result, "evm_snapshot")
	return (*big.Int)(&result), err
}

func (w *Web3Actions) GetNodeInfo(ctx context.Context) (interface{}, error) {
	cmdValue := "hardhat_metadata"
	if w.IsAnvilNode {
		cmdValue = "anvil_nodeInfo"
	}
	msg := RpcMessage{
		Method: cmdValue,
		Id:     1,
		Params: []interface{}{},
	}
	var result interface{}
	err := w.C.Client().CallContext(ctx, &result, msg.Method, msg.Params...)
	if err != nil {
		return nil, err
	}
	return result, err
}

func (w *Web3Actions) ResetNetwork(ctx context.Context, rpcUrl string, blockNumber int) error {
	if rpcUrl != "" && blockNumber != 0 {
		args := toForkingArg(rpcUrl, blockNumber)
		msg := RpcMessage{
			Method: w.swapToAnvil("hardhat_reset"),
			Id:     1,
			Params: []interface{}{args},
		}
		return w.C.Client().CallContext(ctx, nil, w.swapToAnvil("hardhat_reset"), msg)
	}
	return w.C.Client().CallContext(ctx, nil, w.swapToAnvil("hardhat_reset"))
}

func (w *Web3Actions) ImpersonateAccount(ctx context.Context, address string) error {
	var result any
	err := w.C.Client().CallContext(ctx, &result, w.swapToAnvil("hardhat_impersonateAccount"), accounts.HexToAddress(address))
	return err
}

func (w *Web3Actions) StopImpersonatingAccount(ctx context.Context, address string) error {
	err := w.C.Client().CallContext(ctx, nil, w.swapToAnvil("hardhat_stopImpersonatingAccount"), accounts.HexToAddress(address))
	return err
}

func (w *Web3Actions) SetNonce(ctx context.Context, address string, nonce hexutil.Big) error {
	err := w.C.Client().CallContext(ctx, nil, w.swapToAnvil("hardhat_setNonce"), accounts.HexToAddress(address), nonce.String())
	return err
}

func (w *Web3Actions) SetCode(ctx context.Context, address string, bytes string) error {
	err := w.C.Client().CallContext(ctx, nil, w.swapToAnvil("hardhat_setCode"), accounts.HexToAddress(address), bytes)
	return err
}

func (w *Web3Actions) SetBalance(ctx context.Context, address string, balance hexutil.Big) error {
	err := w.C.Client().CallContext(ctx, nil, w.swapToAnvil("hardhat_setBalance"), accounts.HexToAddress(address), balance)
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

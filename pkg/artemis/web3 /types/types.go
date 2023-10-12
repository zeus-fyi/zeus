package web3_types

import (
	"math/big"

	"github.com/zeus-fyi/gochain/web3/accounts"
)

type CallMsg struct {
	From      *accounts.Address // the sender of the 'transaction'
	To        *accounts.Address // the destination contract (nil for contract creation)
	Gas       uint64            // if 0, the call executes with near-infinite gas
	GasPrice  *big.Int          // wei <-> gas exchange ratio
	GasTipCap *big.Int
	GasFeeCap *big.Int
	Value     *big.Int // amount of wei sent along with the call
	Data      []byte   // input data, usually an ABI-encoded contract method invocation
}

type Snapshot struct {
	Number  uint64                        `json:"number"`
	Hash    accounts.Hash                 `json:"hash"`
	Signers map[accounts.Address]uint64   `json:"signers"`
	Voters  map[accounts.Address]struct{} `json:"voters"`
	Votes   []*Vote                       `json:"votes"`
	Tally   map[accounts.Address]Tally    `json:"tally"`
}

type Vote struct {
	Signer    accounts.Address `json:"signer"`
	Block     uint64           `json:"block"`
	Address   accounts.Address `json:"address"`
	Authorize bool             `json:"authorize"`
}

type Tally struct {
	Authorize bool `json:"authorize"`
	Votes     int  `json:"votes"`
}

type ID struct {
	NetworkID   *big.Int      `json:"network_id"`
	ChainID     *big.Int      `json:"chain_id"`
	GenesisHash accounts.Hash `json:"genesis_hash"`
}

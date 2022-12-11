package client_consts

// TODO others
const (
	Geth       = "geth"
	Nethermind = "nethermind"
)

var GethBeaconPorts = []string{"8545:8545"}

func IsExecClient(name string) bool {
	switch name {
	case Geth, Nethermind:
		return true
	default:
		return false
	}
}

type ExecClientSyncStatus struct {
	JsonRPC string `json:"jsonrpc"`
	Id      int    `json:"id"`
	Result  bool   `json:"result"`
}

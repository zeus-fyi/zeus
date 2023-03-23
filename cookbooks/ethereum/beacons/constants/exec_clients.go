package client_consts

// TODO others
const (
	ZeusConfigMapExecClient = "cm-exex-client"
	ZeusExecClient          = "zeus-exec-client"
	Erigon                  = "erigon"
	Geth                    = "geth"
	Nethermind              = "nethermind"
)

var GethBeaconPorts = []string{"8545:8545"}

func IsExecClient(name string) bool {
	switch name {
	case Geth, Nethermind, Erigon:
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

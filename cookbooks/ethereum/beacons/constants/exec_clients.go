package client_consts

// TODO others
const (
	Geth       = "geth"
	Nethermind = "nethermind"
)

func IsExecClient(name string) bool {
	switch name {
	case Geth, Nethermind:
		return true
	default:
		return false
	}
}

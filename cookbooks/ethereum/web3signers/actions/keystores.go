package ethereum_web3signer_actions

import (
	"context"

	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
)

type Web3SignerKeystores struct {
	Keystores          []string `json:"keystores"`
	Passwords          []string `json:"passwords"`
	SlashingProtection string   `json:"slashing_protection,omitempty"`
}

func (k *Web3SignerKeystores) ReadKeystore(ctx context.Context, p filepaths.Path) {
	//m := make(map[string]interface{})
	b := p.ReadFileInPath()

	if len(k.Keystores) <= 0 {
		k.Keystores = []string{}
	}

	k.Keystores = append(k.Keystores, string(b))
}

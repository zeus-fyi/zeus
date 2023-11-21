package ethereum_web3signer_actions

import (
	"context"
	"os"

	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	strings_filter "github.com/zeus-fyi/zeus/pkg/utils/strings"
)

type Web3SignerKeystores struct {
	Keystores          []string `json:"keystores"`
	Passwords          []string `json:"passwords"`
	SlashingProtection string   `json:"slashing_protection,omitempty"`
}

type Pubkeys []string

func (k *Web3SignerKeystores) ReadKeystoreDirAndAppendPw(ctx context.Context, p filepaths.Path, pw string) {
	f := strings_filter.FilterOpts{
		StartsWithAnyOfThese: []string{"keystore"},
	}
	p.FilterFiles = &f
	err := p.WalkAndApplyFuncToFileType(".json", k.ReadKeystoreAndAppend)
	if err != nil {
		return
	}
	k.Passwords = make([]string, len(k.Keystores))
	for i, _ := range k.Keystores {
		k.Passwords[i] = pw
	}
}

func (k *Web3SignerKeystores) ReadKeystoreAndAppend(fp string) error {
	if len(k.Keystores) <= 0 {
		k.Keystores = []string{}
	}
	b, err := os.ReadFile(fp)
	if err != nil {
		return err
	}
	k.Keystores = append(k.Keystores, string(b))
	return err
}

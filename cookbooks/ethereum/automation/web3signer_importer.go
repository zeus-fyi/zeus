package ethereum_automation_cookbook

type Web3SignerKeystores struct {
	Keystores          []string `json:"keystores"`
	Passwords          []string `json:"passwords"`
	SlashingProtection string   `json:"slashing_protection,omitempty"`
}

func ReadKeystores() Web3SignerKeystores {
	// TODO
	return Web3SignerKeystores{}
}

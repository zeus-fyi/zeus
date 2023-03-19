package ethereum_mev_cookbooks

import (
	"context"
	"fmt"
	"strings"
)

var (
	DefaultBase = []string{"-c"}
)

func GetMevBoostArgs(ctx context.Context, network string, relays RelaysEnabled) []string {
	baseArgs := DefaultBase
	baseCmd := fmt.Sprintf("/app/mev-boost -%s -json -addr=0.0.0.0:18550", strings.ToLower(network))
	relayURLs := relays.GetRelays(network)
	for _, relayURL := range relayURLs {
		baseCmd += fmt.Sprintf(" -relay=%s", relayURL)
	}
	return append(baseArgs, baseCmd)
}

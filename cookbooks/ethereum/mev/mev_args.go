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
	if relayURLs != nil {
		baseCmd += " -relays=" + strings.Join(relayURLs, ",")
		baseArgs = append(baseArgs, baseCmd)
	}
	return baseArgs
}

package ethereum_mev_cookbooks

import (
	"context"
	"strings"
)

var (
	DefaultBase = []string{"-c"}
)

func GetMevBoostArgs(ctx context.Context, network string, relays RelaysEnabled) []string {
	baseArgs := DefaultBase
	baseCmd := "/app/mev-boost"
	relayURLs := relays.GetRelays(network)
	if relayURLs != nil {
		baseCmd += " -relays=" + strings.Join(relayURLs, ",")
		baseArgs = append(baseArgs, baseCmd)
	}
	return baseArgs
}

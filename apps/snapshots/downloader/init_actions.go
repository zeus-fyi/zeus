package snapshot_init

import (
	"context"

	init_jwt "github.com/zeus-fyi/zeus/pkg/aegis/jwt"
	"github.com/zeus-fyi/zeus/pkg/utils/ephemery_reset"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
)

type WorkloadInfo struct {
	WorkloadType string // eg, validatorClient, beaconExecClient, beaconConsensusClient
	Protocol     string // eg. eth, cosmos,	etc
	Network      string // eg. mainnet, theta-testnet-001, etc
	DataDir      filepaths.Path
}

func InitWorkloadAction(ctx context.Context, w WorkloadInfo) {
	switch w.Protocol {
	case "cosmos":
		CosmosStartup(ctx, w)
	case "eth", "ethereum":
		if useDefaultToken {
			_ = init_jwt.SetTokenToDefault(Workload.DataDir, "jwt.hex", jwtToken)
		}
		switch w.Network {
		case "ephemery", "ephemeral":
			// do something
			ephemery_reset.ExtractAndDecEphemeralTestnetConfig(Workload.DataDir, clientName)
		}
	}
}

package snapshot_init

import (
	"context"
	"encoding/json"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	init_jwt "github.com/zeus-fyi/zeus/pkg/aegis/jwt"
	"github.com/zeus-fyi/zeus/pkg/utils/ephemery_reset"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	resty_base "github.com/zeus-fyi/zeus/zeus/z_client/base"
)

type WorkloadInfo struct {
	WorkloadType string // eg, validatorClient, beaconExecClient, beaconConsensusClient
	Protocol     string // eg. eth, cosmos,	etc
	Network      string // eg. mainnet, theta-testnet-001, etc
	DataDir      filepaths.Path
}

type Payload struct {
	Body echo.Map `json:"body"`
}

func InitWorkloadAction(ctx context.Context, w WorkloadInfo) {
	switch w.WorkloadType {
	case "send-payload":
		log.Info().Interface("payloadBasePath", payloadBasePath).Interface("payloadPostPath", payloadPostPath).Msg("sending payload")
		payl := w.DataDir.ReadFileInPath()
		if len(payl) <= 0 {
			panic("no payload found")
		}
		pl := &Payload{
			Body: echo.Map{},
		}
		err := json.Unmarshal(payl, &pl.Body)
		if err != nil {
			panic(err)
		}
		rb := resty_base.GetBaseRestyClient(payloadBasePath, bearer)
		resp, err := rb.R().SetBody(pl).Post(payloadPostPath)
		if err != nil {
			log.Err(err).Interface("pl", pl).Interface("resp", resp).Msg("error sending payload")
			panic(err)
		}
		if resp.StatusCode() >= 400 {
			panic(resp.Status())
		}
	}
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
	case "sui":
		SuiStartup(ctx, w)
	}
}

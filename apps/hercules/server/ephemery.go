package hercules_server

import (
	"path"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog/log"
	client_consts "github.com/zeus-fyi/zeus/cookbooks/ethereum/beacons/constants"
	"github.com/zeus-fyi/zeus/pkg/utils/ephemery_reset"
)

// TODO refactor
func EphemeryInit(network string) {
	if network == "ephemery" {
		genesisPath := dataDir.DirIn
		switch clientName {
		case client_consts.Lighthouse:
			genesisPath = path.Join(genesisPath, "/testnet")
		default:
		}
		ok, _ := ephemery_reset.Exists(path.Join(genesisPath, "/retention.vars"))
		if ok {
			kt := ephemery_reset.ExtractResetTime(path.Join(genesisPath, "/retention.vars"))
			go func(timeBeforeKill int64) {
				log.Info().Msgf("killing ephemeral infra due to genesis reset after %d seconds", timeBeforeKill)
				time.Sleep(time.Duration(timeBeforeKill) * time.Second)
				rc := resty.New()
				// assumes you have the default choreography sidecar in your namespace cluster
				_, err := rc.R().Get("http://zeus-choreography:9999/delete/pods")
				if err != nil {
					log.Err(err)
				}
			}(kt)
		}

		// TODO, for others with own setups, not sure if delay is needed, may depend on user specific possible race conditions

		time.Sleep(3 * time.Second)

	}
}

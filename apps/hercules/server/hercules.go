package hercules_server

import (
	"os"
	"path"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	hercules_router "github.com/zeus-fyi/hercules/api"
	v1_common_routes "github.com/zeus-fyi/hercules/api/v1/common"
	client_consts "github.com/zeus-fyi/zeus/cookbooks/ethereum/beacons/constants"
	aegis_inmemdbs "github.com/zeus-fyi/zeus/pkg/aegis/inmemdbs"
	"github.com/zeus-fyi/zeus/pkg/utils/ephemery_reset"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
)

var (
	cfg        = Config{}
	clientName string
	network    string
	env        string
	dataDir    filepaths.Path
)

func Hercules() {
	cfg.Host = "0.0.0.0"
	srv := NewHerculesServer(cfg)
	log.Info().Msg("Hercules: server starting")

	// Request a bucket url from us, or use your own source and add here
	v1_common_routes.CommonManager.Path = dataDir
	srv.E = hercules_router.Routes(srv.E)

	aegis_inmemdbs.InitValidatorDB()

	// TODO refactor
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
				os.Exit(1)
			}(kt)
		}
	}
	srv.Start()
}

func init() {
	viper.AutomaticEnv()
	Cmd.Flags().StringVar(&cfg.Port, "port", "9003", "server port")
	Cmd.Flags().StringVar(&dataDir.DirIn, "dataDirIn", "/data", "data directory location")
	Cmd.Flags().StringVar(&env, "env", "local", "environment")
	Cmd.Flags().StringVar(&network, "network", "", "network")
	Cmd.Flags().StringVar(&clientName, "clientName", "", "client name")
}

// Cmd represents the base command when called without any subcommands
var Cmd = &cobra.Command{
	Use:   "Web3 Middleware & Infra Management",
	Short: "A web3 middleware and infra manager",
	Run: func(cmd *cobra.Command, args []string) {
		Hercules()
	},
}

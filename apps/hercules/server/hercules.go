package hercules_server

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	hercules_router "github.com/zeus-fyi/hercules/api"
	aegis_inmemdbs "github.com/zeus-fyi/zeus/pkg/aegis/inmemdbs"
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

	srv.E = hercules_router.Routes(srv.E)
	aegis_inmemdbs.InitValidatorDB()
	StartAndConfigClientNetworkSettings(clientName, network)
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

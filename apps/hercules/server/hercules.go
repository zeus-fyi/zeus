package hercules_server

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	hercules_router "github.com/zeus-fyi/hercules/api"
	v1_common_routes "github.com/zeus-fyi/hercules/api/v1/common"
	aegis_inmemdbs "github.com/zeus-fyi/zeus/pkg/aegis/inmemdbs"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
)

var cfg = Config{}
var env string
var dataDir filepaths.Path
var bucketURL string

func Hercules() {
	cfg.Host = "0.0.0.0"
	srv := NewHerculesServer(cfg)
	log.Info().Msg("Hercules: server starting")

	// Request a bucket url from us, or use your own source and add here
	v1_common_routes.CommonManager.BucketURL = bucketURL
	v1_common_routes.CommonManager.Path = dataDir
	srv.E = hercules_router.Routes(srv.E, dataDir)

	aegis_inmemdbs.InitValidatorDB()
	srv.Start()
}

func init() {
	viper.AutomaticEnv()
	Cmd.Flags().StringVar(&cfg.Port, "port", "9003", "server port")
	Cmd.Flags().StringVar(&dataDir.DirIn, "dataDirIn", "/data", "data directory location")
	Cmd.Flags().StringVar(&env, "env", "local", "environment")
	// uses a default token for demo, set your own jwt for production usage if desired

	// you can contact us for a free bucket url to download blockchain data to bootstart your node
	Cmd.Flags().StringVar(&bucketURL, "bucketUrl", "", "presigned bucket url for downloading chain snapshot data")
}

// Cmd represents the base command when called without any subcommands
var Cmd = &cobra.Command{
	Use:   "Web3 Middleware & Infra Management",
	Short: "A web3 middleware and infra manager",
	Run: func(cmd *cobra.Command, args []string) {
		Hercules()
	},
}

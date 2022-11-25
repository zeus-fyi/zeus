package hercules_server

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	v1_common_routes "github.com/zeus-fyi/hercules/api/v1/common"
	hercules_jwt "github.com/zeus-fyi/hercules/pkg/jwt"

	hercules_router "github.com/zeus-fyi/hercules/api"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
)

var cfg = Config{}
var env string
var dataDir filepaths.Path
var jwtToken string
var useDefaultToken bool
var bucketURL string

func Hercules() {
	if useDefaultToken {
		_ = hercules_jwt.SetTokenToDefault(dataDir, "jwt.hex", jwtToken)
	}
	cfg.Host = "0.0.0.0"
	srv := NewHerculesServer(cfg)
	log.Info().Msg("Hercules: server starting")

	// Request a bucket url from us, or use your own source and add here
	v1_common_routes.CommonManager.BucketURL = bucketURL
	v1_common_routes.CommonManager.Path = dataDir
	srv.E = hercules_router.Routes(srv.E, dataDir)
	srv.Start()
}

func init() {
	viper.AutomaticEnv()
	Cmd.Flags().StringVar(&cfg.Port, "port", "9003", "server port")
	Cmd.Flags().StringVar(&dataDir.DirIn, "dataDirIn", "/data", "data directory location")
	Cmd.Flags().StringVar(&env, "env", "local", "environment")
	// uses a default token for demo, set your own jwt for production usage if desired
	Cmd.Flags().StringVar(&jwtToken, "jwt", "0x6ad1acdc50a4141e518161ab2fe2bf6294de4b4d48bf3582f22cae8113f0cadc", "set jwt in datadir")
	Cmd.Flags().BoolVar(&useDefaultToken, "useDefaultToken", true, "use default jwt token")

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

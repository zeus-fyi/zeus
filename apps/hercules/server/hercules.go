package hercules_server

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	hercules_jwt "github.com/zeus-fyi/hercules/pkg/jwt"

	hercules_router "github.com/zeus-fyi/hercules/api"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
)

var cfg = Config{}
var env string
var dataDir filepaths.Path
var jwtToken string
var useDefaultToken bool

func Hercules() {
	if useDefaultToken {
		_ = hercules_jwt.CheckIfJwtTokenExistsElseWriteDefault(dataDir, "jwt.hex", jwtToken)
	}
	cfg.Host = "0.0.0.0"
	srv := NewHerculesServer(cfg)
	log.Info().Msg("Hercules: server starting")
	srv.E = hercules_router.Routes(srv.E, dataDir)
	srv.Start()
}

func init() {
	viper.AutomaticEnv()
	Cmd.Flags().StringVar(&cfg.Port, "port", "9003", "server port")
	Cmd.Flags().StringVar(&dataDir.DirIn, "dataDirIn", "/data", "data directory location")
	Cmd.Flags().StringVar(&env, "env", "local", "environment")
	// uses a default token for demo, set your own jwt for production usage if desired
	Cmd.Flags().StringVar(&jwtToken, "jwt", "fb495bff094bbeab0481e40608bf5a8e90737bd2c88405e537505e9a17795918", "set jwt in datadir")
	Cmd.Flags().BoolVar(&useDefaultToken, "useDefaultToken", true, "use default jwt token")
}

// Cmd represents the base command when called without any subcommands
var Cmd = &cobra.Command{
	Use:   "Web3 Middleware & Infra Management",
	Short: "A web3 middleware and infra manager",
	Run: func(cmd *cobra.Command, args []string) {
		Hercules()
	},
}

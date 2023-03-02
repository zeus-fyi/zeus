package hades_server

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	hades_api "github.com/zeus-fyi/hades/api"
	hades_core "github.com/zeus-fyi/zeus/pkg/hades/core"
)

var (
	cfg = Config{
		Host:  "",
		Port:  "",
		Name:  "",
		Hades: hades_core.Hades{},
	}
	env string
)

func Hades() {
	cfg.Host = "0.0.0.0"
	srv := NewHadesServer(cfg)
	ctx := context.Background()
	switch env {
	case "production":
		// Add your production k8s config here
	case "local":
		cfg.Hades.ConnectToK8s()
	}
	srv.E = hades_api.Routes(srv.E, srv.Hades)
	// Start server
	log.Ctx(ctx).Info().Msgf("Hades: %s server starting", env)
	srv.Start()
}

func init() {
	viper.AutomaticEnv()
	Cmd.Flags().StringVar(&cfg.Port, "port", "8888", "server port")
	Cmd.Flags().StringVar(&env, "env", "local", "environment")
}

// Cmd represents the base command when called without any subcommands
var Cmd = &cobra.Command{
	Use:   "Hades kubernetes automation server",
	Short: "Hades kubernetes automation server",
	Run: func(cmd *cobra.Command, args []string) {
		Hades()
	},
}

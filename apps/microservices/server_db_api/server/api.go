package echo_server_template

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfg = Config{}

func Api() {
	cfg.Host = "0.0.0.0"
	srv := NewEchoServerTemplate(cfg)
	// Echo instance
	srv.E = v1_echo_server_template.Routes(srv.E)
	// Start server
	srv.Start()
}

func init() {
	viper.AutomaticEnv()
	Cmd.Flags().StringVar(&cfg.Port, "port", "9090", "server port")
}

// Cmd represents the base command when called without any subcommands
var Cmd = &cobra.Command{
	Use:   "Use as a base skeleton for generating microservices",
	Short: "Base echo server template",
	Run: func(cmd *cobra.Command, args []string) {
		Api()
	},
}

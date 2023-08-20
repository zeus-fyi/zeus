package load_simulator

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	v1_load_simulator "github.com/zeus-fyi/zeus/load-simulator/api/v1"
)

var cfg = Config{}

func Api() {
	cfg.Host = "0.0.0.0"
	srv := NewEchoServerTemplate(cfg)
	// Echo instance
	srv.E = v1_load_simulator.Routes(srv.E)
	// Start server
	srv.Start()
}

func init() {
	viper.AutomaticEnv()
	Cmd.Flags().StringVar(&cfg.Port, "port", "8888", "server port")
}

// Cmd represents the base command when called without any subcommands
var Cmd = &cobra.Command{
	Use:   "Use to simulate load on a server",
	Short: "Base echo load simulation template",
	Run: func(cmd *cobra.Command, args []string) {
		Api()
	},
}

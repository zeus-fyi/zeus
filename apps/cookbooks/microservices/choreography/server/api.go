package choreography

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	v1_choreography "github.com/zeus-fyi/zeus/choreography/api/v1"
	zeus_client "github.com/zeus-fyi/zeus/zeus/client"
	pods_client "github.com/zeus-fyi/zeus/zeus/client/workloads/pods"
)

var (
	cfg    = Config{}
	bearer string
)

func Api() {
	cfg.Host = "0.0.0.0"
	srv := NewChoreography(cfg)
	v1_choreography.PodsClient = pods_client.NewPodsClientFromZeusClient(zeus_client.NewDefaultZeusClient(bearer))
	srv.E = v1_choreography.Routes(srv.E)
	srv.Start()
}

func init() {
	viper.AutomaticEnv()
	Cmd.Flags().StringVar(&cfg.Port, "port", "9999", "server port")
	// injected values
	Cmd.Flags().StringVar(&bearer, "bearer", "", "bearer")
	Cmd.Flags().StringVar(&v1_choreography.CloudCtxNs.CloudProvider, "cloud-provider", "", "cloud-provider")
	Cmd.Flags().StringVar(&v1_choreography.CloudCtxNs.Context, "ctx", "", "context")
	Cmd.Flags().StringVar(&v1_choreography.CloudCtxNs.Namespace, "ns", "", "namespace")
	Cmd.Flags().StringVar(&v1_choreography.CloudCtxNs.Region, "region", "", "region")
}

// Cmd represents the base command when called without any subcommands
var Cmd = &cobra.Command{
	Use:   "Use as a base skeleton for generating choreography actions",
	Short: "Base echo server template for choreography",
	Run: func(cmd *cobra.Command, args []string) {
		Api()
	},
}

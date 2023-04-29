package snapshot_init

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	preSignedURL    string
	onlyIfEmptyDir  bool
	compressionType string
	clientName      string
	jwtToken        string
	useDefaultToken bool
	Workload        WorkloadInfo
)

func StartUp() {
	ctx := context.Background()
	log.Ctx(ctx).Info().Interface("workload", Workload).Msg("Downloader: WorkloadInfo")
	InitWorkloadAction(ctx, Workload)
	ChainDownload(ctx, Workload)
}

func init() {
	viper.AutomaticEnv()
	Cmd.Flags().StringVar(&Workload.DataDir.DirIn, "dataDir", "/data", "data directory location")
	Cmd.Flags().StringVar(&Workload.WorkloadType, "workload-type", "", "workloadType") // eg validatorClient
	Cmd.Flags().StringVar(&Workload.Network, "network", "", "network")                 // eg mainnet, testnet
	Cmd.Flags().StringVar(&Workload.Protocol, "protocol", "", "protocol")              // eg eth, cosmos, etc

	Cmd.Flags().StringVar(&preSignedURL, "downloadURL", "", "use a presigned bucket url")
	Cmd.Flags().BoolVar(&onlyIfEmptyDir, "onlyIfEmptyDir", true, "only download & extract if the datadir is empty")
	Cmd.Flags().StringVar(&compressionType, "compressionExtension", ".tar.lz4", "compression type")
	Cmd.Flags().StringVar(&clientName, "clientName", "", "client name")
	Cmd.Flags().StringVar(&jwtToken, "jwt", "0x6ad1acdc50a4141e518161ab2fe2bf6294de4b4d48bf3582f22cae8113f0cadc", "set jwt in datadir")
	Cmd.Flags().BoolVar(&useDefaultToken, "useDefaultToken", true, "use default jwt token")
}

// Cmd represents the base command when called without any subcommands
var Cmd = &cobra.Command{
	Use:   "Downloads and extracts blockchain data and configs to your dataDir",
	Short: "Blockchain node data download procedure",
	Run: func(cmd *cobra.Command, args []string) {
		StartUp()
	},
}

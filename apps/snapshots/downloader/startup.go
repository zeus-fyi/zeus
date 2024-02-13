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
	bearer          string
	jwtToken        string
	payloadPostPath string
	payloadBasePath string
	useDefaultToken bool
	Workload        WorkloadInfo
)

func StartUp() {
	log.Info().Msg("Snapshots: starting")
	ctx := context.Background()
	log.Info().Interface("workload", Workload).Msg("Downloader: WorkloadInfo")

	log.Info().Msg("InitWorkloadAction: starting")
	InitWorkloadAction(ctx, Workload)
	log.Info().Msg("InitWorkloadAction: done")
	log.Info().Msg("Download: starting")
	Download(ctx, Workload)
	log.Info().Msg("Download: done")
}

func init() {
	viper.AutomaticEnv()
	Cmd.Flags().StringVar(&Workload.DataDir.DirIn, "dataDir", "/data", "data directory location")
	Cmd.Flags().StringVar(&Workload.WorkloadType, "workload-type", "", "workloadType") // eg validatorClient
	Cmd.Flags().StringVar(&Workload.Network, "network", "", "network")                 // eg mainnet, testnet
	Cmd.Flags().StringVar(&Workload.Protocol, "protocol", "", "protocol")              // eg eth, cosmos, etc
	Cmd.Flags().StringVar(&bearer, "bearer", "", "bearer")
	Cmd.Flags().StringVar(&payloadPostPath, "payload-post-path", "", "payload post path")
	Cmd.Flags().StringVar(&payloadBasePath, "payload-base-path", "", "payload base path")

	Cmd.Flags().StringVar(&Workload.DataDir.FnIn, "fi", "", "file input name")

	Cmd.Flags().StringVar(&preSignedURL, "downloadURL", "", "use a presigned bucket url")
	Cmd.Flags().BoolVar(&onlyIfEmptyDir, "onlyIfEmptyDir", true, "only download & extract if the datadir is empty")
	Cmd.Flags().StringVar(&compressionType, "compressionExtension", ".tar.lz4", "compression type")
	Cmd.Flags().StringVar(&clientName, "clientName", "", "client name")
	Cmd.Flags().StringVar(&jwtToken, "jwt", "0x6ad1acdc50a4141e518161ab2fe2bf6294de4b4d48bf3582f22cae8113f0cadc", "set jwt in datadir")
	Cmd.Flags().BoolVar(&useDefaultToken, "useDefaultToken", true, "use default jwt token")
}

// Cmd represents the base command when called without any subcommands
var Cmd = &cobra.Command{
	Use:   "Downloads and extracts data or uploads it, eg configs to your dataDir",
	Short: "Init container data download/upload procedures",
	Run: func(cmd *cobra.Command, args []string) {
		StartUp()
	},
}

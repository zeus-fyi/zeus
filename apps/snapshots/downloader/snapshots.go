package snapshot_init

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zeus-fyi/snapshots/config_fetching"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
)

var (
	dataDir         filepaths.Path
	preSignedURL    string
	onlyIfEmptyDir  bool
	compressionType string
	clientName      string
)

func StartUp() {
	ChainDownload()
	// the below uses a switch case to download if an ephemeralClientName is used
	config_fetching.ExtractAndDecEphemeralTestnetConfig(dataDir, clientName)
}

func init() {
	viper.AutomaticEnv()
	Cmd.Flags().StringVar(&dataDir.DirIn, "dataDir", "/data", "data directory location")
	Cmd.Flags().StringVar(&preSignedURL, "downloadURL", "", "use a presigned bucket url")
	Cmd.Flags().BoolVar(&onlyIfEmptyDir, "onlyIfEmptyDir", true, "only download & extract if the datadir is empty")
	Cmd.Flags().StringVar(&compressionType, "compressionExtension", ".tar.lz4", "compression type")
	Cmd.Flags().StringVar(&clientName, "clientName", "geth", "client name")
}

// Cmd represents the base command when called without any subcommands
var Cmd = &cobra.Command{
	Use:   "Downloads and extracts blockchain data and configs to your dataDir",
	Short: "Blockchain node data download procedure",
	Run: func(cmd *cobra.Command, args []string) {
		StartUp()
	},
}

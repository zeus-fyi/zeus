package snapshot_init

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

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
	Use:   "Downloads and extracts blockchain data to your dataDir",
	Short: "Blockchain node data download procedure",
	Run: func(cmd *cobra.Command, args []string) {
		ChainDownload()
	},
}

package snapshot_init

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zeus-fyi/snapshots/config_fetching"
	init_jwt "github.com/zeus-fyi/zeus/pkg/aegis/jwt"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
)

var (
	dataDir         filepaths.Path
	preSignedURL    string
	onlyIfEmptyDir  bool
	compressionType string
	clientName      string
	jwtToken        string
	useDefaultToken bool
)

func StartUp() {
	if useDefaultToken {
		_ = init_jwt.SetTokenToDefault(dataDir, "jwt.hex", jwtToken)
	}
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

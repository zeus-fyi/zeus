package config_fetching

import (
	"context"
	"path"

	"github.com/rs/zerolog/log"
	beacon_cookbooks "github.com/zeus-fyi/zeus/cookbooks/ethereum/beacons"
	"github.com/zeus-fyi/zeus/pkg/poseidon"
	"github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/compression"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
)

var dataDir filepaths.Path

func ExtractAndDecEphemeralTestnetConfig(clientName string) {
	switch clientName {
	case beacon_cookbooks.LighthouseEphemeral:
		dataDir.DirIn = path.Join(dataDir.DirIn, "/configs/testnet")
		dataDir.DirOut = dataDir.DirIn
	case beacon_cookbooks.GethEphemeral:
		// placing a genesis.json file directly in the datadir path should set the chain to the expected value
		dataDir.DirOut = dataDir.DirIn
	case "test":
		dataDir.DirIn = "."
		dataDir.DirOut = "configs"
	default:
		return
	}
	ctx := context.Background()
	url := GetLatestReleaseConfigDownloadURL()
	dataDir.FnIn = ephemeralTestnetFile
	err := poseidon.DownloadFile(ctx, dataDir.DirIn, url)
	if err != nil {
		log.Ctx(ctx).Panic().Err(err).Msg("DownloadFile")
	}
	dec := compression.NewCompression()
	err = dec.UnGzip(&dataDir)
	if err != nil {
		log.Ctx(ctx).Panic().Err(err).Msg("UnGzip")
	}
	// cleans up, by deleting the compressed file
	err = dataDir.RemoveFileInPath()
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("RemoveFileInPath")
	}
}

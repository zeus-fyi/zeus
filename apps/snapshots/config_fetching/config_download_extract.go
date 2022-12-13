package config_fetching

import (
	"context"
	"path"

	"github.com/rs/zerolog/log"
	"github.com/zeus-fyi/zeus/pkg/poseidon"
	"github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/compression"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
)

var dataDir filepaths.Path

func ExtractAndDecEphemeralTestnetConfig() {
	ctx := context.Background()
	url := GetLatestReleaseConfigDownloadURL()
	dataDir.DirOut = path.Join(dataDir.DirIn, "/configs/testnet")
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

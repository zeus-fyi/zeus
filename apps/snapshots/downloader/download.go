package snapshot_init

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/zeus-fyi/zeus/pkg/poseidon"
	"github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/compression"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	"github.com/zeus-fyi/zeus/pkg/utils/host_info"
)

var dataDir filepaths.Path
var preSignedURL string
var onlyIfEmptyDir bool
var compressionType string
var clientName string

func ChainDownload() {
	ctx := context.Background()
	if len(preSignedURL) == 0 {
		log.Ctx(ctx).Info().Msg("No download url provided, skipping snapshot download")
		return
	}
	stats, err := host_info.GetDiskUsageStats(ctx, dataDir.DirIn)
	if err != nil {
		log.Ctx(ctx).Panic().Err(err).Msg("GetDiskUsageStats")
	}
	switch {
	// just approximates empty as <= 1% disk usage in dataDir
	case onlyIfEmptyDir && stats.UsedPercent <= float64(1):
		err = poseidon.DownloadSnapshot(ctx, dataDir.DirIn, preSignedURL)
		if err != nil {
			log.Ctx(ctx).Panic().Err(err).Msg("DownloadSnapshot")
		}

		switch compressionType {
		case ".tar.lz4":
			dec := compression.NewCompression()
			dataDir.DirOut = dataDir.DirIn
			dataDir.FnIn = clientName + compressionType
			err = dec.Lz4Decompress(&dataDir)
			if err != nil {
				log.Ctx(ctx).Panic().Err(err).Msg("Lz4Decompress")
			}

			// cleans up, by deleting the compressed file
			err = dataDir.RemoveFileInPath()
			if err != nil {
				log.Ctx(ctx).Err(err).Msg("RemoveFileInPath")
			}
		}
	}
}

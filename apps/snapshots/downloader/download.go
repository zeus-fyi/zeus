package snapshot_init

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/zeus-fyi/zeus/pkg/utils/host_info"
)

func Download(ctx context.Context, w WorkloadInfo) {
	stats, serr := host_info.GetDiskUsageStats(ctx, w.DataDir.DirIn)
	if serr != nil {
		log.Panic().Err(serr).Msg("GetDiskUsageStats")
	}
	log.Info().Msgf("disk usage stats: %+v", stats)
	log.Info().Float64("disk usage stats.UsedPercent", stats.UsedPercent)
	if onlyIfEmptyDir && stats.UsedPercent <= float64(1) {
		switch w.Protocol {
		case "eth", "ethereum":
			EthereumChainDownload(ctx, w)
		case "sui":
			log.Info().Msg("SuiDownloadSnapshotS3: starting")
			err := SuiDownloadSnapshotS3(w)
			if err != nil {
				log.Err(err).Interface("w", w).Msg("SuiDownloadSnapshotS3")
				panic(err)
			}
			log.Info().Msg("SuiDownloadSnapshotS3: done")
		}
	}
}

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
	if onlyIfEmptyDir && stats.UsedPercent <= float64(1) {
		switch w.Protocol {
		case "eth", "ethereum":
			EthereumChainDownload(ctx, w)
		case "sui":
			SuiStartup(ctx, w)
			err := SuiDownloadSnapshotS3(w)
			if err != nil {
				panic(err)
			}
		}
	}
}

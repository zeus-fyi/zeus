package snapshot_init

import "context"

func Download(ctx context.Context, w WorkloadInfo) {
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

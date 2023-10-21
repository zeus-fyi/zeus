package snapshot_init

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/compression"
	resty_base "github.com/zeus-fyi/zeus/zeus/z_client/base"
)

const poseidonEndpoint = "https://poseidon.zeus.fyi"

func EthereumChainDownload(ctx context.Context, w WorkloadInfo) {
	if w.WorkloadType == "beacon" {
		if len(bearer) <= 0 {
			return
		}
		rc := resty_base.GetBaseRestyClient(poseidonEndpoint, bearer)
		switch w.Protocol {
		case "eth", "ethereum":
			switch w.Network {
			case "mainnet":
				switch clientName {
				case "reth":
					// todo: add reth download
				case "geth":
					_, err := rc.R().SetResult(&preSignedURL).Get("/v1/ethereum/mainnet/geth")
					if err != nil {
						log.Err(err).Msg("geth preSignedURL")
					}
				case "lighthouse":
					_, err := rc.R().SetResult(&preSignedURL).Get("/v1/ethereum/mainnet/lighthouse")
					if err != nil {
						log.Err(err).Msg("lighthouse preSignedURL")
					}
				}
			}
		}
	}
	if len(preSignedURL) <= 0 {
		log.Ctx(ctx).Info().Msg("No download url provided, skipping snapshot download")
		return
	}

	switch compressionType {
	case ".tar.lz4":
		dec := compression.NewCompression()
		w.DataDir.DirOut = w.DataDir.DirIn
		w.DataDir.FnIn = clientName + compressionType
		err := dec.Lz4Decompress(&w.DataDir)
		if err != nil {
			log.Ctx(ctx).Panic().Err(err).Msg("Lz4Decompress")
		}
		// cleans up, by deleting the compressed file
		err = w.DataDir.RemoveFileInPath()
		if err != nil {
			log.Ctx(ctx).Err(err).Msg("RemoveFileInPath")
		}
	}
}

package snapshot_init

import (
	"context"
)

const (
	suiMainnetSnapshotS3 = "s3://mysten-mainnet-snapshots/"
	suiTestnetSnapshotS3 = "s3://mysten-testnet-snapshots/"
)

// TODO: finish setup for sui s3 downloads, and genesis

func SuiStartup(ctx context.Context, w WorkloadInfo) {
	switch w.Network {
	case "devnet":
		// curl -fLJO https://github.com/MystenLabs/sui-genesis/raw/main/devnet/genesis.blob
	case "mainnet":
		// curl -fLJO https://github.com/MystenLabs/sui-genesis/raw/main/mainnet/genesis.blob
	case "testnet":
		// curl -fLJO https://github.com/MystenLabs/sui-genesis/raw/main/testnet/genesis.blob
	}
}

/*

https://docs.sui.io/build/snapshot

db-checkpoint-config:
  perform-db-checkpoints-at-epoch-end: true
  perform-index-db-checkpoints-at-epoch-end: true
  object-store-config:
    object-store: "S3"
    bucket: "<BUCKET-NAME>"
    aws-access-key-id: “<ACCESS-KEY>”
    aws-secret-access-key: “<SHARED-KEY>”
    aws-region: "<BUCKET-REGION>"
    object-store-connection-limit: 20
*/

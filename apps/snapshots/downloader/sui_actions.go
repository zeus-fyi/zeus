package snapshot_init

import (
	"context"
	"fmt"
	"os/exec"
	"time"

	"github.com/cavaliergopher/grab/v3"
	"github.com/rs/zerolog/log"
)

const (
	suiMainnetSnapshotS3 = "s3://mysten-mainnet-snapshots/"
	suiTestnetSnapshotS3 = "s3://mysten-testnet-snapshots/"
)

func SuiStartup(ctx context.Context, w WorkloadInfo) {
	// mainnet default
	urlPath := "https://github.com/MystenLabs/sui-genesis/raw/main/mainnet/genesis.blob"

	w.DataDir.FnIn = "genesis.blob"
	if w.DataDir.FileInPathExists() {
		log.Info().Msg("genesis.blob already exists, skipping download")
		return
	}
	switch w.Network {
	case "devnet":
		//urlPath = "https://github.com/MystenLabs/sui-genesis/raw/main/devnet/genesis.blob"
	case "mainnet":
		urlPath = "https://github.com/MystenLabs/sui-genesis/raw/main/mainnet/genesis.blob"
	case "testnet":
		urlPath = "https://github.com/MystenLabs/sui-genesis/raw/main/testnet/genesis.blob"
	}

	err := DownloadGenesisBlob(w, urlPath)
	if err != nil {
		panic(err)
	}
}

func SuiDownloadSnapshotS3(w WorkloadInfo) error {
	switch w.WorkloadType {
	case "full":
	case "validator":
	default:
		log.Warn().Msg("SuiDownloadSnapshotS3: workload type not supported and/or provided")
		return nil
	}

	s3 := ""
	switch w.Network {
	case "mainnet":
		s3 = suiMainnetSnapshotS3
	case "testnet":
		s3 = suiTestnetSnapshotS3
	default:
		log.Warn().Msg("SuiDownloadSnapshotS3: network type not supported and/or provided")
		return nil
	}
	// Form the S3 path for the snapshot
	// Execute AWS CLI command to download the snapshot
	cmd := exec.Command(
		"aws",
		"s3",
		"cp",
		s3,
		w.DataDir.DirIn,
		"--recursive",
		"--no-sign-request",
	)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error downloading snapshot from S3: %v", err)
	}

	// Update ownership of the directory
	chownCmd := exec.Command("sudo", "chown", "-R", "sui:sui", w.DataDir.DirIn)
	err = chownCmd.Run()
	if err != nil {
		err = fmt.Errorf("error changing ownership of the directory: %v", err)
		log.Err(err).Msg("DownloadSnapshotS3")
		return err
	}
	return nil
}

func DownloadGenesisBlob(w WorkloadInfo, blobURL string) error {
	// download procedure
	client := grab.NewClient()
	// Downloads to your datadir
	req, err := grab.NewRequest(w.DataDir.DirIn, blobURL)
	if err != nil {
		log.Err(err).Msgf("DownloadChainSnapshotRequest: NewRequest, %s", blobURL)
		return err
	}
	// start download
	fmt.Printf("Downloading %v...\n", req.URL())
	resp := client.Do(req)
	fmt.Printf("  %v\n", resp.HTTPResponse.Status)

	// set to any desired max time
	timer := time.NewTicker(5 * time.Minute)
	defer timer.Stop()
	select {
	case <-timer.C:
		fmt.Printf("  transferred %v / %v bytes (%.2f%%)\n",
			resp.BytesComplete(),
			resp.Size(),
			100*resp.Progress())
		return nil
	case <-resp.Done:
		// download is complete
		err = resp.Err()
		if err != nil {
			log.Err(err).Msg("DownloadChainSnapshotRequest")
			return err
		}
	}
	return nil
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

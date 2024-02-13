package snapshot_init

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"path"
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
	// Check if the directory exists, if not create it
	if _, err := os.Stat(w.DataDir.DirIn); os.IsNotExist(err) {
		log.Info().Msg("Directory does not exist, creating it")
		if err = os.MkdirAll(w.DataDir.DirIn, 0755); err != nil { // 0755 is commonly used permissions
			log.Fatal().Err(err).Msg("Failed to create directory")
			return
		}
	}

	switch w.Network {
	case "devnet":
		urlPath = "https://github.com/MystenLabs/sui-genesis/raw/main/devnet/genesis.blob"
	case "testnet":
		urlPath = "https://github.com/MystenLabs/sui-genesis/raw/main/testnet/genesis.blob"
	case "mainnet":
		urlPath = "https://github.com/MystenLabs/sui-genesis/raw/main/mainnet/genesis.blob"
	}

	err := DownloadGenesisBlob(w, urlPath)
	if err != nil {
		log.Warn().Err(err).Interface("url", urlPath).Interface("w", w).Msg("SuiStartup: error downloading genesis.blob")
		panic(err)
	}
}

func SuiDownloadSnapshotS3(w WorkloadInfo) error {
	dbPathExt := w.DataDir.DirIn
	switch w.WorkloadType {
	case "full":
		dbPathExt = path.Join(w.DataDir.DirIn, "full_node_db/live")
	case "validator":
		dbPathExt = path.Join(w.DataDir.DirIn, "live")
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
	case "devnet":
		return nil
	default:
		log.Warn().Msg("SuiDownloadSnapshotS3: network type not supported and/or provided")
		return nil
	}

	log.Info().Msgf("SuiDownloadSnapshotS3: downloading snapshot from %s", s3)
	// Form the S3 path for the snapshot
	// Execute AWS CLI command to download the snapshot

	// Capture stdout and stderr
	var out bytes.Buffer
	var stderr bytes.Buffer

	cmd := exec.Command(
		"aws",
		"s3",
		"cp",
		s3,
		dbPathExt,
		"--recursive",
		"--no-sign-request",
	)
	log.Info().Msgf("SuiDownloadSnapshotS3: downloading snapshot using aws cli cmd:", cmd.String())

	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Warn().Err(err).Str("stdout", out.String()).Str("stderr", stderr.String()).Msg("error downloading snapshot from S3")
		log.Err(err).Str("stdout", out.String()).Str("stderr", stderr.String()).Msg("error downloading snapshot from S3")
		return fmt.Errorf("error downloading snapshot from S3: %v", err)
	}
	return nil
}

func DownloadGenesisBlob(w WorkloadInfo, blobURL string) error {
	// download procedure
	client := grab.NewClient()
	// Downloads to your datadir
	req, err := grab.NewRequest(w.DataDir.DirIn, blobURL)
	if err != nil {
		log.Err(err).Interface("w", w).Msgf("DownloadGenesisBlob: NewRequest, %s", blobURL)
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
			log.Err(err).Interface("w", w).Msg("DownloadGenesisBlob")
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

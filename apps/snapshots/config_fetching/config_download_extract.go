package config_fetching

import (
	"context"
	"os"
	"os/exec"
	"path"

	"github.com/rs/zerolog/log"
	beacon_cookbooks "github.com/zeus-fyi/zeus/cookbooks/ethereum/beacons"
	validator_cookbooks "github.com/zeus-fyi/zeus/cookbooks/ethereum/validators"
	"github.com/zeus-fyi/zeus/pkg/poseidon"
	"github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/compression"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
)

func ExtractAndDecEphemeralTestnetConfig(dataDir filepaths.Path, clientName string) {
	switch clientName {
	case beacon_cookbooks.LighthouseEphemeral, validator_cookbooks.ValidatorClusterClassName:
		dataDir.DirIn = "/data/testnet"
		dataDir.DirOut = path.Join(dataDir.DirIn)
		log.Info().Interface("dataDir", dataDir).Msg("ExtractAndDecEphemeralTestnetConfig: LighthouseEphemeral")
	case beacon_cookbooks.GethEphemeral:
		// placing a genesis.json file directly in the datadir path should set the chain to the expected value
		dataDir.DirOut = dataDir.DirIn
		log.Info().Interface("dataDir", dataDir).Msg("ExtractAndDecEphemeralTestnetConfig: GethEphemeral")
	case "test":
		dataDir.DirIn = "."
		dataDir.DirOut = "./testnet"
	default:
		return
	}
	ctx := context.Background()
	log.Info().Interface("dataDir.DirOut", dataDir.DirOut)
	if _, zerr := os.Stat(dataDir.DirOut); os.IsNotExist(zerr) {
		_ = os.MkdirAll(dataDir.DirOut, 0700) // Create your dir
	}
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

	if clientName == beacon_cookbooks.GethEphemeral {
		cmd := exec.Command("geth", "--datadir", dataDir.DirIn, "init", path.Join(dataDir.DirIn, "genesis.json"))
		err = cmd.Run()
		if err != nil {
			log.Ctx(ctx).Panic().Err(err).Msg("setting geth genesis")
		}
	}
}

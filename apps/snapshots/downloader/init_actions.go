package snapshot_init

import (
	"context"
	"io"
	"os"
	"os/exec"

	"github.com/rs/zerolog/log"
	init_jwt "github.com/zeus-fyi/zeus/pkg/aegis/jwt"
	"github.com/zeus-fyi/zeus/pkg/utils/ephemery_reset"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
)

const (
	comsosGenesisFilePath = "/etc/gaiad/genesis.json"
)

type WorkloadInfo struct {
	WorkloadType string // eg, validatorClient, beaconExecClient, beaconConsensusClient
	Protocol     string // eg. eth, cosmos,	etc
	Network      string // eg. mainnet, theta-testnet-001, etc
	DataDir      filepaths.Path
}

func InitWorkloadAction(ctx context.Context, w WorkloadInfo) {
	switch w.Protocol {
	case "cosmos":
		switch w.Network {
		case "theta-testnet-001":
			p := filepaths.Path{
				DirIn: "/config",
				FnIn:  "genesis.json",
			}
			ok := p.FileInPathExists()
			if !ok {
				chainID := "theta-testnet-001"
				cmd := exec.Command("gaiad", "--home", "/", "--chain-id", chainID, "init", "public-testnet")
				err := cmd.Run()
				if err != nil {
					log.Ctx(ctx).Panic().Err(err).Msg("setting cosmos testnet genesis")
					panic(err)
				}
				err = p.RemoveFileInPath()
				if err != nil {
					panic(err)
				}
				sourceFile, err := os.Open(comsosGenesisFilePath)
				if err != nil {
					panic(err)
				}
				defer sourceFile.Close()
				destFile, err := os.Create(p.FileInPath())
				if err != nil {
					panic(err)
				}
				defer destFile.Close()
				_, err = io.Copy(destFile, sourceFile)
				if err != nil {
					panic(err)
				}
				err = os.Chmod(p.FileInPath(), 0644)
				if err != nil {
					panic(err)
				}
			}
		}
	case "eth":
		if useDefaultToken {
			_ = init_jwt.SetTokenToDefault(Workload.DataDir, "jwt.hex", jwtToken)
		}
		switch w.Network {
		case "ephemery":
			// do something
			ephemery_reset.ExtractAndDecEphemeralTestnetConfig(Workload.DataDir, clientName)
		}
	}
}

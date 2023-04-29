package snapshot_init

import (
	"bytes"
	"context"
	"io"
	"os"
	"os/exec"

	"github.com/pelletier/go-toml"
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
				log.Ctx(ctx).Info().Msg("init cosmos testnet genesis")
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
				seeds := "639d50339d7045436c756a042906b9a69970913f@seed-01.theta-testnet.polypore.xyz:26656,3e506472683ceb7ed75c1578d092c79785c27857@seed-02.theta-testnet.polypore.xyz:26656"
				err = cosmosTomlOverride("/config/config.toml", "seeds", seeds)
				if err != nil {
					panic(err)
				}
				err = cosmosTomlOverride("/config/app.toml", "minimum-gas-prices", "0.0025uatom")
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

func cosmosTomlOverride(filename, key, newValue string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	tree, err := toml.LoadBytes(data)
	if err != nil {
		return err
	}
	tree.Set(key, newValue)
	var buffer bytes.Buffer
	encoder := toml.NewEncoder(&buffer)
	err = encoder.Encode(tree)
	if err != nil {
		return err
	}
	err = os.WriteFile(filename, buffer.Bytes(), 0644)
	if err != nil {
		return err
	}
	return nil
}

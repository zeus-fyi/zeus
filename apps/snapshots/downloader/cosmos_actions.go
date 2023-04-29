package snapshot_init

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"regexp"
	"strconv"

	"github.com/go-resty/resty/v2"
	"github.com/pelletier/go-toml"
	"github.com/rs/zerolog/log"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
)

const (
	comsosGenesisFilePath     = "/etc/gaiad/genesis.json"
	cosmosTestnetStateSyncRPC = "https://rpc.state-sync-01.theta-testnet.polypore.xyz:443,https://rpc.state-sync-02.theta-testnet.polypore.xyz:443"
)

func CosmosStartup(ctx context.Context, w WorkloadInfo) {
	switch w.Network {
	case "theta-testnet-001":
		p := filepaths.Path{
			DirIn: "/config",
			FnIn:  "genesis.json",
		}
		ok := p.FileInPathExists()
		if !ok {
			log.Info().Msg("init cosmos testnet genesis")
			chainID := "theta-testnet-001"
			cmd := exec.Command("gaiad", "config", "chain-id", chainID, "--home", "/")
			err := cmd.Run()
			if err != nil {
				log.Ctx(ctx).Panic().Err(err).Msg("setting cosmos chain-id")
				panic(err)
			}
			cmd = exec.Command("gaiad", "config", "keyring-backend", "test", "--home", "/")
			err = cmd.Run()
			if err != nil {
				log.Ctx(ctx).Panic().Err(err).Msg("setting cosmos chain-id")
				panic(err)
			}
			cmd = exec.Command("gaiad", "config", "keyring-backend", "test", "--home", "/")
			err = cmd.Run()
			if err != nil {
				log.Ctx(ctx).Panic().Err(err).Msg("setting cosmos chain-id")
				panic(err)
			}
			cmd = exec.Command("gaiad", "config", "broadcast-mode", "block", "--home", "/")
			err = cmd.Run()
			if err != nil {
				log.Ctx(ctx).Panic().Err(err).Msg("setting cosmos chain-id")
				panic(err)
			}
			cmd = exec.Command("gaiad", "init", "public-testnet", "--home", "/", "--chain-id", chainID)
			err = cmd.Run()
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
			err = cosmosTomlOverride("/config/app.toml", "minimum-gas-prices", "0.0025uatom")
			if err != nil {
				panic(err)
			}
			si := GetStateSyncInfoTestnet(ctx)
			err = applyCosmosTomlOverridesByRegex(si, "", cosmosTestnetStateSyncRPC)
			if err != nil {
				panic(err)
			}
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

type StateSyncInfo struct {
	TrustHash   string `json:"trustHash"`
	TrustHeight string `json:"trustHeight"`
}

func GetStateSyncInfo(ctx context.Context, nodeURL string) StateSyncInfo {
	r := resty.New()
	resp, err := r.R().Get(fmt.Sprintf("%s/block", nodeURL))
	if err != nil {
		log.Ctx(ctx).Panic().Err(err).Msg("failed to get cosmos block info")
		panic(err)
	}
	var result map[string]interface{}
	if err = json.Unmarshal(resp.Body(), &result); err != nil {
		log.Ctx(ctx).Panic().Err(err).Msg("unmarshal cosmos block info")
		panic(err)
	}
	height := result["result"].(map[string]interface{})["block"].(map[string]interface{})["header"].(map[string]interface{})["height"]
	heightInt, err := strconv.Atoi(height.(string))
	if err = json.Unmarshal(resp.Body(), &result); err != nil {
		log.Ctx(ctx).Panic().Err(err).Msg("strconv cosmos block info")
		panic(err)
	}
	trustHeight := heightInt - 1000
	resp, err = r.R().SetQueryParam("height", fmt.Sprintf("%d", trustHeight)).Get(fmt.Sprintf("%s/block", nodeURL))
	if err != nil {
		log.Ctx(ctx).Panic().Err(err).Msg("failed to get cosmos block info")
		panic(err)
	}
	var trustBlock map[string]interface{}
	if err = json.Unmarshal(resp.Body(), &trustBlock); err != nil {
		log.Ctx(ctx).Panic().Err(err).Msg("unmarshal cosmos block info")
		panic(err)
	}
	trustHash := trustBlock["result"].(map[string]interface{})["block_id"].(map[string]interface{})["hash"].(string)
	si := StateSyncInfo{
		TrustHash:   trustHash,
		TrustHeight: fmt.Sprintf("%d", trustHeight),
	}
	return si
}

func GetStateSyncInfoTestnet(ctx context.Context) StateSyncInfo {
	testRPC := "https://rpc.sentry-01.theta-testnet.polypore.xyz"
	si := GetStateSyncInfo(ctx, testRPC)
	return si
}

func cosmosTomlOverrideRegex(filename, key, newValue string, useQuotes bool) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	keyRegex := regexp.MustCompile(`(` + key + `\s*=\s*)(\S.*)`)
	newValueString := newValue
	if useQuotes {
		newValueString = `"` + newValueString + `"`
	}
	replaced := keyRegex.ReplaceAllString(string(data), `$1`+newValueString)

	if replaced == string(data) {
		panic(err)
	}

	err = os.WriteFile(filename, []byte(replaced), 0644)
	if err != nil {
		panic(err)
	}
	return nil
}

func applyCosmosTomlOverridesByRegex(si StateSyncInfo, nodeHome, syncRPC string) error {
	err := cosmosTomlOverrideRegex(nodeHome+"/config/config.toml", "enable", "true", false)
	if err != nil {
		return err
	}
	err = cosmosTomlOverrideRegex(nodeHome+"/config/config.toml", "trust_period", "8h0m0s", true)
	if err != nil {
		return err
	}
	err = cosmosTomlOverrideRegex(nodeHome+"/config/config.toml", "trust_height", si.TrustHeight, false)
	if err != nil {
		return err
	}
	err = cosmosTomlOverrideRegex(nodeHome+"/config/config.toml", "trust_hash", si.TrustHash, true)
	if err != nil {
		return err
	}
	err = cosmosTomlOverrideRegex(nodeHome+"/config/config.toml", "rpc_servers", syncRPC, true)
	if err != nil {
		return err
	}
	return nil
}

package snapshot_init

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog/log"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
)

const (
	comsosGenesisFilePath     = "/etc/gaiad/genesis.json"
	cosmosTestnetStateSyncRPC = "https://rpc.state-sync-01.theta-testnet.polypore.xyz:443,https://rpc.state-sync-02.theta-testnet.polypore.xyz:443"
	cosmosTestnetSeedPeers    = "639d50339d7045436c756a042906b9a69970913f@seed-01.theta-testnet.polypore.xyz:26656,3e506472683ceb7ed75c1578d092c79785c27857@seed-02.theta-testnet.polypore.xyz:26656"
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
			err = replaceLineIfStartsWith("/config/app.toml", "minimum-gas-prices = \"\"", "minimum-gas-prices = \"0.0025uatom\"")
			if err != nil {
				panic(err)
			}
			si := GetStateSyncInfoTestnet(ctx)
			err = replaceLineIfStartsWith("/config/config.toml", "enable = false", "enable = true")
			if err != nil {
				panic(err)
			}
			err = replaceLineIfStartsWith("/config/config.toml", "rpc_servers = \"\"", fmt.Sprintf("rpc_servers = \"%s\"", cosmosTestnetStateSyncRPC))
			if err != nil {
				panic(err)
			}
			err = replaceLineIfStartsWith("/config/config.toml", "trust_height = 0", fmt.Sprintf("trust_height = %s", si.TrustHeight))
			if err != nil {
				panic(err)
			}
			err = replaceLineIfStartsWith("/config/config.toml", "trust_hash = \"\"", fmt.Sprintf("trust_hash = \"%s\"", si.TrustHash))
			if err != nil {
				panic(err)
			}
			err = replaceLineIfStartsWith("/config/config.toml", "trust_period = \"0s\"", fmt.Sprintf("trust_period = \"%s\"", "8h0m0s"))
			if err != nil {
				panic(err)
			}
			err = replaceLineIfStartsWith("/config/config.toml", "seeds = \"\"", fmt.Sprintf("seeds = \"%s\"", cosmosTestnetSeedPeers))
			if err != nil {
				panic(err)
			}
			p = filepaths.Path{
				DirIn: "/config",
				FnIn:  "config.toml",
			}
			log.Info().Interface("config.toml", string(p.ReadFileInPath())).Msg("config.toml")
			p = filepaths.Path{
				DirIn: "/config",
				FnIn:  "app.toml",
			}
			log.Info().Interface("app.toml", string(p.ReadFileInPath())).Msg("config.toml")
		}
	}
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

func replaceLineIfStartsWith(inputFilePath, searchString, replacementString string) error {
	inputFile, err := os.Open(inputFilePath)
	if err != nil {
		return fmt.Errorf("error opening input file: %v", err)
	}
	defer inputFile.Close()
	scanner := bufio.NewScanner(inputFile)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, searchString) {
			lines = append(lines, replacementString)
		} else {
			lines = append(lines, line)
		}
	}
	if err = scanner.Err(); err != nil {
		return fmt.Errorf("error scanning input file: %v", err)
	}
	output := strings.Join(lines, "\n") + "\n"
	if err = os.WriteFile(inputFilePath, []byte(output), 0644); err != nil {
		return fmt.Errorf("error writing to input file: %v", err)
	}
	return nil
}

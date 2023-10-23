package ephemery_reset

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog/log"
	beacon_cookbooks "github.com/zeus-fyi/zeus/cookbooks/ethereum/beacons"
	validator_cookbooks "github.com/zeus-fyi/zeus/cookbooks/ethereum/validators"
	"github.com/zeus-fyi/zeus/pkg/poseidon"
	"github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/compression"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
)

const (
	repoBase    = "pk910/test-testnet-repo"
	newRepoBase = "ephemery-testnet/ephemery-genesis"
	// https://github.com/ephemery-testnet/ephemery-genesis/releases/tag/ephemery-98
	ephemeralTestnetFile = "testnet-all.tar.gz"
)

func ExtractAndDecEphemeralTestnetConfig(dataDir filepaths.Path, clientName string) {
	switch clientName {
	case beacon_cookbooks.LighthouseEphemeral, validator_cookbooks.EphemeryValidatorClusterClassName, "lodestarEphemeral":
		log.Info().Interface("dataDir", dataDir).Msg("ExtractAndDecEphemeralTestnetConfig: LighthouseEphemeral")
		dataDir.DirIn = "/data/testnet"
		dataDir.DirOut = path.Join(dataDir.DirIn)
	case beacon_cookbooks.GethEphemeral:
		log.Info().Interface("dataDir", dataDir).Msg("ExtractAndDecEphemeralTestnetConfig: GethEphemeral")
		// placing a genesis.json file directly in the datadir path should set the chain to the expected value
		dataDir.DirOut = dataDir.DirIn
	case "test":
		dataDir.DirIn = "./data"
		dataDir.DirOut = "./data/testnet"
	default:
		return
	}

	// TODO refactor
	ok, _ := Exists(path.Join(dataDir.DirIn, "/retention.vars"))
	if ok {
		log.Info().Msg("previous genesis artifact for genesis interval found")
		kt := ExtractResetTime(path.Join(dataDir.DirIn, "/retention.vars"))
		log.Info().Int64("seconds until next genesis iteration", kt).Msg("ExtractAndDecEphemeralTestnetConfig: retention.vars")
		if kt <= 0 {
			log.Info().Interface("wipingDirPath", dataDir.DirIn).Msg("wiping datadir in path")
			err := RemoveContents(dataDir.DirIn)
			log.Info().Interface("wipingDirPath", dataDir.DirIn).Msg("wiping datadir in path done")
			if err != nil {
				log.Err(err).Msg("ExtractAndDecEphemeralTestnetConfig: RemoveContents")
			}
		}
	}
	ctx := context.Background()
	log.Info().Interface("dataDir.DirOut", dataDir.DirOut)
	if _, zerr := os.Stat(dataDir.DirOut); os.IsNotExist(zerr) {
		_ = os.MkdirAll(dataDir.DirOut, 0700) // Create your dir
	}
	log.Info().Interface("dataDir.DirIn", dataDir.DirIn)
	if _, zerr := os.Stat(dataDir.DirIn); os.IsNotExist(zerr) {
		_ = os.MkdirAll(dataDir.DirIn, 0700) // Create your dir
	}
	url := GetLatestReleaseConfigDownloadURL()
	dataDir.FnIn = ephemeralTestnetFile
	err := poseidon.DownloadFile(ctx, dataDir.DirIn, url)
	if err != nil {
		log.Ctx(ctx).Panic().Err(err).Msg("DownloadFile")
		panic(err)
	}
	dec := compression.NewCompression()
	err = dec.UnGzip(&dataDir)
	if err != nil {
		log.Ctx(ctx).Panic().Err(err).Msg("UnGzip")
		panic(err)
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
			panic(err)
		}
	}
}

func GetLatestReleaseConfigDownloadURL() string {
	rlNum, err := getLatestTestnetDataReleaseNumber()
	if err != nil {
		panic(err)
	}
	log.Info().Str("rlNum", rlNum).Msg("GetLatestReleaseConfigDownloadURL")
	return fmt.Sprintf("https://github.com/%s/releases/download/%s/%s", newRepoBase, rlNum, ephemeralTestnetFile)
}

func getLatestTestnetDataReleaseNumber() (string, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/releases/latest", newRepoBase)
	r := resty.New()
	resp, err := r.R().
		Get(url)
	var temp map[string]interface{}
	err = json.Unmarshal(resp.Body(), &temp)
	for k, v := range temp {
		if k == "tag_name" || k == "name" {
			log.Info().Interface("v", v).Msg("getLatestTestnetDataReleaseNumber")
			return v.(string), err
		}
	}
	return "", err
}

func ExtractResetTime(path string) int64 {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	var tmp1, tmp2 string

	for scanner.Scan() {
		line := scanner.Text()
		find := strings.Split(line, "export GENESIS_TIMESTAMP=")
		if len(find) > 1 {
			tmp1 = strings.TrimLeft(find[1], `"`)
			tmp1 = strings.TrimRight(tmp1, `"`)
		}
		find2 := strings.Split(line, "export GENESIS_RESET_INTERVAL=")
		if len(find2) > 1 {
			tmp2 = strings.TrimLeft(find2[1], `"`)
			tmp2 = strings.TrimRight(tmp2, `"`)
		}
		lines = append(lines, line)
	}

	genesisTime, err := strconv.ParseInt(tmp1, 10, 64)
	if err != nil {
		panic(err)
	}

	resetInterval, err := strconv.ParseInt(tmp2, 10, 64)
	if err != nil {
		panic(err)
	}
	if scanner.Err() != nil {
		panic(scanner.Err())
	}
	killTime := genesisTime + resetInterval - time.Now().Unix()
	log.Info().Int64("genesisTime", genesisTime).Int64("resetInterval", resetInterval).Int64("killTime", killTime).Msg("ExtractResetTime")
	return killTime
}

func Exists(name string) (bool, error) {
	_, err := os.Stat(name)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return false, err
}

func RemoveContents(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		log.Err(err).Interface("dir", dir).Msg("RemoveContents: os.Open(dir)")
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		log.Err(err).Interface("dir", dir).Msg("RemoveContents: d.Readdirnames(-1)")
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			log.Err(err).Interface("dir", dir).Interface("name", name).Msg("RemoveContents: os.RemoveAll(filepath.Join(dir, name))")
			return err
		}
	}
	return nil
}

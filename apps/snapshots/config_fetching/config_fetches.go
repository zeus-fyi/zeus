package config_fetching

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

const (
	repoBase             = "pk910/test-testnet-repo"
	ephemeralTestnetFile = "testnet-all.tar.gz"
)

func GetLatestReleaseConfigDownloadURL() string {
	rlNum, err := getLatestTestnetDataReleaseNumber()
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("https://github.com/%s/releases/download/%s/%s", repoBase, rlNum, ephemeralTestnetFile)
}

func getLatestTestnetDataReleaseNumber() (string, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/releases/latest", repoBase)
	r := resty.New()
	resp, err := r.R().
		Get(url)
	var temp map[string]interface{}
	err = json.Unmarshal(resp.Body(), &temp)
	for k, v := range temp {
		if k == "tag_name" {
			return v.(string), err
		}
	}
	return "", err
}

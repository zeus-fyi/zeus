package api_calls

import (
	"fmt"

	"github.com/zeus-fyi/zeus/demos/api_calls/endpoints"
	"github.com/zeus-fyi/zeus/test/configs"
)

func ReadTopologiesMetadataAPICall() error {
	cfg := configs.InitLocalTestConfigs()
	client := GetBaseRestyClient()
	resp, err := client.R().
		SetAuthToken(cfg.Bearer).
		Get(endpoints.InfraReadTopologyV1Path)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	PrintRespJson(resp.Body())
	return err
}

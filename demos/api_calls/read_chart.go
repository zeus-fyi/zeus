package api_calls

import (
	"fmt"

	"github.com/zeus-fyi/zeus/demos/api_calls/endpoints"
	"github.com/zeus-fyi/zeus/test/configs"
)

type TopologyReadRequest struct {
	TopologyID int `json:"topologyID"`
}

func ReadChartAPICall() error {
	cfg := configs.InitLocalTestConfigs()
	tar := TopologyReadRequest{
		TopologyID: 0,
	}
	PrintReqJson(tar)
	client := GetBaseRestyClient()
	resp, err := client.R().
		SetAuthToken(cfg.Bearer).
		SetBody(tar).
		Post(endpoints.InfraReadChartV1Path)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	PrintRespJson(resp.Body())
	return err
}

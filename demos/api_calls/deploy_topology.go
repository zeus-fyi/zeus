package api_calls

import (
	"fmt"

	"github.com/zeus-fyi/zeus/demos/api_calls/endpoints"
	"github.com/zeus-fyi/zeus/test/configs"
)

type TopologyDeployRequest struct {
	TopologyID    int    `json:"topologyID"`
	CloudProvider string `json:"cloudProvider"`
	Region        string `json:"region"`
	Context       string `json:"context"`
	Namespace     string `json:"namespace"`
	Env           string `json:"env"`
}

func DeployDemoProdChartApiCall() error {
	cfg := configs.InitLocalTestConfigs()
	deployKns := TopologyDeployRequest{
		TopologyID:    0,
		CloudProvider: "do",
		Region:        "sfo3",
		Context:       "dev-sfo3-zeus",
		Namespace:     "demo",
		Env:           "dev",
	}

	client := GetBaseRestyClient()
	resp, err := client.R().
		SetAuthToken(cfg.Bearer).
		SetBody(deployKns).
		Post(endpoints.DeployTopologyV1Path)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	PrintRespJson(resp.Body())
	return err
}

package api_calls

import (
	"fmt"

	"github.com/zeus-fyi/zeus/demos/api_calls/endpoints"
	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_req_types"
	"github.com/zeus-fyi/zeus/test/configs"
)

func DeployDemoProdChartApiCall() error {
	cfg := configs.InitLocalTestConfigs()
	deployKns := zeus_req_types.TopologyDeployRequest{
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

package zeus_client

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	zk8s_templates "github.com/zeus-fyi/zeus/zeus/workload_config_drivers/templates"
	zeus_endpoints "github.com/zeus-fyi/zeus/zeus/z_client/endpoints"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_req_types"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_resp_types"
)

func (z *ZeusClient) CreateCluster(ctx context.Context, tar zk8s_templates.Cluster) (zeus_resp_types.DeployStatus, error) {
	z.PrintReqJson(tar)
	tcc := zeus_req_types.TopologyCreateCluster{
		Cluster: tar,
	}
	respJson := zeus_resp_types.DeployStatus{}
	resp, err := z.R().
		SetResult(&respJson).
		SetBody(tcc).
		Post(zeus_endpoints.InfraCreateClusterV1Path)

	if err != nil || resp.StatusCode() >= 400 {
		if err == nil {
			err = fmt.Errorf("non-OK status code: %d", resp.StatusCode())
		}
		log.Err(err).Msg("ZeusClient: Deploy")
		return respJson, err
	}
	z.PrintRespJson(resp.Body())
	return respJson, err
}

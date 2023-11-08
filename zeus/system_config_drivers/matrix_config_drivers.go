package system_config_drivers

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	zeus_cluster_config_drivers "github.com/zeus-fyi/zeus/zeus/cluster_config_drivers"
	cluster_node_resources "github.com/zeus-fyi/zeus/zeus/cluster_resources/nodes"
	zeus_client "github.com/zeus-fyi/zeus/zeus/z_client"
	zeus_endpoints "github.com/zeus-fyi/zeus/zeus/z_client/endpoints"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_resp_types"
)

type MatrixDefinition struct {
	zeus_client.ZeusClient
	MatrixName string

	// multi cluster setup, eg 10 ethereum beacons, at supplied cloud ctx ns locations
	Clusters []zeus_cluster_config_drivers.ClusterDefinition
	Nodes    cluster_node_resources.NodesGroup
}

// todo, finish this

func (z *MatrixDefinition) RegisterMatrixDefinition(ctx context.Context, tar any) (any, error) {
	z.PrintReqJson(tar)

	respJson := zeus_resp_types.DeployStatus{}
	resp, err := z.R().
		SetResult(&respJson).
		SetBody(tar).
		Post(zeus_endpoints.InfraCreateMatrixV1Path)

	if err != nil || (resp.StatusCode() >= 400) {
		if err == nil {
			err = fmt.Errorf("non-OK status code: %d", resp.StatusCode())
		}
		log.Ctx(ctx).Err(err).Msg("ZeusClient: RegisterMatrixDefinition")
		return respJson, err
	}
	z.PrintRespJson(resp.Body())
	return respJson, err
}

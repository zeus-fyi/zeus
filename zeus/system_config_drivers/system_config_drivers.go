package system_config_drivers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
	zeus_client "github.com/zeus-fyi/zeus/zeus/client"
	zeus_endpoints "github.com/zeus-fyi/zeus/zeus/client/endpoints"
	"github.com/zeus-fyi/zeus/zeus/client/zeus_resp_types"
	cluster_node_resources "github.com/zeus-fyi/zeus/zeus/cluster_resources"
)

type SystemDefinition struct {
	zeus_client.ZeusClient
	Id         int
	SystemName string

	// large scale infra setup multi-region, multi-cloud,
	// eg 10 ethereum beacons, 3 databases, 5 validator clusters, etc
	// at supplied cloud ctx ns locations
	Matrices []MatrixDefinition
	Nodes    cluster_node_resources.NodesGroup
}

func (z *SystemDefinition) RegisterSystemDefinition(ctx context.Context, tar any) (any, error) {
	z.PrintReqJson(tar)

	respJson := zeus_resp_types.DeployStatus{}
	resp, err := z.R().
		SetResult(&respJson).
		SetBody(tar).
		Post(zeus_endpoints.InfraCreateSystemV1Path)

	if err != nil || (resp.StatusCode() != http.StatusAccepted && resp.StatusCode() != http.StatusOK) {
		log.Ctx(ctx).Err(err).Msg("ZeusClient: RegisterSystemDefinition")
		if err == nil {
			err = fmt.Errorf("non-OK status code: %d", resp.StatusCode())
		}
		return respJson, err
	}
	z.PrintRespJson(resp.Body())
	return respJson, err
}

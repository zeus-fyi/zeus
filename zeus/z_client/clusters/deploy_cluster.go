package zk8s_clusters

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	zeus_client "github.com/zeus-fyi/zeus/zeus/z_client"
	zeus_endpoints "github.com/zeus-fyi/zeus/zeus/z_client/endpoints"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_req_types"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_resp_types"
)

func DeployCluster(ctx context.Context, z zeus_client.ZeusClient, tar zeus_req_types.ClusterTopologyDeployRequest) (zeus_resp_types.ClusterStatus, error) {
	z.PrintReqJson(tar)
	respJson := zeus_resp_types.ClusterStatus{}
	resp, err := z.R().
		SetResult(&respJson).
		SetBody(tar).
		Post(zeus_endpoints.DeployClusterTopologyV1Path)

	if err != nil || resp.StatusCode() >= 400 {
		if err == nil {
			err = fmt.Errorf("non-OK status code: %d", resp.StatusCode())
		}
		log.Err(err).Msg("ZeusClient: DeployCluster")
		return respJson, err
	}
	z.PrintRespJson(resp.Body())
	return respJson, err
}

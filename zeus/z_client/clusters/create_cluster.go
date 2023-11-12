package zk8s_clusters

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	zeus_cluster_config_drivers "github.com/zeus-fyi/zeus/zeus/cluster_config_drivers"
	zeus_client "github.com/zeus-fyi/zeus/zeus/z_client"
	zeus_endpoints "github.com/zeus-fyi/zeus/zeus/z_client/endpoints"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_resp_types"
)

func CreateCluster(ctx context.Context, z zeus_client.ZeusClient, cluster zeus_cluster_config_drivers.Cluster) (zeus_resp_types.DeployStatus, error) {
	z.PrintReqJson(cluster)
	respJson := zeus_resp_types.DeployStatus{}
	resp, err := z.R().
		SetResult(&respJson).
		SetBody(cluster).
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

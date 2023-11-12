package zeus_client

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	zeus_endpoints "github.com/zeus-fyi/zeus/zeus/z_client/endpoints"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_req_types"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_resp_types/topology_workloads"
)

func (z *ZeusClient) ClusterCreate(ctx context.Context, tar zeus_req_types.TopologyCreateClusterClassRequest) (topology_workloads.TopologyCreateClassResponse, error) {
	z.PrintReqJson(tar)
	respJson := topology_workloads.TopologyCreateClassResponse{}
	resp, err := z.R().
		SetResult(&respJson).
		SetBody(&tar).
		Post(zeus_endpoints.InfraCreateClassV1Path)

	if err != nil || resp.StatusCode() >= 400 {
		if err == nil {
			err = fmt.Errorf("non-OK status code: %d", resp.StatusCode())
		}
		log.Err(err).Msg("ZeusClient: CreateClass")
		return respJson, err
	}
	z.PrintRespJson(resp.Body())
	return respJson, err
}

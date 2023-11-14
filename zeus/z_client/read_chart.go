package zeus_client

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/zeus-fyi/zeus/zeus/workload_config_drivers/topology_workloads"
	zeus_endpoints "github.com/zeus-fyi/zeus/zeus/z_client/endpoints"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_req_types"
)

func (z *ZeusClient) ReadChart(ctx context.Context, tar zeus_req_types.TopologyRequest) (topology_workloads.TopologyBaseInfraWorkload, error) {
	respJson := topology_workloads.TopologyBaseInfraWorkload{}
	z.PrintReqJson(tar)
	resp, err := z.R().
		SetResult(&respJson).
		SetBody(tar).
		Post(zeus_endpoints.InfraReadChartV1Path)

	if err != nil || resp.StatusCode() >= 400 {
		if err == nil {
			err = fmt.Errorf("non-OK status code: %d", resp.StatusCode())
		}
		log.Err(err).Msg("ZeusClient: ReadChart")
		return respJson, err
	}
	z.PrintRespJson(resp.Body())
	return respJson, err
}

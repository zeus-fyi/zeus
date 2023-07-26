package zeus_client

import (
	"context"
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
	zeus_endpoints "github.com/zeus-fyi/zeus/zeus/z_client/endpoints"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_req_types"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_resp_types/topology_workloads"
)

func (z *ZeusClient) ReadChart(ctx context.Context, tar zeus_req_types.TopologyRequest) (topology_workloads.TopologyBaseInfraWorkload, error) {
	respJson := topology_workloads.TopologyBaseInfraWorkload{}
	z.PrintReqJson(tar)
	resp, err := z.R().
		SetResult(&respJson).
		SetBody(tar).
		Post(zeus_endpoints.InfraReadChartV1Path)

	if err != nil || resp.StatusCode() != http.StatusOK {
		if err == nil {
			err = fmt.Errorf("non-OK status code: %d", resp.StatusCode())
		}
		log.Ctx(ctx).Err(err).Msg("ZeusClient: ReadChart")
		return respJson, err
	}
	z.PrintRespJson(resp.Body())
	return respJson, err
}

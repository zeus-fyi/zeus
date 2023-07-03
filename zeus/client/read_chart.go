package zeus_client

import (
	"context"
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
	zeus_endpoints "github.com/zeus-fyi/zeus/zeus/client/endpoints"
	"github.com/zeus-fyi/zeus/zeus/client/zeus_req_types"
	"github.com/zeus-fyi/zeus/zeus/client/zeus_resp_types/topology_workloads"
)

func (z *ZeusClient) ReadChart(ctx context.Context, tar zeus_req_types.TopologyRequest) (topology_workloads.TopologyBaseInfraWorkload, error) {
	respJson := topology_workloads.TopologyBaseInfraWorkload{}
	z.PrintReqJson(tar)
	resp, err := z.R().
		SetResult(&respJson).
		SetBody(tar).
		Post(zeus_endpoints.InfraReadChartV1Path)

	if err != nil || resp.StatusCode() != http.StatusOK {
		log.Ctx(ctx).Err(err).Msg("ZeusClient: ReadChart")
		if err == nil {
			err = fmt.Errorf("non-OK status code: %d", resp.StatusCode())
		}
		return respJson, err
	}
	z.PrintRespJson(resp.Body())
	return respJson, err
}

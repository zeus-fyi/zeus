package zeus_client

import (
	"context"
	"net/http"

	"github.com/rs/zerolog/log"
	zeus_endpoints "github.com/zeus-fyi/zeus/pkg/zeus/client/endpoints"
	"github.com/zeus-fyi/zeus/pkg/zeus/client/req_types"
	"github.com/zeus-fyi/zeus/pkg/zeus/client/resp_types/topology_workloads"
)

func (z *ZeusClient) ReadChart(ctx context.Context, tar req_types.TopologyRequest) (topology_workloads.TopologyBaseInfraWorkload, error) {
	respJson := topology_workloads.TopologyBaseInfraWorkload{}
	z.PrintReqJson(tar)
	resp, err := z.R().
		SetResult(&respJson).
		SetBody(tar).
		Post(zeus_endpoints.InfraReadChartV1Path)

	if err != nil || resp.StatusCode() != http.StatusOK {
		log.Ctx(ctx).Err(err).Msg("ZeusClient: ReadChart")
		return respJson, err
	}
	z.PrintRespJson(resp.Body())
	return respJson, err
}

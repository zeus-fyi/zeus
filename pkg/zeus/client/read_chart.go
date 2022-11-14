package zeus_client

import (
	"context"
	"net/http"

	"github.com/rs/zerolog/log"
	zeus_endpoints "github.com/zeus-fyi/zeus/pkg/zeus/client/endpoints"
	"github.com/zeus-fyi/zeus/pkg/zeus/client/req_types"
	"github.com/zeus-fyi/zeus/pkg/zeus/client/resp_types"
)

func (z *ZeusClient) ReadChart(ctx context.Context, tar req_types.TopologyRequest) (resp_types.TopologyChartWorkload, error) {
	respJson := resp_types.TopologyChartWorkload{}
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

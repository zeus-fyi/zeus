package zeus_client

import (
	"context"
	"net/http"

	"github.com/rs/zerolog/log"
	zeus_endpoints "github.com/zeus-fyi/zeus/pkg/zeus/client/endpoints"
	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_req_types"
	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_resp_types/live_workload_query"
)

func (z *ZeusClient) ReadNamespaceWorkload(ctx context.Context, tar zeus_req_types.TopologyCloudCtxNsQueryRequest) (live_workload_query.NamespaceWorkload, error) {
	z.PrintReqJson(tar)
	respJson := live_workload_query.NamespaceWorkload{}

	resp, err := z.R().
		SetResult(&respJson).
		SetBody(tar).
		Post(zeus_endpoints.ReadWorkloadV1Path)

	if err != nil || resp.StatusCode() != http.StatusOK {
		log.Ctx(ctx).Err(err).Msg("ZeusClient: ReadNamespaceWorkload")
		return respJson, err
	}
	z.PrintRespJson(resp.Body())
	return respJson, err
}

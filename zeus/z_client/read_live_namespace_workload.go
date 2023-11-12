package zeus_client

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	zeus_endpoints "github.com/zeus-fyi/zeus/zeus/z_client/endpoints"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_req_types"

	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_resp_types/live_workload_query"
)

func (z *ZeusClient) ReadNamespaceWorkload(ctx context.Context, tar zeus_req_types.TopologyCloudCtxNsQueryRequest) (live_workload_query.NamespaceWorkload, error) {
	z.PrintReqJson(tar)
	respJson := live_workload_query.NamespaceWorkload{}

	resp, err := z.R().
		SetResult(&respJson).
		SetBody(tar).
		Post(zeus_endpoints.ReadWorkloadV1Path)

	if err != nil || resp.StatusCode() >= 400 {
		if err == nil {
			err = fmt.Errorf("non-OK status code: %d", resp.StatusCode())
		}
		log.Err(err).Msg("ZeusClient: ReadNamespaceWorkload")
		return respJson, err
	}
	z.PrintRespJson(resp.Body())
	return respJson, err
}

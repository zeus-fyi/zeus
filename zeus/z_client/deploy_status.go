package zeus_client

import (
	"context"
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
	zeus_endpoints "github.com/zeus-fyi/zeus/zeus/z_client/endpoints"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_req_types"

	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_resp_types"
)

func (z *ZeusClient) ReadDeployStatusUpdates(ctx context.Context, tar zeus_req_types.TopologyRequest) (zeus_resp_types.TopologyDeployStatuses, error) {
	z.PrintReqJson(tar)
	respJson := zeus_resp_types.TopologyDeployStatuses{}
	resp, err := z.R().
		SetResult(&respJson.Slice).
		SetBody(tar).
		Post(zeus_endpoints.DeployStatusV1Path)

	if err != nil || resp.StatusCode() != http.StatusOK {
		if err == nil {
			err = fmt.Errorf("non-OK status code: %d", resp.StatusCode())
		}
		log.Ctx(ctx).Err(err).Msg("ZeusClient: ReadDeployStatusUpdates")
		return respJson, err
	}
	z.PrintRespJson(resp.Body())
	return respJson, err
}

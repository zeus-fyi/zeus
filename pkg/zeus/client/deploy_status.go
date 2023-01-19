package zeus_client

import (
	"context"
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
	zeus_endpoints "github.com/zeus-fyi/zeus/pkg/zeus/client/endpoints"
	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_req_types"
	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_resp_types"
)

func (z *ZeusClient) ReadDeployStatusUpdates(ctx context.Context, tar zeus_req_types.TopologyRequest) (zeus_resp_types.TopologyDeployStatuses, error) {
	z.PrintReqJson(tar)
	respJson := zeus_resp_types.TopologyDeployStatuses{}
	resp, err := z.R().
		SetResult(&respJson.Slice).
		SetBody(tar).
		Post(zeus_endpoints.DeployStatusV1Path)

	if err != nil || resp.StatusCode() != http.StatusOK {
		log.Ctx(ctx).Err(err).Msg("ZeusClient: ReadDeployStatusUpdates")
		if err == nil {
			err = fmt.Errorf("non-OK status code: %d", resp.StatusCode())
		}
		return respJson, err
	}
	z.PrintRespJson(resp.Body())
	return respJson, err
}

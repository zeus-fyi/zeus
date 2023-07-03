package zeus_client

import (
	"context"
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
	zeus_endpoints "github.com/zeus-fyi/zeus/zeus/client/endpoints"
	"github.com/zeus-fyi/zeus/zeus/client/zeus_req_types"
	"github.com/zeus-fyi/zeus/zeus/client/zeus_resp_types"
)

// TODO validate ns '[a-z0-9]([-a-z0-9]*[a-z0-9])?')

func (z *ZeusClient) Deploy(ctx context.Context, tar zeus_req_types.TopologyDeployRequest) (zeus_resp_types.DeployStatus, error) {
	z.PrintReqJson(tar)

	respJson := zeus_resp_types.DeployStatus{}
	resp, err := z.R().
		SetResult(&respJson).
		SetBody(tar).
		Post(zeus_endpoints.DeployTopologyV1Path)

	if err != nil || resp.StatusCode() != http.StatusAccepted {
		log.Ctx(ctx).Err(err).Msg("ZeusClient: Deploy")
		if err == nil {
			err = fmt.Errorf("non-OK status code: %d", resp.StatusCode())
		}
		return respJson, err
	}
	z.PrintRespJson(resp.Body())
	return respJson, err
}

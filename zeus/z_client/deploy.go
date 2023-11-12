package zeus_client

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	zeus_endpoints "github.com/zeus-fyi/zeus/zeus/z_client/endpoints"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_req_types"

	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_resp_types"
)

// TODO validate ns '[a-z0-9]([-a-z0-9]*[a-z0-9])?')

func (z *ZeusClient) Deploy(ctx context.Context, tar zeus_req_types.TopologyDeployRequest) (zeus_resp_types.DeployStatus, error) {
	z.PrintReqJson(tar)

	respJson := zeus_resp_types.DeployStatus{}
	resp, err := z.R().
		SetResult(&respJson).
		SetBody(tar).
		Post(zeus_endpoints.DeployTopologyV1Path)

	if err != nil || resp.StatusCode() >= 400 {
		if err == nil {
			err = fmt.Errorf("non-OK status code: %d", resp.StatusCode())
		}
		log.Err(err).Msg("ZeusClient: Deploy")
		return respJson, err
	}
	z.PrintRespJson(resp.Body())
	return respJson, err
}

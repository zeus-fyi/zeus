package zeus_client

import (
	"context"
	"net/http"

	"github.com/rs/zerolog/log"
	zeus_endpoints "github.com/zeus-fyi/zeus/pkg/zeus/client/endpoints"
	"github.com/zeus-fyi/zeus/pkg/zeus/client/req_types"
	"github.com/zeus-fyi/zeus/pkg/zeus/client/resp_types"
)

func (z *ZeusClient) Deploy(ctx context.Context, tar req_types.TopologyDeployRequest) (resp_types.DeployStatus, error) {
	z.PrintReqJson(tar)

	respJson := resp_types.DeployStatus{}
	resp, err := z.R().
		SetResult(&respJson).
		SetBody(tar).
		Post(zeus_endpoints.DeployTopologyV1Path)

	if err != nil || resp.StatusCode() != http.StatusOK {
		log.Ctx(ctx).Err(err).Msg("ZeusClient: Deploy")
		return respJson, err
	}
	z.PrintRespJson(resp.Body())
	return respJson, err
}

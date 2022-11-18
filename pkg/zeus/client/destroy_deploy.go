package zeus_client

import (
	"context"
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_req_types"
	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_resp_types"

	zeus_endpoints "github.com/zeus-fyi/zeus/pkg/zeus/client/endpoints"
)

func (z *ZeusClient) DestroyDeploy(ctx context.Context, tar zeus_req_types.TopologyDeployRequest) (zeus_resp_types.DeployStatus, error) {
	z.PrintReqJson(tar)

	respJson := zeus_resp_types.DeployStatus{}
	resp, err := z.R().
		SetResult(&respJson).
		SetBody(tar).
		Post(zeus_endpoints.DestroyDeployInfraV1Path)

	if err != nil || resp.StatusCode() != http.StatusOK {
		log.Ctx(ctx).Err(err).Msg("ZeusClient: DestroyDeploy")
		return respJson, err
	}
	z.PrintRespJson(resp.Body())
	return respJson, err
}

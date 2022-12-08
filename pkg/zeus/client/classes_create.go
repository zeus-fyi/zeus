package zeus_client

import (
	"context"
	"errors"
	"net/http"

	"github.com/rs/zerolog/log"
	zeus_endpoints "github.com/zeus-fyi/zeus/pkg/zeus/client/endpoints"
	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_req_types"
	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_resp_types/topology_workloads"
)

func (z *ZeusClient) CreateClass(ctx context.Context, tar zeus_req_types.TopologyCreateOrAddBasesToClassesRequest) (topology_workloads.TopologyCreateClassResponse, error) {
	z.PrintReqJson(tar)
	respJson := topology_workloads.TopologyCreateClassResponse{}
	resp, err := z.R().
		SetResult(&respJson).
		SetBody(&tar).
		Post(zeus_endpoints.InfraCreateClassV1Path)

	if err != nil || resp.StatusCode() != http.StatusOK {
		log.Ctx(ctx).Err(err).Msg("ZeusClient: CreateClass")
		if resp.StatusCode() == http.StatusBadRequest {
			err = errors.New("bad request")
		}
		return respJson, err
	}
	z.PrintRespJson(resp.Body())
	return respJson, err
}

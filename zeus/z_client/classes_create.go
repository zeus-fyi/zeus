package zeus_client

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_req_types"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_resp_types/topology_workloads"

	zeus_endpoints "github.com/zeus-fyi/zeus/zeus/z_client/endpoints"
)

func (z *ZeusClient) CreateClass(ctx context.Context, tar zeus_req_types.TopologyCreateClusterClassRequest) (topology_workloads.TopologyCreateClassResponse, error) {
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
		if err == nil {
			err = fmt.Errorf("non-OK status code: %d", resp.StatusCode())
		}
		return respJson, err
	}
	z.PrintRespJson(resp.Body())
	return respJson, err
}

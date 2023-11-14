package zeus_client

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	zeus_endpoints "github.com/zeus-fyi/zeus/zeus/z_client/endpoints"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_resp_types"
)

func (z *ZeusClient) ReadTopologiesOrgCloudCtxNs(ctx context.Context) (zeus_resp_types.TopologiesOrgCloudCtxNsSlice, error) {
	respJson := zeus_resp_types.TopologiesOrgCloudCtxNsSlice{}
	resp, err := z.R().
		SetResult(&respJson).
		Get(zeus_endpoints.InfraReadOrgTopologiesV1Path)

	if err != nil || resp.StatusCode() > 400 {
		if err == nil {
			err = fmt.Errorf("non-OK status code: %d", resp.StatusCode())
		}
		log.Err(err).Msg("ZeusClient: ReadTopologiesOrgCloudCtxNs")
		return nil, err
	}
	z.PrintRespJson(resp.Body())
	return respJson, err
}

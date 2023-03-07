package zeus_client

import (
	"context"
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
	zeus_endpoints "github.com/zeus-fyi/zeus/pkg/zeus/client/endpoints"
	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_resp_types"
)

func (z *ZeusClient) ReadTopologiesOrgCloudCtxNs(ctx context.Context) (zeus_resp_types.TopologiesOrgCloudCtxNsSlice, error) {
	respJson := zeus_resp_types.TopologiesOrgCloudCtxNsSlice{}
	resp, err := z.R().
		SetResult(&respJson).
		Get(zeus_endpoints.InfraReadOrgTopologiesV1Path)

	if err != nil || resp.StatusCode() != http.StatusOK {
		log.Ctx(ctx).Err(err).Msg("ZeusClient: ReadTopologiesOrgCloudCtxNs")
		if err == nil {
			err = fmt.Errorf("non-OK status code: %d", resp.StatusCode())
		}
		return nil, err
	}
	z.PrintRespJson(resp.Body())
	return respJson, err
}

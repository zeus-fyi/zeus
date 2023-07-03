package zeus_client

import (
	"context"
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
	zeus_endpoints "github.com/zeus-fyi/zeus/zeus/client/endpoints"
	"github.com/zeus-fyi/zeus/zeus/client/zeus_resp_types"
)

func (z *ZeusClient) ReadTopologies(ctx context.Context) (zeus_resp_types.ReadTopologiesMetadataGroup, error) {
	respJson := zeus_resp_types.ReadTopologiesMetadataGroup{}
	resp, err := z.R().
		SetResult(&respJson.Slice).
		Get(zeus_endpoints.InfraReadTopologyV1Path)

	if err != nil || resp.StatusCode() != http.StatusOK {
		log.Ctx(ctx).Err(err).Msg("ZeusClient: ReadTopologies")
		if err == nil {
			err = fmt.Errorf("non-OK status code: %d", resp.StatusCode())
		}
		return respJson, err
	}
	z.PrintRespJson(resp.Body())
	return respJson, err
}

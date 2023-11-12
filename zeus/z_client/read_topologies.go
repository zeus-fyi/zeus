package zeus_client

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	zeus_endpoints "github.com/zeus-fyi/zeus/zeus/z_client/endpoints"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_resp_types"
)

func (z *ZeusClient) ReadTopologies(ctx context.Context) (zeus_resp_types.ReadTopologiesMetadataGroup, error) {
	respJson := zeus_resp_types.ReadTopologiesMetadataGroup{}
	resp, err := z.R().
		SetResult(&respJson.Slice).
		Get(zeus_endpoints.InfraReadTopologyV1Path)

	if err != nil || resp.StatusCode() >= 400 {
		if err == nil {
			err = fmt.Errorf("non-OK status code: %d", resp.StatusCode())
		}
		log.Err(err).Msg("ZeusClient: ReadTopologies")
		return respJson, err
	}
	z.PrintRespJson(resp.Body())
	return respJson, err
}

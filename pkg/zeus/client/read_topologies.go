package zeus_client

import (
	"context"
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/zeus-fyi/zeus/pkg/zeus/client/resp_types"

	zeus_endpoints "github.com/zeus-fyi/zeus/pkg/zeus/client/endpoints"
)

func (z *ZeusClient) ReadTopologies(ctx context.Context) (resp_types.ReadTopologiesMetadataGroup, error) {
	respJson := resp_types.ReadTopologiesMetadataGroup{}
	resp, err := z.R().
		SetResult(&respJson.Slice).
		Get(zeus_endpoints.InfraReadTopologyV1Path)

	if err != nil || resp.StatusCode() != http.StatusOK {
		log.Ctx(ctx).Err(err).Msg("ZeusClient: ReadTopologies")
		return respJson, err
	}
	z.PrintRespJson(resp.Body())
	return respJson, err
}

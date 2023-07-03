package zeus_client

import (
	"context"
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
	zeus_endpoints "github.com/zeus-fyi/zeus/zeus/client/endpoints"
	zeus_config_map_reqs "github.com/zeus-fyi/zeus/zeus/client/zeus_req_types/config_maps"
)

func (z *ZeusClient) SwapConfigMapKeys(ctx context.Context, par zeus_config_map_reqs.ConfigMapActionRequest) ([]byte, error) {
	par.Action = zeus_config_map_reqs.KeySwapAction
	resp, err := z.R().
		SetBody(par).
		Post(zeus_endpoints.ConfigMapsActionV1Path)

	if err != nil || resp.StatusCode() != http.StatusOK {
		log.Ctx(ctx).Err(err).Msg("ZeusClient: SwapConfigMapKeys")
		if err == nil {
			err = fmt.Errorf("non-OK status code: %d", resp.StatusCode())
		}
		return resp.Body(), err
	}
	z.PrintRespJson(resp.Body())
	return resp.Body(), err
}

// SetOrCreateKeyFromConfigMapKey keyOne=keyToCopy, keyTwo=keyToSetOrCreateFromCopy
func (z *ZeusClient) SetOrCreateKeyFromConfigMapKey(ctx context.Context, par zeus_config_map_reqs.ConfigMapActionRequest) ([]byte, error) {
	par.Action = zeus_config_map_reqs.SetOrCreateKeyFromExisting
	resp, err := z.R().
		SetBody(par).
		Post(zeus_endpoints.ConfigMapsActionV1Path)

	if err != nil || resp.StatusCode() != http.StatusOK {
		log.Ctx(ctx).Err(err).Msg("ZeusClient: SetOrCreateKeyFromConfigMapKey")
		if err == nil {
			err = fmt.Errorf("non-OK status code: %d", resp.StatusCode())
		}
		return resp.Body(), err
	}
	z.PrintRespJson(resp.Body())
	return resp.Body(), err
}

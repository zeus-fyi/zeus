package pods_client

import (
	"context"
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
	zeus_endpoints "github.com/zeus-fyi/zeus/zeus/z_client/endpoints"
	zeus_pods_reqs "github.com/zeus-fyi/zeus/zeus/z_client/zeus_req_types/pods"
)

func (z *PodsClient) DeletePods(ctx context.Context, par zeus_pods_reqs.PodActionRequest) ([]byte, error) {
	par.Action = zeus_pods_reqs.DeleteAllPods
	resp, err := z.R().
		SetBody(par).
		Post(zeus_endpoints.PodsActionV1Path)

	if err != nil || resp.StatusCode() != http.StatusOK {
		log.Ctx(ctx).Err(err).Msg("ZeusClient: DeletePods")
		if err == nil {
			err = fmt.Errorf("non-OK status code: %d", resp.StatusCode())
		}
		return resp.Body(), err
	}
	z.PrintRespJson(resp.Body())
	return resp.Body(), err
}

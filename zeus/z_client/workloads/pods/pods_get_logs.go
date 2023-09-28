package pods_client

import (
	"context"
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
	zeus_endpoints "github.com/zeus-fyi/zeus/zeus/z_client/endpoints"
	zeus_pods_reqs "github.com/zeus-fyi/zeus/zeus/z_client/zeus_req_types/pods"
)

// GetPodLogs will use this filter by default unless you provide an override
// filter request.FilterOpts.StartsWith = request.PodName
func (z *PodsClient) GetPodLogs(ctx context.Context, par zeus_pods_reqs.PodActionRequest) ([]byte, error) {
	par.Action = zeus_pods_reqs.GetPodLogs
	resp, err := z.R().
		SetBody(par).
		Post(zeus_endpoints.PodsActionV1Path)
	if err != nil || resp.StatusCode() != http.StatusOK {
		if err == nil {
			err = fmt.Errorf("GetPodLogs: non-OK status code: %d", resp.StatusCode())
		}
		log.Err(err).Msg("ZeusClient: GetPodLogs")
		return resp.Body(), err
	}
	if z.PrintResp {
		z.PrintRespJson(resp.Body())
	}
	return resp.Body(), err
}

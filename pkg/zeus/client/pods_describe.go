package zeus_client

import (
	"context"
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
	zeus_endpoints "github.com/zeus-fyi/zeus/pkg/zeus/client/endpoints"
	zeus_pods_reqs "github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_req_types/pods"
	v1 "k8s.io/api/core/v1"
)

func (z *ZeusClient) GetPods(ctx context.Context, par zeus_pods_reqs.PodActionRequest) (*v1.PodList, error) {
	par.Action = zeus_pods_reqs.GetPods

	pl := &v1.PodList{}
	resp, err := z.R().
		SetBody(par).
		SetResult(pl).
		Post(zeus_endpoints.PodsActionV1Path)

	if err != nil || resp.StatusCode() != http.StatusOK {
		log.Ctx(ctx).Err(err).Msg("ZeusClient: GetPods")
		if err == nil {
			err = fmt.Errorf("non-OK status code: %d", resp.StatusCode())
		}
		return nil, err
	}
	z.PrintRespJson(resp.Body())
	return pl, err
}

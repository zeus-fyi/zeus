package beacon_actions

import (
	"context"

	zeus_pods_reqs "github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_req_types/pods"
)

func (b *BeaconActionsClient) RestartConsensusClientPods(ctx context.Context, par zeus_pods_reqs.PodActionRequest) ([]byte, error) {
	par.PodName = b.ConsensusClient
	par.ContainerName = b.ConsensusClient
	b.PrintReqJson(par)
	resp, err := b.DeletePods(ctx, par)
	b.PrintRespJson(resp)
	return resp, err
}

func (b *BeaconActionsClient) RestartExecClientPods(ctx context.Context, par zeus_pods_reqs.PodActionRequest) ([]byte, error) {
	par.PodName = b.ExecClient
	par.ContainerName = b.ExecClient
	b.PrintReqJson(par)
	resp, err := b.DeletePods(ctx, par)
	b.PrintRespJson(resp)
	return resp, err
}

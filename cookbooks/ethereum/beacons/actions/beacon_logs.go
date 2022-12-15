package beacon_actions

import (
	"context"
	"path"

	"github.com/rs/zerolog/log"
	strings_filter "github.com/zeus-fyi/zeus/pkg/utils/strings"
	zeus_pods_reqs "github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_req_types/pods"
	v1 "k8s.io/api/core/v1"
)

func (b *BeaconActionsClient) PrintConsensusClientPodLogs(ctx context.Context, par zeus_pods_reqs.PodActionRequest) ([]byte, error) {
	b.PrintReqJson(par)
	par.PodName = "zeus-consensus-client"
	par.ContainerName = "zeus-consensus-client"
	filter := strings_filter.FilterOpts{Contains: "zeus-consensus-client"}
	logOpts := &v1.PodLogOptions{Container: "zeus-consensus-client"}
	par.LogOpts = logOpts
	par.FilterOpts = &filter

	resp, err := b.GetPodLogs(ctx, par)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("PrintConsensusClientPodLogs: GetPodLogs")
		return nil, err
	}
	b.PrintPath.FnOut = b.ConsensusClient + "_logs"
	b.PrintPath.DirOut = path.Join(b.PrintPath.DirIn, "/consensus_client")
	err = b.PrintPath.Print(resp, "json")
	return resp, err
}

func (b *BeaconActionsClient) PrintExecClientPodLogs(ctx context.Context, par zeus_pods_reqs.PodActionRequest) ([]byte, error) {
	b.PrintReqJson(par)
	par.PodName = "zeus-exec-client"
	par.ContainerName = "zeus-exec-client"
	logOpts := &v1.PodLogOptions{Container: "zeus-exec-client"}
	par.LogOpts = logOpts

	filter := strings_filter.FilterOpts{Contains: "zeus-exec-client"}
	par.FilterOpts = &filter
	resp, err := b.GetPodLogs(ctx, par)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("PrintExecClientPodLogs: GetPodLogs")
		return nil, err
	}
	b.PrintPath.FnOut = b.ExecClient + "_logs"
	b.PrintPath.DirOut = path.Join(b.PrintPath.DirIn, "/exec_client")
	err = b.PrintPath.Print(resp, "logs")
	return resp, err
}

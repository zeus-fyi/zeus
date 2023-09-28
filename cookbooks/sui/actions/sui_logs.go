package sui_actions

import (
	"context"
	"path"

	"github.com/rs/zerolog/log"
	strings_filter "github.com/zeus-fyi/zeus/pkg/utils/strings"
	zeus_pods_reqs "github.com/zeus-fyi/zeus/zeus/z_client/zeus_req_types/pods"
	v1 "k8s.io/api/core/v1"
)

func (s *SuiActionsClient) GetLogs(ctx context.Context, par zeus_pods_reqs.PodActionRequest) ([]byte, error) {
	par.ContainerName = "sui"
	filter := strings_filter.FilterOpts{Contains: "sui"}
	logOpts := &v1.PodLogOptions{Container: par.ContainerName}
	par.LogOpts = logOpts
	par.FilterOpts = &filter

	resp, err := s.GetPodLogs(ctx, par)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("SuiActionsClient: GetPodLogs")
		return nil, err
	}
	s.PrintPath.FnOut = "sui_logs"
	s.PrintPath.DirOut = path.Join(s.PrintPath.DirIn, "/sui")
	err = s.PrintPath.Print(resp, "json")
	return resp, err
}

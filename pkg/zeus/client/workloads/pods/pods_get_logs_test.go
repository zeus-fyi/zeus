package pods_client

import (
	zeus_pods_reqs "github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_req_types/pods"
	v1 "k8s.io/api/core/v1"
)

func (t *PodsClientTestSuite) TestGetPodsLogs() {
	deployKnsReq.Namespace = "ethereum"

	tailLines := int64(100)
	par := zeus_pods_reqs.PodActionRequest{
		TopologyDeployRequest: deployKnsReq,
		Action:                zeus_pods_reqs.GetPodLogs,
		PodName:               "zeus-geth-0",
		ContainerName:         "geth",
		FilterOpts:            nil,
		ClientReq:             nil,
		LogOpts:               &v1.PodLogOptions{Container: "geth", TailLines: &tailLines},
		DeleteOpts:            nil,
	}
	resp, err := t.ZeusTestClient.GetPodLogs(ctx, par)
	t.Require().Nil(err)
	t.Assert().NotEmpty(resp)
}

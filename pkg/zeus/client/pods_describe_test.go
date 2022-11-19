package zeus_client

import (
	zeus_pods_reqs "github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_req_types/pods"
)

func (t *ZeusClientTestSuite) TestGetPods() {
	deployKnsReq.Namespace = "ethereum"
	par := zeus_pods_reqs.PodActionRequest{
		TopologyDeployRequest: deployKnsReq,
		Action:                zeus_pods_reqs.GetPods,
	}
	resp, err := t.ZeusTestClient.GetPods(ctx, par)
	t.Require().Nil(err)
	t.Assert().NotEmpty(resp)
}

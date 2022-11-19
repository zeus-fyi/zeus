package zeus_client

import (
	zeus_pods_reqs "github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_req_types/pods"
)

func (t *ZeusClientTestSuite) TestDeletePods() {
	deployKnsReq.Namespace = "ethereum"
	par := zeus_pods_reqs.PodActionRequest{
		TopologyDeployRequest: deployKnsReq,
		Action:                zeus_pods_reqs.DeleteAllPods,
		PodName:               "zeus-geth-0",
	}
	err := t.ZeusTestClient.DeletePods(ctx, par)
	t.Require().Nil(err)
}

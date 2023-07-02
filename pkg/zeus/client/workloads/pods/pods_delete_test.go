package pods_client

import (
	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_common_types"
	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_req_types"
	zeus_pods_reqs "github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_req_types/pods"
)

// set your own topologyID here after uploading a chart workload
var deployKnsReq = zeus_req_types.TopologyDeployRequest{
	TopologyID: 1669101767430968000,
	CloudCtxNs: topCloudCtxNs,
}

// directs your api request to the right location
var topCloudCtxNs = zeus_common_types.CloudCtxNs{
	CloudProvider: "do",
	Region:        "sfo3",
	Context:       "do-sfo3-dev-do-sfo3-zeus",
	Namespace:     "demo", // set with your own namespace
	Env:           "dev",
}

// s
func (t *PodsClientTestSuite) TestDeletePods() {
	deployKnsReq.Namespace = "ethereum"
	par := zeus_pods_reqs.PodActionRequest{
		TopologyDeployRequest: deployKnsReq,
		Action:                zeus_pods_reqs.DeleteAllPods,
		PodName:               "zeus-geth-0",
	}
	_, err := t.ZeusTestClient.DeletePods(ctx, par)
	t.Require().Nil(err)
}

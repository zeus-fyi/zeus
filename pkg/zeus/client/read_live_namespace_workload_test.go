package zeus_client

import "github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_req_types"

func (t *ZeusClientTestSuite) TestReadLiveNamespaceWorkload() {
	tar := zeus_req_types.TopologyCloudCtxNsQueryRequest{
		CloudProvider: deployKnsReq.CloudProvider,
		Region:        deployKnsReq.Region,
		Context:       deployKnsReq.Context,
		Namespace:     deployKnsReq.Namespace,
		Env:           deployKnsReq.Env,
	}
	resp, err := t.ZeusTestClient.ReadNamespaceWorkload(ctx, tar)
	t.Require().Nil(err)
	t.Assert().NotEmpty(resp)
}

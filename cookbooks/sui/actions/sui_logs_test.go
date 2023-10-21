package sui_actions

import (
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_common_types"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_req_types"
	zeus_pods_reqs "github.com/zeus-fyi/zeus/zeus/z_client/zeus_req_types/pods"
)

func (t *SuiActionsCookbookTestSuite) TestGetLogs() {
	cloudCtxNs := zeus_common_types.CloudCtxNs{
		CloudProvider: "aws",
		Region:        "us-west-1",
		Context:       "zeus-us-west-1",
		Namespace:     "sui-03e7d0b6",
	}

	basePar := zeus_pods_reqs.PodActionRequest{
		TopologyDeployRequest: zeus_req_types.TopologyDeployRequest{
			CloudCtxNs: cloudCtxNs,
		},
	}
	resp, err := t.su.GetLogs(ctx, basePar)
	t.Assert().Nil(err)
	t.Assert().NotNil(resp)
}

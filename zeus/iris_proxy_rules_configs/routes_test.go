package iris_proxy_rules_configs

import (
	"context"

	hestia_req_types "github.com/zeus-fyi/zeus/pkg/hestia/client/req_types"
)

var ctx = context.Background()

func (t *IrisConfigTestSuite) TestCreateRoutingEndpoints() {
	rr := hestia_req_types.IrisOrgGroupRoutesRequest{
		Routes: []string{"https://zeus.fyi", "https://artemis.zeus.fyi"},
	}
	err := t.IrisClient.CreateRoutingEndpoints(ctx, rr)
	t.NoError(err)
}

func (t *IrisConfigTestSuite) TestReadRoutingEndpoints() {
	resp, err := t.IrisClient.ReadRoutingEndpoints(ctx)
	t.NoError(err)
	t.NotNil(resp)
}

func (t *IrisConfigTestSuite) TestDeleteRoutingEndpoints() {
	rr := hestia_req_types.IrisOrgGroupRoutesRequest{
		Routes: []string{"https://zeus.fyi", "https://artemis.zeus.fyi"},
	}
	resp, err := t.IrisClient.DeleteRoutingEndpoints(ctx, rr)
	t.NoError(err)
	t.NotNil(resp)
}

package iris_programmable_proxy

import hestia_req_types "github.com/zeus-fyi/zeus/pkg/hestia/client/req_types"

func (t *IrisConfigTestSuite) TestCreateRoutingGroup() {
	rr := hestia_req_types.IrisOrgGroupRoutesRequest{
		GroupName: "testGroupZ",
		Routes:    []string{"https://zeus.fyi", "https://artemis.zeus.fyi"},
	}
	err := t.IrisClient.CreateRoutingGroup(ctx, rr)
	t.NoError(err)
}

func (t *IrisConfigTestSuite) TestReadRoutingGroup() {
	resp, err := t.IrisClient.ReadRoutingGroupEndpoints(ctx, "testGroupZ")
	t.NoError(err)
	t.NotNil(resp)
}

func (t *IrisConfigTestSuite) TestReadAllRoutingGroups() {
	resp, err := t.IrisClient.ReadAllRoutingGroupsEndpoints(ctx)
	t.NoError(err)
	t.NotNil(resp)

}
func (t *IrisConfigTestSuite) TestUpdateRoutingGroup() {
	rr := hestia_req_types.IrisOrgGroupRoutesRequest{
		GroupName: "testGroupZ",
		Routes:    []string{"https://zeus.fyi"},
	}
	err := t.IrisClient.CreateRoutingGroup(ctx, rr)
	t.NoError(err)
}

func (t *IrisConfigTestSuite) TestDeleteRoutingGroup() {
	rr := hestia_req_types.IrisOrgGroupRoutesRequest{
		GroupName: "testGroupZ",
	}
	err := t.IrisClient.DeleteRoutingGroupEndpoints(ctx, rr)
	t.NoError(err)
}

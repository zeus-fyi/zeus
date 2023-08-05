package iris_programmable_proxy

import (
	"context"
	"fmt"
	"time"

	web3_actions "github.com/zeus-fyi/gochain/web3/client"
	hestia_req_types "github.com/zeus-fyi/zeus/pkg/hestia/client/req_types"
	resty_base "github.com/zeus-fyi/zeus/zeus/z_client/base"
)

func (t *IrisConfigTestSuite) TestRPCLoadBalancing() {
	routeGroup := "quicknode-mainnet"
	path := fmt.Sprintf("https://iris.zeus.fyi/v1/router")
	path = fmt.Sprintf("http://localhost:8080/v1/router")

	web3a := web3_actions.NewWeb3ActionsClient(path)
	web3a.AddRoutingGroupHeader(routeGroup)
	web3a.AddBearerToken(t.IrisClient.Token)
	web3a.Dial()
	reqCount := 4
	defer web3a.Close()
	for i := 0; i < reqCount; i++ {
		resp, err := web3a.C.BlockNumber(context.Background())
		t.NoError(err)
		t.NotNil(resp)
		fmt.Println(resp)
	}
}

func (t *IrisConfigTestSuite) TestGetLoadBalancing() {
	routeGroup := "olympus"
	path := fmt.Sprintf("https://iris.zeus.fyi/v1/router")
	//path = fmt.Sprintf("http://localhost:8080/v1/router")
	routeOne := "https://hestia.zeus.fyi/health"
	routeTwo := "https://iris.zeus.fyi/health"

	err := t.IrisClientProd.UpdateRoutingGroupEndpoints(ctx, hestia_req_types.IrisOrgGroupRoutesRequest{
		GroupName: routeGroup,
		Routes: []string{
			"https://hestia.zeus.fyi/health",
			"https://iris.zeus.fyi/health",
		},
	})
	t.Nil(err)
	r := resty_base.GetBaseRestyClient(path, t.IrisClientProd.Token)
	r.SetRoutingGroupHeader(routeGroup)
	reqCount := 4
	m := make(map[string]int)
	m[routeOne] = 0
	m[routeTwo] = 0
	for i := 0; i < reqCount; i++ {
		resp1, err1 := r.R().Get(path)
		t.NoError(err1)
		t.NotNil(resp1)

		selectedHeader := resp1.Header().Get("X-Selected-Route")
		t.NotEmpty(selectedHeader)
		fmt.Println(selectedHeader)

		m[selectedHeader]++
	}

	t.GreaterOrEqual(m[routeOne], 1)
	t.GreaterOrEqual(m[routeTwo], 1)

	rr := hestia_req_types.IrisOrgGroupRoutesRequest{
		Routes: []string{routeOne},
	}
	resp, err := t.IrisClientProd.DeleteRoutingEndpoints(ctx, rr)
	t.NoError(err)
	t.NotNil(resp)

	m[routeOne] = 0
	m[routeTwo] = 0

	// gives time for the routing group to update
	time.Sleep(5 * time.Second)
	for i := 0; i < reqCount; i++ {
		resp2, err2 := r.R().Get(path)
		t.NoError(err2)
		t.NotNil(resp)

		selectedHeader := resp2.Header().Get("X-Selected-Route")
		t.NotEmpty(selectedHeader)
		fmt.Println(selectedHeader)

		m[selectedHeader]++
	}
	t.Assert().Zero(m[routeOne])
	t.Assert().Equal(reqCount, m[routeTwo])
}

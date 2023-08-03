package iris_programmable_proxy

import (
	"context"
	"fmt"

	web3_actions "github.com/zeus-fyi/gochain/web3/client"
)

func (t *IrisConfigTestSuite) TestRPCLoadBalancing() {
	routeGroup := "quicknode-mainnet"
	path := fmt.Sprintf("https://iris.zeus.fyi/v1/router/group?routeGroup=%s", routeGroup)
	//path = fmt.Sprintf("http://localhost:8080/v1/router/group?routeGroup=%s", routeGroup)

	web3a := web3_actions.NewWeb3ActionsClient(path)
	web3a.AddBearerToken(t.IrisClient.Token)
	web3a.Dial()
	defer web3a.Close()
	for i := 0; i < 10; i++ {
		resp, err := web3a.C.BlockNumber(context.Background())
		t.NoError(err)
		t.NotNil(resp)
		fmt.Println(resp)
	}
}

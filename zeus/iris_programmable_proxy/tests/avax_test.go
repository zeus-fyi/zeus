package iris_tests_suite

import "fmt"

func (t *IrisConfigTestSuite) TestAvaxEndpoints() {
	groupName := "avalanche-mainnet"
	t.IrisClient.SetRoutingGroupHeader(groupName)
	t.IrisClient.Header.Set("Content-Type", "application/json")
	t.IrisClient.SetDebug(false)

	// C-Chain
	payload := `{
    "jsonrpc": "2.0",
    "method": "eth_chainId",
    "params": [],
    "id": 1}`
	resp, err := t.IrisClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		Post("/ext/bc/C/rpc")

	t.Require().NoError(err)
	t.Require().NotNil(resp)
	fmt.Println(resp.String())

	// X-Chain
	payload = `{
	"jsonrpc":"2.0",
		"id"     : 1,
		"method" :"avm.getBalance",
		"params" :{
		"address":"X-avax1pue5luvh6klhjkq8zk5zltxk84asvcnznsauxa",
			"assetID": "2pYGetDWyKdHxpFxh2LHeoLNCH6H5vxxCxHQtFnnFaYxLsqtHC"
	}
}`
	resp, err = t.IrisClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		Post("/ext/bc/X")

	t.Require().NoError(err)
	t.Require().NotNil(resp)
	fmt.Println(resp.String())

	// P-Chain
	payload = `{
	"jsonrpc":"2.0",
			"id"     :1,
			"method" :"platform.getHeight",
			"params": {}
	}`

	resp, err = t.IrisClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		Post("/ext/bc/P")

	t.Require().NoError(err)
	t.Require().NotNil(resp)
	fmt.Println(resp.String())

	// Index API
	payload = `{
	"jsonrpc":"2.0",
		"id"     :1,
		"method" :"index.getContainerByIndex",
		"params": {
		"index":0,
			"encoding":"hex"
	}
}`
	resp, err = t.IrisClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		Post("/ext/index/X/tx")

	t.Require().NoError(err)
	t.Require().NotNil(resp)
	fmt.Println(resp.String())

	// Debug API
	payload = `{"method":"debug_traceBlockByHash","params":["0x3e56c97d34f03b1369c351fa6c9f57c8bfa987c7da40964fab981303e0ef5849", {"tracer": "callTracer"}],"id":1,"jsonrpc":"2.0"}}`
	resp, err = t.IrisClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		Post("/ext/bc/C/rpc")

	t.Require().NoError(err)
	t.Require().NotNil(resp)
	fmt.Println(resp.String())

}

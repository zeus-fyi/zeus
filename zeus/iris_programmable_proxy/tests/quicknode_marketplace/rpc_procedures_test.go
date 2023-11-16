package iris_quicknode

import "fmt"

func (t *IrisConfigTestSuite) TestEthMaxBlockAggReduce() {
	groupName := "ethereum-mainnet"
	t.IrisClient.SetRoutingGroupHeader(groupName)
	t.IrisClient.Header.Set("Content-Type", "application/json")

	payload := `{
		"jsonrpc": "2.0",
		"procedure": "eth_maxBlockAggReduce",
		"method": "eth_getBlockByNumber",
		"params": ["latest", true],
		"id": 1
	}`
	resp, err := t.IrisClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		Post("/")

	t.Require().NoError(err)
	t.Require().NotNil(resp)
	fmt.Println(resp.String())
}

func (t *IrisConfigTestSuite) TestNearMaxBlockAggReduce() {
	groupName := "near-mainnet"
	t.IrisClient.SetRoutingGroupHeader(groupName)
	t.IrisClient.Header.Set("Content-Type", "application/json")
	payload := `{
		"jsonrpc": "2.0",
		"procedure": "near_maxBlockAggReduce",
		"method": "block",
		"params": {
			"finality": "final"
		},
		"id": 1
	}`
	resp, err := t.IrisClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		Post("/")

	t.Require().NoError(err)
	t.Require().NotNil(resp)
	fmt.Println(resp.String())
}

func (t *IrisConfigTestSuite) TestAvaxMaxBlockAggReduce() {
	groupName := "avax-mainnet"
	t.IrisClient.SetRoutingGroupHeader(groupName)
	t.IrisClient.Header.Set("Content-Type", "application/json")

	payload := `{
			"jsonrpc": "2.0",
			"procedure": "avax_maxBlockAggReduce",
			"method": "eth_getBlockByNumber",
			"params": ["latest", true],
			"id": 1
		}`
	resp, err := t.IrisClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		Post("/")

	t.Require().NoError(err)
	t.Require().NotNil(resp)
	fmt.Println(resp.String())
}

func (t *IrisConfigTestSuite) TestBtcMaxBlockAggReduce() {
	groupName := "btc-mainnet"
	t.IrisClient.SetRoutingGroupHeader(groupName)
	t.IrisClient.Header.Set("Content-Type", "application/json")

	payload := `{
		"jsonrpc": "2.0",
		"procedure": "btc_maxBlockAggReduce",
		"method": "getdifficulty",
		"id": 1
	}`
	resp, err := t.IrisClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		Post("/")

	t.Require().NoError(err)
	t.Require().NotNil(resp)
	fmt.Println(resp.String())
}

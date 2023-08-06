package iris_tests_suite

import "fmt"

func (t *IrisConfigTestSuite) TestSolanaEndpoints() {
	groupName := "solana-mainnet"
	t.IrisClient.SetRoutingGroupHeader(groupName)
	t.IrisClient.Header.Set("Content-Type", "application/json")
	t.IrisClient.SetDebug(false)

	payload := `{"jsonrpc":"2.0","id":1, "method":"getEpochInfo"}`
	resp, err := t.IrisClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		Post("/")

	t.Require().NoError(err)
	t.Require().NotNil(resp)
	fmt.Println(resp.String())

	payload = `{"jsonrpc":"2.0", "id":1, "method":"getTokenSupply", "params": ["7xKXtg2CW87d97TXJSDpbD5jBkheTqA83TZRuJosgAsU"]}`
	resp, err = t.IrisClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		Post("/")

	t.Require().NoError(err)
	t.Require().NotNil(resp)
	fmt.Println(resp.String())
}

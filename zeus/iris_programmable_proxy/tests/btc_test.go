package iris_tests_suite

import "fmt"

func (t *IrisConfigTestSuite) TestBtcEndpoints() {
	groupName := "btc-mainnet"
	t.IrisClient.SetRoutingGroupHeader(groupName)
	t.IrisClient.Header.Set("Content-Type", "application/json")
	t.IrisClient.SetDebug(false)

	payload := `{"method": "getblock", "params": ["00000000c937983704a73af28acdec37b049d214adbda81d7e2a3dd146f6ed09"]}`
	resp, err := t.IrisClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		Post("/")

	t.Require().NoError(err)
	t.Require().NotNil(resp)
	fmt.Println(resp.String())
}

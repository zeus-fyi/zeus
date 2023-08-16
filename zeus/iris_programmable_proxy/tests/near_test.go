package iris_tests_suite

import "fmt"

func (t *IrisConfigTestSuite) TestNearEndpoints() {
	groupName := "near-mainnet"
	t.IrisClient.SetRoutingGroupHeader(groupName)
	t.IrisClient.Header.Set("Content-Type", "application/json")
	t.IrisClient.SetDebug(false)

	payload := `{"method": "block","params": {"finality": "final"},"id":1,"jsonrpc":"2.0"}`
	resp, err := t.IrisClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		Post("/")

	t.Require().NoError(err)
	t.Require().NotNil(resp)
	fmt.Println(resp.String())
}

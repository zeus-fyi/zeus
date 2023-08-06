package iris_tests_suite

import "fmt"

func (t *IrisConfigTestSuite) TestAvaxEndpoints() {
	groupName := "avalanche-mainnet"
	t.IrisClient.SetRoutingGroupHeader(groupName)
	t.IrisClient.Header.Set("Content-Type", "application/json")
	t.IrisClient.SetDebug(false)
	resp, err := t.IrisClient.Client.R().Get("")
	t.Require().NoError(err)
	t.Require().NotNil(resp)
	fmt.Println(resp.String())
}

package iris_tests_suite

import "fmt"

func (t *IrisConfigTestSuite) TestAptosEndpoints() {
	groupName := "aptos-mainnet"
	t.IrisClient.SetRoutingGroupHeader(groupName)
	t.IrisClient.Header.Set("Content-Type", "application/json")
	t.IrisClient.SetDebug(true)
	resp, err := t.IrisClient.Client.R().Get("/v1/transactions/by_hash/0xbe9e71660e128e0e3e1f082c394f7b1bd2f4cb9c52207fe63cf4c8e7eb080e9d/")
	t.Require().NoError(err)
	t.Require().NotNil(resp)
	fmt.Println(resp.String())
}

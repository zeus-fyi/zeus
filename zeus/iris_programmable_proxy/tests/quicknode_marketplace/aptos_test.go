package iris_quicknode

import "fmt"

func (t *IrisConfigTestSuite) TestAptosEndpoints() {
	groupName := "aptos-mainnet"
	t.IrisClient.SetRoutingGroupHeader(groupName)
	t.IrisClient.Header.Set("Content-Type", "application/json")
	t.IrisClient.SetDebug(false)
	resp, err := t.IrisClient.Client.R().Get("/v1/transactions/by_hash/0xbe9e71660e128e0e3e1f082c394f7b1bd2f4cb9c52207fe63cf4c8e7eb080e9d/")
	t.Require().NoError(err)
	t.Require().NotNil(resp)
	fmt.Println(resp.String())

	ok := resp.Header().Get("X-Aptos-Epoch")
	t.Require().NotEmpty(ok)
	fmt.Println(ok)

	ok = resp.Header().Get("X-Aptos-Block-Height")
	t.Require().NotEmpty(ok)
	fmt.Println(ok)

	ok = resp.Header().Get("X-Aptos-Oldest-Block-Height")
	t.Require().NotEmpty(ok)
	fmt.Println(ok)

	pathSuffix := "/v1/tables/0x982be82410ddc6480dc8976e2866a2c7d162ee9d02d97ee842ac38b0b105086/item"

	payload := `{
		"key": {
			"collection":"Aptos Zero",
			"creator":"0xabf3630d0532fef81dfe610dd4def095070d91e344d475051e1c49da5e6d51c3",
			"name": "Aptos Zero: 1"
		},
		"key_type": "0x3::token::TokenDataId",
		"value_type": "0x3::token::TokenData"
	}`

	resp, err = t.IrisClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		Post(pathSuffix)

	t.Require().NoError(err)
	t.Require().NotNil(resp)
	fmt.Println(resp.String())
}

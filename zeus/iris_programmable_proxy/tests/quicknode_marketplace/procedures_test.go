package iris_quicknode

import (
	"fmt"

	iris_operators "github.com/zeus-fyi/zeus/pkg/iris/operators"
	iris_programmable_proxy_v1_beta "github.com/zeus-fyi/zeus/zeus/iris_programmable_proxy/v1beta"
)

func (t *IrisConfigTestSuite) TestMaxBlockAggProcedure() {
	groupName := "ethereum-mainnet"
	t.IrisClient.SetRoutingGroupHeader(groupName)

	payload := `{
				  "id": 1,
				  "jsonrpc": "2.0",
				  "method": "eth_blockNumber"
				}`

	m := map[string]interface{}{}

	resp, err := t.IrisClient.R().
		SetHeader("Content-Type", "application/json").
		SetResult(&m).
		SetBody(payload).
		Post("/")

	t.Require().NoError(err)
	t.Require().NotNil(resp)

	extMap := iris_programmable_proxy_v1_beta.IrisRoutingResponseETL{
		ExtractionKey: "result",
		DataType:      "",
	}

	extMap.ExtractKeyValue(m)
	t.Assert().NotEmpty(extMap.Value)
	fmt.Println(extMap)

	agg := iris_operators.Aggregation{
		Operator: "max",
		DataType: "int",
	}

	err = agg.AggregateOn(extMap.Value)
	t.Require().NoError(err)

	fmt.Println(agg)
}

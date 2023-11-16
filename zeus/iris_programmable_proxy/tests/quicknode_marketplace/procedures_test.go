package iris_quicknode

import (
	"fmt"

	iris_operators "github.com/zeus-fyi/zeus/pkg/iris/operators"
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

	source := resp.Header().Get("X-Selected-Route")
	t.Assert().NotEmpty(source)

	extResp := iris_operators.IrisRoutingResponseETL{
		Source:        source,
		ExtractionKey: "result",
		DataType:      "",
	}

	extResp.ExtractKeyValue(m)
	t.Assert().NotEmpty(extResp.Value)
	fmt.Println(extResp)

	agg := iris_operators.Aggregation{
		Operator: "max",
		DataType: "int",
	}

	err = agg.AggregateOn(extResp.Value, extResp)
	t.Require().NoError(err)

	fmt.Println(agg)
	t.Assert().NotEmpty(agg.DataSlice)

	fmt.Println(agg.DataSlice)
}

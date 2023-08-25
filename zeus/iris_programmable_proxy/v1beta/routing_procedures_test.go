package iris_programmable_proxy_v1_beta

import (
	"fmt"
	"time"

	"github.com/phf/go-queue/queue"
	iris_operators "github.com/zeus-fyi/zeus/pkg/iris/operators"
)

func (t *IrisConfigTestSuite) TestRoutingProcedure() {
	payload := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "eth_blockNumber",
		"params":  "[]",
		"id":      "1",
	}
	timeOut := time.Second * 3
	rgName := "ethereum-mainnet"
	getBlockHeightStep := BroadcastInstructions{
		RoutingPath:  "/",
		RestType:     "POST",
		MaxDuration:  timeOut,
		MaxTries:     3,
		RoutingTable: rgName,
		Payload:      payload,
	}
	getBlockHeightProcedure := IrisRoutingProcedureStep{
		BroadcastInstructions: getBlockHeightStep,
		TransformSlice: []iris_operators.IrisRoutingResponseETL{
			{
				Source:        "",
				ExtractionKey: "result",
				DataType:      "",
			},
		},
		AggregateMap: map[string]iris_operators.Aggregation{
			"result": {
				Operator:          "max",
				DataType:          "int",
				CurrentMaxInt:     0,
				CurrentMaxFloat64: 0,
			},
		},
	}
	payloadLatestBlock := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "eth_getBlockByNumber",
		"params":  []interface{}{"latest", true},
		"id":      1,
	}
	getBlockStep := BroadcastInstructions{
		RoutingPath:  "/",
		RestType:     "POST",
		MaxDuration:  timeOut,
		MaxTries:     3,
		RoutingTable: rgName,
		Payload:      payloadLatestBlock,
		FanInRules: &FanInRules{
			Rule: FanInRuleFirstValidResponse,
		},
	}
	getBlockProcedure := IrisRoutingProcedureStep{
		BroadcastInstructions: getBlockStep,
		TransformSlice: []iris_operators.IrisRoutingResponseETL{
			{
				Source:        "",
				ExtractionKey: "",
				DataType:      "",
			},
		},
	}
	que := queue.New()
	que.PushBack(getBlockHeightProcedure)
	que.PushBack(getBlockProcedure)
	procedure := IrisRoutingProcedure{
		Name:         MaxBlockAggReduce,
		OrderedSteps: que,
	}
	fmt.Println(procedure)

	//irisV1Beta := NewIrisV1BetaClient(t.IrisClient)
	//err := irisV1Beta.CreateProcedure(context.Background(), procedure)
	//t.Require().NoError(err)
}

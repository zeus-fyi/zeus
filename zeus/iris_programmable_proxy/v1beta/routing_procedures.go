package iris_programmable_proxy_v1_beta

import (
	"reflect"
	"time"

	"github.com/phf/go-queue/queue"
	iris_operators "github.com/zeus-fyi/zeus/pkg/iris/operators"
)

const (
	RequestHeaderRoutingProcedureHeader = "X-Routing-Procedure"

	ResponseHeaderProcedureLatency = "X-Procedure-Latency-Milliseconds"
)

// pre-canned routing procedures for QuickNode marketplace users

const (
	MaxBlockAggReduce = "max-block-agg-reduce"
)

type IrisRoutingProcedure struct {
	Name string `json:"name"`

	OrderedSteps *queue.Queue `json:"steps"`
}

type BroadcastInstructions struct {
	RoutingPath  string        `json:"routingPath"`
	RestType     string        `json:"restType"`
	Payload      any           `json:"payload,omitempty"`
	MaxDuration  time.Duration `json:"maxRuntime,omitempty"`
	MaxTries     int           `json:"maxTries,omitempty"`
	RoutingTable string        `json:"routingTable"`
	FanInRules   *FanInRules   `json:"fanInRules,omitempty"`
}

type FanInRules struct {
	Rule BroadcastRules `json:"rule"`
}

type BroadcastRules string

const (
	FanInRuleFirstValidResponse = "returnFirstValidResponse"
)

// ReturnFirstResult returns the first result from the fan-in that is not an error
func (b BroadcastRules) ReturnFirstResult() string {
	return FanInRuleFirstValidResponse
}

type IrisRoutingProcedureStep struct {
	BroadcastInstructions BroadcastInstructions                 `json:"broadcastInstructions,omitempty"`
	TransformSlice        []IrisRoutingResponseETL              `json:"transformSlice,omitempty"`
	AggregateMap          map[string]iris_operators.Aggregation `json:"aggregateMap,omitempty"`
}

func (r *IrisRoutingProcedureStep) Aggregate() error {
	if len(r.AggregateMap) == 0 {
		return nil
	}
	for _, v := range r.TransformSlice {
		agg, ok := r.AggregateMap[v.ExtractionKey]
		if !ok {
			continue
		}
		err := agg.AggregateOn(v.Value, v)
		if err != nil {
			return err
		}
		r.AggregateMap[v.ExtractionKey] = agg
	}
	return nil
}

type IrisRoutingResponseETL struct {
	Source        string `json:"source"`
	ExtractionKey string `json:"extractionKey"`
	DataType      string `json:"dataType"`
	Value         any    `json:"result"`
}

func (r *IrisRoutingResponseETL) ExtractKeyValue(m map[string]any) {
	if r.ExtractionKey == "" {
		r.Value = m
		return
	}
	r.Value = m[r.ExtractionKey]
	if r.Value == nil {
		return
	}
	r.DataType = reflect.TypeOf(r.Value).String()
}

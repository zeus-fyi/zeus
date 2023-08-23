package iris_programmable_proxy_v1_beta

import (
	"reflect"
	"time"

	iris_operators "github.com/zeus-fyi/zeus/pkg/iris/operators"
)

const (
	RoutingProcedureHeader = "X-Routing-Procedure"
)

// pre-canned routing procedures for QuickNode marketplace users

const (
	MaxBlockAggReduce = "max-block-agg-reduce"
)

type IrisRoutingProcedure struct {
	Name string `json:"name"`

	OrderedSteps []IrisRoutingProcedureStep `json:"steps"`
}

type BroadcastInstructions struct {
	RoutingPath  string        `json:"routingPath"`
	RestType     string        `json:"restType"`
	Payload      any           `json:"payload,omitempty"`
	MaxDuration  time.Duration `json:"maxRuntime"`
	MaxTries     int           `json:"maxTries"`
	RoutingTable string        `json:"routingTable"`
}

type IrisRoutingProcedureStep struct {
	BroadcastInstructions BroadcastInstructions                 `json:"broadcastInstructions,omitempty"`
	TransformSlice        []IrisRoutingResponseETL              `json:"transformSlice,omitempty"`
	AggregateMap          map[string]iris_operators.Aggregation `json:"aggregateMap,omitempty"`
	NextProcedure         *IrisRoutingProcedure                 `json:"nextProcedure,omitempty"`
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
	r.Value = m[r.ExtractionKey]
	r.DataType = reflect.TypeOf(r.Value).String()
}

func (r *IrisRoutingProcedure) GetNextProcedure() *IrisRoutingProcedure {
	for _, step := range r.OrderedSteps {
		if step.NextProcedure != nil {
			return step.NextProcedure
		}
	}
	return nil
}

// gets key's value from response body, then does operation on it

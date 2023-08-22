package iris_programmable_proxy_v1_beta

import "time"

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
	MaxDuration  time.Duration `json:"maxRuntime"`
	MaxTries     int           `json:"maxTries"`
	RoutingTable string        `json:"routingTable"`
}

type IrisRoutingProcedureStep struct {
	BroadcastInstructions BroadcastInstructions             `json:"broadcastInstructions,omitempty"`
	TransformMap          map[string]IrisRoutingResponseETL `json:"transformMap,omitempty"`
	AggregateMap          map[string]IrisRoutingResponseETL `json:"aggregateMap,omitempty"`
	NextProcedure         *IrisRoutingProcedure             `json:"nextProcedure,omitempty"`
}

type IrisRoutingResponseETL struct {
	//Operations []Operation `json:"operation"`
	DataType string `json:"dataType"`
	Result   any    `json:"result"`
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

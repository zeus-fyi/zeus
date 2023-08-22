package iris_programmable_proxy_v1_beta

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

type IrisRoutingProcedureStep struct {
	RoutingPath string `json:"routingPath"`
	RestType    string `json:"restType"`
	Payload     any    `json:"payload"`

	TransformMap  map[string]IrisRoutingResponseETL `json:"transformMap"`
	AggregateMap  map[string]IrisRoutingResponseETL `json:"aggregateMap"`
	NextProcedure IrisRoutingProcedure              `json:"nextProcedure"`
}

type IrisRoutingResponseETL struct {
	//Operations []Operation `json:"operation"`
	DataType string `json:"dataType"`
	Result   any    `json:"result"`
}

// gets key's value from response body, then does operation on it

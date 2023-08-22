```go
type IrisRoutingProcedureStep struct {
RoutingPath string `json:"routingPath"`
RestType    string `json:"restType"`
Payload     any    `json:"payload"`

TransformMap  map[string]IrisRoutingResponseETL `json:"transformMap"`
AggregateMap  map[string]IrisRoutingResponseETL `json:"aggregateMap"`
NextProcedure IrisRoutingProcedure              `json:"nextProcedure"`
}
```

Order of ETL Operations:

1. the keys of the TransformMap will run the ETL operations in the order they are listed and the results will be passed
   to the next step
2. the keys of the AggregateMap will run the ETL operations in the order they are listed

```go

TransformMap  map[string]IrisRoutingResponseETL `json:"transformMap"`


```
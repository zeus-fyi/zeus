### Routing Procedures

Routing Procedures are a way to create a stored procedure function that can be called by the router to perform
ETL and aggregation operations on the results of the broadcast requests with only a single request to the router.

This feature is currently in beta, and is subject to change

```go
type IrisRoutingProcedureStep struct {
BroadcastInstructions BroadcastInstructions                   `json:"broadcastInstructions,omitempty"`
TransformSlice        []iris_operators.IrisRoutingResponseETL `json:"transformSlice,omitempty"`
AggregateMap          map[string]iris_operators.Aggregation   `json:"aggregateMap,omitempty"`
}
```

Order of ETL Operations:

1. the elements in the TransformSlice will run the ETL operations in the order they are listed and the results will be
   passed if applicable
   to the next step, eg. if the filter is a max aggregate, it will pass the max endpoints to the next step, and discard
   the rest
2. the elements in the AggregateMap are then passed to the next step, eg. if the aggregate map has a max aggregate,
   it will pass the endpoints meeting the max value reference point to the next step, and discard the rest

```go
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
FanInRuleFirstValidResponse = "returnOnFirstSuccess"
)
```

Only a limited number of broadcast instructions are supported at this time.

The `FanInRules` struct is used to specify:

1. how to handle the responses from the broadcast requests, "returnOnFirstSuccess" will return the first
   successful response, meaning a status code less than 400.

Dynamic Routing

This is a feature that allows you to dynamically perform ETL and aggregation window filtering on route requests from
headers in the request. This is useful for when you want to perform different ETL and aggregation operations on the
as if it were a stored procedure function without having to create a stored procedure function.

This example shows you how to filter out endpoints that are not the max block number, and then only send the
eth_getBlockByNumber request to the filtered endpoints, and returns the first successful response containing
the full block payload at the max seen block number.

```shell
curl --location 'https://iris.zeus.fyi/v1/router' \
--header 'X-Load-Balancing-Strategy: Adaptive' \
--header 'X-Adaptive-Metrics-Key: Ethereum' \
--header 'X-Route-Group: ethereum-mainnet' \
--header 'X-Routing-Procedure: <any>' \
--header 'X-Agg-Op: max' \
--header 'X-Agg-Key: result' \
--header 'X-Agg-Filter-Fan-In: returnOnFirstSuccess' \
--header 'X-Agg-Key-Value-Data-Type: int' \
--header 'X-Agg-Filter-Payload: {"jsonrpc":"2.0","method":"eth_getBlockByNumber","params":["latest", true],"id":1}' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer <YOUR-BEARER-TOKEN>' \
--data '{
    "jsonrpc": "2.0",
    "method": "eth_blockNumber",
    "params": [],
    "id": 1
}'
```

### Initial Request

#### --data payload

- the data payload will be sent to the table endpoints specified in the X-Route-Group header, and will be used for the
  first broadcast request

### Headers

#### X-Routing-Procedure [any(string)]:

- set this value with any non-empty key to enable dynamic routing

#### X-Agg-Op [max(int, float64, hexstr(int))]:

- set this value to the aggregation operation you want to perform on the results of the broadcast requests
- only the max aggregation operation is officially supported at this time
- unofficially, you can set this value to any of the aggregation operations supported in [pkg/iris/operators]
- it converts hexstr to int as long as you set the X-Agg-Key-Value-Data-Type header to int

#### X-Agg-Key-Value-Data-Type [int, float64]:

- set this value to the data type of the key-value you want to aggregate on

#### X-Agg-Key [key(string)]:

- set this value to the key-value you want to aggregate on
- for ethereum, eth_blockNumber returns a map with blockNumber(hexstr) = map[result], so you can set this value to "
  result" to aggregate on the block number

#### X-Agg-Filter-Fan-In [returnOnFirstSuccess(string)]:

- set this value to the fan-in rule you want to apply to the results of the broadcast requests
- only the returnOnFirstSuccess fan-in rule is officially supported at this time, which will return the first successful
  response, meaning a status code less than 400

#### X-Agg-Filter-Payload [payload(any, map[string]interface{})]:

- set this value to the payload you want to send with to the filtered aggregate routes
- e.g. for ethereum: {"jsonrpc":"2.0","method":"eth_getBlockByNumber","params":["latest", true],"id":1}
- in the example, this is because I'm using eth_blockNumber to get the latest block number, and then I aggregate only
  endpoints that have the max block number found, and then only sending the eth_getBlockByNumber to the filtered
  endpoints

#### Upcoming Features:

- support for more creating & using stored procedures
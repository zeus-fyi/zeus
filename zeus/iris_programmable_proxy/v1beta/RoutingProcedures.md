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
   passed to the aggregates or otherwise next processing step
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
    FanInRuleReturnAllResponses = "returnAllSuccessful"
)

// ReturnFirstResultOnSuccess returns the first result from the fan-in that is not an error
func (b BroadcastRules) ReturnFirstResultOnSuccess() string {
    return FanInRuleFirstValidResponse
}

// ReturnResultsOnSuccess returns all results from the fan-in that are not errors that complete before any timeouts occur
// this is the default behavior
func (b BroadcastRules) ReturnResultsOnSuccess() string {
  return FanInRuleReturnAllResponses
}

type IrisRoutingProcedureStep struct {
  BroadcastInstructions BroadcastInstructions                   `json:"broadcastInstructions,omitempty"`
  TransformSlice        []iris_operators.IrisRoutingResponseETL `json:"transformSlice,omitempty"`
  AggregateMap          map[string]iris_operators.Aggregation   `json:"aggregateMap,omitempty"`
}

// This is the pattern of expected order processing you should expect during execution
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
```

Only a limited number of broadcast instructions are supported at this time.

The `FanInRules` struct is used to specify:

1. how to handle the responses from the broadcast requests, "returnOnFirstSuccess" will return the first
   successful response, meaning a status code less than 400.

### Dynamic ETL Function Routing

This is a feature that allows you to dynamically perform ETL and aggregation window filtering on route requests using only
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
--header 'X-Agg-Key-Value-Data-Type: int' \
--header 'X-Agg-Filter-Fan-In: returnOnFirstSuccess' \
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

### Key Concepts 

The Dynamic ETL function is purposely limited to two procedure steps, since it makes much more sense to use durable execution functions, or message streaming for long-running and/or more complex multi-stage, multi-state procedures, which
you'll be able to use in a later release.

These headers aren't strictly required

```
--header 'X-Load-Balancing-Strategy: Adaptive' \
--header 'X-Adaptive-Metrics-Key: Ethereum' \
```

This header is required. Currently you can set it to any value, other than `'max-block-agg-reduce'` which is currently a reserved keyword procedure, for it to take effect. 
```
--header 'X-Routing-Procedure: <any>' \
```

All the endpoints you specify requests to be sent to, either directly via a routing table lookup, or those left after an aggregation window will have their requests made in concurrent go-routines to minimize latency & max throughput speed.

You can use one or two stages that work like this:

#### Stage One (X-Agg prefixed)

Only one transformation & one aggregation in stage one is officially supported while in v1Beta. Unofficially, you should
be able to chain arbitrary amounts within reasonable boundaries, eg don't use 100 stages, but if you need 1-10 and
expect your procedure to complete within 60s or so, then you should be fine. Additionally, only `POST` requests are
officially supported at this time, but unofficially you can use any HTTP method that Iris
supports: `GET, POST, PUT, DELETE`

you provide:
  1. The routing group table header you want to send requests
  
  `--header 'X-Route-Group: ethereum-mainnet'`

  2. The group aggregation technique header.

  `--header 'X-Agg-Op: max'`

  3. The key that should be used to extract values to use the aggregate function on.

  `--header 'X-Agg-Key: result'`

  3a. Your initial data payload is what's sent to the initial routing group table

  `--data payload`

  in our example we use this initial payload

  Request
```json
{
    "jsonrpc": "2.0",
    "method": "eth_blockNumber",
    "params": [],
    "id": 1
}
```

  We use the `result` key to extract `"0x112927b"`
 
  Response
```json
{
    "id": 1,
    "jsonrpc": "2.0",
    "result": "0x112927b"
}
```

Effectively, this is creating this struct for you, and is setting the extraction key to `result`

```go
type IrisRoutingResponseETL struct {
	Source        string `json:"source"`
	ExtractionKey string `json:"extractionKey"`
	DataType      string `json:"dataType"`
	Value         any    `json:"result"`
}
```
When a `map[string]any` payload is seen, it will run this extraction method using the key you specify. During the extraction stage, the source value is also set to the endpoint where the data was retrieved from.

```go
func (r *IrisRoutingResponseETL) ExtractKeyValue(m map[string]any) {
  r.ExtractionKey = "result"
	r.Value = m[r.ExtractionKey]
	if r.Value == nil {
		return
	}
	r.DataType = reflect.TypeOf(r.Value).String()
}

  // r.Value now is = 0x112927b in our example
```
  3b. Your requests are sent in parallel goroutines which then aggregate the max int value from the extraction key you provide.

   `--header 'X-Agg-Key-Value-Data-Type: int'`

  In the example, we specify that we want the extracted value from result: "0x112927b", to be treated as an int, since the aggregator is able to convert string decimals and hexstr numeric values that are "0x" prefixed without 
  having to do any additional steps. 

```go
func (a *Aggregation) AggregateMaxInt(x int, y IrisRoutingResponseETL) error {
	a.Operator = Max
	if len(a.DataSlice) == 0 || x >= a.CurrentMaxInt {
		if x > a.CurrentMaxInt {
			a.CurrentMaxInt = x
			a.CurrentMinInt = x
			a.DataSlice = []IrisRoutingResponseETL{y} // Keep only the new maximum value
		} else {
			a.DataSlice = append(a.DataSlice, y) // Append the value if it's equal to the current maximum
		}
	}
	return nil
}
```
In the example shown when the aggregate for the first stage completes, it will contain a data slice with all the IrisRoutingResponseETL responses from endpoints with max seen endpoints.

eg. if you had 3 endpoints in your routing table, and they had the following responses

  `e1 - blockHeight - 100`
  `e2 - blockHeight - 101`
  `e3 - blockHeight - 101`

After the aggregate function completes, you will have your endpoints reduced to [e2, e3], which can then be used for the next procedure stage.

#### Stage Two (X-Agg-Filter prefixed)

This stage is for the final request & response formatting and it has access to the results of the first stage for
continued processing steps. ETL transforms
are currently not officially supported, and only fan-in: returnOnFirstSuccess is officially supported at this time
during v1 beta. Unofficially, you can use any of the fan-in rules supported in [pkg/iris/fanin]

  1.  In our example we use the following header for the final response. This effectively means that we'll send a request to any remaining endpoints that meet the max block number aggregation threshold from the previous stage.
      Since nodes that have the same block height should have the same expected result for the block value, we'll just return the first successful response. The default behavior is to process all results and forward all successful ones.

`--header 'X-Agg-Filter-Fan-In: returnOnFirstSuccess' \`

  2. Since we want to send a different payload to the aggregated endpoints, you can specify the payload in the header directly like below. Larger payloads will require created a stored procedure. More details on creating these coming in a later release.

`--header 'X-Agg-Filter-Payload: {"jsonrpc":"2.0","method":"eth_getBlockByNumber","params":["latest", true],"id":1}' \`

#### Summary

So in our example this is what gets executed:

Stage One:
  - we concurrently send the eth_blockNumber request to all the endpoints in our table, by default it will aggregate all the successful response
  - since we specify a max aggregate on the block nubmer value, our final result are the endpoints at the max block number seen.

Stage Two:
  - we concurrently send the eth_getBlockByNumber request to all the endpoints we aggregated in stage one, and we return the first successful response we see, meaning status code 2xx.

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
  result" to aggregate on the block number.
- nested key-values are not officially supported.

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
- this is will default to a POST request if provided and not empty, and other REST types are not officially supported at this time 

#### Upcoming Features:

- support for more creating & using stored procedures

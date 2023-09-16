---
sidebar_position: 1
---

# Adaptive Load Balancer

#### Exclusively running Iris services through the QuickNode Marketplace until our v1 release later this year.

## Overview

We use ZU to denote Zeus compute units.
These units are derived from server bandwidth, cpu, memory usage and average and peak traffic usage, and operational
costs for development & maintenance.

<b>4 ZU</b> per request & response.<br/>
<b>1 ZU</b> per 1 kB<br/><br/>
<b>10 ZU minimum for round trip request </b><br/><br/>

Each user can store up to 1000 endpoints for free, you can then use these endpoints to create routing group tables via
API or UI Dashboard

## QuickNode Marketplace Users

QuickNode marketplace customers will use an SSO link from QuickNode and be directed to their load balancing stored
endpoints
table, you'll then be able to generate an API key from the access panel.

### Additional Benefits:

QuickNode users will have their endpoints automatically registered with the load balancer, and will have automatically
generated routing group tables based on the network-chain type for that endpoint. E.g. ethereum mainnet endpoints will
be
automatically registered to the ethereum-mainnet routing group table. These won't count against your
routing table limits unless you make any manual changes to them.

### Lite

```text
250M ZU per $99
Up to 25k ZU/s ~ 2.5k req/s
Up to ~25M requests (with responses)
25 Custom Routing Group Tables
Round Robin 
```

### Standard

```text
1B ZU per $299
Up to 50k ZU/s, ~5k req/s
Up to ~ 100M  requests (with responses)
50 Custom Routing Group Tables
Embedded Routing Procedures
Priority Score Weight Tuning
Adaptive Load Balancing + Round Robin
```

### Performance

```text
3B ZU per $999, 
Up to 100k ZU/s, ~10k req/s
Up to ~300M requests (with responses)
250 Custom Routing Group Tables
Embedded Routing Procedures
Priority Score Weight Tuning
Adaptive Load Balancing + Round Robin
```

Need more? Send us a message at support@zeus.fyi

## Setup

This is a guide to help you set up your own programmable proxy for the Iris Load Balancer.
Prerequisites: You'll need to generate an API key from the access panel if you don't have an existing one.

### Part A. Configuration Setup: Registering your endpoints & routing tables

You'll use the Hestia endpoint to make any configuration changes to your routing groups. You'll have a separate one
to use for the actual load balancer.

```go
const HestiaEndpoint = "https://hestia.zeus.fyi"

const (
EthereumValidatorsCreateServiceRequestV1Path = "/v1/validators/service/create"

IrisCreateRoutesPath = "/v1/iris/routes/create"

IrisCreateProcedurePath = "/v1/iris/procedure/create"

IrisReadRoutesPath = "/v1/iris/routes/read"
IrisDeleteRoutesPath = "/v1/iris/routes/delete"

IrisReadGroupRoutesPath = "/v1/iris/routes/group/:groupName/read"
IrisUpdateGroupRoutesPath = "/v1/iris/routes/group/:groupName/update"

IrisCreateGroupRoutesPath = "/v1/iris/routes/groups/create"
IrisReadGroupsRoutesPath = "/v1/iris/routes/groups/read"
IrisDeleteGroupRoutesPath = "/v1/iris/routes/groups/delete"
)
```

Source: pkg/hestia/client/endpoints/endpoints.go

```text    
    // POST request to register new endpoints
    const HestiaEndpoint = "https://hestia.zeus.fyi/v1/iris/routes/create"
    
    // POST request: to create a routing group from a list of stored endpoints
    const IrisCreateGroupRoutesPath = "https://hestia.zeus.fyi/v1/iris/routes/groups/create"
```

### Step One: Register new endpoints

Note that only https routes are supported, http routes will be ignored.

POST request to register new endpoints

```text    
    const HestiaEndpoint = "https://hestia.zeus.fyi/v1/iris/routes/create"    
```

### Step One Payload Example:

```json
{
  "routes": [
    "https://alarmingly-bitter-lambos.quiknode.pro/743c191e-31b5-11ee-be56-0242ac120002/",
    "https://shockingly-bitter-lambos.quiknode.pro/743c191e-31b5-11ee-be56-0242ac120003/"
  ]
}
```

### Step One Curl Example:

```sh
curl --location 'https://hestia.zeus.fyi/v1/iris/routes/create' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer YOUR_BEARER_TOKEN' \
--data '{
  "routes": [
    "https://alarmingly-bitter-lambos.quiknode.pro/743c191e-31b5-11ee-be56-0242ac120002/",
    "https://shockingly-bitter-lambos.quiknode.pro/743c191e-31b5-11ee-be56-0242ac120003/"
  ]
}'
```

### Step Two: Register a routing group table from saved endpoints

POST request: to create a routing group from a list of stored endpoints

```text    
    const IrisCreateGroupRoutesPath = "https://hestia.zeus.fyi/v1/iris/routes/groups/create"
```

### Step Two Payload Example:

```json
{
  "groupName": "quicknode-mainnet",
  "routes": [
    "https://alarmingly-bitter-lambos.quiknode.pro/743c191e-31b5-11ee-be56-0242ac120002/",
    "https://shockingly-bitter-lambos.quiknode.pro/743c191e-31b5-11ee-be56-0242ac120003/"
  ]
}
```

### Step Two Curl Example:

```sh
curl --location 'https://hestia.zeus.fyi/v1/iris/routes/groups/create' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer YOUR_BEARER_TOKEN' \
--data '{
  "groupName": "quicknode-mainnet",
  "routes": [
    "https://alarmingly-bitter-lambos.quiknode.pro/743c191e-31b5-11ee-be56-0242ac120002/",
    "https://shockingly-bitter-lambos.quiknode.pro/743c191e-31b5-11ee-be56-0242ac120003/"
  ]
}'
```

### Part B. Using the Programmable Proxy

You'll use the API bearer token that you generate from the Access panel to authenticate with the load balancer. </b
You then use the name of your route table group as a query parameter like the below,
and it will round-robin the requests between the endpoints in that group table.

```go
package iris_programmable_proxy

const IrisEndpoint = "https://iris.zeus.fyi"


IrisClientProd = Iris{
resty_base.GetBaseRestyClient("https://iris.zeus.fyi", tc.Bearer),
}

routeGroup := "quicknode-mainnet"

Add HEADER "X-Route-Group" with value "quicknode-mainnet"
path := "https://iris.zeus.fyi/v1/router"
```

### Curl Example:

```shell
curl --location 'https://iris.zeus.fyi/v1/router' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer YOUR-BEARER-TOKEN' \
--header 'X-Route-Group: quicknode-mainnet' \
--data '{"jsonrpc":"2.0","method":"eth_getBlockByNumber","params":["latest", true],"id":1}'
```

### Part C. Using the Adaptive Load Balancer

```go
package iris_programmable_proxy


routeGroup := "quicknode-mainnet"

Add HEADER "X-Route-Group" with value "quicknode-mainnet"

Add HEADER "X-Load-Balancing-Strategy" with value "Adaptive"
Add HEADER "X-Adaptive-Metrics-Key" with value "JSON-RPC" (or "Other Metric Keys...")
- JSON-RPC is a reserved value for json-rpc based POST apis, it collects stats by the method value in the json rpc POST request

/*
   Once you have ~20 or so request samples for the same method, the load balancer will start to use the adaptive strategy automatically and
   manage the routing group table for you based on the best predicted performing endpoint for that method that's available.

      Stats will only persist for one hour since the last API call for that method, so you'll need to keep making requests to keep the stats.
   It doesn't take long, only ~20 samples per metric to trend towards a near optimal routing group table from scratch, so it's really not a big deal to reset the stats.
*/
path := "https://iris.zeus.fyi/v1/router"
*/
```

### Curl Example:

```sh
curl --location ‘https://iris.zeus.fyi/v1/router’ \
--header ‘Content-Type: application/json’ \
--header ‘Authorization: Bearer YOUR-BEARER-TOKEN’ \
--header ‘X-Route-Group: ethereum-mainnet’ \
--header ‘X-Load-Balancing-Strategy: Adaptive’ \
--header ‘X-Adaptive-Metrics-Key: JSON-RPC’ \
--data ‘{“jsonrpc”:“2.0”,“method”:“eth_getBlockByNumber”,“params”:[“latest”, true],“id”:1}’
```

You can also check out our round-robin load_balancing_test.go for an example of how to use the programmable proxy to
query
the block number from a routing group of ethereum node urls endpoints.

### Part D. Using Procedures

To use embedded protocol procedures you only need to add the key value to your payload. In this example, to use the max
block procedure for Ethereum, which polls your routing table for the current block number, and then forwards your
request to the endpoints returning the highest block number seen and then returns the first successful response.

"procedure": "eth_maxBlockAggReduce"

```
curl --location 'https://iris.zeus.fyi/v1/router' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer YOUR-BEARER-TOKEN' \
--header 'X-Route-Group: ethereum-mainnet' \
--data '{
    "jsonrpc": "2.0",
    "procedure": "eth_maxBlockAggReduce",
    "method": "eth_getBlockByNumber",
    "params": ["latest", true],
    "id": 1
}'
```
# Iris Load Balancer Programmable Proxy: QuickNode Marketplace

#### Exclusively running Iris services through the QuickNode Marketplace until our v1 release later this year.


## Overview

We use ZU to denote Zeus compute units. These derived units are composed from server bandwidth, cpu, memory usage and average and peak traffic usage, and operational costs for development & maintenance. 

4 ZU per request & response. 
1 ZU per 1 kB, minimum 1 ZU per req/resp

Each user can store up to 1000 endpoints for free, you can then use these endpoints to create routing group tables via API or UI Dashboard

10 ZU minimum size for round trip request & response, up to 1kB payload request and up to 1kB response

### Free
```text
50M ZU
Up to 1k ZU/s ~ 100 req/s
Up to ~5M requests (with responses)
1 Routing Group Table
```
### Standard
```text
1B ZU per $299
Up to 25k ZU/s, ~2.5k req/s
Up to ~ 100M  requests (with responses)
50 Routing Group Tables
```
### Performance
```text
3B ZU per $999, 
Up to 50k ZU/s, ~5k req/s
Up to ~300M requests (with responses)
250 Routing Group Tables
```
Need more? Send us a message at support@zeus.fyi

## Setup

This is a guide to help you set up your own programmable proxy for the Iris Load Balancer.

Prerequisites: You'll need to generate an API key from the access panel if you don't have an existing one.

QuickNode marketplace customers will use an SSO link from QuickNode and be directed to their load balancing stored endpoints 
table, you'll then be able to generate an API key from the access panel.

![accessPanel](https://github.com/zeus-fyi/zeus/assets/17446735/c54a01e0-91fa-48a0-9fba-ff55050848eb)

![Screenshot 2023-08-02 at 10 28 22 PM](https://github.com/zeus-fyi/zeus/assets/17446735/5e61cb1b-f051-408d-8964-82c6835c11f4)
Part A. Configuration Setup Complete

```go
/*
    You'll use the Hestia endpoint to make any configuration changes to your routing groups. You'll have a separate one
    to use for the actual load balancer.
 */
const HestiaEndpoint = "https://hestia.zeus.fyi"
```

Complete list of endpoints: pkg/hestia/client/endpoints/endpoints.go
```text    
    // POST request to register new endpoints
    const HestiaEndpoint = "https://hestia.zeus.fyi/v1/iris/routes/create"
    
    // POST request: to create a routing group from a list of stored endpoints
    const IrisCreateGroupRoutesPath = "https://hestia.zeus.fyi/v1/iris/routes/groups/create"
```

Step One: Register new endpoints
```text    
    // POST request to register new endpoints
    const HestiaEndpoint = "https://hestia.zeus.fyi/v1/iris/routes/create"    
```

Step One Payload Example:
```json
{
  "routes": [
    "https://alarmingly-bitter-lambos.quiknode.pro/743c191e-31b5-11ee-be56-0242ac120002/",
    "https://eth-mainnet.g.alchemy.com/v2/cdVqiD1o29wcb558mc9g74c2l"
  ]
}
```
Step One Curl Example:
```sh
curl --location 'https://hestia.zeus.fyi/v1/iris/routes/create' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer YOUR_BEARER_TOKEN' \
--data '{
  "routes": [
    "https://alarmingly-bitter-lambos.quiknode.pro/743c191e-31b5-11ee-be56-0242ac120002/",
    "https://eth-mainnet.g.alchemy.com/v2/cdVqiD1o29wcb558mc9g74c2l"
  ]
}'
```
Step Two: Register a routing group table from saved endpoints

```text    
    // POST request: to create a routing group from a list of stored endpoints
    const IrisCreateGroupRoutesPath = "https://hestia.zeus.fyi/v1/iris/routes/groups/create"
```

Step Two Payload Example:
```json
{
  "groupName": "quicknode-mainnet",
  "routes": [
    "https://alarmingly-bitter-lambos.quiknode.pro/743c191e-31b5-11ee-be56-0242ac120002/",
    "https://eth-mainnet.g.alchemy.com/v2/cdVqiD1o29wcb558mc9g74c2l"
  ]
}
```
Step Two Curl Example:
```sh
curl --location 'https://hestia.zeus.fyi/v1/iris/routes/groups/create' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer YOUR_BEARER_TOKEN' \
--data '{
  "groupName": "quicknode-mainnet",
  "routes": [
    "https://alarmingly-bitter-lambos.quiknode.pro/743c191e-31b5-11ee-be56-0242ac120002/",
    "https://eth-mainnet.g.alchemy.com/v2/cdVqiD1o29wcb558mc9g74c2l"
  ]
}'
```
Part B. Using the Programmable Proxy

```go
package iris_programmable_proxy

const IrisEndpoint = "https://iris.zeus.fyi"

/*
    You'll use the API bearer token that you generate from the Access panel to authenticate with the load balancer.
 */

IrisClientProd = Iris{
    resty_base.GetBaseRestyClient("https://iris.zeus.fyi", tc.Bearer),
}

/*
    You then use the name of your route table group as a query parameter like the below,
    and it will round-robin the requests between the endpoints in that group table. 
 */

routeGroup := "quicknode-mainnet"
path := fmt.Sprintf("https://iris.zeus.fyi/v1/router/group?routeGroup=%s", routeGroup)

// e.g. https://iris.zeus.fyi/v1/router/group?routeGroup=quicknode-mainnet
```

You can also check out our round-robin load_balancing_test.go for an example of how to use the programmable proxy to query 
the block number from a routing group of ethereum node urls endpoints.


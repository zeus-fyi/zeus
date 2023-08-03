Iris Load Balancer Programmable Proxy: QuickNode Marketplace Setup

Exclusively running Iris services through the QuickNode Marketplace until our v1 release later this year.

This is a guide to help you set up your own programmable proxy for the Iris Load Balancer.

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

/*
   You can see our round-robin load_balancing_test.go for an example of how to use the programmable proxy to query 
    the block number from a routing group of ethereum node urls endpoints.
*/
```
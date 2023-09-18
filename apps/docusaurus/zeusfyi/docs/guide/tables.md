---
sidebar_position: 1
---

# Tables

### Registering your endpoints & routing tables

You'll use the Hestia endpoint to make any configuration changes to your routing groups. You'll have a separate one
to use for the actual load balancer.

```go
const (
HestiaEndpoint = "https://hestia.zeus.fyi"

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

POST request to register new endpoints

```go    
const HestiaEndpoint = "https://hestia.zeus.fyi/v1/iris/routes/create"
```

POST request: to create a routing group from a list of stored endpoints

```go
const IrisCreateGroupRoutesPath = "https://hestia.zeus.fyi/v1/iris/routes/groups/create"
```

### Step One: Register new endpoints

Note that only https routes are supported, http routes will be ignored.

POST request to register new endpoints

```go    
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
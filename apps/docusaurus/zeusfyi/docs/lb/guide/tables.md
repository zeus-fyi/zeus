---
sidebar_position: 1
---

# Tables

### Adding routes via the dashboard

![addendpoints](https://github.com/zeus-fyi/zeus/assets/17446735/98eb761a-9ab3-4eff-9c2a-65c9c957f13f)

### Adding new tables via the dashboard

You must select at least 1 route when creating a new table, otherwise it will not be created. When a table contains
no endpoints it is automatically deleted.

![addgroups](https://github.com/zeus-fyi/zeus/assets/17446735/a499b3f2-de4d-4ec8-8a5e-652412dea4f7)

### Adding existing routes to a table via the dashboard

When you are in the table view, meaning you have selected a table from the list of tables (e.g. not the all, or unused
group),
you can add existing routes to the table by selecting the Add Group Endpoints button, and then it will toggle
the table view to all available routes that are not already in the table. You can then select the routes you want to add
to the table, and then press Add for it to take effect. Refresh the page to see the changes.

![addgroupendpoints](https://github.com/zeus-fyi/zeus/assets/17446735/cd95f2c7-a819-434c-b2bb-7746b264b40d)

![addgroupendpointspt2](https://github.com/zeus-fyi/zeus/assets/17446735/b3ca33d5-d731-4ace-b10c-fe3954913b44)

### Removing existing table routes via the dashboard

When you select table endpoints while in the default view, you'll be given the option to delete routes.

![deletetableroutes](https://github.com/zeus-fyi/zeus/assets/17446735/9c621ee9-26f4-4074-96d3-8758276f4cd9)

### Removing existing tables via the dashboard

When you delete all routes in a given table, it will also delete the table.

### Registering your endpoints & routing tables via API

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
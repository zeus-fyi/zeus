---
sidebar_position: 1
displayed_sidebar: zK8s
---

# zK8s Client #

The zK8s client is used for interacting with our cloud platform's core infrastructure apis. You can override the values
in the test files
we have and set to your own following these steps

1. In /test/configs -> create a config.yaml using the sample-config.yaml as a reference, the config.yaml should be in
   .gitignore by default so it doesn't commit your tokens
2. Add your bearer token to this config, otherwise config it directly in the client
3. Override the zeus_client_test variables that are used to point to your desired chart and kubernetes location
4. Then run the test, the zeus_client_integrated_test will upload and then query for the uploaded chart and then print
   it.

## Table of Contents

1. [Imports](#imports)
2. [Data Structures](#data-structures)
   - [ZeusClient](#zeusclient)
3. [Functions](#functions)
   - [NewZeusClient](#newzeusclient)
   - [NewDefaultZeusClient](#newdefaultzeusclient)

---

## Imports

```go
import resty_base "github.com/zeus-fyi/zeus/zeus/z_client/base"
```

---

## Data Structures

### `ZeusClient`

Provides a client for interacting with Zeus, wrapping the base resty client.

**Fields:**

- `Resty`: The underlying REST client.

---

## Functions

### `NewZeusClient(baseURL, bearer string) ZeusClient`

Creates a new Zeus client.

**Parameters:**

- `baseURL`: The base URL for the Zeus service.
- `bearer`: The bearer token for authentication.

**Returns:**

- A new `ZeusClient` instance.

### `NewDefaultZeusClient(bearer string) ZeusClient`

Creates a new default Zeus client with the predefined Zeus endpoint.

**Parameters:**

- `bearer`: The bearer token for authentication.

**Returns:**

- A new `ZeusClient` instance.

---

## Base API Endpoints Overview

### Upload Infrastructure

```go
// InfraCreateV1Path uploads and saves a kubernetes app workload
const InfraCreateV1Path = "/v1/infra/create"
```

You're limited to one service, one config map, one ingress, and one of either a stateful set or a deployment per
infrastructure topology. If you need more resources just create another topology and deploy both of them to the same
namespace location. Later on you'll be able to create topology classes where you can append these groups into a single
higher level topology.

```go
type TopologyBaseInfraWorkload struct {
*v1core.Service       `json:"service"`
*v1core.ConfigMap     `json:"configMap"`
*v1.Deployment        `json:"deployment"` // Only 1 StatefulSet, or 1 Deployment, not both
*v1.StatefulSet       `json:"statefulSet"` // Only 1 StatefulSet, or 1 Deployment, not both
*v1networking.Ingress `json:"ingress"`
}
```

See example in chart_upload_test.go
Gzips the k8s workload, fills out the form params, and uploads via API

### Read Infrastructure Chart

```go
// InfraReadChartV1Path reads the chart workload you uploaded
const InfraReadChartV1Path = "/v1/infra/read/chart"

// Request 
type TopologyRequest struct {
TopologyID int `json:"topologyID"`
}

// Response 
type TopologyCreateResponse struct {
TopologyID int `json:"topologyID"`
}

```

See example in read_chart_test.go
Query this endpoint to read the stored k8s workload associated with its topology id.

### Read High Level Metadata for Uploaded Topologies

```go
// InfraReadTopologyV1Path reads the metadata for your uploaded topologies
const InfraReadTopologyV1Path = "/v1/infra/read/topologies"

// Response
type ReadTopologiesMetadata struct {
TopologyID       int            `json:"topologyID"`
TopologyName     string         `json:"topologyName"`
ChartName        string         `json:"chartName"`
ChartVersion     string         `json:"chartVersion"`
ChartDescription sql.NullString `json:"chartDescription"`
}

type ReadTopologiesMetadataGroup struct {
Slice []ReadTopologiesMetadata
}
```

Example Response

```json
  [
  {
    "topology_id": 1668066250334934000,
    "topology_name": "demo",
    "chart_name": "demo",
    "chart_version": "v0.0.1668066250013676081",
    "chart_description": {
      "String": "",
      "Valid": false
    }
  },
  {
    "topology_id": 1668062631385564001,
    "topology_name": "demo topology",
    "chart_name": "demo chart",
    "chart_version": "v0.0.1668062631136840081",
    "chart_description": {
      "String": "",
      "Valid": false
    }
  }
]
```

### Deploys Topology

```go
// DeployTopologyV1Path deploys topology
const DeployTopologyV1Path = "/v1/deploy"

// Location for API to send request
type CloudCtxNs struct {
CloudProvider string `json:"cloudProvider"`
Region        string `json:"region"`
Context       string `json:"context"`
Namespace     string `json:"namespace"`
Env           string `json:"env"`
}

// Request Struct
type TopologyDeployRequest struct {
TopologyID int `json:"topologyID"`
zeus_common_types.CloudCtxNs
}
```

Post to this endpoint to deploy this infrastructure topology

### Destroy Topology

```go
// DestroyDeployInfraV1Path destroys topology, in other words uninstalls the app
const DestroyDeployInfraV1Path = "/v1/deploy/destroy"

// Location for API to send request
type CloudCtxNs struct {
CloudProvider string `json:"cloudProvider"`
Region        string `json:"region"`
Context       string `json:"context"`
Namespace     string `json:"namespace"`
Env           string `json:"env"`
}

// Request Struct
type TopologyDeployRequest struct {
TopologyID int `json:"topologyID"`
zeus_common_types.CloudCtxNs
}
```

Post to this endpoint to destroy this infrastructure topology.

### Read Live Namespace Deployed Kubernetes Workloads

```go
// ReadWorkload reads all the statefulsets, services, ingresses, deployments, configmaps, and pods in a namespace.
const ReadWorkload = "/v1/workload/read"

// req type
type TopologyCloudCtxNsQueryRequest struct {
zeus_common_types.CloudCtxNs
}

// response type
type NamespaceWorkload struct {
*v1.PodList               `json:"podList"`
*v1.ServiceList           `json:"serviceList"`
*v1networking.IngressList `json:"ingressList"`
*v1apps.StatefulSetList   `json:"statefulSetList"`
*v1apps.DeploymentList    `json:"deploymentList"`
*v1.ConfigMapList         `json:"configMapList"`
}
```

See example in read_live_namespace_workload.go

### Read Deployment Status Updates

```go
// DeployStatusV1Path gets the topology deployment status updates

const DeployStatusV1Path = "/v1/deploy/status"

// Request
type TopologyRequest struct {
TopologyID int `json:"topologyID"`
}

// Response
type TopologyDeployStatuses struct {
Slice []TopologyDeployStatus
}

type TopologyDeployStatus struct {
TopologyID     int       `json:"topologyID"`
TopologyName   string    `json:"topologyName"`
TopologyStatus string    `json:"topologyStatus"`
UpdatedAt      time.Time `json:"updatedAt"`
CloudCtxNs     zeus_common_types.CloudCtxNs
}
```

See example in deploy_status_test.go




## API Endpoints

#### Upload Infrastructure

```go
// InfraCreateV1Path uploads and saves a kubernetes app workload
const InfraCreateV1Path = "/v1/infra/create"
```

See example in chart_upload_test.go
Gzips the k8s workload, fills out the form params, and uploads via API

#### Read Infrastructure Chart

```go
// InfraReadChartV1Path reads the chart workload you uploaded
const InfraReadChartV1Path = "/v1/infra/read/chart"

// Request 
type TopologyRequest struct {
    TopologyID int `json:"topologyID"`
}

// Response 
type TopologyCreateResponse struct {
    TopologyID int `db:"topology_id" json:"topologyID"`
}

```

See example in read_chart_test.go
Query this endpoint to read the stored k8s workload associated with its topology id.

#### Read High Level Metadata for Uploaded Topologies
```go
// InfraReadTopologyV1Path reads the metadata for your uploaded topologies
const InfraReadTopologyV1Path = "/v1/infra/read/topologies"

// Response
type ReadTopologiesMetadata struct {
    TopologyID       int            `db:"topology_id" json:"topologyID"`
    TopologyName     string         `db:"topology_name" json:"topologyName"`
    ChartName        string         `db:"chart_name" json:"chartName"`
    ChartVersion     string         `db:"chart_version" json:"chartVersion"`
    ChartDescription sql.NullString `db:"chart_description" json:"chartDescription"`
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

#### Deploys Topology

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

#### Destroy Topology

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

#### Read Live Namespace Deployed Kubernetes Workloads
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

#### Read Deployment Status Updates
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




## API Endpoints

#### Upload Infrastructure

```go
// InfraCreateV1Path uploads and saves a kubernetes app workload
const InfraCreateV1Path = "/v1/infra/create"
```

See example in upload_chart.go

Gzips the k8s workload, fills out the form params, and uploads via API

#### Read Infrastructure Chart

```go
// InfraReadChartV1Path reads the chart workload you uploaded
const InfraReadChartV1Path = "/v1/infra/read/chart"

// Request Object
type TopologyReadRequest struct {
    TopologyID int `json:"topologyID"`
}
```

See example in read_chart.go

Query this endpoint to read the stored k8s workload associated with its topology id.

#### Read High Level Metadata for Uploaded Topologies

```go
// InfraReadTopologyV1Path reads the metadata for your uploaded topologies
const InfraReadTopologyV1Path = "/v1/infra/read/topologies"
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

// Request Struct
type TopologyDeployRequest struct {
    TopologyID    int    `json:"topologyID"`
    CloudProvider string `json:"cloudProvider"`
    Region        string `json:"region"`
    Context       string `json:"context"`
    Namespace     string `json:"namespace"`
    Env           string `json:"env"`
}
```

Post to this endpoint to deploy this infrastructure topology

#### Destroy Topology

```go
// DestroyDeployInfraV1Path destroys topology, in other words uninstalls the app
const DestroyDeployInfraV1Path = "/v1/deploy/destroy"
```

Post to this endpoint to destroy this infrastructure topology, coming online soon.

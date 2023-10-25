---
sidebar_position: 1
displayed_sidebar: zK8s
---

# Overview

## Infra Configuration via Code + Orchestration

We've merged Kubernetes & Temporal orchestration and added state management using a relational database to manage
distributed infra setups. We exposed control of this system orchestrator to an SDK and UI that allows you to glue
sophisticated distributed systems together in a fraction of the time it took previously which lets you focus on research
and product development instead of figuring out how to build a distributed system or understand Kubernetes.

### Overview

1. Automates translation of kubernetes yaml configurations into representative SQL models
2. Users upload these infrastructure configurations via API where they are stored in the DB
3. Users can then query the contents of these infrastructure components, deploy, mutate or destroy them on demand

### Currently Supported Workload Types

1. Deployments
2. StatefulSets
3. Services
4. ConfigMaps
5. Ingresses
6. ServiceMonitors
7. Secrets

+ Node Tainting Automation & Infra Provisioning (Servers & SSDs) Deployable & Scalable on Demand.

### Pods Controller

1. GetPods
2. GetPodLogs
3. PortforwardReqToPods
4. DeletePods

Not every possible field type is supported, but the most common ones are, and even a decent amount of the uncommon ones.
If you find a field you need isn't supported please send us an email at support@zeus.fyi

![sys](https://user-images.githubusercontent.com/17446735/236006394-d7782657-c2a8-4ee6-a53c-a55f07dbc2b8.png)

Miro: https://miro.com/app/board/uXjVPMNYAhc=/?share_link_id=723624775048

## Topology Class Hierarchy Definitions

### Network-Orchestration Topology ###

One or many network matrix system topologies that are combined with orchestration workflows from Artemis, and Zeus, to
build complex control flows and/or sequenced network states. This could be an enterprise fleet of infrastructure, a
complex devops operation done at scale, or a built-in orchestration flow for automating maintenance procedures & service
restoration and sending them notifications and prompts.

At this stage you can manage complex topologies of infrastructure at the highest of scale over numerous regions and
cloud providers, and include workflow hooks & event triggers, and long running control loops and batch processes. Which
lets you control your environment on demand with complete automation.

### Matrix Topology ###

A multi-component cluster topology that accomplishes one or more system components on its own and combined with a Zeus
injection deploys this topology onto the network. Some but not all of these topologies can be stacked with another
cluster topology or a base topology to create a higher level system component.

It can be any combination of lower level system topologies or components. At this stage you can now include server
resource requirements & tainting rules.

### Cluster Topology ###

A fully working single component cluster topology that accomplishes one system component on its own and combined with a
Zeus injection deploys this topology onto the network. Some but not all of these topologies can be stacked with another
cluster topology or a base topology to create a higher level system component.

### Base Topology ###

A fully working single cluster topology that needs at least one other Base Topology to create a higher level component.
An example would be deploying an execution client by itself post-merge on ethereum. It would be able to download chain
data, but it wouldn’t be able to fulfill a useful purpose without another piece e.g. a consensus client component.

Not deployable on its own, a mix of these is combined to create a deployable topology

### Infrastructure ###

An abstract atomic infrastructure base that needs a Skeleton and Configuration to create a Base Topology

### Configuration ###

An abstract atomic configuration base that needs an Infrastructure Base and Skeleton to create a Base Topology

### Skeleton ###

An abstract atomic component base that needs additional pieces to create deployable infrastructure like config map,
docker image links, etc. Needs an Infrastructure and Configuration Base to create a Base Topology

## How This Looks in Code

### Uniform Building Blocks

Clusters are built from uniform building blocks and are referenced by the TopologyBaseInfraWorkload struct below. The
only limitation is that you cannot add a StatefulSet and a Deployment in one workload block.
Just add another workload block if you need the additional type.

```go
    type TopologyBaseInfraWorkload struct {
        *v1core.Service       `json:"service"`
        *v1core.ConfigMap     `json:"configMap"`
        *v1.Deployment        `json:"deployment"`
        *v1.StatefulSet       `json:"statefulSet"`
        *v1networking.Ingress `json:"ingress"`
        *v1sm.ServiceMonitor  `json:"serviceMonitor"`
    }
```

### Infra Routing

##### To Infinity and Beyond To The Multiverse Cloud

Use the struct below to indicate where the infra should be deployed. CloudProvider and Region are specific to
the cloud provider host, eg. Digital Ocean, GCP, etc. Context, Namespace are specific to Kubernetes. The Env tag is only
for your reference.

```go
type CloudCtxNs struct {
    CloudProvider string `json:"cloudProvider"`
    Region        string `json:"region"`
    Context       string `json:"context"`
    Namespace     string `json:"namespace"`
    Env           string `json:"env"`
}
```

### Cluster Definitions

```go
type ClusterDefinition struct {
    ClusterClassName string
    CloudCtxNs       zeus_common_types.CloudCtxNs
    ComponentBases   map[string]ComponentBaseDefinition
}
// Methods
GenerateDeploymentRequest() -> zeus_req_types.ClusterTopologyDeployRequest
GenerateSkeletonBaseCharts() -> ([]ClusterSkeletonBaseDefinition, error)
BuildClusterDefinitions() -> GeneratedClusterCreationRequests
UploadChartsFromClusterDefinition(ctx context.Context, z zeus_client.ZeusClient, print bool)
->([]zeus_resp_types.TopologyCreateResponse, error)

pkg/zeus/cluster_config_drivers
```

A fully working single component cluster topology that accomplishes one system component on its own and combined with a
Zeus injection deploys this topology onto the network.

```go
type GeneratedClusterCreationRequests struct {
    ClusterClassRequest    zeus_req_types.TopologyCreateClusterClassRequest
    ComponentBasesRequests zeus_req_types.TopologyCreateOrAddComponentBasesToClassesRequest
    SkeletonBasesRequests  []zeus_req_types.TopologyCreateOrAddSkeletonBasesToClassesRequest
}
// Methods
CreateClusterClassDefinitions(ctx context.Context, z zeus_client.ZeusClient)
```

These helpers methods allow you to quickly create new cluster classes, and evaluate the configs of iterative designs.

### Base Topology

```go
type ComponentBaseDefinition struct {
    SkeletonBases map[string]ClusterSkeletonBaseDefinition
}
```

A fully working single cluster topology that needs at least one other Base Topology to create a higher level component.
An example would be deploying an execution client by itself post-merge on ethereum. It would be able to download chain
data, but it wouldn't be able to fulfill a useful purpose without another piece e.g. a consensus client component.

### Skeleton Base

```go
type ClusterSkeletonBaseDefinition struct {
    SkeletonBaseChart         zeus_req_types.TopologyCreateRequest
    SkeletonBaseNameChartPath filepaths.Path
    
    Workload             topology_workloads.TopologyBaseInfraWorkload
    TopologyConfigDriver *zeus_topology_config_drivers.TopologyConfigDriver
}
```

An abstract atomic component base that needs additional pieces to create deployable infrastructure like config map,
docker image links, etc. Needs an Infrastructure and Configuration Base to create a Base Topology

### Infrastructure Base

```go
type TopologyBaseInfraWorkload struct {
    *v1core.Service       `json:"service"`
    *v1core.ConfigMap     `json:"configMap"`
    *v1.Deployment        `json:"deployment"`
    *v1.StatefulSet       `json:"statefulSet"`
    *v1networking.Ingress `json:"ingress"`
    *v1sm.ServiceMonitor  `json:"serviceMonitor"`
}
```

An abstract atomic infrastructure base that needs a Skeleton and Configuration to create a Base Topology

### Configuration Base

```go
type TopologyConfigDriver struct {
    *IngressDriver
    *StatefulSetDriver
    *ServiceDriver
    *DeploymentDriver
    *ServiceMonitorDriver
    *ConfigMapDriver
}
pkg/zeus/workload_config_drivers
```

An abstract atomic configuration base that needs an Infrastructure Base and Skeleton to create a Base Topology. You use
these to override and extend the workload types that match the drivers.

![Screenshot 2022-11-17 at 8 09 48 PM](https://user-images.githubusercontent.com/17446735/202614955-2708063e-1547-4dae-9332-f712102c287e.png)

## Full Code Example - Setting up Hades Microservice

### Hades Microservice

In this subsection we'll show you how to build a cluster definition for a microservice called Hades, which is an
external service
api server that can be used to let Zeus orchestrate and manage your infrastructure on demand.

The code references a template folder which contains these files:

- cm-hades.yaml
- deployment.yaml
- ingress.yaml
- service.yaml

Which is a generic microservice template. You only need to change the docker image to swap your app into this and set
the config map to your start up command.
You can even do that from the UI on the platform.

If you want to deploy your own API microservice, you can use this template as a starting point.

#### cm-hades.yaml

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: cm-hades
data:
  start.sh: |-
    #!/bin/sh
    exec hades --env="production"
```

Change "hades" in the above to the name of your binary in your docker container, and update the startup command.

#### deployment.yaml

Change the below 8888 to the port your app runs on, leaving the other fields as is works for most cases.

```yaml
            - name: "http"
              containerPort: 8888
              protocol: "TCP"
```

#### See the full code at

`cookbooks/hades`

https://github.com/zeus-fyi/zeus
```go
var (
HadesCloudCtxNs = zeus_common_types.CloudCtxNs{
    CloudProvider: "do",
    Region:        "sfo3",
    Context:       "do-nyc1-do-nyc1-zeus-demo",
    Namespace:     "hades",
    Env:           "production",
}
HadesClusterDefinition = zeus_cluster_config_drivers.ClusterDefinition{
    ClusterClassName: "hades",
    CloudCtxNs:       HadesCloudCtxNs,
    ComponentBases: map[string]zeus_cluster_config_drivers.ComponentBaseDefinition{
        "hades": HadesComponentBase,
    },
}
HadesComponentBase = zeus_cluster_config_drivers.ComponentBaseDefinition{
    SkeletonBases: map[string]zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
        "hades": HadesSkeletonBaseConfig,
    },
}
HadesSkeletonBaseConfig = zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
    SkeletonBaseChart:         zeus_req_types.TopologyCreateRequest{},
    SkeletonBaseNameChartPath: HadesChartPath,
}
HadesChartPath = filepaths.Path{
    PackageName: "",
    DirIn:       "./hades/infra",
    DirOut:      "./hades/outputs",
    FnIn:        "hades", // filename for your gzip workload
    FnOut:       "",
    Env:         "",
})
```

This test case shows how to build a cluster definition, and upload it to Zeus. You can use this as a template to build
your own cluster definitions.

```go
type HadesCookbookTestSuite struct {
    test_suites.BaseTestSuite
    ZeusTestClient zeus_client.ZeusClient
}

func (t *HadesCookbookTestSuite) TestClusterSetup() {
    gcd := HadesClusterDefinition.BuildClusterDefinitions()
    t.Assert().NotEmpty(gcd)
    fmt.Println(gcd)
    
    gdr := HadesClusterDefinition.GenerateDeploymentRequest()
    t.Assert().NotEmpty(gdr)
    fmt.Println(gdr)
    
    sbDefs, err := HadesClusterDefinition.GenerateSkeletonBaseCharts()
    t.Require().Nil(err)
    t.Assert().NotEmpty(sbDefs)
}

func (t *HadesCookbookTestSuite) SetupTest() {
    cookbooks.ChangeToCookbookDir()
}

func TestHadesCookbookTestSuite(t *testing.T) {
    suite.Run(t, new(HadesCookbookTestSuite))
}
```

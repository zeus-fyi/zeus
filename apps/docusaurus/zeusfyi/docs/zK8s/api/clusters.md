# Clusters

This package, `zeus_cluster_config_drivers`, provides definitions and functionalities related to the configuration of
clusters in Zeus.

## Table of Contents

1. [Imports](#imports)
2. [Data Structures](#data-structures)
   - [ClusterDefinition](#clusterdefinition)
   - [ComponentBaseDefinition](#componentbasedefinition)
   - [ClusterSkeletonBaseDefinition](#clusterskeletonbasedefinition)
3. [Functions](#functions)
   - [UploadChartsFromClusterDefinition](#uploadchartsfromclusterdefinition)
   - [GenerateDeploymentRequest](#generatedeploymentrequest)
   - [GenerateSkeletonBaseCharts](#generateskeletonbasecharts)

---

## Imports

```go
type ClusterDefinition struct {
ClusterClassName          string
CloudCtxNs                zeus_common_types.CloudCtxNs
ComponentBases            map[string]ComponentBaseDefinition
FilterSkeletonBaseUploads *strings_filter.FilterOpts
DisablePrint              bool
UseEmbeddedWorkload       bool
}

type ComponentBaseDefinition struct {
SkeletonBases map[string]ClusterSkeletonBaseDefinition
}

type ClusterSkeletonBaseDefinition struct {
SkeletonBaseChart         zeus_req_types.TopologyCreateRequest
SkeletonBaseNameChartPath filepaths.Path

TopologyID           int
Workload             topology_workloads.TopologyBaseInfraWorkload
TopologyConfigDriver *zeus_topology_config_drivers.TopologyConfigDriver
}

type ClusterSkeletonBaseDefinitions []ClusterSkeletonBaseDefinition

```

---

## Data Structures

### `ClusterDefinition`

Defines the class name, context, and components for a cluster.

**Fields:**

- `ClusterClassName`: String name of the cluster class.
- `CloudCtxNs`: The cloud context namespace.
- `ComponentBases`: Map with the structure `{string: ComponentBaseDefinition}`.
- `FilterSkeletonBaseUploads`: Options to filter skeleton base uploads.
- `DisablePrint`: Boolean flag to disable print.
- `UseEmbeddedWorkload`: Boolean flag to use the embedded workload.

### `ComponentBaseDefinition`

Definition for base components in a cluster.

**Fields:**

- `SkeletonBases`: Map with the structure `{string: ClusterSkeletonBaseDefinition}`.

### `ClusterSkeletonBaseDefinition`

Details of the skeleton base of a cluster component.

**Fields:**

- `SkeletonBaseChart`: Topology create request.
- `SkeletonBaseNameChartPath`: Path to the chart name.
- `TopologyID`: Identification number for the topology.
- `Workload`: Defines the topology base infrastructure workload.
- `TopologyConfigDriver`: Pointer to the configuration driver for the topology.

---

## Functions

### `UploadChartsFromClusterDefinition`

Uploads charts from a given cluster definition.

**Parameters:**

- `ctx`: Context of the current execution.
- `z`: Zeus client.
- `print`: Boolean flag to print.

**Returns:**

- Array of topology create responses.
- Possible error.

### `GenerateDeploymentRequest`

Generates a request for deploying a cluster topology.

**Returns:**

- `ClusterTopologyDeployRequest`.

### `GenerateSkeletonBaseCharts`

Generates the base charts for skeletons in a cluster.

**Returns:**

- Array of `ClusterSkeletonBaseDefinition`.
- Possible error.

---


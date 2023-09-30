---
sidebar_position: 1
displayed_sidebar: zK8s
---

# zK8s Client #

The zK8s client is used for interacting with our cloud platform's core infrastructure apis and services.

## Quickstart ##

You can override the values in the test files we have and set to your own following these steps

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

# zK8s API Requests

This package, `zeus_req_types`, provides various request structures used to interact with Zeus services.

## Table of Contents

1. [Topology Requests](#topology-requests)
    - [TopologyCreateRequest](#topologycreaterequest)
    - [TopologyDeployRequest](#topologydeployrequest)
    - [TopologyRequest](#topologyrequest)
    - [TopologyCloudCtxNsQueryRequest](#topologycloudctxnsqueryrequest)
    - [ClusterTopologyDeployRequest](#clustertopologydeployrequest)
    - [ClusterTopology](#clustertopology)
    - [ClusterTopologies](#clustertopologies)
2. [Class Creation Requests](#class-creation-requests)
    - [TopologyCreateClusterClassRequest](#topologycreateclusterclassrequest)
    - [TopologyCreateOrAddComponentBasesToClassesRequest](#topologycreateoraddcomponentbasestoclassesrequest)
    - [TopologyCreateOrAddSkeletonBasesToClassesRequest](#topologycreateoraddskeletonbasestoclassesrequest)

---

## Imports

```go
package zeus_req_types

import "github.com/zeus-fyi/zeus/zeus/z_client/zeus_common_types"

type TopologyCreateRequest struct {
	TopologyName     string `json:"topologyName"`
	ChartName        string `json:"chartName"`
	ChartDescription string `json:"chartDescription,omitempty"`
	Version          string `json:"version"`

	ClusterClassName  string `json:"clusterClassName,omitempty"`
	ComponentBaseName string `json:"componentBaseName,omitempty"`
	SkeletonBaseName  string `json:"skeletonBaseName,omitempty"`
	Tag               string `json:"tag,omitempty"`
}

type TopologyDeployRequest struct {
	TopologyID                   int `json:"topologyID"`
	zeus_common_types.CloudCtxNs `json:"cloudCtxNs"`

	SecretRef                       string `json:"secretRef,omitempty"`
	RequestChoreographySecretDeploy bool   `json:"requestChoreographySecretDeploy,omitempty"`
}

type TopologyRequest struct {
	TopologyID int `json:"topologyID"`
}

type TopologyCloudCtxNsQueryRequest struct {
	zeus_common_types.CloudCtxNs
}

type ClusterTopologyDeployRequest struct {
	ClusterClassName             string   `json:"clusterClassName"`
	SkeletonBaseOptions          []string `json:"skeletonBaseOptions"`
	AppTaint                     bool     `json:"appTaint,omitempty"`
	zeus_common_types.CloudCtxNs `json:"cloudCtxNs"`
}

type ClusterTopology struct {
	ClusterClassName string              `json:"clusterClassName"`
	Topologies       []ClusterTopologies `json:"topologies"`
}

type ClusterTopologies struct {
	TopologyID       int    `json:"topologyID"`
	SkeletonBaseName string `json:"skeletonBaseName"`
	Tag              string `json:"tag"`
}

type TopologyCreateClusterClassRequest struct {
	ClusterClassName string `json:"clusterClassName"`
}

type TopologyCreateOrAddComponentBasesToClassesRequest struct {
	ClusterClassName   string   `json:"clusterClassName,omitempty"`
	ComponentBaseNames []string `json:"componentBaseNames,omitempty"`
}

type TopologyCreateOrAddSkeletonBasesToClassesRequest struct {
	ClusterClassName  string   `json:"clusterClassName"`
	ComponentBaseName string   `json:"componentBaseName,omitempty"`
	SkeletonBaseNames []string `json:"skeletonBaseNames,omitempty"`
}

```

## Topology Requests

### `TopologyCreateRequest`

Request structure for creating a new topology.

**Fields:**

- `TopologyName`: Name of the topology.
- `ChartName`: Name of the chart.
- `ChartDescription`: (Optional) Description of the chart.
- `Version`: Version information for the topology.
- `ClusterClassName`: (Optional) Name of the cluster class.
- `ComponentBaseName`: (Optional) Base name for the component.
- `SkeletonBaseName`: (Optional) Name for the skeleton base.
- `Tag`: (Optional) Tag for the topology.

### `TopologyDeployRequest`

Request structure for deploying a topology.

**Fields:**

- `TopologyID`: Identification number for the topology.
- `CloudCtxNs`: Cloud context namespace.
- `SecretRef`: (Optional) Reference for the secret.
- `RequestChoreographySecretDeploy`: Boolean flag indicating if choreography secret deployment is requested.

### `TopologyRequest`

Basic request structure with topology ID.

**Fields:**

- `TopologyID`: Identification number for the topology.

### `TopologyCloudCtxNsQueryRequest`

Request structure for querying based on cloud context namespace.

**Fields:**

- `CloudCtxNs`: Cloud context namespace.

### `ClusterTopologyDeployRequest`

Request structure for deploying cluster topology.

**Fields:**

- `ClusterClassName`: Name of the cluster class.
- `SkeletonBaseOptions`: List of options for skeleton base.
- `AppTaint`: Boolean flag to indicate application taint.
- `CloudCtxNs`: Cloud context namespace.

### `ClusterTopology`

Structure representing cluster topology details.

**Fields:**

- `ClusterClassName`: Name of the cluster class.
- `Topologies`: List of `ClusterTopologies`.

### `ClusterTopologies`

Structure representing details of cluster topologies.

**Fields:**

- `TopologyID`: Identification number for the topology.
- `SkeletonBaseName`: Name for the skeleton base.
- `Tag`: Tag for the topology.

---

## Class Creation Requests

### `TopologyCreateClusterClassRequest`

Request structure to create a cluster class.

**Fields:**

- `ClusterClassName`: Name of the cluster class.

### `TopologyCreateOrAddComponentBasesToClassesRequest`

Request structure to create or add component bases to classes.

**Fields:**

- `ClusterClassName`: (Optional) Name of the cluster class.
- `ComponentBaseNames`: (Optional) List of component base names.

### `TopologyCreateOrAddSkeletonBasesToClassesRequest`

Request structure to create or add skeleton bases to classes.

**Fields:**

- `ClusterClassName`: Name of the cluster class.
- `ComponentBaseName`: (Optional) Base name for the component.
- `SkeletonBaseNames`: (Optional) List of names for the skeleton bases.

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

---
sidebar_position: 3
displayed_sidebar: zK8s
---

# Pods #

This documentation provides a detailed overview of the API structures and constants related to pod actions.

## Table of Contents

1. [Imports](#imports)
2. [Constants](#constants)
3. [API Details](#api-details)
   - [PodActionRequest](#podactionrequest)
   - [ClientRequest](#clientrequest)
4. [zK8s Go Client Package](#zk8s-pods-go-client-package-comprehensive-guide)
---


## Imports

```go
package zeus_pods_reqs

const (
   DeleteAllPods                = "delete-all"
   GetPodLogs                   = "logs"
   GetPods                      = "describe"
   PortForwardToAllMatchingPods = "port-forward-all"
   DescribeAudit                = "describe-audit"
)

type PodActionRequest struct {
   zeus_req_types.TopologyDeployRequest
   Action        string `json:"action"`
   PodName       string `json:"podName,omitempty"`
   ContainerName string `json:"containerName,omitempty"`

   FilterOpts *strings_filter.FilterOpts
   ClientReq  *ClientRequest
   LogOpts    *v1.PodLogOptions
   DeleteOpts *metav1.DeleteOptions
}

type ClientRequest struct {
   MethodHTTP      string
   Endpoint        string
   Ports           []string
   Payload         any
   EndpointHeaders map[string]string
}

```
## Constants

- `DeleteAllPods`: Used to denote the action of deleting all pods. Value: `"delete-all"`
- `GetPodLogs`: Used for retrieving pod logs. Value: `"logs"`
- `GetPods`: Used to describe the pods. Value: `"describe"`
- `PortForwardToAllMatchingPods`: Represents the action of port forwarding to all matching pods. Value: `"port-forward-all"`
- `DescribeAudit`: Represents the action of describing an audit. Value: `"describe-audit"`

---

## API Details

### `PodActionRequest`

This structure contains details related to the request for performing an action on a pod.

**Fields:**

- `TopologyDeployRequest`: (Inherits from `zeus_req_types.TopologyDeployRequest`) Contains the details of the topology deployment request.
- `Action`: A string representing the action to be performed.
- `PodName`: (Optional) The name of the pod.
- `ContainerName`: (Optional) The name of the container within the pod.
- `FilterOpts`: Filtering options for the action.
- `ClientReq`: A request from the client.
- `LogOpts`: Options for logging related to the pod.
- `DeleteOpts`: Options for deleting the pod.

### `ClientRequest`

Details of the request made by a client.

**Fields:**

- `MethodHTTP`: The HTTP method used (e.g., `GET`, `POST`).
- `Endpoint`: The endpoint to which the request is made.
- `Ports`: An array of strings representing the ports.
- `Payload`: Any payload data to be sent with the request.
- `EndpointHeaders`: A map containing headers for the endpoint.

---

## zK8s Pods Go Client Package: Comprehensive Guide

The `pods_client` package focuses on enabling interactions related to pods using the Zeus client. This guide offers a detailed overview of the structures, functionalities, and purposes of various components.

## Table of Contents

1. [Imports](#imports)
2. [Structures](#structures)
   - [PodsClient](#podsclient)
3. [Functions](#functions)
   - [NewPodsClient](#newpodsclient)
   - [NewPodsClientFromZeusClient](#newpodsclientfromzeusclient)
   - [DeletePods](#deletepods)
   - [GetPods](#getpods)
   - [GetPodsAudit](#getpodsaudit)
   - [GetPodLogs](#getpodlogs)
   - [PortForwardReqToPods](#portforwardreqtopods)

---

## Imports

```go
package pods_client

import (
   zeus_client "github.com/zeus-fyi/zeus/zeus/z_client"
   resty_base "github.com/zeus-fyi/zeus/zeus/z_client/base"
)

type PodsClient struct {
   zeus_client.ZeusClient
}

func NewPodsClient(baseURL, bearer string) PodsClient {
   p := PodsClient{}
   p.Resty = resty_base.GetBaseRestyClient(baseURL, bearer)

   return p
}

func NewPodsClientFromZeusClient(z zeus_client.ZeusClient) PodsClient {
   p := PodsClient{
      ZeusClient: z,
   }
   return p
}
```

This package imports the primary Zeus client, enabling high-level interactions with Zeus services, and the resty base to facilitate RESTful operations.

---

## Structures

### `PodsClient`

This structure serves as an enhancement of the `ZeusClient`, specializing in pod-related operations.

**Fields:**

- `ZeusClient`: An embedded field, it gives `PodsClient` all the capabilities of the main Zeus client while allowing for the addition of pod-specific methods.

---

## Functions

### `NewPodsClient(baseURL, bearer string) PodsClient`

Initializes and returns a new instance of `PodsClient` using the given base URL and bearer authentication token.

- **Usage:** Ideal for scenarios where there's a need to establish a new client connection to interact with Zeus services regarding pod operations.

### `NewPodsClientFromZeusClient(z zeus_client.ZeusClient) PodsClient`

Generates a `PodsClient` instance using an existing `ZeusClient`.

- **Usage:** Beneficial when there's an existing Zeus client in context, and you want to leverage its properties and configurations for pod operations without reinitializing.

### `DeletePods(ctx context.Context, par zeus_pods_reqs.PodActionRequest) ([]byte, error)`

Performs deletion of pods based on the provided action request. The method internally sets the action to `DeleteAllPods`.

- **Usage:** Used when you need to clear resources or handle errors that necessitate removing pods. Always ensure you're deleting the right resources to avoid unwanted data loss.

### `GetPods(ctx context.Context, par zeus_pods_reqs.PodActionRequest) (*v1.PodList, error)`

Fetches a list of pods based on the criteria mentioned in the action request. The action is internally set to `GetPods`.

- **Usage:** A fundamental method for monitoring and management, allowing system administrators or services to pull the current state and details of running pods.

### `GetPodsAudit(ctx context.Context, par zeus_pods_reqs.PodActionRequest) (zeus_pods_resp.PodsSummary, error)`

Gathers a summary report or audit of pod states based on the action request, where the action is set to `DescribeAudit`.

- **Usage:** Particularly useful for generating reports, monitoring, or when conducting periodic reviews of pod states and behaviors in the infrastructure.

### `GetPodLogs(ctx context.Context, par zeus_pods_reqs.PodActionRequest) ([]byte, error)`

Retrieves logs from the specified pods. If no custom filter is provided, it defaults to filtering by `PodName`.

- **Usage:** Essential for debugging, monitoring, or any operations that require insights from pod logs. Always ensure you have adequate permissions to access logs for security reasons.

### `PortForwardReqToPods(ctx context.Context, par zeus_pods_reqs.PodActionRequest) (zeus_pods_resp.ClientResp, error)`

Initiates port forwarding to the selected pods. If you provide a pod name without a custom filter, the default filter applied will be `StartsWith` using `PodName`.

- **Usage:** Useful for debugging, direct pod interactions, or when setting up specific networking routes to pods for service access or data transfer.

---

Note: It's imperative to understand the ramifications of operations, especially delete actions, on production systems. Always follow best practices and guidelines to avoid adverse impacts.

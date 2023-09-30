---
sidebar_position: 3
displayed_sidebar: zK8s
---

# Pods #

This documentation provides a detailed overview of the API structures and constants related to pod actions.

## Table of Contents

1. [Imports](#imports)
2. [Constants](#constants)
3. [Structures](#structures)
   - [PodActionRequest](#podactionrequest)
   - [ClientRequest](#clientrequest)

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


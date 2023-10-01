---
sidebar_position: 4
displayed_sidebar: zK8s
---

# Workloads

## Configuration Drivers Documentation

The `zeus_topology_config_drivers` package provides a collection of drivers tailored to modify or augment Kubernetes
configurations. This guide delves deep into the structure, functionalities, and purposes of the components.

## Table of Contents

1. [Imports](#imports)
2. [Structures](#structures)
    - [ConfigMapDriver](#configmapdriver)
    - [ContainerDriver](#containerdriver)
    - [DeploymentDriver](#deploymentdriver)
    - [IngressDriver](#ingressdriver)
    - [PersistentVolumeClaimsConfigDriver](#persistentvolumeclaimsconfigdriver)
3. [Functions and Methods](#functions-and-methods)
    - [ConfigMapDriver](#configmapdriver)
    - [ContainerDriver](#containerdriver)
    - [DeploymentDriver](#deploymentdriver)
    - [IngressDriver](#ingressdriver)
    - [PersistentVolumeClaimsConfigDriver](#persistentvolumeclaimsconfigdriver)

---

## Imports

```go
package zeus_topology_config_drivers

import "github.com/zeus-fyi/zeus/zeus/z_client/zeus_resp_types/topology_workloads"

type TopologyConfigDriver struct {
   *IngressDriver
   *StatefulSetDriver
   *ServiceDriver
   *DeploymentDriver
   *ServiceMonitorDriver
   *ConfigMapDriver
}

func (t *TopologyConfigDriver) SetCustomConfig(inf *topology_workloads.TopologyBaseInfraWorkload) {
   if inf.Ingress != nil && t.IngressDriver != nil {
      t.SetIngressConfigs(inf.Ingress)
   }
   if inf.StatefulSet != nil && t.StatefulSetDriver != nil {
      t.SetStatefulSetConfigs(inf.StatefulSet)
   }
   if inf.Deployment != nil && t.DeploymentDriver != nil {
      t.SetDeploymentConfigs(inf.Deployment)
   }
   if inf.Service != nil && t.ServiceDriver != nil {
      t.SetServiceConfigs(inf.Service)
   }
   if inf.ServiceMonitor != nil && t.ServiceMonitorDriver != nil {
      t.SetServiceMonitorConfigs(inf.ServiceMonitor)
   }
   if inf.ConfigMap != nil && t.ConfigMapDriver != nil {
      t.SetConfigMaps(inf.ConfigMap)
   }
}

type ConfigMapDriver struct {
   v1.ConfigMap
   // swap key for values key in the configmap
   SwapKeys map[string]string
}

type ContainerDriver struct {
   IsAppendContainer bool
   IsInitContainer   bool
   v1Core.Container
   AppendEnvVars []v1Core.EnvVar
}

type DeploymentDriver struct {
   ReplicaCount     *int32
   ContainerDrivers map[string]ContainerDriver
}

type IngressDriver struct {
   v1.Ingress
   Host         string
   NginxAuthURL string
}

type PersistentVolumeClaimsConfigDriver struct {
   AppendPVC                    map[string]bool
   PersistentVolumeClaimDrivers map[string]v1.PersistentVolumeClaim
}

type ServiceMonitorDriver struct {
   v1.ServiceMonitor
}

type ServiceDriver struct {
   v1.Service
   ExtendPorts []v1.ServicePort
}

type StatefulSetDriver struct {
   ReplicaCount     *int32
   ContainerDrivers map[string]ContainerDriver
   PVCDriver        *PersistentVolumeClaimsConfigDriver
}
```

---

## Structures

### `ConfigMapDriver`

This structure extends the Kubernetes `ConfigMap` with an additional functionality to swap specific keys with their
corresponding values.

**Fields:**

- `ConfigMap`: The Kubernetes `ConfigMap` object.
- `SwapKeys`: A map denoting which keys should be swapped for new keys in the configmap.

### `ContainerDriver`

Configures container properties and allows appending of environment variables.

**Fields:**

- `IsAppendContainer`: Specifies if the container should be appended.
- `IsInitContainer`: Denotes if the container is an initialization container.
- `Container`: The Kubernetes `Container` object.
- `AppendEnvVars`: List of environment variables to append.

### `DeploymentDriver`

Deals with the configuration of Kubernetes deployments.

**Fields:**

- `ReplicaCount`: Desired replica count.
- `ContainerDrivers`: A map of container drivers.

### `IngressDriver`

Manages the configuration of Kubernetes ingress resources.

**Fields:**

- `Ingress`: The Kubernetes `Ingress` object.
- `Host`: The host domain.
- `NginxAuthURL`: Nginx authentication URL.

### `PersistentVolumeClaimsConfigDriver`

Facilitates configurations around persistent volume claims in Kubernetes.

**Fields:**

- `AppendPVC`: Determines which PVCs should be appended.
- `PersistentVolumeClaimDrivers`: A map of PVC drivers.

---

## Functions and Methods

### ConfigMapDriver:

#### `SetConfigMaps(cmap *v1.ConfigMap)`

Modifies a given `ConfigMap` object by swapping specific keys and adding new data.

### ContainerDriver:

#### `SetContainerConfigs(cont *v1Core.Container)`

Updates the configuration of a given container based on the `ContainerDriver`.

#### `CreateEnvVarKeyValue(k, v string) v1Core.EnvVar`

Creates an environment variable from a given key and value.

#### `MakeEnvVar(name, key, localObjRef string) v1Core.EnvVar`

Generates an environment variable with specific references.

### DeploymentDriver:

#### `NewDeploymentDriver() DeploymentDriver`

Creates a new instance of `DeploymentDriver`.

#### `SetDeploymentConfigs(dep *v1.Deployment)`

Modifies a given deployment's configurations based on the `DeploymentDriver`.

### IngressDriver:

#### `SetIngressConfigs(ing *v1.Ingress)`

Updates a given ingress's configurations based on the `IngressDriver`.

### PersistentVolumeClaimsConfigDriver:

#### `CustomPVCS(pvcs []v1.PersistentVolumeClaim) []v1.PersistentVolumeClaim`

Customizes a slice of `PersistentVolumeClaim` based on the driver's configurations.

---

This guide provides in-depth documentation of the `zeus_topology_config_drivers` package. Understanding each component
and its operation helps in efficient utilization and management of Kubernetes resources.

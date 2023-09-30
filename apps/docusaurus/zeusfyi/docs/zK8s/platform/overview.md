---
sidebar_position: 1
displayed_sidebar: zK8s
---

# Overview

## Infra configuration via code + orchestration

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

## Cluster Topology Class Hierarchy Definitions: Highest to Lowest Level

### Network Orchestration Topology ###

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

### Infrastructure Base ###

An abstract atomic infrastructure base that needs a Skeleton and Configuration to create a Base Topology

### Configuration Base ###

An abstract atomic configuration base that needs an Infrastructure Base and Skeleton to create a Base Topology

### Skeleton Base ###

An abstract atomic component base that needs additional pieces to create deployable infrastructure like config map,
docker image links, etc. Needs an Infrastructure and Configuration Base to create a Base Topology

![Screenshot 2022-11-17 at 8 09 48 PM](https://user-images.githubusercontent.com/17446735/202614955-2708063e-1547-4dae-9332-f712102c287e.png)

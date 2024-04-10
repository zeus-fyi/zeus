# Documentation ##

https://docs.zeus.fyi

## Mockingbird Front End GitHub

https://github.com/zeus-fyi/mockingbird

## Mockingbird

Mockingbird is a time series controlled AI systems coordinator & workflow executor system, data indexer and searcher,
that builds control
loops and hierarchical planning rules around higher order goals & objectives using automated evaluation systems and
model-to-model communication via synchronized quorums.

### Usage

Mockingbird is currently a model based supervisor entity, and soon, it will be a Model Supervised Intelligent Network, and now a product you can use, like ChatGPT 4, but with power steering and cruise control.

### Quickstart

https://github.com/zeus-fyi/zeus/assets/17446735/af178092-e47c-4782-ab91-985ef2e57086

### Costs

Mockingbird is free to use for the first 3 days without a credit card with a valid login account. Though we may restrict account usage if we see abuse. If you connect your own OpenAI API key, you will be responsible for any costs associated with that, otherwise we will use our API key and then charge you 1-to-1 for the cost of the API usage that we incur, however you may have your usage throttled if you exceed our usage limits while using our API key.

### Features Available via Self Service UI

- Social Media Indexing (Twitter, Reddit, Discord)
- Retrievals (API, Social Media)
- Analysis/Aggregation Tasks
- Integrated Time Window Search by Aggregation Windows
- Multi-model communication and I/O
- Time series controlled AI workflows, with dynamic time step control
- Workflow + relationship builder (UI)
- Integrated OpenAI + Social integration APIs, natively via Platform Secrets
- Run Scheduler/Controller
- Run Step-by-Step Cycle Execution History
- Run Token Usage
- Human-in-the-loop control trigger actions like API requests
- Automated evaluation systems
- Integrated JSON schema builder

### Features Available To Researchers & Partners

- Additional indexing sources (News, Blogs, Forums, etc)
- Adaptive time series control via Metrics + Evals (PromQL, Adaptive)
- Analysis cache feedback searching (eg. reviewing previous model generated analysis to steer actions)
- Objective building for complex decision making, and milestone tracking
- Decision tree generation for structured decision making
- Unrestricted trigger actions (eg. outside of Human-in-the-loop control)
- Experimental distributed executive task planner
- Tasking single-to-multi-model task planning
- Tasking scheduling, recursive workflows, objective overrides
- Shared workflows contexts, selective contexts, searchable contexts
- Scenario driven workflow generator
- Multi-Model-Single-Model communication via quorums
- Multi-Model-Multi-Model communication via quorums + task planning + synchronization
- Ranked/Promotable-Model-Leader-Model-Follower
- Multi-Model Team vs Multi-Model Team Competitive Scenarios
- Multi-Modal (Text, Image, Video, Audio) analysis
- AI driven compute infrastructure, building, purchasing, maintenance, & monitoring systems for Kubernetes + Cloud
  Vendor APIs
- PagerDuty integration

### Unveiling the Next Generation of AI-Powered Workflow Automation

https://medium.zeus.fyi/unveiling-the-next-generation-of-ai-powered-workflow-automation-1f957bc20d3e

## zK8s == Kubernetes + Zeus
Here we overview the core concepts needed to understand how you can build, deploy, configure K8s apps using Zeus, with a full walkthrough example of how we created an Ethereum beacon.

https://medium.com/@zeusfyi/zeus-k8s-infra-as-code-concepts-47e690c6e3c5

## Hosted Docusaurus in 5 Minutes and Under $10/month
Developing Kubernetes applications is often complex and time-consuming, with a steep learning curve. However, the advent of zK8s is changing the game, allowing developers to build Kubernetes apps with pure Go code that’s functional, testable, and effortlessly turned into deployable infrastructure.

```go
type WorkloadDefinition struct {
   WorkloadName string                    `json:"workloadName"`
   ReplicaCount int                       `json:"replicaCount"`
   Containers   zk8s_templates.Containers `json:"containers"`
   FilePath     filepaths.Path            `json:"-"`
}

type Containers map[string]Container

type Container struct {
   IsInitContainer bool        `json:"isInitContainer"`
   ImagePullPolicy string      `json:"imagePullPolicy,omitempty"`
   DockerImage     DockerImage `json:"dockerImage"`
}

type DockerImage struct {
   ImageName            string               `json:"imageName"`
   Cmd                  string               `json:"cmd"`
   Args                 string               `json:"args"`
   ResourceRequirements ResourceRequirements `json:"resourceRequirements,omitempty"`
   EnvVars              []EnvVar             `json:"envVars,omitempty"`
   Ports                []Port               `json:"ports,omitempty"`
   VolumeMounts         []VolumeMount        `json:"volumeMounts,omitempty"`
}
```
### How we built a Docusaurus deployment from this.
```go
wd := zeus_cluster_config_drivers.WorkloadDefinition{
    WorkloadName: "docusaurus-template",
    ReplicaCount: 1,
    Containers: zk8s_templates.Containers{
     docusaurusTemplate: zk8s_templates.Container{
      ImagePullPolicy: "Always",
      DockerImage: zk8s_templates.DockerImage{
         ImageName: "docker.io/zeusfyi/docusaurus-template:latest",
         ResourceRequirements: zk8s_templates.ResourceRequirements{
            CPU:    "100m",
            Memory: "500Mi",
         },
         Ports: []zk8s_templates.Port{
           {
             Name:               "http",
             Number:             "3000",
             Protocol:           "TCP",
             IngressEnabledPort: true,
             ProbeSettings: zk8s_templates.ProbeSettings{
                UseForLivenessProbe:  true,
                UseForReadinessProbe: true,
                UseTcpSocket:         true,
          },
         },
        },
      },
    },
  }
}
```

### Tutorial & zK8s Overview
https://medium.zeus.fyi/hosted-docusaurus-in-5-minutes-and-under-10-month-af999d7ef90a

## Authenticated API in 5 Minutes
Step by step tutorial using our UI

https://medium.com/@zeusfyi/zeus-ui-no-code-kubernetes-authenticated-api-tutorial-c468d5ef0446

## zK8s Apps & Clients ##
zK8s is an expressive language for cloud infrastructure, used for building, assembling, and keeping them running over their entire lifecycle. Enabling cost efficient, effortless large scale infra automation, coordination, customization, and control.

#### ```cluster_config_drivers ```
#### ```system_config_drivers ```
#### ```workload_config_drivers ```

Workflow & Proxy Programmable Automation (Rolling releases coming through end of year)

#### ```artemis_workflows ```
#### ```iris_programmable_proxy ```

#### API Endpoints 

Documentation and code examples are found here
[API_README.md](https://github.com/zeus-fyi/zeus/blob/main/pkg/zeus/API_README.md)

How to use the test suite to setup your own api calls
[README.md](https://github.com/zeus-fyi/zeus/blob/main/pkg/zeus/README.md)

The test directory contains useful mocks and tools for interacting with the API. It also contains a useful
config-sample.yaml, convert this to config.yaml and set your bearer token here, which then allows you to
use the demo code to create your first api request in seconds

## Overview

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

Not every possible field type is supported, but the most common ones are, and even a decent amount of the uncommon ones. If you find a field you need isn't supported please send us an email at support@zeus.fyi

### Hades Library

Hades is used to interact with Kubernetes workloads via API, and can apply saved Zeus workloads & cookbooks onto your
own in house infrastructure.

### Hera Client

This client uses the OpenAI API to generate code with AI. This service is available at OpenAI cost, so just pay for the
token cost, otherwise it is free to use.

## Cookbooks ##

Contains common web2 & web3 recipes for deploying applications, and managing infrastructure using zK8s or vanilla
Kubernetes.

### Zeus UI Highlights

![Screenshot 2023-04-11 at 1 58 00 PM](https://user-images.githubusercontent.com/17446735/231288647-e3bf8db7-67c9-4e0c-af38-c0f1091da726.png)

![Screenshot 2023-04-11 at 1 57 46 PM](https://user-images.githubusercontent.com/17446735/231288668-f8c147fe-d06b-49e1-8313-a7453cbe6f19.png)

![Screenshot 2023-04-07 at 6 46 08 PM](https://user-images.githubusercontent.com/17446735/231288683-a350f36b-d103-428f-88b3-eac80742a9c4.png)

![Screenshot 2023-04-05 at 2 11 33 PM](https://user-images.githubusercontent.com/17446735/231288689-f970cd81-76b3-4b85-9241-3f30ad7c80b9.png) 


# Recent Articles ##

## Introducing Serverless EVMs
Foundry’s Anvil as an EVM Simulation Environment on Demand.

Democratizing access to simulation technology at scale  by lowering the cost of simulation to near zero.
For safer smart contracts, less exploits, more robust engineering by using isolated ephemeral environments per each test.
![Screenshot 2023-11-02 at 11 52 38AM](https://github.com/zeus-fyi/zeus/assets/17446735/5d6576d7-a734-48e7-a2e4-a8a0fd379752)

https://medium.zeus.fyi/introducing-serverless-evms-0035549a5e7f

## Secret Sauce: Running Ethereum Validators at Scale for Pennies
Why run them any other way after you’ve read this?

![Screenshot 2023-11-03 at 2 06 34PM](https://github.com/zeus-fyi/zeus/assets/17446735/b710b0ae-a7ca-468e-ba9d-0dce65d964b9)

https://medium.com/zeusfyi/secret-sauce-running-ethereum-validators-at-scale-for-pennies-24ac8dad4efd

## Automating Validator Status Management in Ethereum in SQL
Using SQL Triggers and Relational States Compliant with Enforcing Consensus Spec

![Screenshot 2023-11-02 at 11 51 00AM](https://github.com/zeus-fyi/zeus/assets/17446735/4b152ebb-0ff2-4458-a6a9-66542c144530)

https://medium.zeus.fyi/automating-validator-status-management-in-ethereum-in-sql-part-1-464e40e32b0b

## Adaptive RPC Load Balancer Benchmarks
Over several weeks on a production load monitoring Uniswap prices

When should you consider using the Adaptive RPC Load Balancer?​

TLDR:
- You need to scale your application to handle more requests than a single endpoint can handle
- You need to reduce the error rate of your application
- Especially if you are still paying for the request even if it fails
- You need to improve the reliability of your application
- You want to run multi-step procedures in a single request

![weekly](https://github.com/zeus-fyi/zeus/assets/17446735/46aff89c-8035-438c-a5a8-45f6195e02f6)

https://medium.zeus.fyi/adaptive-rpc-load-balancer-benchmarks-c7aa3aa0d42a

## High Performance Disks: NVMe in the Cloud
How to use NVMe disks the right way before you spend $$$$

https://medium.zeus.fyi/high-performance-disks-nvme-in-the-cloud-abb2bfc11fd9

![Screenshot 2023-10-25 at 12 20 37 AM](https://github.com/zeus-fyi/zeus/assets/17446735/5fc8399a-9bcf-4cab-a1c3-0bf2cfdb48d1)

## Show Me the Stats
Optimal adaptive load balancing in stochastic environments.
Recommended reading for scientists, engineers, data driven individuals

![Screenshot 2023-09-14 at 11 12 23 PM](https://github.com/zeus-fyi/zeus/assets/17446735/025d3201-9236-40e9-8723-3a7d2d7a3e0a)

https://medium.com/zeusfyi/show-me-the-stats-6740f8d6d0b7

## Adaptive RPC Load Balancer
Accurate, Reliable, Performant Node Traffic at Web3 Scale

![Screenshot 2023-09-14 at 11 11 30 PM](https://github.com/zeus-fyi/zeus/assets/17446735/802b7670-6b30-4e65-9348-e45e2a0cfcac)

https://medium.com/zeusfyi/adaptive-rpc-load-balancer-on-quicknode-marketplace-e68bb7c9d8ac

![Screenshot 2023-09-14 at 11 13 55 PM](https://github.com/zeus-fyi/zeus/assets/17446735/1d2a263e-5aa7-418c-a0f0-1f497cab0353)


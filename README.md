## Zeus ##

We’ve built technology significantly improving Kubernetes in both scalability, and its ease of use over multi cloud/hybrid/on premise as well as its ability to automate complex devops management procedures and manage distributed cloud infra setups. We built an orchestration engine using Temporal over Kubernetes itself for managing, and controlling infrastructure at scale and mapped all the infrastructure into our Postgres database. This makes the entire infrastructure layout completely searchable and easily mappable to business logic requirements. Our infrastructure is designed specifically to meet the performance requirements for complex cloud distributed systems such as in AI/Ml, Big Data, and managing infra at significant scale, and in targeted areas of web3.

### Zeus UI Highlights

![Screenshot 2023-04-11 at 1 58 00 PM](https://user-images.githubusercontent.com/17446735/231288647-e3bf8db7-67c9-4e0c-af38-c0f1091da726.png)

![Screenshot 2023-04-11 at 1 57 46 PM](https://user-images.githubusercontent.com/17446735/231288668-f8c147fe-d06b-49e1-8313-a7453cbe6f19.png)

![Screenshot 2023-04-07 at 6 46 08 PM](https://user-images.githubusercontent.com/17446735/231288683-a350f36b-d103-428f-88b3-eac80742a9c4.png)

![Screenshot 2023-04-05 at 2 11 33 PM](https://user-images.githubusercontent.com/17446735/231288689-f970cd81-76b3-4b85-9241-3f30ad7c80b9.png)

### Zeus K8s Infra as Code Concepts

Here we overview the core concepts needed to understand how you can build, deploy, configure K8s apps using Zeus, with a full walkthrough example of how we created an Ethereum beacon.

https://medium.com/@zeusfyi/zeus-k8s-infra-as-code-concepts-47e690c6e3c5

## Cookbooks ##

Contains common web2 & web3 building components like ethereum infra setups with customization driven through code, blurring the line between infra configuration and app development, and contains microservice designs & patterns like api servers and injectable choreography, their setups in kubernetes, golang, docker, and startup commands, and useful tools for debugging, interacting, and automating actions.

### Cookbook Structure ###

#### Microservices ###

Contains full kubernetes infra setup templates for microservices, injectable choreography for clusters, and more continually being added.

#### ```zeus/cookbooks/microservices/deployments ```
#### ```zeus/cookbooks/microservices/choreography ```

## Zeus Apps & Clients ##

Core Zeus Infra Automation Client
#### ```pkg/zeus/client```

Powerful Cluster Building, Allowing for Large Scale Infra Automation, Customization, Control

#### ```pkg/zeus/cluster_config_drivers ```
#### ```pkg/zeus/system_config_drivers ```
#### ```pkg/zeus/workload_config_drivers ```

#### API Endpoints 

Documentation and code examples are found here
[API_README.md](https://github.com/zeus-fyi/zeus/blob/main/pkg/zeus/API_README.md)

How to use the test suite to setup your own api calls
[README.md](https://github.com/zeus-fyi/zeus/blob/main/pkg/zeus/README.md)

The test directory contains useful mocks and tools for interacting with the API. It also contains a useful
config-sample.yaml, convert this to config.yaml and set your bearer token here, which then allows you to
use the demo code to create your first api request in seconds

### Hera Client

#### ```pkg/hera/client```

This client uses the OpenAI API to generate code with AI. This service is available at OpenAI cost, so just pay for the token cost, otherwise it is free to use.

### Hades Library

#### ```pkg/hades```

Hades is used to interact with kubernetes workloads via API, and can apply saved Zeus workloads & cookbooks onto your own in house infrastructure.

## Infra configuration is in the 14th century. Introducing the 21st century printing press

We've merged Kubernetes & Temporal orchestration and added state management using a relational database to manage distributed infra setups. We exposed control of this system orchestrator to an SDK and UI that allows you to glue sophisticated distributed systems together in a fraction of the time it took previously which lets you focus on research and product development instead of figuring out how to build a distributed system or understand Kubernetes.

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

### Pods Controller 

1. GetPods
2. GetPodLogs
3. PortforwardReqToPods
4. DeletePods

Not every possible field type is supported, but the most common ones are, and even a decent amount of the uncommon ones.

See demos section for api calls, you'll need to get a bearer token from us first. More documentation to come.
Schedule a demo: https://calendly.com/alex-zeus-fyi/zeus-demo
Request a bearer token, email: alex@zeus.fyi

![Screenshot 2022-11-17 at 8 09 48 PM](https://user-images.githubusercontent.com/17446735/202614955-2708063e-1547-4dae-9332-f712102c287e.png)

### Cluster Topology Class Hierarchy Definitions ###

### Highest to Lowest Level ###

### Deployable Topologies ###

### Network Orchestration Topology ###

One or many network matrix system topologies that are combined with orchestration workflows from Artemis, and Zeus, to build complex control flows and/or sequenced network states. This could be an enterprise fleet of infrastructure, a complex devops operation done at scale, or a simulation of a large testnet, or a built-in orchestration flow for automating web3 smart contract interactions for users and sending them notifications and prompts.

### Matrix Topology ###

A multi-component cluster topology that accomplishes one or more system components on its own and combined with a Zeus injection deploys this topology onto the network. Some but not all of these topologies can be stacked with another cluster topology or a base topology to create a higher level system component.

It can be any combination of lower level system topologies or components.

### Cluster Topology ###

A fully working single component cluster topology that accomplishes one system component on its own and combined with a Zeus injection deploys this topology onto the network. Some but not all of these topologies can be stacked with another cluster topology or a base topology to create a higher level system component.

### Deployable but Incomplete Topologies ###

### Base Topology ###

A fully working single cluster topology that needs at least one other Base Topology to create a higher level component. An example would be deploying an execution client by itself post-merge on ethereum. It would be able to download chain data, but it wouldn’t be able to fulfill a useful purpose without another piece e.g. a consensus client component.  

Not deployable on its own, a mix of these is combined to create a deployable topology

### Infrastructure Base ###
An abstract atomic infrastructure base that needs a Skeleton and Configuration to create a Base Topology

### Configuration Base ###
An abstract atomic configuration base that needs an Infrastructure Base and Skeleton to create a Base Topology

### Skeleton Base ###
An abstract atomic component base that needs additional pieces to create deployable infrastructure like config map, docker image links, etc. Needs an Infrastructure and Configuration Base to create a Base Topology

### Web3 Cookbook 

### Artemis ###

#### ```artemis.zeus.fyi``` ####
#### ```pkg/artemis/client```

Artemis is a tx orchestrator. It reliably submits & confirms ethereum transactions and logs their receipts. Chain with 
the in memory db for storing web3 signer keys to build highly reliable web3 api actions with other users and smart contracts. You'll need
a bearer token to use this client. A more advanced orchestrator that can handle high volume DeFi trading, which manages nonce sequences, sets up event triggers & scheduling, and has queriable event artifacts is in works, targeted release by end of March.

### Hercules ###

#### ```apps/hercules``` ####
#### ```pkg/hercules/client```

Hercules is web3 middleware that manages web3 infrastructure and connections to other middleware packages. It also contains useful apis to debug and troubleshoot web3 infrastructure.

### Snapshots ###

#### ```apps/snapshots``` ####

Snapshot app is embedded into the hercules docker app, and it can be used as an init container to download snapshot data on new node creation.

#### Ethereum ####

#### ```protocol/components```

Contains smart contract automation kits. This testcase shows a full end-end seed, create, and deposits validators on the Ethereum ephemery testnet.

#### ```cookbooks/ethereum/automation```
#### ```cookbooks/ethereum/automation/deposits_test.go ```

Cookbook items listed by protocol & component class. Eg. Ethereum has a beacon component group. Contains Kubernetes config setup templates. Here's a few example paths. Also contains an actions folder, which does log dumps, pod restarts, configuration changes on demand for k8s applications, and more.

#### ```cookbooks/ethereum/beacons/infra/consensus_client```
#### ```cookbooks/ethereum/beacons/infra/exec_client```
#### ```cookbooks/ethereum/validators/infra/validators```
#### ```cookbooks/ethereum/web3signers/infra/consensys_web3signer```

Complete, and powerful Ethereum infra automation 

#### ```zeus/cookbooks/ethereum/beacons/beacon_cluster_test.go ```
#### ```zeus/cookbooks/ethereum/validators/validator_cluster_test.go ```
#### ```zeus/cookbooks/ethereum/web3signers/web3signer_cluster_test.go ```

See this test case to see how a beacon cluster class was created, and then extended to support choreography to reset configs on a scheduled interval for the Ephemery testnet, and then added validator clients, and then again to add web3signing integration.

### Ethereum staking from your wallets & private secret managers.

Staking, in fact all of web3 is in fact only one thing, and that thing is the ability to sign & transmit signed messages with your keys. We've decoupled that everything else, the high technical barriers, and high costs, so that it no longer requires you to give up access to your validator keys to participate in staking, taking a step towards making validators, defined by their signing keys decentralized at scale. 

Event driven signature automation for Ethereum staking using our synthetic staking beacon, which bundles middleware like slashing protection and mev into the service which streams validator messages to you on demand to serverless functions, hosted signers, or mobile apps, with the benefit of letting you stake from your wallet without anyone having access to your signing or withdrawal keys and without any infrastructure setup, with only a few lines of code.

### How much will staking services cost for Ethereum?

$10/mo solo or large scale enterprise staking for Ethereum per validator.

### How do I setup validators?

You can use our UI which is in beta, which takes you from 0 -> deployed & serviced validators using Zeus without any coding, technical know-how, you get the picture. You can set it up within 5 minutes, and you have ZERO ongoing maintenance after setup, which is a first of its kind when it compared with other validator services you can get.

> https://cloud.zeus.fyi/signup

Networks supported:

#### Mainnet - Request access. Includes MEV.
#### Goerli - Request access. Includes MEV. 
#### Ephemery - All users have access by default.

It also comes with an industry first toolkit of secure secrets key generation and deposit data generation. Your keys never get written to disk, or exposed at any point, and all of the secrets operations and storage happen on YOUR account not ours, we're 100% non-custodial for signing key & withdrawal keys. The next best alternative is pure cold storage generation. Read more here.

https://medium.com/zeusfyi/in-depth-overview-serverless-aws-lambda-for-ethereum-staking-with-zeus-6ad87d3be77f

Read our beta users notes: 

https://docs.google.com/document/d/1ciNOEJNEOkKFfhi0bQyhHydhdfp1zi8V0OHm-0oRN04/edit?usp=sharing

Ephemery explorers

> https://beaconchain.ephemery.pk910.de

> https://explorer.ephemery.pk910.de

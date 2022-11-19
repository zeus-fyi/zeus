## Zeus ##

Zeus is used to create, deploy, and manage infrastructure programmatically via APIs and orchestration. Using this tooling I was able to create an orchestrated setup of a lighthouse-geth beacon in less than an hour. It can reliably install/uninstall kubernetes workloads in seconds. See the test/mocks section for a copy of the configs I used.  

### API Endpoints 

Documentation and code examples are found here
[API_README.md](https://github.com/zeus-fyi/zeus/blob/main/pkg/zeus/API_README.md)

How to use the test suite to setup your own api calls
[README.md](https://github.com/zeus-fyi/zeus/blob/main/pkg/zeus/README.md)

The test directory contains useful mocks and tools for interacting with the API. It also contains a useful
config-sample.yaml, convert this to config.yaml and set your bearer token here, which then allows you to
use the demo code to create your first api request in seconds

### Beta Version Overview 

1. Automates translation of kubernetes yaml configurations into representative SQL models
2. Users upload these infrastructure configurations via API where they are stored in the DB
3. Users can then query the contents of these infrastructure components, deploy, mutate or destroy them on demand
   
### Currently Supported Workload Types

1. Deployments
2. StatefulSets
3. Services
4. ConfigMaps
5. Ingresses

### Pods Controller 

1. GetPods
2. GetPodLogs
3. PortforwardReqToPods
4. DeletePods

Not every possible field type is supported in the MVP version, but the most common ones are.

See demos section for api calls, you'll need to get a bearer token from us first while we're in beta testing. More documentation to come.
Schedule a demo: https://calendly.com/alex-zeus-fyi/zeus-demo

![Screenshot 2022-11-17 at 8 09 48 PM](https://user-images.githubusercontent.com/17446735/202614955-2708063e-1547-4dae-9332-f712102c287e.png)

### Cluster Topology Class Hierarchy Definitions ###

Cluster class creation coming soon for better organziation and easier management of more complex app configurations, here's an overview of how the hiearchy system will work.

### Highest to Lowest Level ###

### Deployable Topologies ###

### Systems Orchestration Topology ###

One or many network matrix system topologies that are combined with orchestration workflows from Artemis (coming soon), and Zeus, to build complex control flows and/or sequenced network states. This could be an enterprise fleet of infrastructure, or a simulation of a large testnet, or a built-in orchstration flow for automating web3 smart contract interactions for users and sending them notifications and prompts.

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

## Hercules ##

Hercules is web3 middleware that manages web3 infrastruture and connections to other middleware packages, such as web3signer, chain snapshot downloading, setting up mev-boost, orchestrating transcations and client switching, key generation and management, and much more coming soon.

### Upcoming Features Overview 

#### Timeline ~ End of Nov 22'

Automated blockchain snapshot downloads and setup for a lighthouse-geth beacon.

#### Timeline ~ Early January 23'

### Automated web3 infrastructure setup

Completely orchestrated and automated web3 infra. Sets up infrastructure on demand, automates run books, sets up mev-boost, web3signer, adds snapshot chain download integration, automates devops that's done by hand today, automates interactions for seeding and withdrawing validators, automates upgrades, automates configuration setup and verification, automates notifications and rewards info. Enables web3 staking infrastructure to be portable across cloud, and for vendor switching on demand. Starting with Ethereum.

#### Timeline ~ Early January 23'

### Automated web3 interactions and orchestrations 
 
```
IN: Params(Contract ABI, Address, Network, UsersAddresses)

OUT: Array(Funcs, Params to Tune/Approve)
```
Set up an orchestrated highly reliable web3 action such as creating a validator, withdrawing from a smart contract,  or
sending transactions or notifications based on event triggers such as smart contract state changes. 

#### Timeline ~ Q1/Q2 23'

### Automated web3 network setup for large scale private network testing

Create network from scratch that can replicate the size of mainnet, starting with Ethereum.

```
Runs Genesis -> Seeds Validators -> Deploys Validator Infra -> Metrics/Data
```



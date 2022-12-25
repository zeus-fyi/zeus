## Zeus ##

Zeus is an evolution of web container orchestration into web system building and orchestration. It changes the paradigm into one that unifies configuration with the underlying orchestrator, instead of decoupling them into confusing messes like helm and terraform. It's designed to solve complex infra configuration, operation, and building bottlenecks as dramatically as those the printing press solved for the book industry, and is capable of building systems at the scale and complexity of AWS by unifying multi-network kubernetes node pools.

By unifying the infra ecosystem tools of today it allows the acceleration of infra development exponentially overtime with each new open source cookbook & system template allowing you to glue sophisticated distributed systems together in zero time, which lets small decentralized solo & small team builders compete directly against large incumbents at the edge of technology by erasing the need for large investment in infra & devops departments that are needed now to even get started.

Zeus will you build & operate complex web systems as complex as AWS in timeframes that are unheard of today, and is doing this for web3 infra right now.

## Cookbooks ##

Contains common web3 components, their setups in kubernetes, and startup commands, and useful tools for debugging, interacting, and automating actions.

### Cookbook Structure ###

#### ```protocol/components```

Cookbook items listed by protocol & component class. Eg. Ethereum has a beacon component group. 

#### ```../../infra```

Contains kubernetes config setup templates

#### ```../../constants```
#### ```../../actions```
#### ```../../logs```

Example actions do log dumps, pod restarts, configuration changes on demand for k8s applications, and more.

### ```zeus/cookbooks/ethereum/beacons/beacon_cluster_test.go ```

See this test case to see how a beacon cluster class was created. Use a versioned release tag to ensure consistent stable setups if using any of the examples, the main branch is frequently in flux.

#### Ethereum ####

Contains full kubernetes infra setup for a lighthouse-geth beacon with snapshot download capability, and common interactions for developing & debugging an ethereum beacon. 

## Zeus Apps & Clients ##

### Zeus Client

#### ```pkg/zeus/client```

#### API Endpoints 

Documentation and code examples are found here
[API_README.md](https://github.com/zeus-fyi/zeus/blob/main/pkg/zeus/API_README.md)

How to use the test suite to setup your own api calls
[README.md](https://github.com/zeus-fyi/zeus/blob/main/pkg/zeus/README.md)

The test directory contains useful mocks and tools for interacting with the API. It also contains a useful
config-sample.yaml, convert this to config.yaml and set your bearer token here, which then allows you to
use the demo code to create your first api request in seconds

### Artemis ###

#### ```artemis.zeus.fyi``` ####
#### ```pkg/artemis/client```

Artemis is a tx orchestrator. It reliably submits & confirms ethereum transactions and logs their receipts. Chain with 
the in memory db for storing web3 signer keys to build highly reliable web3 api actions with other users and smart contracts. You'll need
a bearer token to use this client. A more advanced orchestrator that can handle high volume DeFi trading, which manages nonce sequences, sets up event triggers & scheduling, and has queriable event artifacts is in works, targeted release by end of Feb.

### Aegis ###

#### ```pkg/aegis/inmemdbs```
#### ```pkg/crypto```

Aegis is a library for securely managing crypto keys. It currently contains an in memory database for storing ethereum validators
and for storing ecdsa, eth1 wallet keys. 

### Hercules ###

#### ```apps/hercules``` ####
#### ```pkg/hercules/client```

Hercules is web3 middleware that manages web3 infrastructure and connections to other middleware packages, such as web3signer, chain snapshot downloading, setting up mev-boost, orchestrating transactions and client switching, key generation and management, and much more coming soon.

It also contains useful apis to debug and troubleshoot web3 infrastructure.

```go
    // manage jwt tokens
	e.POST("/jwt/create", hercules_jwt_route.JwtHandler)
	e.POST("/jwt/replace", hercules_jwt_route.JwtReplaceHandler)
	
    // used to initiate downloads for blockchain db snapshots 
	e.POST("/snapshot/download", hercules_chain_snapshots.DownloadChainSnapshotHandler)

    // used to suspend, start, kill, resume processes inside the container
	e.POST("/routines/suspend", hercules_routines.SuspendRoutineHandler)
	e.POST("/routines/start", hercules_routines.StartAppRoutineHandler)
	e.POST("/routines/resume", hercules_routines.ResumeProcessRoutineHandler)
	e.POST("/routines/kill", hercules_routines.KillProcessRoutineHandler)

    // used to wipe your data dir contents
	e.POST("/routines/disk/wipe", hercules_routines.WipeDiskHandler)

    // used to query disk & memory stats in your container
	e.GET("/host/disk", host.GetDiskStatsHandler)
	e.GET("/host/memory", host.GetMemStatsHandler)
	
    // import to secure in memdb
	e.POST("/import/validators", aegis.ImportValidatorsHandler)
```

### Snapshots ###

#### ```apps/snapshots``` ####

Snapshot app is embedded into the hercules docker app, and it can be used as an init container to download snapshot data on new node creation.

### Hades Library

#### ```pkg/hades```

Hades is used to interact with kubernetes workloads via API, and can apply saved Zeus workloads & cookbooks onto your own in house infrastructure.

## Zeus Users ##

### Beacon API ###

Users with bearer tokens are able to access our common mainnet Ethereum beacon API at https://eth.zeus.fyi, which supports both consensus client & exec client requests.

The ephemeral beacon is open to anyone, no auth is requried. Large Eth quantities for testing is available on request.

### https://eth.ephemeral.zeus.fyi

### Beacon Indexer ###

#### ```https://apollo.eth.zeus.fyi```
#### ```pkg/apollo```

Users with bearer tokens are able to access our common beacon balance & status indexer at https://apollo.eth.zeus.fyi, which indexes mainnet validator balances and statuses. It only indexes from epoch 169,000+. Once the DB reaches near capacity it removes the trailing 5k or so epochs, all the previous data is archived and accessible by request. It contains an updatedAt field for validator statuses, so you can tell how recent the status update was. It follows head behind ~10 epochs, tracking the finalized checkpoint range plus some small margin. The API is limited to 10k records per request. Requesting >10k, or requesting an epoch that isn't indexed yet will usually result in a null response. Better error messages will come soon though. See the pkg section for the apollo client which shows you how to use it, and makes it easy to integrate directly using the client.

### Snapshot Downloads ###

Snapshot download urls for mainnet geth & lighthouse available on request.

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

### Pods Controller 

1. GetPods
2. GetPodLogs
3. PortforwardReqToPods
4. DeletePods

Not every possible field type is supported in the MVP version, but the most common ones are.

See demos section for api calls, you'll need to get a bearer token from us first. More documentation to come.
Schedule a demo: https://calendly.com/alex-zeus-fyi/zeus-demo
Request a bearer token, email: alex@zeus.fyi

![Screenshot 2022-11-17 at 8 09 48 PM](https://user-images.githubusercontent.com/17446735/202614955-2708063e-1547-4dae-9332-f712102c287e.png)

### Cluster Topology Class Hierarchy Definitions ###

Network class creation coming soon for better organziation and easier management of more complex ecosystem configurations, here's an overview of how the hierarchy system will work.

### Highest to Lowest Level ###

### Deployable Topologies ###

### Network Orchestration Topology ###

One or many network matrix system topologies that are combined with orchestration workflows from Artemis, and Zeus, to build complex control flows and/or sequenced network states. This could be an enterprise fleet of infrastructure, a complex devops operation done at scale, or a simulation of a large testnet, or a built-in orchstration flow for automating web3 smart contract interactions for users and sending them notifications and prompts.

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

## Upcoming Features Overview ##

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



## Zeus ##

Zeus is an evolution of web container orchestration into web ecosystems orchestration. It changes the paradigm into one that unifies configuration with the underlying orchestrator, instead of decoupling them into confusing messes like Helm, Terraform, and GitOps, while also reducing the operational complexity of building with Kubernetes significantly and is capable of building systems at the scale and complexity of AWS (without the terrible UX/UI) by unifying multi-network Kubernetes node pools with SQL driven relationship building & querying. 

We're not just a web3 company, we have a lot of experience in crypto cloud tech so that's why it's our first product line theme, we'll be offering many  web2 focused products by the end of the year in addition to advancing our web3 products.

### Event driven signature automation for Ethereum staking, and web3 dapp interactions

Event driven signature automation for Ethereum staking using our synthetic staking beacon, an in-house technology invention that dramatically lowers the infrastucture costs of traditional enterprise staking architectures by 100x+ and bundles middleware like slashing protection and mev into the service. Which also comes with the added benefit of letting you stake from your wallet without anyone having access to your signing or withdrawal keys and without any infrastructure setup, with only a few lines of code. For those who do want hosted cloud signers, you'll have flexible 1-click style deployable hosted signers.

#### How much will staking services cost for Ethereum?

$10/mo solo or large scale enterprise staking for Ethereum per validator, all inclusive, and unlike everyone else in the industry, we're not taking a cut of your staking yield, or locking you in to any contracts, smart contracts, or anything else. We'd rather give you a better experience in web3, one that is fun to interact with, helps you build communities, helps attracts new users to the web3 ecosystem, simply by making it user-friendly, and we're betting that is worth a lot more than extracting as much money as we can from our users.

### How do I setup validators?

Roughly these are the main steps (if using the AWS Lambda option, for other options send us a message)

1. Generate ethereum keystores & deposit them. See cookbooks/ethereum/automation, or use the Ethereum staking repo below to generate them.

https://github.com/ethereum/staking-deposit-cli

2. Decrypt your keystores and then re-encrypt them with your Age encryption key into a keystores.tar.gz.age file, then zip it into a keystores.zip file.

See the test cases here for how to do this:
#### apps/cookbooks/serverless/ethereum/bls_signatures/inmemfs/inmemfs_test.go

3. Create a secret that stores your Age public and private key values in AWS secret manager. 

Secret Name: Choose any name, you'll need this name to send us.

In AWS Secret Manager
   a. Secret Key: (Your Age public key, starts with age...)
   b. Secret Value: (Your Age private key, starts with AGE-SECRET-KEY-...)

4. Create the AWS Lambda function. See this google doc for details. You can use the main.zip file pre-built in the below directory or build it yourself.

#### build_artifacts/serverless/bls_signatures/main.zip

https://docs.google.com/document/d/1Owp7nK6470WuOHPpFU1i7l4F_liCXzQ9-TfDRm0ndBw/edit?usp=sharing

5. Send a validator service request. 

See pkg/hestia, which provides a client that calls the API, and shows how you post a request to the endpoint below to service new validators in a test case.

#### https://hestia.zeus.fyi/v1/validators/service/create 

## Cookbooks ##

Contains common web2 & web3 building components like ethereum infra setups with customization driven through code, blurring the line between infra configuration and app development, and contains microservice designs & patterns like api servers and injectable choreography, their setups in kubernetes, golang, docker, and startup commands, and useful tools for debugging, interacting, and automating actions.

### Cookbook Structure ###

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

### Artemis ###

#### ```artemis.zeus.fyi``` ####
#### ```pkg/artemis/client```

Artemis is a tx orchestrator. It reliably submits & confirms ethereum transactions and logs their receipts. Chain with 
the in memory db for storing web3 signer keys to build highly reliable web3 api actions with other users and smart contracts. You'll need
a bearer token to use this client. A more advanced orchestrator that can handle high volume DeFi trading, which manages nonce sequences, sets up event triggers & scheduling, and has queriable event artifacts is in works, targeted release by end of March.

### Aegis ###

#### ```pkg/aegis/inmemdbs```
#### ```pkg/crypto```

Aegis is a library for securely managing crypto keys. It currently contains an in memory database for storing ethereum validators
and for storing ecdsa, eth1 wallet keys. 

### Hercules ###

#### ```apps/hercules``` ####
#### ```pkg/hercules/client```

Hercules is web3 middleware that manages web3 infrastructure and connections to other middleware packages. It also contains useful apis to debug and troubleshoot web3 infrastructure.

### Snapshots ###

#### ```apps/snapshots``` ####

Snapshot app is embedded into the hercules docker app, and it can be used as an init container to download snapshot data on new node creation.

### Hades Library

#### ```pkg/hades```

Hades is used to interact with kubernetes workloads via API, and can apply saved Zeus workloads & cookbooks onto your own in house infrastructure.

## Zeus Users ##

### Beacon API ###

Our ephemeral ethereum testnet beacon is open to anyone, no auth is required. Large Eth quantities for testing is available on request.

### https://eth.ephemeral.zeus.fyi

### Beacon Indexer ###

#### ```https://apollo.eth.zeus.fyi```
#### ```pkg/apollo```

Users with bearer tokens are able to access our common beacon balance & status indexer at https://apollo.eth.zeus.fyi, which indexes mainnet validator balances and statuses. It only indexes from epoch 169,000+. Once the DB reaches near capacity it removes the trailing 5k or so epochs, all the previous data is archived and accessible by request. It contains an updatedAt field for validator statuses, so you can tell how recent the status update was. It follows head behind ~10 epochs, tracking the finalized checkpoint range plus some small margin. The API is limited to 10k records per request. Requesting >10k, or requesting an epoch that isn't indexed yet will usually result in a null response. Better error messages will come soon though. See the pkg section for the apollo client which shows you how to use it, and makes it easy to integrate directly using the client.

### Snapshot Downloads ###

Snapshot download urls for mainnet geth & lighthouse available on request.

## Infra configuration is in the 14th century. Introducing the 21st century printing press

By unifying the infra ecosystem tools of today it allows the acceleration of infra development exponentially over time with each new open source cookbook & system template allowing you to glue sophisticated distributed systems together in zero time, which lets decentralized solo & small team builders compete directly against large incumbents at the edge of technology by erasing the need for large investment in infra & devops departments that are needed now to even get started.

Lastly, by solving infra configuration, mobility, and orchestration through remote API driven configuration, it allows you to run sophisticated web apps in virtually zero time on far more cost efficient bare metal cloud providers using commoditized managed kubernetes services and simple middleware we provide, which is up to 6x more cost efficient than AWS, GCP, and comparable cloud companies. Take a look for yourself, then ask yourself why you want to make Jeff Bezos richer than he already is?

#### You have alternatives

##### https://www.ibm.com/cloud/kubernetes-service
##### https://us.ovhcloud.com/public-cloud/kubernetes
##### https://www.linode.com/products/kubernetes
##### https://www.digitalocean.com/products/kubernetes
##### https://www.vultr.com/kubernetes

#### It doesn't take long to figure it out, here's some pricing links to help you out

##### https://www.ovhcloud.com/en/public-cloud/prices/
##### https://www.ibm.com/cloud/virtual-servers/pricing
##### https://instances.vantage.sh/

AWS, GCP, Azure, type cloud companies purposely have highly confusing pricing models to mislead you on costs using data transfer bills, among other sales tactics like free initial cloud usage until you're locked into their ecosystem complexity, exactly when you realize the cloud spend is outrageous. The vast majority of enterprise users simply need RAM, CPU, Bandwidth Traffic, and a few disk options like RAID setups, and HDD, SSD, NVMes. All major cloud companies have similar performance & online SLAs. Why would you want to spend 100+ engineering hours figuring out some obscure EC2 instance number that they deprecate in a year anyway? The 0.001% of people that have that need already know what they want.

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


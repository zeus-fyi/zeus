---
sidebar_position: 1
displayed_sidebar: zK8s
---

# Overview

## Cookbooks ##

Contains common microservice designs & patterns like api servers and injectable choreography, their setups in
Kubernetes,
Golang, Docker, and startup commands, and useful tools for debugging, interacting, and automating actions.

## Directory ##

### Docusaurus ###

#### ```apps/docusaurus ```

Docusaurus is a static-site generator. It builds a single-page application with fast client-side navigation,
leveraging the full power of React to make your site interactive. It provides out-of-the-box documentation features
but can be used to create any kind of site (personal website, product, blog, marketing landing pages, etc).

#### ```cookbooks/docusaurus ```

Contains full Kubernetes infra setup for Docusaurus, which is the site generator used to build our docs site. Which you
can deploy on our platform via UI, API, or any place that supports Kubernetes.

#### ```docker/docusaurus ```

You can use this docker setup as a template to build your own docusaurus site.

### Microservices ###

Contains full Kubernetes infra setup templates for microservices, simulating loads, injectable choreography for
clusters, and more continually being added.

#### ```microservices/choreography ```

#### ```microservices/deployments ```

#### ```microservices/load_simulator ```

### Hades ###

#### ```cookbooks/hades ```

Contains a minimal Kubernetes Api server for external orchestration by us, or can be entirely used by yourself to apply
templates & cookbooks via API.

### Redis ###

#### ```cookbooks/redis ```

Contains full Kubernetes infra setup for open source BSD 3-Clause version of Redis with master-replica and cluster
configurations. We also include our Redis t-digest integration in our default Docker image.
You can find the pre-built bundle on our Docker repo zeusfyi/redis:latest

### KeyDB ###

#### ```cookbooks/keydb ```

Contains full Kubernetes infra setup for keyDB, which is a forked multi-threaded version of Redis that is maintained by
Snap Inc.

## Web3

### Avax ####

#### ```avax/node``` ####

Contains Avalanche node setups for use on AWS, DigitalOcean, GCP on Kubernetes

### Cosmos ####

#### ```cosmos/node``` ####

Contains Cosmos node setups for use on AWS, DigitalOcean, GCP on Kubernetes. We also have a detailed Medium post
guide [How To Build with zK8s](https://medium.com/zeusfyi/how-to-build-on-zeus-f1e40e529377)

### Sui ####

#### ```sui/infra``` ####

Contains multiple Sui node setups for use on AWS, DigitalOcean, GCP, on local NVMe disks as well as Block Storage.
Additionally, contains NVMe setup instructions
for your own kubernetes environments.

#### ```sui/actions``` ####

Includes a client that's setup to interact with the Sui node inside the cluster via port-forwarding.

### Ethereum ####

#### ```protocol/components```

Contains smart contract automation kits. This testcase shows a full end-end seed, create, and deposits validators on the
Ethereum ephemery testnet.

#### ```ethereum/automation```

#### ```ethereum/automation/deposits_test.go ```

Cookbook items listed by protocol & component class. Eg. Ethereum has a beacon component group. Contains Kubernetes
config setup templates. Here's a few example paths. Also contains an actions folder, which does log dumps, pod restarts,
configuration changes on demand for k8s applications, and more.

#### ```beacons/infra/consensus_client```

#### ```beacons/infra/exec_client```

#### ```validators/infra/validators```

#### ```web3signers/infra/consensys_web3signer```

Complete, and powerful Ethereum infra automation

#### ```beacons/beacon_cluster_test.go ```

#### ```validators/validator_cluster_test.go ```

#### ```web3signers/web3signer_cluster_test.go ```

See this test case to see how a beacon cluster class was created, and then extended to support choreography to reset
configs on a scheduled interval for the Ephemery testnet, and then added validator clients, and then again to add
web3signing integration.

### Artemis ###

#### ```artemis.zeus.fyi``` ####

#### ```artemis/client```

Our web3 Artemis client is a tx orchestrator. It reliably submits & confirms ethereum transactions and logs their
receipts. Chain with
the in memory db for storing web3 signer keys to build highly reliable web3 api actions with other users and smart
contracts. You'll need
a bearer token to use this client. A more advanced orchestrator that can handle high volume DeFi trading & simulation,
which manages nonce sequences, sets up event triggers & scheduling, and has queriable event artifacts is available only
via enterprise licensing or other pre-arranged agreements at the moment.

### Hercules ###

#### ```apps/hercules``` ####

#### ```pkg/hercules/client```

Hercules is middleware that manages web infrastructure and connections to other middleware packages. It also contains
useful apis to debug and troubleshoot web infrastructure.

### Snapshots ###

#### ```apps/snapshots``` ####

Snapshot app is embedded into the hercules docker app, and it can be used as an init container to download snapshot data
on new node creation.

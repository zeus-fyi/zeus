# Web3 #

## Avax ##

#### ```avax/node``` ####

Contains Avalanche node setups for use on AWS, DigitalOcean, GCP on Kubernetes

## Cosmos ##

#### ```cosmos/node``` ####

Contains Cosmos node setups for use on AWS, DigitalOcean, GCP on Kubernetes. We also have a detailed Medium post
guide [How To Build with zK8s](https://medium.com/zeusfyi/how-to-build-on-zeus-f1e40e529377)

## Sui ##

#### ```sui/infra``` ####

Contains multiple Sui node setups for use on AWS, DigitalOcean, GCP, on local NVMe disks as well as Block Storage.
Additionally, contains NVMe setup instructions
for your own kubernetes environments.

#### ```sui/actions``` ####

Includes a client that's setup to interact with the Sui node inside the cluster via port-forwarding.

## Ethereum ##

#### ```protocol/components```

Contains smart contract automation kits. This testcase shows a full end-end seed, create, and deposits validators on the
Ethereum ephemery testnet.

#### Automation & Libraries

#### ```ethereum/automation```

#### ```ethereum/automation/deposits_test.go ```

Cookbook items listed by protocol & component class. Eg. Ethereum has a beacon component group. Contains Kubernetes
config setup templates. Here's a few example paths. Also contains an actions folder, which does log dumps, pod restarts,
configuration changes on demand for k8s applications, and more.

#### ```validators/infra/validators```

#### ```web3signers/infra/consensys_web3signer```

Complete, and powerful Ethereum infra automation

### Beacons ###

#### ```beacons/infra/consensus_client```

#### ```beacons/infra/exec_client```

### Validators ###

#### ```beacons/beacon_cluster_test.go ```

#### ```validators/validator_cluster_test.go ```

### Web3signer ###

#### ```web3signers/web3signer_cluster_test.go ```

See this test case to see how a beacon cluster class was created, and then extended to support choreography to reset
configs on a scheduled interval for the Ephemery testnet, and then added validator clients, and then again to add
web3signing integration.

## Hercules ##

#### ```apps/hercules``` ####

#### ```pkg/hercules/client```

Hercules is middleware that manages web infrastructure and connections to other middleware packages. It also contains
useful apis to debug and troubleshoot web infrastructure.

## Snapshots ##

#### ```apps/snapshots``` ####

Snapshot app is embedded into the hercules docker app, and it can be used as an init container to download snapshot data
on new node creation.

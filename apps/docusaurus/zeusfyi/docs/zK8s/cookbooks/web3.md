# Web3 #

## Avax ##

#### ```avax/node``` ####

Contains Avalanche node setups for use on AWS, DigitalOcean, GCP, and OvhCloud on Kubernetes

## Cosmos ##

#### ```cosmos/node``` ####

Contains Cosmos node setups for use on AWS, DigitalOcean, GCP, and OvhCloud on Kubernetes. We also have a detailed Medium post
guide [How To Build with zK8s](https://medium.com/zeusfyi/how-to-build-on-zeus-f1e40e529377)

## Sui ##

#### ```sui/infra``` ####

Contains multiple Sui node setups for use on AWS, DigitalOcean, GCP, on local NVMe disks as well as Block Storage.
Additionally, contains NVMe setup instructions for self-managed kubernetes environments.

#### ```sui/actions``` ####

Includes a client that's setup to interact with the Sui node inside the cluster via port-forwarding.

## Anvil (Foundry) ##

#### ```ethereum/foundry```

Contains anvil server setup for use on AWS, DigitalOcean, GCP, and OvhCloud on Kubernetes

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

## Ethereum Staking ##

Staking, in fact all of web3 is in fact only one thing, and that thing is the ability to sign & transmit signed messages
with your keys. We've decoupled that everything else, the high technical barriers, and high costs, so that it no longer
requires you to give up access to your validator keys to participate in staking, taking a step towards making
validators, defined by their signing keys decentralized at scale.

Event driven signature automation for Ethereum staking using our synthetic staking beacon, which bundles middleware like
slashing protection and mev into the service which streams validator messages to you on demand to serverless functions,
hosted signers, or mobile apps, with the benefit of letting you stake from your wallet without anyone having access to
your signing or withdrawal keys and without any infrastructure setup, with only a few lines of code.

### How much will staking services cost for Ethereum?

$10/mo solo or large scale enterprise staking for Ethereum per validator.

### How do I setup validators?

You can use our UI which is in beta, which takes you from 0 -> deployed & serviced validators using Zeus without any
coding, technical know-how, you get the picture. You can set it up within 5 minutes, and you have ZERO ongoing
maintenance after setup, which is a first of its kind when it compared with other validator services you can get.

> https://cloud.zeus.fyi/signup

Networks supported:

#### Mainnet - Request access. Includes MEV.

#### Goerli - Request access. Includes MEV.

#### Ephemery - All users have access by default.

It also comes with an industry first toolkit of secure secrets key generation and deposit data generation. Your keys
never get written to disk, or exposed at any point, and all of the secrets operations and storage happen on YOUR account
not ours, we're 100% non-custodial for signing key & withdrawal keys. The next best alternative is pure cold storage
generation. Read more here.

https://medium.com/zeusfyi/in-depth-overview-serverless-aws-lambda-for-ethereum-staking-with-zeus-6ad87d3be77f

Read our beta users notes:

https://docs.google.com/document/d/1ciNOEJNEOkKFfhi0bQyhHydhdfp1zi8V0OHm-0oRN04/edit?usp=sharing

Ephemery explorers

> https://beaconchain.ephemery.pk910.de

> https://explorer.ephemery.pk910.de


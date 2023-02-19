## Zeus ##

Zeus is an evolution of web container orchestration into web ecosystems orchestration. It changes the paradigm into one that unifies configuration with the underlying orchestrator, instead of decoupling them into confusing messes like Helm, Terraform, and GitOps, while also reducing the operational complexity of building with Kubernetes significantly and is capable of building systems at the scale and complexity of AWS (without the terrible UX/UI) by unifying multi-network Kubernetes node pools with SQL driven relationship building & querying. 

We're not just a web3 company, we have a lot of experience in crypto cloud tech so that's why it's our first product line theme, we'll be offering many  web2 focused products by the end of the year in addition to advancing our web3 products. 

Here are some early demo videos of infrastructure building and interfacing with the pods controller (pre-cluster & matrix version). We'll replace them with higher quality demos of the latest cluster building suite over the next few months.

https://drive.google.com/drive/folders/1tACuTADLbmYT7vS8qBqY_l5buZh_cn__?usp=sharing

### Ethereum staking from your wallets & private secret managers.

Event driven signature automation for Ethereum staking using our synthetic staking beacon, an in-house technology invention that dramatically lowers the infrastucture costs of traditional enterprise staking architectures by 100x+ and bundles middleware like slashing protection and mev into the service. Which also comes with the added benefit of letting you stake from your wallet without anyone having access to your signing or withdrawal keys and without any infrastructure setup, with only a few lines of code. For those who do want hosted cloud signers, you'll have flexible 1-click style deployable hosted signers.

### How much will staking services cost for Ethereum?

$10/mo solo or large scale enterprise staking for Ethereum per validator.

### How do I setup validators?

See details here:```builds/serverless/README.md```

You can use our automation setup to setup your own validators in 1 click, or step by step using the config value

```text
You can copy the sample-config.yaml into a new config.yaml and use that, or you can use the cli flag directly.
      --automation-steps string          AUTOMATION_STEPS: select which steps to automate and which order, using a comma separated list of numbers. default is all steps in order (default "all")
```
```text
######################################
##    Automation Step Selection     ##
######################################

# you can replace all with a comma separated list of steps to run
# e.g. only run step 7 to verify the lambda function
# or only step 9 to send the validator deposits

# STEP "1", "createSecretsAndStoreInAWS"
# STEP "2", "createInternalLambdaUser"
# STEP "3", "generateValidatorDeposits"
# STEP "4", "createLambdaFunctionKeystoresLayer"
# STEP "5", "createLambdaFunction"
# STEP "6", "createExternalLambdaUser"
# STEP "7", "verifyLambdaFunction"
# STEP "8", "createValidatorServiceRequestOnZeus"
# STEP "9", "sendValidatorDeposits"

# ACTIONS keywords
# all        - will run steps 1-9
# serverless - will run steps 1-7

# AUTOMATION_STEPS: set this config variable with keywords

# ACTIONS keywords
# all        - will run steps 1-9
# serverless - will run steps 1-7

# HELPERS: keywords
# use these keywords to fetch the secrets from aws secret manager and print them to the console

# getAgeEncryptionKeySecret
# getMnemonicHDWalletPasswordSecret
# getExternalLambdaAccessKeys

# updateLambdaKeystoresLayerToLatest - this will update the lambda function with the latest keystores layer, if you run createLambdaFunctionKeystoresLayer
# again it will create a new layer with the latest keystores.zip, so you need to run this to update the lambda function with the latest layer (e.g. if you want to add more keystores)

#############################################
## serverless automation minimum settings  ##
#############################################

You will need these values to be set at minimum to run the aws automation for steps 1-7, this account should have permissions to create lambda functions, users, roles, policies, and secrets in secret manager.
      --aws-account-number string        AWS_ACCOUNT_NUMBER: aws account number
      --aws-access-key string            AWS_ACCESS_KEY: aws access key, which needs permissions to create iam users, roles, policies, secrets, and lambda functions and layers
      --aws-secret-key string            AWS_SECRET_KEY: aws secret key

#############################################
## validator service request via API      ##
#############################################

You will need these values to be set at minimum to run the automation for step 8. Send us a message to get a bearer token
      --bearer string                    BEARER: bearer token for validator service on zeus
      --fee-recipient string             FEE_RECIPIENT_ADDR: fee recipient address for validators service on zeus

#############################################
## validator deposit submission            ##
#############################################      
      
You will need these values to be set at minimum to run the automation for step 9. 
    --eth1-addr-priv-key string        ETH1_PRIVATE_KEY: eth1 address private key for submitting deposits

If you are using this for a network other than ephemery you will need to update these values as well
      --network string                   NETWORK: network to run on mainnet, goerli, ephemery, etc (default "ephemery")
      --node-url string                  NODE_URL: beacon for getting network data for validator deposit generation & submitting deposits (default "https://eth.ephemeral.zeus.fyi") 

#############################################
## makefile commands                       ##
#############################################      

build.staking.cli:
	go build -o ./builds/serverless/bin/serverless ./builds/serverless

# generates new mnemonic, age encryption key, uses default hd password if none provided, and creates keystores
# zipped age encrypted file for serverless app --keygen true/false will toggle new keygen creation

VALIDATORS_COUNT := 0
AUTOMATION_STEPS := serverless
serverless.automation:
	./builds/serverless/bin/serverless --validator-count $(VALIDATORS_COUNT) --automation-steps $(AUTOMATION_STEPS)

serverless.validator.gen:
	./builds/serverless/bin/serverless --validator-count $(VALIDATORS_COUNT) --automation-steps generateValidatorDeposits

serverless.verify:
	./builds/serverless/bin/serverless --automation-steps verifyLambdaFunction

serverless.service:
	./builds/serverless/bin/serverless --automation-steps createValidatorServiceRequestOnZeus

ETH1_PRIV_KEY := ""
# you will need an eth1 address and it must have 32 Eth + gas fees to deposit per validator
serverless.submit.deposits:
	./builds/serverless/bin/serverless --keygen false --submit-deposits true --eth1-addr-priv-key $(ETH1_PRIV_KEY) --automation-steps sendValidatorDeposits

AWS_ACCOUNT_NUMBER:= ""
AWS_ACCESS_KEY := ""
AWS_SECRET_KEY := ""
BEARER := ""

serverless.deploy.all.cli:
	./builds/serverless/bin/serverless --aws-account-number $(AWS_ACCOUNT_NUMBER) --aws-access-key $(AWS_ACCESS_KEY) --aws-secret-key $(AWS_SECRET_KEY) --eth1-addr-priv-key $(ETH1_PRIV_KEY) --bearer $(BEARER) --automation-steps all

serverless.deploy.all.config:
	./builds/serverless/bin/serverless --automation-steps all


EXAMPLES:

Full automation using cli params:
    make serverless.deploy.all.cli AWS_ACCOUNT_NUMBER=accountNumber AWS_ACCESS_KEY=accessKey AWS_SECRET_KEY=secretKey ETH1_PRIV_KEY=0xYourPrivateKey BEARER=ZeusBearerToken
    
If you have builds/serverless/config.yaml setup with your values, you can run:
    make serverless.deploy.all.config

If you want to make changes to the build app, you can run this to rebuild the helper app:
    make build.staking.cli
```
For ephemery network deposits you can check the status of your validator here: https://beaconchain.ephemery.pk910.de/validators/eth1deposits

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


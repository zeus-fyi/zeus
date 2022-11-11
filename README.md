# Zeus #

Zeus is used to create, deploy, and manage infrastructure programmatically via APIs and orchestration. 

See demos section for api calls, you'll need to get a bearer token from us first while we're in beta testing. More documentation to come.
Schedule a demo: https://calendly.com/alex-zeus-fyi/zeus-demo

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

Not every possible field type is supported in the MVP version, but the most common ones are.
StatefulSets are still undergoing setup of end to end integration testing.

### Upcoming Features Overview 

#### Timeline ~ Early January 23'

### Automated web3 infrastructure setup

-> Completely orchestrated and automated web3 infra. Sets up infrastructure on demand, automates run books, 
automates devops that's done by hand today, automates interactions for seeding validators, automates upgrades,
automates configuration setup and verification. Enables web3 staking infrastructure to be portable across cloud, 
and for vendor switching on demand. Starting with Ethereum.

#### Timeline ~ Early January 23'

### Automated web3 interactions and orchestrations 
 
```
IN: Params(Contract ABI, Address, Network, UsersAddresses)

OUT: Array(Funcs, Params to Tune/Approve)
```
-> Set up an orchestrated highly reliable web3 action such as creating a validator, withdrawing from a smart contract,  or
sending transactions or notifications based on event triggers such as smart contract state changes. 

#### Timeline ~ Q1/Q2 23'

### Automated web3 network setup for large scale private network testing

Create network from scratch that can replicate the size of mainnet, starting with Ethereum.

```
Runs Genesis -> Seeds Validators -> Deploys Validator Infra -> Metrics/Data
```



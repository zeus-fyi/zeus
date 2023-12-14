---
sidebar_position: 6
displayed_sidebar: zK8s
---

# Ethereum

## Setup

You can use the configs test setup `test/configs` via creating a config.yaml and setting `BEARER: "{YOUR_API_KEY}"`
and add your API key to use any of our test cases on this repo, or for this example directory specifically,
you can replace the bearer token in the ethereum_test.go file with your own.

```go
s.Tc = configs.InitLocalTestConfigs()
s.ZeusTestClient = zeus_client.NewDefaultZeusClient(s.Tc.Bearer)
```

## Default Setup Options

You provide a param to specify consensus client type, and execution client type, and network, and if
you want it to be a private network or not and call the base function to create a cluster definition.

```go
beaconConfig := ethereum_beacon_cookbooks.BeaconConfig{
ConsensusClient:    consensusClient,
ExecClient:         execClient,
Network:            network,
WithIngress:        true,
WithServiceMonitor: false,
WithChoreography:   false,
}
cd := ethereum_beacon_cookbooks.CreateClientClusterDefWithParams(bc)
```

## Customizing the beacon

To customize startup scripts, resource requests, docker container images, etc you can update the constants and
parameters in:

### Consensus Client Config Drivers: `cookbooks/ethereum/beacons/beacon_consensus_config_driver.go`

### Execution Client Config Drivers: `cookbooks/ethereum/beacons/beacon_exec_client_config_driver.go`

# Dashboard Usage

## Apps Page

After you've created your beacon definition, you can navigate to the Apps page to deploy your beacon. Click on the
app row to view the beacon app deployment page.

![AppsPage](https://github.com/zeus-fyi/zeus/assets/17446735/5d0066d0-9e14-4a1b-989f-35ca9f741fd3)

## App Deployment View

Select the servers that make sense for your application, and click the "Deploy" button to deploy your beacon.
It will provision the nodes, and deploy the beacon app to the nodes. It summarizes your resource requests per
workload, and displays the total cost, including any block storage costs.

![TaintedServerDeployment](https://github.com/zeus-fyi/zeus/assets/17446735/dc968bcf-c124-4df0-908e-bc0358b51ddc)

### Navigating to your cluster devops page

You can navigate to all apps with the same class name as your beacon using the cluster view. Find your beacon and click
the row button to view the beacon page. Delete will delete the beacon and the infra in that namespace,
but it does not deprovision servers. You can do that in the Compute page.

![ClusterView](https://github.com/zeus-fyi/zeus/assets/17446735/569f0daa-04dd-457b-a32f-ed57351d1f7b)

You can use the dashboard UI to view the status of your beacon and the nodes in your cluster. It's also useful for
rapid development & debugging.

![Dashboard](https://github.com/zeus-fyi/zeus/assets/17446735/30869445-89b9-4bd6-bf1f-28c8154fd17f)

## Rapid Development

The zK8s App page button will take you to the infra config page.

![ConfigChangesUI](https://github.com/zeus-fyi/zeus/assets/17446735/69aec498-3679-4e74-ab84-acab0c5fb54f)

You can make changes to the config and click the "Apply" button to update your beacon configuration. Back on the beacon
page, you can click the "Deploy Latest" button to upgrade your beacon to the new config.

## Beacon API

If you selected an ingress for your beacon, the namespace of your cluster will be the same as the beacon url path

    namespace: ethereum-beacon-mainnet-lighthouse-geth-7627d4b5
    domain: zeus.fyi
    url: ethereum-beacon-mainnet-lighthouse-geth-7627d4b5.zeus.fyi
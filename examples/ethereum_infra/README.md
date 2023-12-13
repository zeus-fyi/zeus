# Ethereum Infra Examples

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
you want it to be a private network or not and call the base function.

```go
    cd := ethereum_beacon_cookbooks.GetClientClusterDef(consensusClient, execClient, network, true)
```

## Customizing the beacon

To customize startup scripts, resource requests, docker container images, etc you can update the constants and
parameters in:

### Consensus Client Config Drivers: `cookbooks/ethereum/beacons/beacon_consensus_config_driver.go`

### Execution Client Config Drivers: `cookbooks/ethereum/beacons/beacon_exec_client_config_driver.go`

## Dashboard UI Usage

### Navigating to your beacon

You can navigate to all apps with the same class name as your beacon using the cluster view. Find your beacon and click
the row button to view the beacon page. Delete will delete the beacon and the infra in that namespace,
but it does not deprovision nodes. You can do that in the Compute page.

![ClusterView](https://github.com/zeus-fyi/zeus/assets/17446735/569f0daa-04dd-457b-a32f-ed57351d1f7b)

You can use the dashboard UI to view the status of your beacon and the nodes in your cluster. It's also useful for
rapid development & debugging.

![Dashboard](https://github.com/zeus-fyi/zeus/assets/17446735/30869445-89b9-4bd6-bf1f-28c8154fd17f)

## Rapid Development

The zK8s App page button will take you to the infra config page.

![ConfigChangesUI](https://github.com/zeus-fyi/zeus/assets/17446735/69aec498-3679-4e74-ab84-acab0c5fb54f)

You can make changes to the config and click the "Apply" button to update your beacon configuration. Back on the beacon
page, you can click the "Deploy Latest" button to upgrade your beacon to the new config.
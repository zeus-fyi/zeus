package ethereum_infra_examples

import (
	"context"
	"fmt"

	ethereum_beacon_cookbooks "github.com/zeus-fyi/zeus/cookbooks/ethereum/beacons"
	client_consts "github.com/zeus-fyi/zeus/cookbooks/ethereum/beacons/constants"
	hestia_req_types "github.com/zeus-fyi/zeus/pkg/hestia/client/req_types"
	zeus_cluster_config_drivers "github.com/zeus-fyi/zeus/zeus/cluster_config_drivers"
)

var ctx = context.Background()

func (s *EthereumInfraExamplesTestSuite) TestEthereumBeacon() {
	consensusClient := client_consts.Lighthouse
	// client_consts.Lodestar
	execClient := client_consts.Geth
	// client_consts.Reth
	network := hestia_req_types.Mainnet
	bc := ethereum_beacon_cookbooks.BeaconConfig{
		ConsensusClient:    consensusClient,
		ExecClient:         execClient,
		Network:            network,
		WithIngress:        true,
		WithServiceMonitor: false,
		WithChoreography:   false,
	}
	cd := ethereum_beacon_cookbooks.CreateClientClusterDefWithParams(bc)
	s.Assert().NotEmpty(cd)
	/*
		To customize startup scripts, resource requests, docker containers, etc you can update the constants and parameters in:

		Consensus Client Config Drivers: cookbooks/ethereum/beacons/beacon_consensus_config_driver.go
		Execution Client Config Drivers: cookbooks/ethereum/beacons/beacon_exec_client_config_driver.go
	*/
	// creates new class. no-op if class exists already
	s.testCreateEthereumBeaconClass(cd, true)
	_, err := cd.UploadChartsFromClusterDefinition(ctx, s.ZeusTestClient, true)
	s.Require().Nil(err)
}

func (s *EthereumInfraExamplesTestSuite) testCreateEthereumBeaconClass(cd zeus_cluster_config_drivers.ClusterDefinition, createClass bool) {
	if !createClass {
		return
	}
	gcd := cd.BuildClusterDefinitions()
	s.Assert().NotEmpty(gcd)
	fmt.Println(gcd)

	ccd := gcd.CreateClusterClassDefinitions(context.Background(), s.ZeusTestClient)
	s.Require().Nil(ccd)
}

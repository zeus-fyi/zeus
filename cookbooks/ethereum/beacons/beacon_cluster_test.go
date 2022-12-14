package beacon_cookbooks

import (
	"context"

	"github.com/zeus-fyi/zeus/cookbooks"
	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_req_types"
)

func (t *BeaconCookbookTestSuite) TestClusterDeploy() {
	ctx := context.Background()
	resp, err := t.ZeusTestClient.DeployCluster(ctx, cd)
	t.Require().Nil(err)
	t.Assert().NotEmpty(resp)
}

// Follow this order of commands to create a beacon class with infra, then use the above ^ to deploy it
// cd is the cluster definition
func (t *BeaconCookbookTestSuite) EndToEndTest() {
	// ethereumBeacons is a reserved keyword, so it can be global to our stored config we maintain.
	// you can replace the below with your own setup by changing the class name and following the tests.
	t.TestCreateClusterClass()
	t.TestCreateClusterBase()
	t.TestCreateClusterSkeletonBases()
	t.TestUploadBeaconCharts()
}
func (t *BeaconCookbookTestSuite) TestCreateClusterClass() {
	ctx := context.Background()
	cookbooks.ChangeToCookbookDir()

	cc := zeus_req_types.TopologyCreateOrAddBasesToClassesRequest{
		ClassName: "ethereumBeacons",
	}
	resp, err := t.ZeusTestClient.CreateClass(ctx, cc)
	t.Require().Nil(err)
	t.Assert().NotEmpty(resp)
}

func (t *BeaconCookbookTestSuite) TestCreateClusterBase() {
	ctx := context.Background()
	basesInsert := []string{"executionClient", "consensusClient"}
	cc := zeus_req_types.TopologyCreateOrAddBasesToClassesRequest{
		ClassName:      "ethereumBeacons",
		ClassBaseNames: basesInsert,
	}
	_, err := t.ZeusTestClient.AddBasesToClass(ctx, cc)
	t.Require().Nil(err)
}

func (t *BeaconCookbookTestSuite) TestCreateClusterSkeletonBases() {
	ctx := context.Background()
	basesInsert := []string{"gethHercules"}
	cc := zeus_req_types.TopologyCreateOrAddBasesToClassesRequest{
		ClassName:      "executionClient",
		ClassBaseNames: basesInsert,
	}
	_, err := t.ZeusTestClient.AddSkeletonBasesToClass(ctx, cc)

	basesInsert = []string{"lighthouseHercules"}
	cc = zeus_req_types.TopologyCreateOrAddBasesToClassesRequest{
		ClassName:      "consensusClient",
		ClassBaseNames: basesInsert,
	}
	_, err = t.ZeusTestClient.AddSkeletonBasesToClass(ctx, cc)
	t.Require().Nil(err)
}

func (t *BeaconCookbookTestSuite) TestUploadBeaconCharts() {
	ctx := context.Background()
	// Consensus
	resp, err := t.ZeusTestClient.UploadChart(ctx, beaconConsensusClientChartPath, consensusClientChart)
	t.Require().Nil(err)
	t.Assert().NotZero(resp.TopologyID)

	DeployConsensusClientKnsReq.TopologyID = resp.TopologyID
	tar := zeus_req_types.TopologyRequest{TopologyID: DeployConsensusClientKnsReq.TopologyID}
	chartResp, err := t.ZeusTestClient.ReadChart(ctx, tar)
	t.Require().Nil(err)
	t.Assert().NotEmpty(chartResp)

	err = chartResp.PrintWorkload(beaconConsensusClientChartPath)
	t.Require().Nil(err)

	// Exec
	resp, err = t.ZeusTestClient.UploadChart(ctx, beaconExecClientChartPath, execClientChart)
	t.Require().Nil(err)
	t.Assert().NotZero(resp.TopologyID)

	DeployExecClientKnsReq.TopologyID = resp.TopologyID
	tar = zeus_req_types.TopologyRequest{TopologyID: DeployExecClientKnsReq.TopologyID}
	chartResp, err = t.ZeusTestClient.ReadChart(ctx, tar)
	t.Require().Nil(err)
	t.Assert().NotEmpty(chartResp)

	err = chartResp.PrintWorkload(beaconExecClientChartPath)
	t.Require().Nil(err)
}

func (t *BeaconCookbookTestSuite) TestEphemeralStakingBeaconConfig() {
	ConfigEphemeralLighthouseGethStakingBeacon()
}

package beacon_cookbooks

import (
	"context"

	"github.com/zeus-fyi/zeus/cookbooks"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_req_types"
)

// ethereumBeacons is a reserved keyword, so it can be global to our stored config we maintain.
// you can replace the below with your own setup by changing the class name and following the tests.
var (
	className      = "ethereumEphemeralBeacons"
	execBases      = []string{"gethHercules"}
	consensusBases = []string{"lighthouseHercules"}
)

func (t *BeaconCookbookTestSuite) TestClusterDeploy() {
	ctx := context.Background()
	switch className {
	case "ethereumEphemeralBeacons":
		cd.ClusterName = className
		cd.Namespace = "ephemeral"
	}
	resp, err := t.ZeusTestClient.DeployCluster(ctx, cd)
	t.Require().Nil(err)
	t.Assert().NotEmpty(resp)
}

func (t *BeaconCookbookTestSuite) TestClusterDestroy() {
	ctx := context.Background()

	knsReq := DeployConsensusClientKnsReq
	switch className {
	case "ethereumEphemeralBeacons":
		cd.ClusterName = className
		knsReq.Namespace = "ephemeral"
	}
	resp, err := t.ZeusTestClient.DestroyDeploy(ctx, knsReq)
	t.Require().Nil(err)
	t.Assert().NotEmpty(resp)
}

// Follow this order of commands to create a beacon class with infra, then use the above ^ to deploy it
// cd is the cluster definition

func (t *BeaconCookbookTestSuite) TestEndToEnd() {
	t.TestCreateClusterClass()
	t.TestCreateClusterBase()
	t.TestCreateClusterSkeletonBases()

	switch className {
	case "ethereumBeacons":
		t.TestUploadStandardBeaconCharts()
	case "ethereumEphemeralBeacons":
		consensusClientChart.ClusterBaseName = className
		execClientChart.ClusterBaseName = className
		t.TestUploadEphemeralStakingBeaconConfig()
	}
}

func (t *BeaconCookbookTestSuite) TestCreateClusterClass() {
	ctx := context.Background()
	cookbooks.ChangeToCookbookDir()

	cc := zeus_req_types.TopologyCreateOrAddBasesToClassesRequest{
		ClassName: className,
	}
	resp, err := t.ZeusTestClient.CreateClass(ctx, cc)
	t.Require().Nil(err)
	t.Assert().NotEmpty(resp)
}

func (t *BeaconCookbookTestSuite) TestCreateClusterBase() {
	ctx := context.Background()
	basesInsert := []string{"executionClient", "consensusClient"}
	cc := zeus_req_types.TopologyCreateOrAddBasesToClassesRequest{
		ClassName:      className,
		ClassBaseNames: basesInsert,
	}
	_, err := t.ZeusTestClient.AddBasesToClass(ctx, cc)
	t.Require().Nil(err)
}

func (t *BeaconCookbookTestSuite) TestCreateClusterSkeletonBases() {
	ctx := context.Background()

	cc := zeus_req_types.TopologyCreateOrAddBasesToClassesRequest{
		ClassName:      "executionClient",
		ClassBaseNames: execBases,
	}
	_, err := t.ZeusTestClient.AddSkeletonBasesToClass(ctx, cc)

	cc = zeus_req_types.TopologyCreateOrAddBasesToClassesRequest{
		ClassName:      "consensusClient",
		ClassBaseNames: consensusBases,
	}
	_, err = t.ZeusTestClient.AddSkeletonBasesToClass(ctx, cc)
	t.Require().Nil(err)
}

func (t *BeaconCookbookTestSuite) TestUploadBeaconCharts(consensusChartPath, execChartPath filepaths.Path) {
	ctx := context.Background()
	// Consensus
	resp, err := t.ZeusTestClient.UploadChart(ctx, consensusChartPath, consensusClientChart)
	t.Require().Nil(err)
	t.Assert().NotZero(resp.TopologyID)

	DeployConsensusClientKnsReq.TopologyID = resp.TopologyID
	tar := zeus_req_types.TopologyRequest{TopologyID: DeployConsensusClientKnsReq.TopologyID}
	chartResp, err := t.ZeusTestClient.ReadChart(ctx, tar)
	t.Require().Nil(err)
	t.Assert().NotEmpty(chartResp)

	err = chartResp.PrintWorkload(consensusChartPath)
	t.Require().Nil(err)

	// Exec
	resp, err = t.ZeusTestClient.UploadChart(ctx, execChartPath, execClientChart)
	t.Require().Nil(err)
	t.Assert().NotZero(resp.TopologyID)

	DeployExecClientKnsReq.TopologyID = resp.TopologyID
	tar = zeus_req_types.TopologyRequest{TopologyID: DeployExecClientKnsReq.TopologyID}
	chartResp, err = t.ZeusTestClient.ReadChart(ctx, tar)
	t.Require().Nil(err)
	t.Assert().NotEmpty(chartResp)

	err = chartResp.PrintWorkload(execChartPath)
	t.Require().Nil(err)
}

func (t *BeaconCookbookTestSuite) TestUploadStandardBeaconCharts() {
	t.TestUploadBeaconCharts(beaconConsensusClientChartPath, beaconExecClientChartPath)
}
func (t *BeaconCookbookTestSuite) TestUploadEphemeralStakingBeaconConfig() {
	cp := beaconConsensusClientChartPath
	cp.DirOut = "./ethereum/beacons/infra/processed_consensus_client"

	ep := beaconExecClientChartPath
	ep.DirOut = "./ethereum/beacons/infra/processed_exec_client"
	ConfigEphemeralLighthouseGethStakingBeacon(cp, ep)

	cp.DirIn = cp.DirOut
	ep.DirIn = ep.DirOut
	t.TestUploadBeaconCharts(cp, ep)
}

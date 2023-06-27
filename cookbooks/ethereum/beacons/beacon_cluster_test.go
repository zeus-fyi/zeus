package ethereum_beacon_cookbooks

import (
	"context"
	"fmt"

	"github.com/zeus-fyi/zeus/cookbooks"
	client_consts "github.com/zeus-fyi/zeus/cookbooks/ethereum/beacons/constants"
	choreography_cookbooks "github.com/zeus-fyi/zeus/cookbooks/microservices/choreography"
	hestia_req_types "github.com/zeus-fyi/zeus/pkg/hestia/client/req_types"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_req_types"
)

var (
	clusterClassName       = "ethereumEphemeralBeacons"
	execSkeletonBases      = []string{"gethHercules"}
	consensusSkeletonBases = []string{"lighthouseHercules"}
	ingressBaseName        = []string{"beaconIngress"}

	ctx = context.Background()
)

func (t *BeaconCookbookTestSuite) TestClusterDeployV2() {
	cd := BeaconClusterDefinition
	_, err := cd.UploadChartsFromClusterDefinition(ctx, t.ZeusTestClient, true)
	t.Require().Nil(err)

	cdep := cd.GenerateDeploymentRequest()
	_, err = t.ZeusTestClient.DeployCluster(ctx, cdep)
	t.Require().Nil(err)
}

func (t *BeaconCookbookTestSuite) TestClusterSetupV2() {
	cd := BeaconClusterDefinition
	gcd := cd.BuildClusterDefinitions()
	t.Assert().NotEmpty(gcd)
	fmt.Println(gcd)

	gdr := cd.GenerateDeploymentRequest()
	t.Assert().NotEmpty(gdr)
	fmt.Println(gdr)

	sbDefs, err := cd.GenerateSkeletonBaseCharts()
	t.Require().Nil(err)
	t.Assert().NotEmpty(sbDefs)
}

func (t *BeaconCookbookTestSuite) TestClusterSetupWithCfgDriver() {
	cd := GetClientClusterDef(client_consts.Lighthouse, client_consts.Geth, hestia_req_types.Mainnet)
	gcd := cd.BuildClusterDefinitions()
	t.Assert().NotEmpty(gcd)
	fmt.Println(gcd)

	gdr := cd.GenerateDeploymentRequest()
	t.Assert().NotEmpty(gdr)
	fmt.Println(gdr)

	sbDefs, err := cd.GenerateSkeletonBaseCharts()
	t.Require().Nil(err)
	t.Assert().NotEmpty(sbDefs)
	_, err = cd.UploadChartsFromClusterDefinition(ctx, t.ZeusTestClient, true)
	t.Require().Nil(err)
}

func (t *BeaconCookbookTestSuite) TestLighthouseRethClusterSetupWithCfgDriver() {
	cd := GetClientClusterDef(client_consts.Lighthouse, client_consts.Reth, hestia_req_types.Mainnet)
	gcd := cd.BuildClusterDefinitions()
	t.Assert().NotEmpty(gcd)
	fmt.Println(gcd)

	gdr := cd.GenerateDeploymentRequest()
	t.Assert().NotEmpty(gdr)
	fmt.Println(gdr)

	sbDefs, err := cd.GenerateSkeletonBaseCharts()
	t.Require().Nil(err)
	t.Assert().NotEmpty(sbDefs)
	_, err = cd.UploadChartsFromClusterDefinition(ctx, t.ZeusTestClient, true)
	t.Require().Nil(err)
}

func (t *BeaconCookbookTestSuite) TestClusterDefinitionCreationV2() {
	cd := GetClientClusterDef(client_consts.Lighthouse, client_consts.Geth, hestia_req_types.Mainnet)
	gcd := cd.BuildClusterDefinitions()
	t.Assert().NotEmpty(gcd)
	fmt.Println(gcd)

	err := gcd.CreateClusterClassDefinitions(context.Background(), t.ZeusTestClient)
	t.Require().Nil(err)
}

func (t *BeaconCookbookTestSuite) TestLighthouseRethClusterDefinitionCreation() {
	cd := GetClientClusterDef(client_consts.Lighthouse, client_consts.Reth, hestia_req_types.Mainnet)
	gcd := cd.BuildClusterDefinitions()
	t.Assert().NotEmpty(gcd)
	fmt.Println(gcd)

	err := gcd.CreateClusterClassDefinitions(context.Background(), t.ZeusTestClient)
	t.Require().Nil(err)
}

func (t *BeaconCookbookTestSuite) TestClusterDeploy() {
	switch clusterClassName {
	case "ethereumEphemeralBeacons":
		Cd.ClusterClassName = clusterClassName
		Cd.Namespace = "ephemeral"
		// choreography is a reserved skeleton base keyword that will always deploy a choreography secret to support actions
		Cd.SkeletonBaseOptions = append(Cd.SkeletonBaseOptions, choreography_cookbooks.GenericChoreographyChart.SkeletonBaseName)
		Cd.SkeletonBaseOptions = append(Cd.SkeletonBaseOptions, ingressBaseName...)
	}
	resp, err := t.ZeusTestClient.DeployCluster(ctx, Cd)
	t.Require().Nil(err)
	t.Assert().NotEmpty(resp)
}

func (t *BeaconCookbookTestSuite) TestClusterDestroy() {
	knsReq := DeployConsensusClientKnsReq
	switch clusterClassName {
	case "ethereumEphemeralBeacons":
		Cd.ClusterClassName = clusterClassName
		knsReq.Namespace = "ephemeral"
	case "ethereumEphemeralValidatorCluster":
		Cd.ClusterClassName = clusterClassName
		knsReq.Namespace = "ephemeral-staking"
	}
	resp, err := t.ZeusTestClient.DestroyDeploy(ctx, knsReq)
	t.Require().Nil(err)
	t.Assert().NotEmpty(resp)
}

// Follow this order of commands to create a beacon class with infra, then use the above ^ to deploy it
// Cd is the cluster definition

func (t *BeaconCookbookTestSuite) TestEndToEnd() {
	t.TestCreateClusterClass()
	t.TestCreateClusterBase()
	t.TestCreateClusterSkeletonBases()

	switch clusterClassName {
	case "ethereumBeacons":
		t.TestUploadStandardBeaconCharts()
	case "ethereumEphemeralValidatorCluster":
		ConsensusClientChart.ClusterClassName = clusterClassName
		ExecClientChart.ClusterClassName = clusterClassName
		t.TestUploadEphemeralBeaconConfig(false)
	case "ethereumEphemeralBeacons":
		ConsensusClientChart.ClusterClassName = clusterClassName
		ExecClientChart.ClusterClassName = clusterClassName
		t.TestUploadEphemeralBeaconConfig(true)
	}
}

func (t *BeaconCookbookTestSuite) TestCreateClusterClass() {
	cookbooks.ChangeToCookbookDir()

	cc := zeus_req_types.TopologyCreateClusterClassRequest{
		ClusterClassName: clusterClassName,
	}
	resp, err := t.ZeusTestClient.CreateClass(ctx, cc)
	t.Require().Nil(err)
	t.Assert().NotEmpty(resp)
}

func (t *BeaconCookbookTestSuite) TestCreateClusterBase() {
	basesInsert := []string{
		"executionClient",
		"zeusConsensusClient",
		"beaconIngress",
		"serviceMonitorConsensusClient",
		"serviceMonitorExecClient",
		"choreography"}
	cc := zeus_req_types.TopologyCreateOrAddComponentBasesToClassesRequest{
		ClusterClassName:   clusterClassName,
		ComponentBaseNames: basesInsert,
	}
	_, err := t.ZeusTestClient.AddComponentBasesToClass(ctx, cc)
	t.Require().Nil(err)
}

func (t *BeaconCookbookTestSuite) TestCreateClusterSkeletonBases() {
	cc := zeus_req_types.TopologyCreateOrAddSkeletonBasesToClassesRequest{
		ClusterClassName:  clusterClassName,
		ComponentBaseName: "executionClient",
		SkeletonBaseNames: execSkeletonBases,
	}
	_, err := t.ZeusTestClient.AddSkeletonBasesToClass(ctx, cc)
	t.Require().Nil(err)

	cc = zeus_req_types.TopologyCreateOrAddSkeletonBasesToClassesRequest{
		ClusterClassName:  clusterClassName,
		ComponentBaseName: "zeusConsensusClient",
		SkeletonBaseNames: consensusSkeletonBases,
	}
	_, err = t.ZeusTestClient.AddSkeletonBasesToClass(ctx, cc)
	t.Require().Nil(err)

	ing := zeus_req_types.TopologyCreateOrAddSkeletonBasesToClassesRequest{
		ClusterClassName:  clusterClassName,
		ComponentBaseName: "beaconIngress",
		SkeletonBaseNames: ingressBaseName,
	}
	_, err = t.ZeusTestClient.AddSkeletonBasesToClass(ctx, ing)
	t.Require().Nil(err)

	cc = zeus_req_types.TopologyCreateOrAddSkeletonBasesToClassesRequest{
		ClusterClassName:  clusterClassName,
		ComponentBaseName: "choreography",
		SkeletonBaseNames: []string{"choreography"},
	}
	_, err = t.ZeusTestClient.AddSkeletonBasesToClass(ctx, cc)
	t.Require().Nil(err)

	cc = zeus_req_types.TopologyCreateOrAddSkeletonBasesToClassesRequest{
		ClusterClassName:  clusterClassName,
		ComponentBaseName: "choreography",
		SkeletonBaseNames: []string{"choreography"},
	}
	_, err = t.ZeusTestClient.AddSkeletonBasesToClass(ctx, cc)
	t.Require().Nil(err)

	cc = zeus_req_types.TopologyCreateOrAddSkeletonBasesToClassesRequest{
		ClusterClassName:  clusterClassName,
		ComponentBaseName: "serviceMonitorConsensusClient",
		SkeletonBaseNames: []string{"serviceMonitorConsensusClient"},
	}
	_, err = t.ZeusTestClient.AddSkeletonBasesToClass(ctx, cc)
	t.Require().Nil(err)

	cc = zeus_req_types.TopologyCreateOrAddSkeletonBasesToClassesRequest{
		ClusterClassName:  clusterClassName,
		ComponentBaseName: "serviceMonitorExecClient",
		SkeletonBaseNames: []string{"serviceMonitorExecClient"},
	}
	_, err = t.ZeusTestClient.AddSkeletonBasesToClass(ctx, cc)
	t.Require().Nil(err)
}

func (t *BeaconCookbookTestSuite) TestUploadBeaconCharts(consensusChartPath, execChartPath, ingChartPath filepaths.Path, withIngress bool) {
	// Consensus
	resp, err := t.ZeusTestClient.UploadChart(ctx, consensusChartPath, ConsensusClientChart)
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
	resp, err = t.ZeusTestClient.UploadChart(ctx, execChartPath, ExecClientChart)
	t.Require().Nil(err)
	t.Assert().NotZero(resp.TopologyID)

	DeployExecClientKnsReq.TopologyID = resp.TopologyID
	tar = zeus_req_types.TopologyRequest{TopologyID: DeployExecClientKnsReq.TopologyID}
	chartResp, err = t.ZeusTestClient.ReadChart(ctx, tar)
	t.Require().Nil(err)
	t.Assert().NotEmpty(chartResp)

	err = chartResp.PrintWorkload(execChartPath)
	t.Require().Nil(err)

	if withIngress {
		// Ingress
		resp, err = t.ZeusTestClient.UploadChart(ctx, ingChartPath, IngressChart)
		t.Require().Nil(err)
		t.Assert().NotZero(resp.TopologyID)

		DeployExecClientKnsReq.TopologyID = resp.TopologyID
		tar = zeus_req_types.TopologyRequest{TopologyID: DeployExecClientKnsReq.TopologyID}
		chartResp, err = t.ZeusTestClient.ReadChart(ctx, tar)
		t.Require().Nil(err)
		t.Assert().NotEmpty(chartResp)

		err = chartResp.PrintWorkload(ingChartPath)
		t.Require().Nil(err)
	}
}

func (t *BeaconCookbookTestSuite) TestUploadStandardBeaconCharts() {
	t.TestUploadBeaconCharts(BeaconConsensusClientChartPath, BeaconExecClientChartPath, IngressChartPath, true)
}

func (t *BeaconCookbookTestSuite) TestUploadEphemeralBeaconStakingConfig() {
	clusterClassName = "ethereumEphemeralValidatorCluster"
	t.TestUploadEphemeralBeaconConfig(false)
}

func (t *BeaconCookbookTestSuite) TestUploadEphemeralStandardBeaconConfig() {
	t.TestUploadEphemeralBeaconConfig(true)
}

func (t *BeaconCookbookTestSuite) TestUploadEphemeralBeaconConfig(withIngress bool) {
	ConsensusClientChart.ClusterClassName = clusterClassName
	ExecClientChart.ClusterClassName = clusterClassName
	IngressChart.ClusterClassName = clusterClassName

	cp := BeaconConsensusClientChartPath
	cp.DirOut = "./ethereum/beacons/infra/processed_consensus_client"

	ep := BeaconExecClientChartPath
	ep.DirOut = "./ethereum/beacons/infra/processed_exec_client"

	ing := IngressChartPath
	ing.DirOut = "./ethereum/beacons/infra/processed_beacon_ingress"

	ConfigEphemeralLighthouseGethBeacon(cp, ep, ing, withIngress)

	cp.DirIn = cp.DirOut
	ep.DirIn = ep.DirOut
	ing.DirIn = ing.DirOut
	t.TestUploadBeaconCharts(cp, ep, ing, withIngress)

	// choreography option
	choreography_cookbooks.GenericChoreographyChart.ClusterClassName = clusterClassName
	resp, err := t.ZeusTestClient.UploadChart(ctx, choreography_cookbooks.GenericDeploymentChartPath, choreography_cookbooks.GenericChoreographyChart)
	t.Require().Nil(err)
	t.Assert().NotZero(resp.TopologyID)

}

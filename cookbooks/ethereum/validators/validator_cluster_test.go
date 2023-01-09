package validator_cookbooks

import (
	"context"

	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_req_types"
	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_resp_types/topology_workloads"
)

func (t *ValidatorCookbookTestSuite) TestClusterDeploy() {
	t.TestUploadValidatorClientCharts()
	ctx := context.Background()
	resp, err := t.ZeusTestClient.DeployCluster(ctx, ValidatorClusterDefinition)
	t.Require().Nil(err)
	t.Assert().NotEmpty(resp)
}

func (t *ValidatorCookbookTestSuite) TestClusterDestroy() {
	ctx := context.Background()
	knsReq := DeployConsensusValidatorClientKnsReq
	resp, err := t.ZeusTestClient.DestroyDeploy(ctx, knsReq)
	t.Require().Nil(err)
	t.Assert().NotEmpty(resp)
}

func (t *ValidatorCookbookTestSuite) TestCreateClusterBase() {
	ctx := context.Background()
	basesInsert := []string{"executionClient", "consensusClient", consensusValidatorClientComponentBaseName, ChoreographySkeletonBase}
	cc := zeus_req_types.TopologyCreateOrAddComponentBasesToClassesRequest{
		ClusterClassName:   EphemeryValidatorClusterClassName,
		ComponentBaseNames: basesInsert,
	}
	_, err := t.ZeusTestClient.AddComponentBasesToClass(ctx, cc)
	t.Require().Nil(err)
}

func (t *ValidatorCookbookTestSuite) TestUploadValidatorClientCharts() {
	ctx := context.Background()
	// Consensus

	inf := topology_workloads.NewTopologyBaseInfraWorkload()
	err := ValidatorsChartPath.WalkAndApplyFuncToFileType(".yaml", inf.DecodeK8sWorkload)
	t.Require().Nil(err)
	EphemeralValidatorClientLighthouseConfig(inf)

	resp, err := t.ZeusTestClient.UploadChart(ctx, ValidatorsChartPath, ValidatorsChart)
	t.Require().Nil(err)
	t.Assert().NotZero(resp.TopologyID)

	DeployConsensusValidatorClientKnsReq.TopologyID = resp.TopologyID
	tar := zeus_req_types.TopologyRequest{TopologyID: DeployConsensusValidatorClientKnsReq.TopologyID}
	chartResp, err := t.ZeusTestClient.ReadChart(ctx, tar)
	t.Require().Nil(err)
	t.Assert().NotEmpty(chartResp)

	err = chartResp.PrintWorkload(ValidatorsChartPath)
	t.Require().Nil(err)
}

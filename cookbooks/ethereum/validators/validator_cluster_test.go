package validator_cookbooks

import (
	"context"

	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_req_types"
)

func (t *ValidatorCookbookTestSuite) TestClusterDeploy() {
	ctx := context.Background()
	resp, err := t.ZeusTestClient.DeployCluster(ctx, cd)
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
	basesInsert := []string{"executionClient", "consensusClient", consensusValidatorClientComponentBaseName}
	cc := zeus_req_types.TopologyCreateOrAddComponentBasesToClassesRequest{
		ClusterClassName:   ValidatorClusterClassName,
		ComponentBaseNames: basesInsert,
	}
	_, err := t.ZeusTestClient.AddComponentBasesToClass(ctx, cc)
	t.Require().Nil(err)
}

func (t *ValidatorCookbookTestSuite) TestUploadValidatorClientCharts() {
	ctx := context.Background()
	// Consensus
	resp, err := t.ZeusTestClient.UploadChart(ctx, validatorsChartPath, validatorsChart)
	t.Require().Nil(err)
	t.Assert().NotZero(resp.TopologyID)

	DeployConsensusValidatorClientKnsReq.TopologyID = resp.TopologyID
	tar := zeus_req_types.TopologyRequest{TopologyID: DeployConsensusValidatorClientKnsReq.TopologyID}
	chartResp, err := t.ZeusTestClient.ReadChart(ctx, tar)
	t.Require().Nil(err)
	t.Assert().NotEmpty(chartResp)

	err = chartResp.PrintWorkload(validatorsChartPath)
	t.Require().Nil(err)
}

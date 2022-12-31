package web3signer_cookbooks

import (
	"context"

	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_req_types"
)

func (t *Web3SignerCookbookTestSuite) TestClusterDeploy() {
	ctx := context.Background()
	resp, err := t.ZeusTestClient.DeployCluster(ctx, cd)
	t.Require().Nil(err)
	t.Assert().NotEmpty(resp)
}

func (t *Web3SignerCookbookTestSuite) TestClusterDestroy() {
	ctx := context.Background()
	knsReq := DeployWeb3SignerKnsReq
	resp, err := t.ZeusTestClient.DestroyDeploy(ctx, knsReq)
	t.Require().Nil(err)
	t.Assert().NotEmpty(resp)
}

func (t *Web3SignerCookbookTestSuite) TestCreateClusterBase() {
	ctx := context.Background()
	basesInsert := []string{web3SignerComponentBaseName}
	cc := zeus_req_types.TopologyCreateOrAddComponentBasesToClassesRequest{
		ClusterClassName:   EphemeryWeb3SignerClusterClassName,
		ComponentBaseNames: basesInsert,
	}
	_, err := t.ZeusTestClient.AddComponentBasesToClass(ctx, cc)
	t.Require().Nil(err)
}

func (t *Web3SignerCookbookTestSuite) TestUploadValidatorClientCharts() {
	ctx := context.Background()
	// Consensus
	resp, err := t.ZeusTestClient.UploadChart(ctx, web3SignerChartPath, web3SignerChart)
	t.Require().Nil(err)
	t.Assert().NotZero(resp.TopologyID)

	DeployWeb3SignerKnsReq.TopologyID = resp.TopologyID
	tar := zeus_req_types.TopologyRequest{TopologyID: DeployWeb3SignerKnsReq.TopologyID}
	chartResp, err := t.ZeusTestClient.ReadChart(ctx, tar)
	t.Require().Nil(err)
	t.Assert().NotEmpty(chartResp)

	err = chartResp.PrintWorkload(web3SignerChartPath)
	t.Require().Nil(err)
}

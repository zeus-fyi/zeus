package deployment_cookbooks

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/zeus-fyi/zeus/cookbooks"
	"github.com/zeus-fyi/zeus/test/configs"
	"github.com/zeus-fyi/zeus/test/test_suites"
	zeus_client "github.com/zeus-fyi/zeus/zeus/z_client"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_req_types"
)

type DeploymentsCookbookTestSuite struct {
	test_suites.BaseTestSuite
	ZeusTestClient zeus_client.ZeusClient
}

func (t *DeploymentsCookbookTestSuite) TestDeployMicroservice() {
	ctx := context.Background()
	resp, err := t.ZeusTestClient.Deploy(ctx, genericDeploymentKnsReq)
	t.Require().Nil(err)
	t.Assert().NotEmpty(resp)
}

func (t *DeploymentsCookbookTestSuite) TestDestroyMicroservice() {
	ctx := context.Background()
	resp, err := t.ZeusTestClient.DestroyDeploy(ctx, genericDeploymentKnsReq)
	t.Require().Nil(err)
	t.Assert().NotEmpty(resp)
}

func (t *DeploymentsCookbookTestSuite) TestUploadCharts() {
	ctx := context.Background()
	resp, err := t.ZeusTestClient.UploadChart(ctx, genericDeploymentChartPath, genericDeploymentChart)
	t.Require().Nil(err)
	t.Assert().NotZero(resp.TopologyID)

	genericDeploymentKnsReq.TopologyID = resp.TopologyID
	tar := zeus_req_types.TopologyRequest{TopologyID: genericDeploymentKnsReq.TopologyID}
	chartResp, err := t.ZeusTestClient.ReadChart(ctx, tar)
	t.Require().Nil(err)
	t.Assert().NotEmpty(chartResp)

	err = chartResp.PrintWorkload(genericDeploymentChartPath)
	t.Require().Nil(err)
}

func (t *DeploymentsCookbookTestSuite) SetupTest() {
	// points dir to test/configs
	tc := configs.InitLocalTestConfigs()

	// uses the bearer token from test/configs/config.yaml
	t.ZeusTestClient = zeus_client.NewDefaultZeusClient(tc.Bearer)
	//t.ZeusTestClient = zeus_client.NewLocalZeusClient(tc.Bearer)
	cookbooks.ChangeToCookbookDir()
}

func TestBeaconCookbookTestSuite(t *testing.T) {
	suite.Run(t, new(DeploymentsCookbookTestSuite))
}

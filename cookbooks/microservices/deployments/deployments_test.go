package deployments_cookbook

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/zeus-fyi/zeus/cookbooks"
	zeus_client "github.com/zeus-fyi/zeus/pkg/zeus/client"
	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_req_types"
	"github.com/zeus-fyi/zeus/test/configs"
	"github.com/zeus-fyi/zeus/test/test_suites"
)

type DeploymentsCookbookTestSuite struct {
	test_suites.BaseTestSuite
	ZeusTestClient zeus_client.ZeusClient
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

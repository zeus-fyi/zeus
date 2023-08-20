package load_sim_cookbook

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/zeus-fyi/zeus/cookbooks"
	"github.com/zeus-fyi/zeus/test/configs"
	"github.com/zeus-fyi/zeus/test/test_suites"
	zeus_client "github.com/zeus-fyi/zeus/zeus/z_client"
)

var ctx = context.Background()

type LoadSimCookbookTestSuite struct {
	test_suites.BaseTestSuite
	ZeusTestClient zeus_client.ZeusClient
}

func (t *LoadSimCookbookTestSuite) TestDeploy() {
	resp, err := t.ZeusTestClient.Deploy(ctx, loadSimDeploymentKnsReq)
	t.Require().Nil(err)
	t.Assert().NotEmpty(resp)
}

func (t *LoadSimCookbookTestSuite) TestDestroy() {
	resp, err := t.ZeusTestClient.DestroyDeploy(ctx, loadSimDeploymentKnsReq)
	t.Require().Nil(err)
	t.Assert().NotEmpty(resp)
}

func (t *LoadSimCookbookTestSuite) TestUpload() {
	_, rerr := LoadSimClusterDefinition.UploadChartsFromClusterDefinition(ctx, t.ZeusTestClient, true)
	t.Require().Nil(rerr)
}

func (t *LoadSimCookbookTestSuite) TestCreateClusterClass() {
	cd := LoadSimClusterDefinition
	gcd := cd.BuildClusterDefinitions()
	t.Assert().NotEmpty(gcd)
	fmt.Println(gcd)

	err := gcd.CreateClusterClassDefinitions(context.Background(), t.ZeusTestClient)
	t.Require().Nil(err)
}

func (t *LoadSimCookbookTestSuite) SetupTest() {
	// points dir to test/configs
	tc := configs.InitLocalTestConfigs()

	// uses the bearer token from test/configs/config.yaml
	t.ZeusTestClient = zeus_client.NewDefaultZeusClient(tc.Bearer)
	//t.ZeusTestClient = zeus_client.NewLocalZeusClient(tc.Bearer)
	cookbooks.ChangeToCookbookDir()
}

func TestLoadSimCookbookTestSuite(t *testing.T) {
	suite.Run(t, new(LoadSimCookbookTestSuite))
}

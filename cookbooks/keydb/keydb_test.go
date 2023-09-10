package zeus_keydb

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/zeus-fyi/zeus/cookbooks"
	"github.com/zeus-fyi/zeus/test/configs"
	"github.com/zeus-fyi/zeus/test/test_suites"
	zeus_client "github.com/zeus-fyi/zeus/zeus/z_client"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_req_types"
)

func (t *KeyDBCookbookTestSuite) TestDeployKeyDB() {
	t.TestUploadKeyDB()
	cdep := keyDBClusterDefinition.GenerateDeploymentRequest()

	_, err := t.ZeusTestClient.DeployCluster(ctx, cdep)
	t.Require().Nil(err)
}

func (t *KeyDBCookbookTestSuite) TestDestroyKeyDB() {
	d := zeus_req_types.TopologyDeployRequest{
		CloudCtxNs: keyDBCloudCtxNs,
	}
	resp, err := t.ZeusTestClient.DestroyDeploy(ctx, d)
	t.Require().Nil(err)
	t.Assert().NotEmpty(resp)
}

func (t *KeyDBCookbookTestSuite) TestUploadKeyDB() {
	_, rerr := keyDBClusterDefinition.UploadChartsFromClusterDefinition(ctx, t.ZeusTestClient, true)
	t.Require().Nil(rerr)
}

func (t *KeyDBCookbookTestSuite) TestCreateClusterClassKeyDB() {
	cd := keyDBClusterDefinition
	gcd := cd.BuildClusterDefinitions()
	t.Assert().NotEmpty(gcd)
	fmt.Println(gcd)

	err := gcd.CreateClusterClassDefinitions(context.Background(), t.ZeusTestClient)
	t.Require().Nil(err)
}

type KeyDBCookbookTestSuite struct {
	test_suites.BaseTestSuite
	ZeusTestClient zeus_client.ZeusClient
}

var ctx = context.Background()

func (t *KeyDBCookbookTestSuite) SetupTest() {
	// points dir to test/configs
	tc := configs.InitLocalTestConfigs()
	t.Tc = tc
	// uses the bearer token from test/configs/config.yaml
	t.ZeusTestClient = zeus_client.NewDefaultZeusClient(tc.Bearer)
	//t.ZeusTestClient = zeus_client.NewZeusClient("http://localhost:9001", tc.Bearer)
	cookbooks.ChangeToCookbookDir()
}

func TestKeyDBCookbookTestSuite(t *testing.T) {
	suite.Run(t, new(KeyDBCookbookTestSuite))
}

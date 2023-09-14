package sui_cookbooks

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

func (t *SuiCookbookTestSuite) TestDeploy() {
	t.TestUpload()
	cdep := suiNodeDefinition.GenerateDeploymentRequest()

	_, err := t.ZeusTestClient.DeployCluster(ctx, cdep)
	t.Require().Nil(err)
}

func (t *SuiCookbookTestSuite) TestDestroy() {
	d := zeus_req_types.TopologyDeployRequest{
		CloudCtxNs: suiCloudCtxNs,
	}
	resp, err := t.ZeusTestClient.DestroyDeploy(ctx, d)
	t.Require().Nil(err)
	t.Assert().NotEmpty(resp)
}

func (t *SuiCookbookTestSuite) TestUpload() {
	_, rerr := suiNodeDefinition.UploadChartsFromClusterDefinition(ctx, t.ZeusTestClient, true)
	t.Require().Nil(rerr)
}

func (t *SuiCookbookTestSuite) TestUploadTestnet() {
	suiNodeDefinition.ComponentBases[Sui] = GetSuiClientNetworkConfigBase(Full, "testnet")
	t.Require().NotNil(suiNodeDefinition.ComponentBases[Sui])
	//_, rerr := suiNodeDefinition.UploadChartsFromClusterDefinition(ctx, t.ZeusTestClient, true)
	//t.Require().Nil(rerr)
}

func (t *SuiCookbookTestSuite) TestCreateClusterClass() {
	cd := suiNodeDefinition
	gcd := cd.BuildClusterDefinitions()
	t.Assert().NotEmpty(gcd)
	fmt.Println(gcd)

	err := gcd.CreateClusterClassDefinitions(context.Background(), t.ZeusTestClient)
	t.Require().Nil(err)
}

type SuiCookbookTestSuite struct {
	test_suites.BaseTestSuite
	ZeusTestClient zeus_client.ZeusClient
}

var ctx = context.Background()

func (t *SuiCookbookTestSuite) SetupTest() {
	// points dir to test/configs
	tc := configs.InitLocalTestConfigs()
	t.Tc = tc
	// uses the bearer token from test/configs/config.yaml
	t.ZeusTestClient = zeus_client.NewDefaultZeusClient(tc.Bearer)
	//t.ZeusTestClient = zeus_client.NewZeusClient("http://localhost:9001", tc.Bearer)
	cookbooks.ChangeToCookbookDir()
}

func TestSuiCookbookTestSuite(t *testing.T) {
	suite.Run(t, new(SuiCookbookTestSuite))
}

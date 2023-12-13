package avax_node_cookbooks

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/zeus-fyi/zeus/cookbooks"
	"github.com/zeus-fyi/zeus/test/configs"
	"github.com/zeus-fyi/zeus/test/test_suites"
	"github.com/zeus-fyi/zeus/zeus/workload_config_drivers/config_overrides"
	zeus_client "github.com/zeus-fyi/zeus/zeus/z_client"
	zk8s_clusters "github.com/zeus-fyi/zeus/zeus/z_client/clusters"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_req_types"
)

var ctx = context.Background()

func (t *AvaxCookbookTestSuite) TestClusterDeploy() {
	infCfg := config_overrides.IngressDriver{NginxAuthURL: t.Tc.Web3SignerAuthURL}
	customIngTc := config_overrides.TopologyConfigDriver{
		IngressDriver: &infCfg,
	}
	AvaxIngressSkeletonBaseConfig.TopologyConfigDriver = &customIngTc
	IngressComponentBase.SkeletonBases["avaxIngress"] = AvaxIngressSkeletonBaseConfig
	AvaxNodeComponentBases["avaxIngress"] = IngressComponentBase
	cd := AvaxNodeClusterDefinition
	_, err := cd.UploadChartsFromClusterDefinition(ctx, t.ZeusTestClient, true)
	t.Require().Nil(err)

	cdep := cd.GenerateDeploymentRequest()
	_, err = zk8s_clusters.DeployCluster(ctx, t.ZeusTestClient, cdep)
	t.Require().Nil(err)
}

func (t *AvaxCookbookTestSuite) TestClusterDestroy() {
	d := zeus_req_types.TopologyDeployRequest{
		CloudCtxNs: AvaxNodeCloudCtxNs,
	}
	resp, err := t.ZeusTestClient.DestroyDeploy(ctx, d)
	t.Require().Nil(err)
	t.Assert().NotEmpty(resp)
}

func (t *AvaxCookbookTestSuite) TestClusterSetup() {
	infCfg := config_overrides.IngressDriver{NginxAuthURL: t.Tc.Web3SignerAuthURL}
	customIngTc := config_overrides.TopologyConfigDriver{
		IngressDriver: &infCfg,
	}
	AvaxIngressSkeletonBaseConfig.TopologyConfigDriver = &customIngTc
	IngressComponentBase.SkeletonBases["avaxIngress"] = AvaxIngressSkeletonBaseConfig
	AvaxNodeComponentBases["avaxIngress"] = IngressComponentBase
	cd := AvaxNodeClusterDefinition
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

func (t *AvaxCookbookTestSuite) TestClusterDefinitionCreation() {
	cd := AvaxNodeClusterDefinition
	gcd := cd.BuildClusterDefinitions()
	t.Assert().NotEmpty(gcd)
	fmt.Println(gcd)

	err := gcd.CreateClusterClassDefinitions(ctx, t.ZeusTestClient)
	t.Require().Nil(err)
}

type AvaxCookbookTestSuite struct {
	test_suites.BaseTestSuite
	ZeusTestClient zeus_client.ZeusClient
}

func (t *AvaxCookbookTestSuite) SetupTest() {
	// points dir to test/configs
	tc := configs.InitLocalTestConfigs()
	t.Tc = tc
	// uses the bearer token from test/configs/config.yaml
	t.ZeusTestClient = zeus_client.NewDefaultZeusClient(tc.Bearer)
	//t.ZeusTestClient = zeus_client.NewZeusClient("http://localhost:9001", tc.Bearer)
	cookbooks.ChangeToCookbookDir()
}

func TestAvaxCookbookTestSuite(t *testing.T) {
	suite.Run(t, new(AvaxCookbookTestSuite))
}

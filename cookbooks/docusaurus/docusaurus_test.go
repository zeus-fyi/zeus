package docusaurus_cookbooks

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/zeus-fyi/zeus/cookbooks"
	"github.com/zeus-fyi/zeus/test/configs"
	"github.com/zeus-fyi/zeus/test/test_suites"
	zeus_client "github.com/zeus-fyi/zeus/zeus/z_client"
)

var ctx = context.Background()

type DocusaurusCookbookTestSuite struct {
	test_suites.BaseTestSuite
	ZeusTestClient zeus_client.ZeusClient
}

func (t *DocusaurusCookbookTestSuite) TestDeployDocusaurus() {
	resp, err := t.ZeusTestClient.Deploy(ctx, docusaurusKnsReq)
	t.Require().Nil(err)
	t.Assert().NotEmpty(resp)
}

func (t *DocusaurusCookbookTestSuite) TestUploadDocusaurus() {
	_, rerr := DocusaurusClusterDefinition.UploadChartsFromClusterDefinition(ctx, t.ZeusTestClient, true)
	t.Require().Nil(rerr)
}

func (t *DocusaurusCookbookTestSuite) TestCreateDocusaurusClass() {
	//cd := DocusaurusClusterDefinition
	//gcd := cd.BuildClusterDefinitions()
	//t.Assert().NotEmpty(gcd)
	//fmt.Println(gcd)
	//
	//err := gcd.CreateClusterClassDefinitions(context.Background(), t.ZeusTestClient)
	//t.Require().Nil(err)

	dockerImage := "docker.io/zeusfyi/docusaurus-template:latest"
	cd := ConfigDocusaurus("docker.io/zeusfyi/docusaurus:latest")

	ch, err := cd.GenerateSkeletonBaseCharts()
	t.Require().Nil(err)
	t.Assert().NotEmpty(ch)

	for _, c := range ch {
		t.Require().Equal(dockerImage, c.Workload.Deployment.Spec.Template.Spec.Containers[0].Image)
	}

}

func (t *DocusaurusCookbookTestSuite) SetupTest() {
	// points dir to test/configs
	tc := configs.InitLocalTestConfigs()

	// uses the bearer token from test/configs/config.yaml
	t.ZeusTestClient = zeus_client.NewDefaultZeusClient(tc.Bearer)
	//t.ZeusTestClient = zeus_client.NewLocalZeusClient(tc.Bearer)
	cookbooks.ChangeToCookbookDir()
}

func TestBeaconCookbookTestSuite(t *testing.T) {
	suite.Run(t, new(DocusaurusCookbookTestSuite))
}

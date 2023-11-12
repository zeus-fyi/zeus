package docusaurus_cookbooks

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/zeus-fyi/zeus/cookbooks"
	"github.com/zeus-fyi/zeus/test/configs"
	"github.com/zeus-fyi/zeus/test/test_suites"
	zeus_cluster_config_drivers "github.com/zeus-fyi/zeus/zeus/cluster_config_drivers"
	zk8s_templates "github.com/zeus-fyi/zeus/zeus/workload_config_drivers/templates"
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

const (
	docusaurus = "docusaurus-template"
)

func (t *DocusaurusCookbookTestSuite) TestCreateDocsClass() {
	gcd := DocusaurusClusterDefinition.BuildClusterDefinitions()
	t.Assert().NotEmpty(gcd)
	fmt.Println(gcd)

	err := gcd.CreateClusterClassDefinitions(context.Background(), t.ZeusTestClient)
	t.Require().Nil(err)
}

func (t *DocusaurusCookbookTestSuite) TestUploadDocusaurus() {
	_, rerr := DocusaurusClusterDefinition.UploadChartsFromClusterDefinition(ctx, t.ZeusTestClient, true)
	t.Require().Nil(rerr)
}

func (t *DocusaurusCookbookTestSuite) TestCreateDocusaurusClass() {
	dockerImage := "docker.io/zeusfyi/docusaurus-template:latest"
	wd := zeus_cluster_config_drivers.WorkloadDefinition{
		WorkloadName: docusaurus,
		ReplicaCount: 1,
		Containers: zk8s_templates.Containers{
			docusaurus: zk8s_templates.Container{
				IsInitContainer: false,
				ImagePullPolicy: "Always",
				DockerImage: zk8s_templates.DockerImage{
					ImageName: dockerImage,
					ResourceRequirements: zk8s_templates.ResourceRequirements{
						CPU:    "100m",
						Memory: "100Mi",
					},
					Ports: []zk8s_templates.Port{
						{
							Name:               "http",
							Number:             "3000",
							Protocol:           "TCP",
							IngressEnabledPort: true,
						},
					},
				},
			},
		},
	}
	cd, err := zeus_cluster_config_drivers.GenerateDeploymentCluster(ctx, wd)
	t.Require().Nil(err)
	t.Assert().NoError(err)
	t.Assert().NotEmpty(cd)

	cd.IngressPaths = map[string]zk8s_templates.IngressPath{
		docusaurus: {
			Path:     "/",
			PathType: "ImplementationSpecific",
		},
	}
	t.Assert().Equal(docusaurus, cd.ClusterName)
	preview, err := zeus_cluster_config_drivers.GenerateSkeletonBaseChartsPreview(ctx, *cd)
	t.Require().Nil(err)
	t.Assert().NoError(err)
	t.Assert().NotEmpty(preview)
	prt := zeus_cluster_config_drivers.PreviewTemplateGeneration(ctx, *cd)
	t.Assert().NotEmpty(prt)
	prt.DisablePrint = true
	prt.UseEmbeddedWorkload = true

	dpr, err := prt.GenerateSkeletonBaseCharts()
	t.Require().Nil(err)
	t.Assert().NotEmpty(dpr)

	gcd := zeus_cluster_config_drivers.CreateGeneratedClusterClassCreationRequest(cd)
	t.Assert().NotEmpty(gcd)
	fmt.Println(gcd)

	//gcdExp := DocusaurusClusterDefinition.BuildClusterDefinitions()
	//t.Assert().NotEmpty(gcdExp)
	//fmt.Println(gcdExp)

	//t.Assert().Equal(gcdExp, gcd)
	//err = gcd.CreateClusterClassDefinitions(ctx, t.ZeusTestClient)
	//t.Require().Nil(err)
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

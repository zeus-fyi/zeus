package ethereum_mev_cookbooks

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/zeus-fyi/zeus/cookbooks"
	"github.com/zeus-fyi/zeus/test/configs"
	"github.com/zeus-fyi/zeus/test/test_suites"
	zeus_client "github.com/zeus-fyi/zeus/zeus/client"
	"github.com/zeus-fyi/zeus/zeus/client/zeus_req_types"
)

type MevCookbookTestSuite struct {
	test_suites.BaseTestSuite
	ZeusTestClient zeus_client.ZeusClient
}

// This assumes a pre-existing cluster class called "goerliValidators", since
// the cluster already exists, you will run into an error if you try to create
// the cluster class again, so you will use the below functions to add the
// component base and skeleton base to the existing cluster class
func (t *MevCookbookTestSuite) TestCreateClusterBase() {
	basesInsert := []string{"mev"}
	cc := zeus_req_types.TopologyCreateOrAddComponentBasesToClassesRequest{
		ClusterClassName:   "goerliValidators",
		ComponentBaseNames: basesInsert,
	}
	_, err := t.ZeusTestClient.AddComponentBasesToClass(ctx, cc)
	t.Require().Nil(err)
}

func (t *MevCookbookTestSuite) TestCreateClusterSkeletonBases() {
	cc := zeus_req_types.TopologyCreateOrAddSkeletonBasesToClassesRequest{
		ClusterClassName:  "goerliValidators",
		ComponentBaseName: "mev",
		SkeletonBaseNames: []string{"mevBoost"},
	}
	_, err := t.ZeusTestClient.AddSkeletonBasesToClass(ctx, cc)
	t.Require().Nil(err)
}

func (t *MevCookbookTestSuite) SetupTest() {
	// points dir to test/configs
	tc := configs.InitLocalTestConfigs()

	// uses the bearer token from test/configs/config.yaml
	t.ZeusTestClient = zeus_client.NewDefaultZeusClient(tc.Bearer)
	//t.ZeusTestClient = zeus_client.NewZeusClient("http://localhost:9001", tc.Bearer)
	cookbooks.ChangeToCookbookDir()
}

func TestMevCookbookTestSuite(t *testing.T) {
	suite.Run(t, new(MevCookbookTestSuite))
}

package zeus_cluster_config_drivers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/zeus-fyi/zeus/cookbooks"
	ethereum_beacon_cookbooks "github.com/zeus-fyi/zeus/cookbooks/ethereum/beacons"
	zeus_client "github.com/zeus-fyi/zeus/pkg/zeus/client"
	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_common_types"
	"github.com/zeus-fyi/zeus/test/configs"
	"github.com/zeus-fyi/zeus/test/test_suites"
)

type ClusterConfigTestSuite struct {
	test_suites.BaseTestSuite
	ZeusTestClient zeus_client.ZeusClient
}

func (t *ClusterConfigTestSuite) TestClusterCreation() {
	cd := ClusterDefinition{
		ClusterClassName: "testCluster",
		CloudCtxNs: zeus_common_types.CloudCtxNs{
			CloudProvider: "do",
			Region:        "nyc1",
			Context:       "do-nyc1-do-nyc1-zeus-demo",
			Namespace:     "web3signer",
			Env:           "dev",
		},
		ComponentBases: make(map[string]ComponentBaseDefinition),
	}

	cd.ComponentBases["executionClient"] = ComponentBaseDefinition{SkeletonBases: make(map[string]ClusterSkeletonBaseDefinition)}
	cd.ComponentBases["executionClient"].SkeletonBases["gethHercules"] = ClusterSkeletonBaseDefinition{
		SkeletonBaseNameChartPath: ethereum_beacon_cookbooks.BeaconExecClientChartPath,
	}
	cd.ComponentBases["consensusClient"] = ComponentBaseDefinition{SkeletonBases: make(map[string]ClusterSkeletonBaseDefinition)}
	cd.ComponentBases["consensusClient"].SkeletonBases["lighthouseHercules"] = ClusterSkeletonBaseDefinition{
		SkeletonBaseNameChartPath: ethereum_beacon_cookbooks.BeaconConsensusClientChartPath,
	}

	gcd := cd.BuildClusterDefinitions()
	t.Assert().NotEmpty(gcd)
	fmt.Println(gcd)

	gdr := cd.GenerateDeploymentRequest()
	t.Assert().NotEmpty(gdr)
	count := 0
	for _, sbo := range gdr.SkeletonBaseOptions {
		if sbo == "gethHercules" {
			count += 1
		}
		if sbo == "lighthouseHercules" {
			count += 10
		}
	}
	t.Assert().Equal(11, count)
	fmt.Println(gdr)
}

func (t *ClusterConfigTestSuite) SetupTest() {
	// points dir to test/configs
	tc := configs.InitLocalTestConfigs()

	t.ZeusTestClient = zeus_client.NewDefaultZeusClient(tc.Bearer)
	// points dir to cookbooks
	cookbooks.ChangeToCookbookDir()
}

func TestClusterConfigTestSuite(t *testing.T) {
	suite.Run(t, new(ClusterConfigTestSuite))
}

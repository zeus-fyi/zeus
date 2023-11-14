package zeus_cluster_config_drivers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/zeus-fyi/zeus/cookbooks"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	"github.com/zeus-fyi/zeus/test/configs"
	"github.com/zeus-fyi/zeus/test/test_suites"
	"github.com/zeus-fyi/zeus/zeus/workload_config_drivers/config_overrides"
	zeus_client "github.com/zeus-fyi/zeus/zeus/z_client"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_common_types"
)

type ClusterConfigTestSuite struct {
	test_suites.BaseTestSuite
	ZeusTestClient zeus_client.ZeusClient
}

var BeaconExecClientChartPath = filepaths.Path{
	PackageName: "",
	DirIn:       "./ethereum/beacons/infra/exec_client",
	DirOut:      "./ethereum/outputs",
	FnIn:        "gethHercules", // filename for your gzip workload
	FnOut:       "",
	Env:         "",
}

var BeaconConsensusClientChartPath = filepaths.Path{
	PackageName: "",
	DirIn:       "./ethereum/beacons/infra/consensus_client",
	DirOut:      "./ethereum/outputs",
	FnIn:        "lighthouseHercules", // filename for your gzip workload
	FnOut:       "",
	Env:         "",
}

var IngressChartPath = filepaths.Path{
	PackageName: "",
	DirIn:       "./ethereum/beacons/infra/ingress",
	DirOut:      "./ethereum/beacons/infra/processed_beacon_ingress",
	FnIn:        "beaconIngress", // filename for your gzip workload
	FnOut:       "",
	Env:         "",
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
		SkeletonBaseNameChartPath: BeaconExecClientChartPath,
	}
	cd.ComponentBases["consensusClient"] = ComponentBaseDefinition{SkeletonBases: make(map[string]ClusterSkeletonBaseDefinition)}
	cd.ComponentBases["consensusClient"].SkeletonBases["lighthouseHercules"] = ClusterSkeletonBaseDefinition{
		SkeletonBaseNameChartPath: BeaconConsensusClientChartPath,
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

	fakeURL := "https://test.zeus.fyi"
	infCfg := config_overrides.IngressDriver{NginxAuthURL: fakeURL}
	customIngTc := config_overrides.TopologyConfigDriver{
		IngressDriver: &infCfg,
	}

	cd.ComponentBases["ingress"] = ComponentBaseDefinition{SkeletonBases: make(map[string]ClusterSkeletonBaseDefinition)}
	cd.ComponentBases["ingress"].SkeletonBases["ingress"] = ClusterSkeletonBaseDefinition{
		SkeletonBaseNameChartPath: IngressChartPath,
		TopologyConfigDriver:      &customIngTc,
	}

	sbDefs, err := cd.GenerateSkeletonBaseCharts()
	t.Require().Nil(err)
	t.Assert().NotEmpty(sbDefs)

	seen := false
	for _, w := range sbDefs {
		t.Require().NotEmpty(w.Workload)
		if w.Workload.Ingress != nil {
			seen = true
			t.Require().NotEmpty(w.Workload.Ingress.Annotations)
			v, ok := w.Workload.Ingress.Annotations["nginx.ingress.kubernetes.io/auth-url"]
			t.Require().True(ok)
			t.Assert().Equal(fakeURL, v)
		}
	}
	t.Assert().True(seen)
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

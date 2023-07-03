package ethereum_mev_cookbooks

import (
	"context"

	hestia_req_types "github.com/zeus-fyi/zeus/pkg/hestia/client/req_types"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_req_types"

	zeus_cluster_config_drivers "github.com/zeus-fyi/zeus/zeus/cluster_config_drivers"
	zeus_topology_config_drivers "github.com/zeus-fyi/zeus/zeus/workload_config_drivers"
	v1Core "k8s.io/api/core/v1"
)

const (
	mevContainerReference = "zeus-mev"
	flashbotsDockerImage  = "flashbots/mev-boost:1.5.0"
)

var (
	ctx                   = context.Background()
	MevSkeletonBaseConfig = zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
		SkeletonBaseChart:         zeus_req_types.TopologyCreateRequest{},
		SkeletonBaseNameChartPath: MevChartPath,
		TopologyConfigDriver: &zeus_topology_config_drivers.TopologyConfigDriver{
			DeploymentDriver: &zeus_topology_config_drivers.DeploymentDriver{
				ContainerDrivers: map[string]zeus_topology_config_drivers.ContainerDriver{
					mevContainerReference: {Container: v1Core.Container{
						Name:  mevContainerReference,
						Image: flashbotsDockerImage,
						Args:  GetMevBoostArgs(ctx, hestia_req_types.Goerli, MevRelays),
					}},
				},
			}},
	}
	MevChartPath = filepaths.Path{
		PackageName: "",
		DirIn:       "./ethereum/mev/infra",
		DirOut:      "./ethereum/mev/infra/processed_mev",
		FnIn:        "mev", // filename for your gzip workload
		FnOut:       "",
		Env:         "",
	}
	MevRelays = RelaysEnabled{
		Flashbots:   true,
		Blocknative: true,
		EdenNetwork: true,
	}
)

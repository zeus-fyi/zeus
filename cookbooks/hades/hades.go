package cookbooks_hades

import (
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	"github.com/zeus-fyi/zeus/zeus/client/zeus_common_types"
	"github.com/zeus-fyi/zeus/zeus/client/zeus_req_types"
	zeus_cluster_config_drivers "github.com/zeus-fyi/zeus/zeus/cluster_config_drivers"
)

var (
	HadesCloudCtxNs = zeus_common_types.CloudCtxNs{
		CloudProvider: "do",
		Region:        "sfo3",
		Context:       "do-nyc1-do-nyc1-zeus-demo",
		Namespace:     "demo", // set with your own namespace
		Env:           "production",
	}
	HadesClusterDefinition = zeus_cluster_config_drivers.ClusterDefinition{
		ClusterClassName: "hades",
		CloudCtxNs:       HadesCloudCtxNs,
		ComponentBases: map[string]zeus_cluster_config_drivers.ComponentBaseDefinition{
			"hades": HadesComponentBase,
		},
	}
	HadesComponentBase = zeus_cluster_config_drivers.ComponentBaseDefinition{
		SkeletonBases: map[string]zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
			"hades": HadesSkeletonBaseConfig,
		},
	}
	HadesSkeletonBaseConfig = zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
		SkeletonBaseChart:         zeus_req_types.TopologyCreateRequest{},
		SkeletonBaseNameChartPath: HadesChartPath,
	}
	HadesChartPath = filepaths.Path{
		PackageName: "",
		DirIn:       "./hades/infra",
		DirOut:      "./hades/outputs",
		FnIn:        "hades", // filename for your gzip workload
		FnOut:       "",
		Env:         "",
	}
)

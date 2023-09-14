package sui_cookbooks

import (
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	zeus_cluster_config_drivers "github.com/zeus-fyi/zeus/zeus/cluster_config_drivers"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_common_types"
)

const (
	Sui       = "sui"
	Full      = "full"
	Validator = "validator"
)

var (
	suiNodeDefinition = zeus_cluster_config_drivers.ClusterDefinition{
		ClusterClassName: Sui,
		CloudCtxNs:       suiCloudCtxNs,
		ComponentBases:   suiComponentBases,
	}
	suiComponentBases = map[string]zeus_cluster_config_drivers.ComponentBaseDefinition{
		Sui: suiMasterComponentBase,
	}
	suiMasterComponentBase = zeus_cluster_config_drivers.ComponentBaseDefinition{
		SkeletonBases: map[string]zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
			Sui: suiSkeletonBaseConfig,
		},
	}
	suiSkeletonBaseConfig = zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
		SkeletonBaseNameChartPath: suiMasterChartPath,
	}
)
var (
	suiCloudCtxNs = zeus_common_types.CloudCtxNs{
		CloudProvider: "do",
		Region:        "sfo3",
		Context:       "do-sfo3-dev-do-sfo3-zeus",
		Namespace:     Sui, // set with your own namespace
		Env:           "production",
	}
	suiMasterChartPath = filepaths.Path{
		PackageName: "",
		DirIn:       "./sui/node/infra",
		DirOut:      "./sui/output",
		FnIn:        Sui, // filename for your gzip workload
		FnOut:       "",
		Env:         "",
	}
)

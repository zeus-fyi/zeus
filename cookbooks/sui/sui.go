package sui_cookbooks

import (
	"strings"

	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	zeus_cluster_config_drivers "github.com/zeus-fyi/zeus/zeus/cluster_config_drivers"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_common_types"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_req_types"
)

const (
	Sui       = "sui"
	Full      = "full"
	Validator = "validator"

	SuiIngress        = "suiIngress"
	SuiServiceMonitor = "suiServiceMonitor"
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
		SkeletonBaseNameChartPath: SuiMasterChartPath,
	}
	suiCloudCtxNs = zeus_common_types.CloudCtxNs{
		CloudProvider: "do",
		Region:        "sfo3",
		Context:       "do-sfo3-dev-do-sfo3-zeus",
		Namespace:     Sui, // set with your own namespace
		Env:           "production",
	}
	SuiMasterChartPath = filepaths.Path{
		PackageName: "",
		DirIn:       "./sui/node/infra",
		DirOut:      "./sui/output",
		FnIn:        Sui, // filename for your gzip workload
		FnOut:       "",
		Env:         "",
	}
	suiIngressComponentBase = zeus_cluster_config_drivers.ComponentBaseDefinition{
		SkeletonBases: map[string]zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
			SuiIngress: suiIngressSkeletonBaseConfig,
		},
	}
	suiIngressSkeletonBaseConfig = zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
		SkeletonBaseChart:         zeus_req_types.TopologyCreateRequest{},
		SkeletonBaseNameChartPath: SuiIngressChartPath,
	}
	SuiIngressChartPath = filepaths.Path{
		PackageName: "",
		DirIn:       "./sui/node/ingress",
		DirOut:      "./sui/output",
		FnIn:        SuiIngress, // filename for your gzip workload
		FnOut:       "",
		Env:         "",
	}
	suiServiceMonitorComponentBase = zeus_cluster_config_drivers.ComponentBaseDefinition{
		SkeletonBases: map[string]zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
			SuiServiceMonitor: suiServiceMonitorSkeletonBaseConfig,
		},
	}
	suiServiceMonitorSkeletonBaseConfig = zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
		SkeletonBaseChart:         zeus_req_types.TopologyCreateRequest{},
		SkeletonBaseNameChartPath: SuiServiceMonitorChartPath,
	}
	SuiServiceMonitorChartPath = filepaths.Path{
		PackageName: "",
		DirIn:       "./sui/node/servicemonitor",
		DirOut:      "./sui/output",
		FnIn:        SuiServiceMonitor, // filename for your gzip workload
		Env:         "",
	}
)

func GetSuiClientClusterDef(cfg SuiConfigOpts) zeus_cluster_config_drivers.ClusterDefinition {
	return zeus_cluster_config_drivers.ClusterDefinition{
		ClusterClassName: strings.ToLower(Sui) + "-" + strings.ToLower(cfg.Network) + "-" + strings.ToLower(cfg.CloudProvider),
		ComponentBases:   GetSuiConfig(cfg),
	}
}

func GetSuiConfig(cfg SuiConfigOpts) map[string]zeus_cluster_config_drivers.ComponentBaseDefinition {
	suiComponentBases = map[string]zeus_cluster_config_drivers.ComponentBaseDefinition{
		Sui: GetSuiClientNetworkConfigBase(cfg),
	}
	if cfg.WithIngress {
		suiComponentBases[SuiIngress] = suiIngressComponentBase
	}
	if cfg.WithServiceMonitor {
		suiComponentBases[SuiServiceMonitor] = suiServiceMonitorComponentBase
	}
	return suiComponentBases
}

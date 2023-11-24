package zeus_jobs

import (
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	zeus_cluster_config_drivers "github.com/zeus-fyi/zeus/zeus/cluster_config_drivers"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_common_types"
)

var (
	jobsClusterDefinition = zeus_cluster_config_drivers.ClusterDefinition{
		ClusterClassName: "job-template",
		CloudCtxNs:       jobsCloudCtxNs,
		ComponentBases:   jobsComponentBases,
	}
	jobsComponentBases = map[string]zeus_cluster_config_drivers.ComponentBaseDefinition{
		"job-template": jobsMasterComponentBase,
	}
	jobsMasterComponentBase = zeus_cluster_config_drivers.ComponentBaseDefinition{
		SkeletonBases: map[string]zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
			"jobs-template": jobsMasterSkeletonBaseConfig,
		},
	}
	jobsMasterSkeletonBaseConfig = zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
		SkeletonBaseNameChartPath: jobsMasterChartPath,
	}
)
var (
	jobsCloudCtxNs = zeus_common_types.CloudCtxNs{
		CloudProvider: "do",
		Region:        "sfo3",
		Context:       "do-sfo3-dev-do-sfo3-zeus",
		Namespace:     "jobs", // set with your own namespace
		Env:           "production",
	}
	jobsMasterChartPath = filepaths.Path{
		PackageName: "",
		DirIn:       "./jobs/infra",
		DirOut:      "./jobs/outputs",
		FnIn:        "jobs-master", // filename for your gzip workload
		FnOut:       "",
		Env:         "",
	}
)

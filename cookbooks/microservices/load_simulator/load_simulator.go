package load_sim_cookbook

import (
	"fmt"
	"time"

	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_common_types"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_req_types"

	zeus_cluster_config_drivers "github.com/zeus-fyi/zeus/zeus/cluster_config_drivers"
)

var (
	LoadSimClusterDefinition = zeus_cluster_config_drivers.ClusterDefinition{
		ClusterClassName: "loadSimulator",
		CloudCtxNs:       LoadSimCloudCtxNs,
		ComponentBases:   LoadSimComponentBases,
	}
	LoadSimCloudCtxNs = zeus_common_types.CloudCtxNs{
		CloudProvider: "do",
		Region:        "nyc1",
		Context:       "do-nyc1-do-nyc1-zeus-demo",
		Namespace:     "load-sim",
		Env:           "production",
	}
	LoadSimComponentBases = map[string]zeus_cluster_config_drivers.ComponentBaseDefinition{
		"load-sim": LoadSimComponentBase,
	}
	LoadSimComponentBase = zeus_cluster_config_drivers.ComponentBaseDefinition{
		SkeletonBases: map[string]zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
			"api": LoadSimSkeletonBaseConfig,
		},
	}
	LoadSimSkeletonBaseConfig = zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
		SkeletonBaseChart:         zeus_req_types.TopologyCreateRequest{},
		SkeletonBaseNameChartPath: loadSimChartPath,
	}
)

// set your own topologyID here after uploading a chart workload
var loadSimDeploymentKnsReq = zeus_req_types.TopologyDeployRequest{
	TopologyID: 0,
	CloudCtxNs: loadSimCloudCtxNs,
}

var loadSimCloudCtxNs = zeus_common_types.CloudCtxNs{
	CloudProvider: "do",
	Region:        "sfo3",
	Context:       "do-sfo3-dev-do-sfo3-zeus",
	Namespace:     "load-sim", // set with your own namespace
	Env:           "production",
}

// chart workload metadata
var loadSimChart = zeus_req_types.TopologyCreateRequest{
	TopologyName:     "load-sim",
	ChartName:        "load-sim",
	ChartDescription: "load-sim",
	Version:          fmt.Sprintf("v0.0.%d", time.Now().Unix()),
}

// DirOut is where it will write a copy of the chart you uploaded, which helps verify the workload is correct
var loadSimChartPath = filepaths.Path{
	PackageName: "",
	DirIn:       "./microservices/load_simulator/infra",
	DirOut:      "./microservices/outputs",
	FnIn:        "load-sim", // filename for your gzip workload
	FnOut:       "",
	Env:         "",
}

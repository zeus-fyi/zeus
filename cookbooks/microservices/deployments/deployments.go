package deployment_cookbooks

import (
	"fmt"
	"time"

	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	zeus_cluster_config_drivers "github.com/zeus-fyi/zeus/pkg/zeus/cluster_config_drivers"
)

var (
	MicroserviceClusterDefinition = zeus_cluster_config_drivers.ClusterDefinition{
		ClusterClassName: "microservice",
		CloudCtxNs:       MicroserviceNodeCloudCtxNs,
		ComponentBases:   MicroserviceComponentBases,
	}
	MicroserviceNodeCloudCtxNs = zeus_common_types.CloudCtxNs{
		CloudProvider: "do",
		Region:        "nyc1",
		Context:       "do-nyc1-do-nyc1-zeus-demo",
		Namespace:     "microservice",
		Env:           "production",
	}
	MicroserviceComponentBases = map[string]zeus_cluster_config_drivers.ComponentBaseDefinition{
		"microservice": MicroserviceComponentBase,
	}
	MicroserviceComponentBase = zeus_cluster_config_drivers.ComponentBaseDefinition{
		SkeletonBases: map[string]zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
			"api": MicroserviceSkeletonBaseConfig,
		},
	}
	MicroserviceSkeletonBaseConfig = zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
		SkeletonBaseChart:         zeus_req_types.TopologyCreateRequest{},
		SkeletonBaseNameChartPath: genericDeploymentChartPath,
	}
)

// set your own topologyID here after uploading a chart workload
var genericDeploymentKnsReq = zeus_req_types.TopologyDeployRequest{
	TopologyID: 0,
	CloudCtxNs: genericDeploymentCloudCtxNs,
}

var genericDeploymentCloudCtxNs = zeus_common_types.CloudCtxNs{
	CloudProvider: "do",
	Region:        "sfo3",
	Context:       "do-sfo3-dev-do-sfo3-zeus",
	Namespace:     "generic-deployment", // set with your own namespace
	Env:           "production",
}

// chart workload metadata
var genericDeploymentChart = zeus_req_types.TopologyCreateRequest{
	TopologyName:     "generic-deployment",
	ChartName:        "generic-deployment",
	ChartDescription: "generic-deployment",
	Version:          fmt.Sprintf("v0.0.%d", time.Now().Unix()),
}

// DirOut is where it will write a copy of the chart you uploaded, which helps verify the workload is correct
var genericDeploymentChartPath = filepaths.Path{
	PackageName: "",
	DirIn:       "./microservices/deployments/infra",
	DirOut:      "./microservices/outputs",
	FnIn:        "generic-deployment", // filename for your gzip workload
	FnOut:       "",
	Env:         "",
}

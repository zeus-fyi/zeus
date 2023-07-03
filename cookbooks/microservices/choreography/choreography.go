package choreography_cookbooks

import (
	"fmt"
	"time"

	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_common_types"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_req_types"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_resp_types/topology_workloads"

	zeus_cluster_config_drivers "github.com/zeus-fyi/zeus/zeus/cluster_config_drivers"
)

var ChoreographyKnsReq = zeus_req_types.TopologyDeployRequest{
	TopologyID: 0,
	CloudCtxNs: GenericChoreographyCloudCtxNs,
}

var ChoreographyComponentBase = zeus_cluster_config_drivers.ComponentBaseDefinition{
	SkeletonBases: map[string]zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
		"choreography": {
			SkeletonBaseNameChartPath: GenericDeploymentChartPath,
			Workload:                  topology_workloads.TopologyBaseInfraWorkload{},
			TopologyConfigDriver:      nil,
		},
	},
}
var GenericChoreographyCloudCtxNs = zeus_common_types.CloudCtxNs{
	CloudProvider: "do",
	Region:        "sfo3",
	Context:       "do-sfo3-dev-do-sfo3-zeus",
	Namespace:     "generic-choreography", // set with your own namespace
	Env:           "production",
}

var GenericChoreographyChart = zeus_req_types.TopologyCreateRequest{
	TopologyName:      "generic-choreography",
	ChartName:         "generic-choreography",
	ChartDescription:  "generic-choreography",
	ClusterClassName:  "choreography",
	ComponentBaseName: "choreography",
	SkeletonBaseName:  "choreography",
	Version:           fmt.Sprintf("v0.0.%d", time.Now().Unix()),
}

// GenericDeploymentChartPath DirOut is where it will write a copy of the chart you uploaded, which helps verify the workload is correct
var GenericDeploymentChartPath = filepaths.Path{
	PackageName: "",
	DirIn:       "./microservices/choreography/infra",
	DirOut:      "./microservices/outputs",
	FnIn:        "generic-choreography", // filename for your gzip workload
	FnOut:       "",
	Env:         "",
}

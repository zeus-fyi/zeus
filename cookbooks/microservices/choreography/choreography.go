package choreography_cookbooks

import (
	"fmt"
	"time"

	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	strings_filter "github.com/zeus-fyi/zeus/pkg/utils/strings"
	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_common_types"
	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_req_types"
)

var ChoreographyKnsReq = zeus_req_types.TopologyDeployRequest{
	TopologyID: 0,
	CloudCtxNs: GenericChoreographyCloudCtxNs,
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
	FilterFiles: strings_filter.FilterOpts{},
}

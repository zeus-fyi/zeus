package deployment_cookbooks

import (
	"fmt"
	"time"

	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	strings_filter "github.com/zeus-fyi/zeus/pkg/utils/strings"
	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_common_types"
	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_req_types"
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
	FilterFiles: strings_filter.FilterOpts{},
}

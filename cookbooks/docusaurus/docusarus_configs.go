package docusaurus_cookbooks

import (
	"fmt"
	"time"

	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_common_types"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_req_types"
)

// set your own topologyID here after uploading a chart workload
var docusaurusKnsReq = zeus_req_types.TopologyDeployRequest{
	TopologyID: 0,
	CloudCtxNs: docusaurusCloudCtxNs,
}

var docusaurusCloudCtxNs = zeus_common_types.CloudCtxNs{
	CloudProvider: "do",
	Region:        "sfo3",
	Context:       "do-sfo3-dev-do-sfo3-zeus",
	Namespace:     "docusaurus", // set with your own namespace
	Env:           "production",
}

var docusaurusChart = zeus_req_types.TopologyCreateRequest{
	TopologyName:     "docusaurus",
	ChartName:        "docusaurus",
	ChartDescription: "docusaurus",
	Version:          fmt.Sprintf("v0.0.%d", time.Now().Unix()),
}

// DocusaurusChartPath is where it will write a copy of the chart you uploaded, which helps verify the workload is correct
var DocusaurusChartPath = filepaths.Path{
	PackageName: "",
	DirIn:       "./docusaurus/infra",
	DirOut:      "./docusaurus/outputs",
	FnIn:        "docusaurus", // filename for your gzip workload
	FnOut:       "",
	Env:         "",
}

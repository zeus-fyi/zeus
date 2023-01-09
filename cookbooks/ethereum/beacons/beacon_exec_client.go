package beacon_cookbooks

import (
	"fmt"
	"time"

	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_req_types"
)

// DeployExecClientKnsReq set your own topologyID here after uploading a chart workload
var DeployExecClientKnsReq = zeus_req_types.TopologyDeployRequest{
	TopologyID: 0,
	CloudCtxNs: BeaconCloudCtxNs,
}

var ExecClientChart = zeus_req_types.TopologyCreateRequest{
	TopologyName:      "gethHercules",
	ChartName:         "gethHercules",
	ChartDescription:  "gethHercules",
	Version:           fmt.Sprintf("gethHerculesv0.0.%d", time.Now().Unix()),
	ClusterClassName:  "ethereumBeacons",
	ComponentBaseName: "executionClient",
	SkeletonBaseName:  "gethHercules",
	Tag:               "latest",
}

var BeaconExecClientChartPath = filepaths.Path{
	PackageName: "",
	DirIn:       "./ethereum/beacons/infra/exec_client",
	DirOut:      "./ethereum/outputs",
	FnIn:        "gethHercules", // filename for your gzip workload
	FnOut:       "",
	Env:         "",
}

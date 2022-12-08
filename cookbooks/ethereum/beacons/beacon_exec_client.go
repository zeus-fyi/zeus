package beacon_cookbook

import (
	"fmt"
	"time"

	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	strings_filter "github.com/zeus-fyi/zeus/pkg/utils/strings"
	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_req_types"
)

// set your own topologyID here after uploading a chart workload
var deployExecClientKnsReq = zeus_req_types.TopologyDeployRequest{
	TopologyID: 0,
	CloudCtxNs: beaconCloudCtxNs,
}

// chart workload metadata
var execClientChart = zeus_req_types.TopologyCreateRequest{
	TopologyName:     "gethHercules",
	ChartName:        "gethHercules",
	ChartDescription: "gethHercules",
	Version:          fmt.Sprintf("gethHerculesv0.0.%d", time.Now().Unix()),
	SkeletonBaseName: "gethHercules",
	Tag:              "latest",
}

var beaconExecClientChartPath = filepaths.Path{
	PackageName: "",
	DirIn:       "./ethereum/beacons/infra/exec_client",
	DirOut:      "./ethereum/outputs",
	FnIn:        "gethHercules", // filename for your gzip workload
	FnOut:       "",
	Env:         "",
	FilterFiles: strings_filter.FilterOpts{},
}

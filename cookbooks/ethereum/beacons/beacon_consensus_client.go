package beacon_cookbooks

import (
	"fmt"
	"time"

	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	strings_filter "github.com/zeus-fyi/zeus/pkg/utils/strings"
	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_req_types"
)

// DeployConsensusClientKnsReq set your own topologyID here after uploading a chart workload
var DeployConsensusClientKnsReq = zeus_req_types.TopologyDeployRequest{
	TopologyID: 0,
	CloudCtxNs: BeaconCloudCtxNs,
}

// chart workload metadata
var consensusClientChart = zeus_req_types.TopologyCreateRequest{
	TopologyName:      "lighthouseHercules",
	ChartName:         "lighthouseHercules",
	ChartDescription:  "lighthouseHercules",
	Version:           fmt.Sprintf("lighthouseHerculesv0.0.%d", time.Now().Unix()),
	ClusterBaseName:   "ethereumBeacons",
	ComponentBaseName: "consensusClient",
	SkeletonBaseName:  "lighthouseHercules",
	Tag:               "latest",
}

// DirOut is where it will write a copy of the chart you uploaded, which helps verify the workload is correct
var beaconConsensusClientChartPath = filepaths.Path{
	PackageName: "",
	DirIn:       "./ethereum/beacons/infra/consensus_client",
	DirOut:      "./ethereum/outputs",
	FnIn:        "lighthouseHercules", // filename for your gzip workload
	FnOut:       "",
	Env:         "",
	FilterFiles: strings_filter.FilterOpts{},
}

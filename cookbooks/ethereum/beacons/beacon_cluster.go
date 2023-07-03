package ethereum_beacon_cookbooks

import (
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	"github.com/zeus-fyi/zeus/zeus/client/zeus_req_types"
)

var Cd = zeus_req_types.ClusterTopologyDeployRequest{
	// ethereumBeacons is a reserved keyword, to make it global, you can replace the below with your own setup
	ClusterClassName:    "ethereumBeacons",
	SkeletonBaseOptions: []string{"gethHercules", "lighthouseHercules"},
	CloudCtxNs:          BeaconCloudCtxNs,
}

var ServiceMonitorChartPath = filepaths.Path{
	PackageName: "",
	DirIn:       "./ethereum/beacons/infra/servicemonitor",
	DirOut:      "./ethereum/beacons/infra/processed_servicemonitor",
	FnIn:        "servicemonitor", // filename for your gzip workload
	FnOut:       "",
	Env:         "",
}

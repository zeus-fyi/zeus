package ethereum_beacon_cookbooks

import "github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_req_types"

var Cd = zeus_req_types.ClusterTopologyDeployRequest{
	// ethereumBeacons is a reserved keyword, to make it global, you can replace the below with your own setup
	ClusterClassName:    "ethereumBeacons",
	SkeletonBaseOptions: []string{"gethHercules", "lighthouseHercules"},
	CloudCtxNs:          BeaconCloudCtxNs,
}

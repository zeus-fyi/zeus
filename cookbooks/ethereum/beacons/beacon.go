package beacon_cookbooks

import (
	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_common_types"
)

const (
	start    = "start"
	download = "download"
)

var BeaconCloudCtxNs = zeus_common_types.CloudCtxNs{
	CloudProvider: "do",
	Region:        "sfo3",
	Context:       "do-sfo3-dev-do-sfo3-zeus",
	Namespace:     "beacon", // set with your own namespace
	Env:           "production",
}

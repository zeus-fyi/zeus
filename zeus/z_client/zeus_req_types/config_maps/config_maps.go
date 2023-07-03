package zeus_config_map_reqs

import (
	strings_filter "github.com/zeus-fyi/zeus/pkg/utils/strings"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_req_types"
)

const (
	KeySwapAction              = "key-swap"
	SetOrCreateKeyFromExisting = "set-or-create-from-key"
)

type ConfigMapActionRequest struct {
	zeus_req_types.TopologyDeployRequest
	Action        string
	ConfigMapName string
	Keys          KeySwap
	FilterOpts    *strings_filter.FilterOpts
}

// KeySwap If using create new key from existing then keyOne=keyToCopy, keyTwo=keyToSetOrCreateFromCopy
type KeySwap struct {
	KeyOne string `json:"keyOne"`
	KeyTwo string `json:"keyTwo"`
}

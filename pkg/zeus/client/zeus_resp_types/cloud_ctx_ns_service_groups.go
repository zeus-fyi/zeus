package zeus_resp_types

import (
	"time"

	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_common_types"
)

type TopologiesOrgCloudCtxNs struct {
	CloudCtxNsID int `json:"cloudCtxNsID"`
	zeus_common_types.CloudCtxNs
	CreatedAt time.Time `json:"createdAt"`
}
type TopologiesOrgCloudCtxNsSlice []TopologiesOrgCloudCtxNs

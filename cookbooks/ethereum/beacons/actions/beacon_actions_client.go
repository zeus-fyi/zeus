package beacon_actions

import (
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	zeus_client "github.com/zeus-fyi/zeus/pkg/zeus/client"
	resty_base "github.com/zeus-fyi/zeus/pkg/zeus/client/base"
	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_req_types"
)

type BeaconActionsClient struct {
	zeus_client.ZeusClient
	BeaconKnsReq    zeus_req_types.TopologyDeployRequest
	PrintPath       filepaths.Path
	ConfigPaths     filepaths.Path
	ConsensusClient string
	ExecClient      string
}

func NewBeaconActionsClient(baseURL, bearer string, kCtxNs zeus_req_types.TopologyDeployRequest) BeaconActionsClient {
	z := BeaconActionsClient{}
	z.BeaconKnsReq = kCtxNs
	z.Resty = resty_base.GetBaseRestyClient(baseURL, bearer)
	return z
}

const ZeusEndpoint = "https://api.zeus.fyi"

func NewDefaultBeaconActionsClient(bearer string, kCtxNs zeus_req_types.TopologyDeployRequest) BeaconActionsClient {
	return NewBeaconActionsClient(ZeusEndpoint, bearer, kCtxNs)
}

const ZeusLocalEndpoint = "http://localhost:9000"

func NewLocalBeaconActionsClient(bearer string, kCtxNs zeus_req_types.TopologyDeployRequest) BeaconActionsClient {
	return NewBeaconActionsClient(ZeusLocalEndpoint, bearer, kCtxNs)
}

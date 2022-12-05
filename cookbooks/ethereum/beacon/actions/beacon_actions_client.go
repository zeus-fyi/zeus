package beacon_actions

import (
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	zeus_client "github.com/zeus-fyi/zeus/pkg/zeus/client"
	resty_base "github.com/zeus-fyi/zeus/pkg/zeus/client/base"
)

type BeaconActionsClient struct {
	zeus_client.ZeusClient
	PrintPath       filepaths.Path
	ConfigPaths     filepaths.Path
	ConsensusClient string
	ExecClient      string
}

func NewBeaconActionsClient(baseURL, bearer string) BeaconActionsClient {
	z := BeaconActionsClient{}
	z.Resty = resty_base.GetBaseRestyTestClient(baseURL, bearer)
	return z
}

const ZeusEndpoint = "https://api.zeus.fyi"

func NewDefaultBeaconActionsClient(bearer string) BeaconActionsClient {
	return NewBeaconActionsClient(ZeusEndpoint, bearer)
}

const ZeusLocalEndpoint = "http://localhost:9000"

func NewLocalBeaconActionsClient(bearer string) BeaconActionsClient {
	return NewBeaconActionsClient(ZeusLocalEndpoint, bearer)
}

package zeus_req_types

import "github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_common_types"

type TopologyCreateRequest struct {
	TopologyName     string `json:"topologyName"`
	ChartName        string `json:"chartName"`
	ChartDescription string `json:"chartDescription,omitempty"`
	Version          string `json:"version"`
}

type TopologyDeployRequest struct {
	TopologyID int `json:"topologyID"`
	zeus_common_types.CloudCtxNs
}

type TopologyRequest struct {
	TopologyID int `json:"topologyID"`
}

type TopologyCloudCtxNsQueryRequest struct {
	zeus_common_types.CloudCtxNs
}

package zeus_req_types

import "github.com/zeus-fyi/zeus/zeus/z_client/zeus_common_types"

type TopologyCreateRequest struct {
	TopologyName     string `json:"topologyName"`
	ChartName        string `json:"chartName"`
	ChartDescription string `json:"chartDescription,omitempty"`
	Version          string `json:"version"`

	ClusterClassName  string `json:"clusterClassName,omitempty"`
	ComponentBaseName string `json:"componentBaseName,omitempty"`
	SkeletonBaseName  string `json:"skeletonBaseName,omitempty"`
	Tag               string `json:"tag,omitempty"`
}

type TopologyDeployRequest struct {
	TopologyID                   int `json:"topologyID"`
	zeus_common_types.CloudCtxNs `json:"cloudCtxNs"`

	SecretRef                       string `json:"secretRef,omitempty"`
	RequestChoreographySecretDeploy bool   `json:"requestChoreographySecretDeploy,omitempty"`
}

type TopologyRequest struct {
	TopologyID int `json:"topologyID"`
}

type TopologyCloudCtxNsQueryRequest struct {
	zeus_common_types.CloudCtxNs
}

type ClusterTopologyDeployRequest struct {
	ClusterClassName             string   `json:"clusterClassName"`
	SkeletonBaseOptions          []string `json:"skeletonBaseOptions"`
	AppTaint                     bool     `json:"appTaint,omitempty"`
	zeus_common_types.CloudCtxNs `json:"cloudCtxNs"`
}

type ClusterTopology struct {
	ClusterClassName string              `json:"clusterClassName"`
	Topologies       []ClusterTopologies `json:"topologies"`
}

type ClusterTopologies struct {
	TopologyID       int    `json:"topologyID"`
	SkeletonBaseName string `json:"skeletonBaseName"`
	Tag              string `json:"tag"`
}

// class creation

type TopologyCreateClusterClassRequest struct {
	ClusterClassName string `json:"clusterClassName"`
}

type TopologyCreateOrAddComponentBasesToClassesRequest struct {
	ClusterClassName   string   `json:"clusterClassName,omitempty"`
	ComponentBaseNames []string `json:"componentBaseNames,omitempty"`
}

type TopologyCreateOrAddSkeletonBasesToClassesRequest struct {
	ClusterClassName  string   `json:"clusterClassName"`
	ComponentBaseName string   `json:"componentBaseName,omitempty"`
	SkeletonBaseNames []string `json:"skeletonBaseNames,omitempty"`
}

package zeus_pods_reqs

import (
	strings_filter "github.com/zeus-fyi/zeus/pkg/utils/strings"
	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_req_types"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PodActionRequest struct {
	zeus_req_types.TopologyDeployRequest
	Action        string `json:"action"`
	PodName       string `json:"podName,omitempty"`
	ContainerName string `json:"containerName,omitempty"`

	FilterOpts *strings_filter.FilterOpts
	ClientReq  *ClientRequest
	LogOpts    *v1.PodLogOptions
	DeleteOpts *metav1.DeleteOptions
}

type ClientRequest struct {
	MethodHTTP      string
	Endpoint        string
	Ports           []string
	Payload         any
	EndpointHeaders map[string]string
}

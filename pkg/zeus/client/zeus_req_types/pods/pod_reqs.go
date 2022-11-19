package zeus_pods_reqs

import (
	strings_filter "github.com/zeus-fyi/zeus/pkg/utils/strings"
	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_req_types"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PodActionRequest struct {
	zeus_req_types.TopologyDeployRequest
	Action        string
	PodName       string
	ContainerName string

	FilterOpts *strings_filter.FilterOpts
	ClientReq  *ClientRequest
	LogOpts    *v1.PodLogOptions
	DeleteOpts *metav1.DeleteOptions
}

type ClientRequest struct {
	MethodHTTP      string
	Endpoint        string
	Ports           []string
	Payload         *string
	PayloadBytes    *[]byte
	EndpointHeaders map[string]string
}

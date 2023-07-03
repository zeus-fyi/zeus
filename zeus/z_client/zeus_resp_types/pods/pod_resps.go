package zeus_pods_resp

import (
	"time"

	v1 "k8s.io/api/core/v1"
)

type ClientResp struct {
	ReplyBodies map[string][]byte
}

func (c *ClientResp) GetAnyValue() []byte {
	for _, v := range c.ReplyBodies {
		return v
	}
	return nil
}

type PodsSummary struct {
	Pods map[string]PodSummary `json:"pods"`
}

type PodSummary struct {
	PodName               string                        `json:"podName"`
	Phase                 string                        `json:"podPhase"`
	Message               string                        `json:"message"`
	Reason                string                        `json:"reason"`
	StartTime             time.Time                     `json:"startTime"`
	PodConditions         []v1.PodCondition             `json:"podConditions"`
	InitContainerStatuses map[string]v1.ContainerStatus `json:"initContainerConditions"`
	ContainerStatuses     map[string]v1.ContainerStatus `json:"containerStatuses"`
}

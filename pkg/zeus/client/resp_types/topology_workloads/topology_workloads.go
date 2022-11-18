package topology_workloads

import (
	v1 "k8s.io/api/apps/v1"
	v1core "k8s.io/api/core/v1"
	v1networking "k8s.io/api/networking/v1"
)

type TopologyBaseInfraWorkload struct {
	*v1core.Service       `json:"service"`
	*v1core.ConfigMap     `json:"configMap"`
	*v1.Deployment        `json:"deployment"`
	*v1.StatefulSet       `json:"statefulSet"`
	*v1networking.Ingress `json:"ingress"`
}

func NewTopologyBaseInfraWorkload() TopologyBaseInfraWorkload {
	k8s := TopologyBaseInfraWorkload{
		StatefulSet: nil,
		Deployment:  nil,
		Service:     nil,
		Ingress:     nil,
		ConfigMap:   nil,
	}
	return k8s
}

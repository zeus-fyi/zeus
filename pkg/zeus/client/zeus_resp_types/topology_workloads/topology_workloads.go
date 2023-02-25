package topology_workloads

import (
	v1sm "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
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
	*v1sm.ServiceMonitor  `json:"serviceMonitor"`
}

func NewTopologyBaseInfraWorkload() TopologyBaseInfraWorkload {
	k8s := TopologyBaseInfraWorkload{
		StatefulSet:    nil,
		Deployment:     nil,
		Service:        nil,
		Ingress:        nil,
		ConfigMap:      nil,
		ServiceMonitor: nil,
	}
	return k8s
}

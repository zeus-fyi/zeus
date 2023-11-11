package live_workload_query

import (
	v1apps "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	v1networking "k8s.io/api/networking/v1"
)

type NamespaceWorkload struct {
	*v1.PodList               `json:"podList,omitempty"`
	*v1.ServiceList           `json:"serviceList,omitempty"`
	*v1networking.IngressList `json:"ingressList,omitempty"`
	*v1apps.StatefulSetList   `json:"statefulSetList,omitempty"`
	*v1apps.DeploymentList    `json:"deploymentList,omitempty"`
	*v1.ConfigMapList         `json:"configMapList,omitempty"`
}

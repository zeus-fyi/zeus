package live_workload_query

import (
	v1apps "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	v1networking "k8s.io/api/networking/v1"
)

type NamespaceWorkload struct {
	*v1.PodList               `json:"podList"`
	*v1.ServiceList           `json:"serviceList"`
	*v1networking.IngressList `json:"ingressList"`
	*v1apps.StatefulSetList   `json:"statefulSetList"`
	*v1apps.DeploymentList    `json:"deploymentList"`
	*v1.ConfigMapList         `json:"configMapList"`
}

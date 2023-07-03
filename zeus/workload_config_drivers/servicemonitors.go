package zeus_topology_config_drivers

import v1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"

type ServiceMonitorDriver struct {
	v1.ServiceMonitor
}

func (sm *ServiceMonitorDriver) SetServiceMonitorConfigs(smn *v1.ServiceMonitor) {
	if smn == nil {
		return
	}
	if sm.Name != "" {
		smn.Name = sm.Name
	}
	if sm.Labels != nil {
		smn.Labels = sm.Labels
	}
	if sm.Annotations != nil {
		smn.Annotations = sm.Annotations
	}
	if sm.Spec.Selector.MatchLabels != nil {
		smn.Spec.Selector.MatchLabels = sm.Spec.Selector.MatchLabels
	}
}

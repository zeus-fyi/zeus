package zeus_topology_config_drivers

import v1 "k8s.io/api/core/v1"

type ConfigMapDriver struct {
}

func (cm *ConfigMapDriver) SetConfigMaps(cmDriver *v1.ConfigMap) {
	if cm == nil {
		return
	}
	// TODO
}

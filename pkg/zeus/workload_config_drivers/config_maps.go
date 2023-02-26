package zeus_topology_config_drivers

import v1 "k8s.io/api/core/v1"

type ConfigMapDriver struct {
	v1.ConfigMap
	// swap key for values in the configmap
	SwapKeys map[string]string
}

func (cm *ConfigMapDriver) SetConfigMaps(cmap *v1.ConfigMap) {
	if cmap == nil {
		return
	}
	if cmap.Data != nil {
		for originalKey, _ := range cmap.Data {
			if swapValue, ok := cm.SwapKeys[originalKey]; ok {
				cmap.Data[originalKey] = swapValue
			}
		}
	}

}

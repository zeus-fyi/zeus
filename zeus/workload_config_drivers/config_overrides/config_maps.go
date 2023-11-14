package config_overrides

import v1 "k8s.io/api/core/v1"

type ConfigMapDriver struct {
	v1.ConfigMap
	// swap key for values key in the configmap
	SwapKeys map[string]string
}

func (cm *ConfigMapDriver) SetConfigMaps(cmap *v1.ConfigMap) {
	if cmap == nil {
		return
	}
	if cmap.Data != nil {
		for swapKey, newContentsKey := range cm.SwapKeys {
			if _, ok := cmap.Data[swapKey]; ok {
				nc := cmap.Data[newContentsKey]
				cmap.Data[swapKey] = nc
			}
		}
	}
	if cmap.Data == nil {
		cmap.Data = make(map[string]string)
	}
	for k, v := range cm.Data {
		cmap.Data[k] = v
	}
}

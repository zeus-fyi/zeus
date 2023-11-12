package zk8s_templates

import (
	"context"

	zeus_topology_config_drivers "github.com/zeus-fyi/zeus/zeus/workload_config_drivers"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetConfigMapTemplate(ctx context.Context, name string) *v1.ConfigMap {
	return &v1.ConfigMap{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ConfigMap",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: GetConfigMapName(ctx, name),
		},
		Data: make(map[string]string),
	}
}

func BuildConfigMapDriver(ctx context.Context, configMap ConfigMap) (zeus_topology_config_drivers.ConfigMapDriver, error) {
	cmDriver := zeus_topology_config_drivers.ConfigMapDriver{
		ConfigMap: v1.ConfigMap{
			Data: make(map[string]string),
		},
	}
	for key, value := range configMap {
		cmDriver.ConfigMap.Data[key] = value
	}
	return cmDriver, nil
}

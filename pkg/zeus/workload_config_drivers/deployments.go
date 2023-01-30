package zeus_topology_config_drivers

import (
	v1 "k8s.io/api/apps/v1"
)

type DeploymentDriver struct {
	ContainerDrivers map[string]ContainerDriver
}

func (d *DeploymentDriver) SetDeploymentConfigs(sts *v1.Deployment) {
	if d == nil {
		return
	}
	for i, c := range sts.Spec.Template.Spec.Containers {
		if v, ok := d.ContainerDrivers[c.Name]; ok {
			v.SetContainerConfigs(&c)
			sts.Spec.Template.Spec.Containers[i] = c
		}
	}
}

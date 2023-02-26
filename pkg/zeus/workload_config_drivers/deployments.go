package zeus_topology_config_drivers

import (
	v1 "k8s.io/api/apps/v1"
)

type DeploymentDriver struct {
	ContainerDrivers map[string]ContainerDriver
}

func NewDeploymentDriver() DeploymentDriver {
	return DeploymentDriver{ContainerDrivers: make(map[string]ContainerDriver)}
}

func (d *DeploymentDriver) SetDeploymentConfigs(dep *v1.Deployment) {
	if dep == nil {
		return
	}
	// TODO refactor into pod template spec config, then share w/sts + here
	// init containers
	for i, c := range dep.Spec.Template.Spec.InitContainers {
		if v, ok := d.ContainerDrivers[c.Name]; ok {
			v.SetContainerConfigs(&c)
			dep.Spec.Template.Spec.InitContainers[i] = c
		}
	}
	// containers
	for i, c := range dep.Spec.Template.Spec.Containers {
		if v, ok := d.ContainerDrivers[c.Name]; ok {
			v.SetContainerConfigs(&c)
			dep.Spec.Template.Spec.Containers[i] = c
		}
	}
}

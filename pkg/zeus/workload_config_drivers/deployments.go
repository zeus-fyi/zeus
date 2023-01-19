package zeus_topology_config_drivers

import (
	v1 "k8s.io/api/apps/v1"
	v1Core "k8s.io/api/core/v1"
)

type DeploymentDriver struct {
	ContainerDrivers map[string]v1Core.Container
}

func (d *DeploymentDriver) SetDeploymentConfigs(sts *v1.Deployment) {
	if d == nil {
		return
	}
	// TODO
}

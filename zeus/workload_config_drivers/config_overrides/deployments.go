package config_overrides

import (
	v1 "k8s.io/api/apps/v1"
)

type DeploymentDriver struct {
	ReplicaCount     *int32
	ContainerDrivers map[string]ContainerDriver
}

func NewDeploymentDriver() DeploymentDriver {
	return DeploymentDriver{ContainerDrivers: make(map[string]ContainerDriver)}
}

func (d *DeploymentDriver) SetDeploymentConfigs(dep *v1.Deployment) {
	if dep == nil {
		return
	}
	if d.ReplicaCount != nil {
		dep.Spec.Replicas = d.ReplicaCount
	}
	for _, contDriver := range d.ContainerDrivers {
		if contDriver.IsAppendContainer {
			if contDriver.IsInitContainer {
				dep.Spec.Template.Spec.InitContainers = append(dep.Spec.Template.Spec.InitContainers, contDriver.Container)
			} else {
				dep.Spec.Template.Spec.Containers = append(dep.Spec.Template.Spec.Containers, contDriver.Container)
			}
		}
	}
	for i, c := range dep.Spec.Template.Spec.InitContainers {
		if v, ok := d.ContainerDrivers[c.Name]; ok {
			v.SetContainerConfigs(&c)
			dep.Spec.Template.Spec.InitContainers[i] = c
		}
	}
	for i, c := range dep.Spec.Template.Spec.Containers {
		if v, ok := d.ContainerDrivers[c.Name]; ok {
			v.SetContainerConfigs(&c)
			dep.Spec.Template.Spec.Containers[i] = c
		}
	}
}

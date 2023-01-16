package zeus_topology_config_drivers

import v1Core "k8s.io/api/core/v1"

type ContainerDriver struct {
	v1Core.Container
}

func (c *ContainerDriver) SetContainerConfigs(cont *v1Core.Container) {
	if len(c.Image) > 0 {
		cont.Image = c.Image
	}
	if c.Env != nil {
		cont.Env = c.Env
	}
	if c.Ports != nil {
		cont.Ports = c.Ports
	}
	if c.Command != nil {
		cont.Command = c.Command
	}
	if c.Args != nil {
		cont.Args = c.Args
	}
}

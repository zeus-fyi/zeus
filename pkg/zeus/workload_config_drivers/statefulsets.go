package zeus_topology_config_drivers

import (
	v1 "k8s.io/api/apps/v1"
	v1Core "k8s.io/api/core/v1"
)

type StatefulSetDriver struct {
	ContainerDrivers map[string]v1Core.Container
}

func (s *StatefulSetDriver) SetStatefulSetConfigs(sts *v1.StatefulSet) {
	if sts == nil {
		return
	}

	// TODO, create container config override in own file
	//
	//for cn, driver := range s.ContainerDrivers {
	//
	//}
	for i, c := range sts.Spec.Template.Spec.Containers {
		if v, ok := s.ContainerDrivers[c.Name]; ok {
			if len(v.Image) > 0 {
				sts.Spec.Template.Spec.Containers[i].Image = v.Image
			}
			if v.Env != nil {
				sts.Spec.Template.Spec.Containers[i].Env = v.Env
			}
			if v.Ports != nil {
				sts.Spec.Template.Spec.Containers[i].Ports = v.Ports
			}
			if v.Command != nil {
				sts.Spec.Template.Spec.Containers[i].Command = v.Command
			}
			if v.Args != nil {
				sts.Spec.Template.Spec.Containers[i].Args = v.Args
			}
		}
	}

}

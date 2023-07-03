package zeus_topology_config_drivers

import (
	v1Core "k8s.io/api/core/v1"
)

type ContainerDriver struct {
	IsAppendContainer bool
	IsInitContainer   bool
	v1Core.Container
	AppendEnvVars []v1Core.EnvVar
}

func (cd *ContainerDriver) SetContainerConfigs(cont *v1Core.Container) {
	if len(cd.Image) > 0 {
		cont.Image = cd.Image
	}
	if cd.Env != nil {
		cont.Env = cd.Env
	}
	if cd.AppendEnvVars != nil {
		cont.Env = append(cont.Env, cd.AppendEnvVars...)
	}
	if cd.Ports != nil {
		cont.Ports = cd.Ports
	}
	if cd.Command != nil {
		cont.Command = cd.Command
	}
	if cd.Args != nil {
		cont.Args = cd.Args
	}
	if cd.Resources.Limits != nil {
		cont.Resources.Limits = cd.Resources.Limits
	}
	if cd.Resources.Requests != nil {
		cont.Resources.Requests = cd.Resources.Requests
	}
	if cd.ImagePullPolicy != "" {
		cont.ImagePullPolicy = cd.ImagePullPolicy
	}
}

func (cd *ContainerDriver) CreateEnvVarKeyValue(k, v string) v1Core.EnvVar {
	return v1Core.EnvVar{
		Name:  k,
		Value: v,
	}
}

func MakeEnvVar(name, key, localObjRef string) v1Core.EnvVar {
	return v1Core.EnvVar{
		Name: name,
		ValueFrom: &v1Core.EnvVarSource{
			FieldRef:         nil,
			ResourceFieldRef: nil,
			ConfigMapKeyRef:  nil,
			SecretKeyRef: &v1Core.SecretKeySelector{
				LocalObjectReference: v1Core.LocalObjectReference{Name: localObjRef},
				Key:                  key,
			},
		},
	}
}

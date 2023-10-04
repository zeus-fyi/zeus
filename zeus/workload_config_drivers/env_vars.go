package zeus_topology_config_drivers

import v1Core "k8s.io/api/core/v1"

func MakeSecretEnvVar(name, key, localObjRef string) v1Core.EnvVar {
	return v1Core.EnvVar{
		Name: name,
		ValueFrom: &v1Core.EnvVarSource{
			SecretKeyRef: &v1Core.SecretKeySelector{
				LocalObjectReference: v1Core.LocalObjectReference{Name: localObjRef},
				Key:                  key,
			},
		},
	}
}

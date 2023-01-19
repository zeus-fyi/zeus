package ethereum_hercules

import v1 "k8s.io/api/core/v1"

const (
	Hercules                    = "hercules"
	HerculesImage               = "zeusfyi/hercules:latest"
	HerculesImagePullPolicy     = "Always"
	HerculesContainerPortNumber = int32(9003)
)

var (
	HerculesContainerDefaultConfig = v1.Container{
		Name:            Hercules,
		Image:           HerculesImage,
		Command:         nil,
		Args:            nil,
		WorkingDir:      "",
		Ports:           []v1.ContainerPort{HerculesPort},
		EnvFrom:         nil,
		Env:             nil,
		Resources:       v1.ResourceRequirements{},
		VolumeMounts:    nil,
		VolumeDevices:   nil,
		LivenessProbe:   nil,
		ReadinessProbe:  nil,
		StartupProbe:    nil,
		ImagePullPolicy: HerculesImagePullPolicy,
	}
	HerculesPort = v1.ContainerPort{
		Name:          Hercules,
		ContainerPort: HerculesContainerPortNumber,
		Protocol:      "TCP",
	}
)

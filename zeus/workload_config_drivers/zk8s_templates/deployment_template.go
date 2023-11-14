package zk8s_templates

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/zeus-fyi/zeus/zeus/workload_config_drivers/config_overrides"
	v1 "k8s.io/api/apps/v1"
	v1Core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetDeploymentTemplate(ctx context.Context, name string) *v1.Deployment {
	return &v1.Deployment{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Deployment",
			APIVersion: "apps/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:   GetDeploymentName(ctx, name),
			Labels: GetLabels(ctx, name),
		},
		Spec: v1.DeploymentSpec{
			Selector: metav1.SetAsLabelSelector(GetSelector(ctx, name)),
			Template: v1Core.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: GetLabels(ctx, name),
				},
				Spec: v1Core.PodSpec{},
			},
			Strategy: v1.DeploymentStrategy{},
		},
	}
}

func BuildDeploymentDriver(ctx context.Context, containers Containers, dep Deployment) (config_overrides.DeploymentDriver, error) {
	rc := int32(dep.ReplicaCount)
	depDriver := config_overrides.DeploymentDriver{
		ReplicaCount:     &rc,
		ContainerDrivers: make(map[string]config_overrides.ContainerDriver),
	}
	for containerName, container := range containers {
		contDriver, err := BuildContainerDriver(ctx, containerName, container)
		if err != nil {
			log.Error().Err(err).Msg("Failed to build container driver")
			return config_overrides.DeploymentDriver{}, err
		}
		depDriver.ContainerDrivers[containerName] = config_overrides.ContainerDriver{
			IsAppendContainer: true,
			IsInitContainer:   container.IsInitContainer,
			Container:         contDriver,
			AppendEnvVars:     nil,
		}
	}
	return depDriver, nil
}

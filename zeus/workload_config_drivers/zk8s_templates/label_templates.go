package zk8s_templates

import (
	"context"
	"fmt"
)

func GetLabels(ctx context.Context, name string) map[string]string {
	labels := map[string]string{
		"app.kubernetes.io/name":       name,
		"app.kubernetes.io/instance":   name,
		"app.kubernetes.io/managed-by": "zeus",
	}
	return labels
}

func GetIngressName(ctx context.Context, name string) string {
	return name
}

func GetIngressSecretName(ctx context.Context, name string) string {
	return fmt.Sprintf("tls-%s", name)
}

func GetIngressHostName(ctx context.Context, name string) string {
	return fmt.Sprintf("%s.zeus.fyi", name)
}

func GetStatefulSetName(ctx context.Context, name string) string {
	return name
}

func GetDeploymentName(ctx context.Context, name string) string {
	return name
}

func GetServiceName(ctx context.Context, name string) string {
	return name
}

func GetConfigMapName(ctx context.Context, name string) string {
	return fmt.Sprintf("cm-%s", name)
}

func GetJobName(ctx context.Context, name string) string {
	return fmt.Sprintf("job-%s", name)
}

func GetCronJobName(ctx context.Context, name string) string {
	return fmt.Sprintf("cronjob-%s", name)
}

func GetSelector(ctx context.Context, name string) map[string]string {
	labels := map[string]string{
		"app.kubernetes.io/name":     name,
		"app.kubernetes.io/instance": name,
	}
	return labels
}

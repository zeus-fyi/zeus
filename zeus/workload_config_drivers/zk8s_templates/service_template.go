package zk8s_templates

import (
	"context"
	"strconv"

	"github.com/rs/zerolog/log"
	zeus_topology_config_drivers "github.com/zeus-fyi/zeus/zeus/workload_config_drivers/config_overrides"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

/*
metadata:
  name: zeus-service
  labels:
    app.kubernetes.io/name: zeus-client
    app.kubernetes.io/instance: zeus-client
    app.kubernetes.io/managed-by: zeus
spec:
  type: ClusterIP
  ports:
  selector:
    app.kubernetes.io/name: zeus-client
    app.kubernetes.io/instance: zeus-client
*/

func GetServiceTemplate(ctx context.Context, name string) *v1.Service {
	labels := GetLabels(ctx, name)
	selectors := GetSelector(ctx, name)
	return &v1.Service{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:   GetServiceName(ctx, name),
			Labels: labels,
		},
		Spec: v1.ServiceSpec{
			Selector: selectors,
			Type:     v1.ServiceTypeClusterIP,
		},
	}
}

func BuildServiceDriver(ctx context.Context, containers Containers) (zeus_topology_config_drivers.ServiceDriver, error) {
	svcDriver := zeus_topology_config_drivers.ServiceDriver{
		Service: v1.Service{
			Spec: v1.ServiceSpec{
				Ports: []v1.ServicePort{},
			},
		},
	}
	var sps []v1.ServicePort
	for _, c := range containers {
		for _, p := range c.DockerImage.Ports {
			numberInt64, err := strconv.ParseInt(p.Number, 10, 32)
			if err != nil {
				log.Error().Err(err).Msg("failed to parse port number")
				return svcDriver, err
			}

			if p.IngressEnabledPort {
				svcDriver.AddNginxTargetPort("http", p.Name)
			} else {
				sps = append(sps, v1.ServicePort{
					Name:       p.Name,
					Port:       int32(numberInt64),
					Protocol:   v1.Protocol(p.Protocol),
					TargetPort: intstr.IntOrString{Type: intstr.String, StrVal: p.Name},
				})
			}
		}
	}
	svcDriver.Service.Spec.Ports = sps
	return svcDriver, nil
}

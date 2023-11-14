package zk8s_templates

import (
	"context"

	"github.com/google/uuid"
	zeus_topology_config_drivers "github.com/zeus-fyi/zeus/zeus/workload_config_drivers/config_overrides"
	v1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

/*
metadata:
  name: "zeus-client"
  annotations:
    cert-manager.io/cluster-issuer: letsencrypt-prod
spec:
  ingressClassName: "nginx"
  tls:
    - secretName: zeus-client-tls
      hosts:
        - host.zeus.fyi
  rules:
    - host: host.zeus.fyi
*/

func GetIngressTemplate(ctx context.Context, name string) *v1.Ingress {
	ingressClassName := "nginx"

	annotations := make(map[string]string)
	annotations["cert-manager.io/cluster-issuer"] = "letsencrypt-prod"

	md := metav1.ObjectMeta{
		Name:        GetIngressName(ctx, name),
		Annotations: annotations,
	}
	return &v1.Ingress{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Ingress",
			APIVersion: "networking.k8s.io/v1",
		},
		ObjectMeta: md,
		Spec: v1.IngressSpec{
			IngressClassName: &ingressClassName,
			TLS: []v1.IngressTLS{{
				SecretName: GetIngressSecretName(ctx, name),
			}},
		},
	}
}

func BuildIngressDriver(ctx context.Context, cbName string, containers Containers, ing Ingress, ip IngressPaths) (zeus_topology_config_drivers.IngressDriver, error) {
	uid := uuid.New()
	ing.Host = GetIngressHostName(ctx, uid.String())
	var httpPaths []v1.HTTPIngressPath
	for _, pa := range ip {
		pt := v1.PathType(pa.PathType)
		appendPath := v1.HTTPIngressPath{
			Path:     pa.Path,
			PathType: &pt,
			Backend: v1.IngressBackend{
				Service: &v1.IngressServiceBackend{
					Name: GetServiceName(ctx, cbName),
					Port: v1.ServiceBackendPort{
						Number: int32(80),
					},
				},
			},
		}
		httpPaths = append(httpPaths, appendPath)
	}

	ingressRuleValue := v1.IngressRuleValue{HTTP: &v1.HTTPIngressRuleValue{Paths: httpPaths}}
	ingDriver := zeus_topology_config_drivers.IngressDriver{
		Ingress: v1.Ingress{
			Spec: v1.IngressSpec{
				TLS: []v1.IngressTLS{{
					Hosts:      []string{ing.Host},
					SecretName: GetIngressSecretName(ctx, uid.String()),
				}},
				Rules: []v1.IngressRule{{
					Host:             ing.Host,
					IngressRuleValue: ingressRuleValue,
				}},
			},
		},
		Host:         ing.Host,
		NginxAuthURL: ing.AuthServerURL,
	}
	return ingDriver, nil
}

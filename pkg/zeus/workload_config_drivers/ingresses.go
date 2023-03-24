package zeus_topology_config_drivers

import v1 "k8s.io/api/networking/v1"

const (
	nginxAuthURLAnnotation = "nginx.ingress.kubernetes.io/auth-url"
)

type IngressDriver struct {
	v1.Ingress
	NginxAuthURL string
}

func (i *IngressDriver) SetIngressConfigs(ing *v1.Ingress) {
	if ing == nil {
		return
	}
	if ing.Annotations == nil {
		ing.Annotations = make(map[string]string)
	}
	tmp := ing.Annotations
	if i.NginxAuthURL != "" {
		tmp[nginxAuthURLAnnotation] = i.NginxAuthURL
	}
	if i.Ingress.Spec.TLS != nil {
		ing.Spec.TLS = i.Ingress.Spec.TLS
	}
	if i.Ingress.Spec.Rules != nil {
		ing.Spec.Rules = i.Ingress.Spec.Rules
	}
	ing.Annotations = tmp
}

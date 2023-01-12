package zeus_topology_config_drivers

import v1 "k8s.io/api/networking/v1"

const (
	nginxAuthURLAnnotation = "nginx.ingress.kubernetes.io/auth-url"
)

type IngressDriver struct {
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
	ing.Annotations = tmp
}

package zeus_topology_config_drivers

import v1 "k8s.io/api/networking/v1"

const (
	nginxAuthURLAnnotation = "nginx.ingress.kubernetes.io/auth-url"
)

type IngressDriver struct {
	v1.Ingress
	Host         string
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
		for ind, _ := range i.Ingress.Spec.TLS {
			ing.Spec.TLS[ind].SecretName = i.Spec.TLS[ind].SecretName
			ing.Spec.TLS[ind].Hosts = i.Spec.TLS[ind].Hosts
		}
	}
	if len(i.Host) > 0 {
		for ind, _ := range ing.Spec.Rules {
			ing.Spec.Rules[ind].Host = i.Host
		}
	}
	if i.Spec.Rules != nil {
		for ind, _ := range i.Spec.Rules {
			ing.Spec.Rules[ind].Host = i.Spec.Rules[ind].Host
			ing.Spec.Rules[ind].HTTP = i.Spec.Rules[ind].HTTP
		}
	}
	ing.Annotations = tmp
}

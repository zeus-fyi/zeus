package ethereum_beacon_cookbooks

import (
	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_resp_types/topology_workloads"
)

const (
	ingressName        = "ephemeral"
	secretName         = "ephemeral-tls"
	host               = "eth.ephemeral.zeus.fyi"
	ephemeralNamespace = "ephemeral"
)

func EphemeralIngressConfig(inf topology_workloads.TopologyBaseInfraWorkload) {
	if inf.Ingress != nil {
		inf.Ingress.ObjectMeta.Name = ingressName
		inf.Ingress.ObjectMeta.Namespace = ephemeralNamespace

		// assumes only one exists
		for i, _ := range inf.Ingress.Spec.TLS {
			inf.Ingress.Spec.TLS[i].SecretName = secretName
			inf.Ingress.Spec.TLS[i].Hosts = []string{host}
		}
		// assumes only one exists
		for i, _ := range inf.Ingress.Spec.Rules {
			inf.Ingress.Spec.Rules[i].Host = host
		}
	}
}

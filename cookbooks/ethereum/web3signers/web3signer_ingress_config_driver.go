package web3signer_cookbooks

import "github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_resp_types/topology_workloads"

const (
	nginxAuthURLAnnotation = "nginx.ingress.kubernetes.io/auth-url"
)

type Web3SignerIngressConfig struct {
	AuthURL string
}

func Web3SignerIngressConfigDriver(inf topology_workloads.TopologyBaseInfraWorkload, cfg Web3SignerIngressConfig) {
	if inf.Ingress != nil {
		if len(inf.Annotations) == 0 {
			inf.Annotations = make(map[string]string)
		}
		tmp := inf.Annotations
		if cfg.AuthURL != "" {
			tmp[nginxAuthURLAnnotation] = cfg.AuthURL
		}
		inf.Annotations = tmp
	}
}

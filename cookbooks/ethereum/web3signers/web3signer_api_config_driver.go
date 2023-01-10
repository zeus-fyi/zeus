package web3signer_cookbooks

import (
	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_resp_types/topology_workloads"
	v1 "k8s.io/api/core/v1"
)

type Web3SignerStatefulSetConfig struct {
	Web3SignerContainerCfg
}

// TODO add ports, cmd & args to config driver

type Web3SignerContainerCfg struct {
	CustomImage string
	EnvVars     []v1.EnvVar
}

func Web3SignerAPIConfigDriver(inf topology_workloads.TopologyBaseInfraWorkload, stsCfg Web3SignerStatefulSetConfig) {
	if inf.StatefulSet != nil {
		for i, c := range inf.StatefulSet.Spec.Template.Spec.Containers {
			if c.Name == web3SignerClient {
				if len(stsCfg.CustomImage) > 0 {
					inf.StatefulSet.Spec.Template.Spec.Containers[i].Image = stsCfg.CustomImage
				} else {
					inf.StatefulSet.Spec.Template.Spec.Containers[i].Image = web3signerDockerImage
				}
				if stsCfg.EnvVars != nil {
					c.Env = stsCfg.EnvVars
				}
			}
		}
	}
}

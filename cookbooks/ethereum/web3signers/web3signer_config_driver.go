package web3signer_cookbooks

import "github.com/zeus-fyi/zeus/zeus/z_client/zeus_resp_types/topology_workloads"

const (
	web3SignerClient      = "zeus-web3signer"
	web3signerDockerImage = "consensys/web3signer:22.11"
)

func EphemeralWeb3SignerConfig(inf topology_workloads.TopologyBaseInfraWorkload, customImage string) {
	if inf.StatefulSet != nil {
		for i, c := range inf.StatefulSet.Spec.Template.Spec.Containers {
			if c.Name == web3SignerClient {
				if len(customImage) > 0 {
					inf.StatefulSet.Spec.Template.Spec.Containers[i].Image = customImage
				} else {
					inf.StatefulSet.Spec.Template.Spec.Containers[i].Image = web3signerDockerImage
				}
			}
		}
	}
}

package web3signer_cookbooks

import (
	web3signer_cmds_ai_generated "github.com/zeus-fyi/zeus/cookbooks/ethereum/web3signers/web3signer_cmds/ai_generated"
	zeus_topology_config_drivers "github.com/zeus-fyi/zeus/pkg/zeus/workload_config_drivers"
	v1 "k8s.io/api/core/v1"
)

func GetWeb3SignerAPIStatefulSetConfig(customImage string) zeus_topology_config_drivers.StatefulSetDriver {
	args, _ := web3signer_cmds_ai_generated.Web3SignerAPICmd.CreateFieldsForCLI("eth2")
	port := v1.ContainerPort{
		Name:          "http",
		ContainerPort: 9000,
	}
	c := v1.Container{
		Name:      web3SignerClient,
		Image:     customImage,
		Command:   []string{"/bin/sh"},
		Args:      args,
		Ports:     []v1.ContainerPort{port},
		Env:       []v1.EnvVar{},
		Resources: v1.ResourceRequirements{},
	}

	sc := zeus_topology_config_drivers.StatefulSetDriver{}
	sc.ContainerDrivers = make(map[string]v1.Container)
	sc.ContainerDrivers[c.Name] = c
	return sc
}

func GetWeb3SignerAPIServiceConfig() zeus_topology_config_drivers.ServiceDriver {
	s := zeus_topology_config_drivers.ServiceDriver{ExtendPorts: []v1.ServicePort{}}
	s.AddNginxTargetPort("nginx", "http")
	return s
}

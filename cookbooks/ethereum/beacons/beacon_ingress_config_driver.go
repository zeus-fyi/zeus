package ethereum_beacon_cookbooks

import (
	"fmt"
	"time"

	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_req_types"
	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_resp_types/topology_workloads"
	zeus_cluster_config_drivers "github.com/zeus-fyi/zeus/pkg/zeus/cluster_config_drivers"
)

const (
	ingressName        = "ephemeral"
	secretName         = "ephemeral-tls"
	host               = "eth.ephemeral.zeus.fyi"
	ephemeralNamespace = "ephemeral"
)

var BeaconIngressSkeletonBaseConfig = zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
	SkeletonBaseChart:         zeus_req_types.TopologyCreateRequest{},
	SkeletonBaseNameChartPath: IngressChartPath,
}

var IngressChart = zeus_req_types.TopologyCreateRequest{
	TopologyName:      "beaconIngress",
	ChartName:         "beaconIngress",
	ChartDescription:  "beaconIngress",
	Version:           fmt.Sprintf("beaconIngress-v.0.%d", time.Now().Unix()),
	SkeletonBaseName:  "beaconIngress",
	ComponentBaseName: "beaconIngress",
	ClusterClassName:  "ethereumBeacon",
	Tag:               "latest",
}

var IngressChartPath = filepaths.Path{
	PackageName: "",
	DirIn:       "./ethereum/beacons/infra/ingress",
	DirOut:      "./ethereum/beacons/infra/processed_beacon_ingress",
	FnIn:        "beaconIngress", // filename for your gzip workload
	FnOut:       "",
	Env:         "",
}

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

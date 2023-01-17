package system_config_drivers

import zeus_cluster_config_drivers "github.com/zeus-fyi/zeus/pkg/zeus/cluster_config_drivers"

type MatrixDefinition struct {
	MatrixName string

	// multi cluster setup, eg 10 ethereum beacons, at supplied cloud ctx ns locations
	Clusters []zeus_cluster_config_drivers.ClusterDefinition
}

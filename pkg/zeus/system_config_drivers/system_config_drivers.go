package system_config_drivers

type SystemDefinition struct {
	SystemName string

	// large scale infra setup multi-region, multi-cloud,
	// eg 10 ethereum beacons, 3 databases, 5 validator clusters, etc
	// at supplied cloud ctx ns locations
	Matrices []MatrixDefinition
}

package routines

func GetProcessName(cn string) string {
	appName := ""
	switch cn {
	case "lighthouse":
		appName = "lighthouse"
	case "prysm":
		appName = "prysm"
	case "lodestar":
		appName = "lodestar"
	case "teku":
		appName = "teku"
	case "nethermind":
		appName = "nethermind"
	case "besu":
		appName = "besu"
	case "geth":
		appName = "geth"
	}
	return appName
}

// TODO others...
func GetProcessPorts(cn string) string {
	port := ""
	switch cn {
	case "lighthouse":
		port = "5052"
	case "geth":
		port = "8545"
	}
	return port
}

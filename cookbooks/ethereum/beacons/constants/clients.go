package client_consts

func GetAnyClientApiPorts(clientName string) []string {
	switch clientName {
	case Lighthouse:
		return LighthouseBeaconPorts
	case Geth:
		return GethBeaconPorts
	}

	return []string{}
}

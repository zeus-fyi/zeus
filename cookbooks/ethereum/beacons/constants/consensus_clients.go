package client_consts

const (
	Lighthouse = "lighthouse"
	Prysm      = "prysm"
	Lodestart  = "lodestar"
	Teku       = "teku"
)

var LighthouseBeaconPorts = []string{"5052:5052"}

func GetClientBeaconPortsHTTP(clientName string) []string {
	switch clientName {
	case Lighthouse:
		return LighthouseBeaconPorts
	}
	return []string{}
}

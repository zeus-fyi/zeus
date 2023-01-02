package client_consts

const (
	Lighthouse = "lighthouse"
	Prysm      = "prysm"
	Lodestar   = "lodestar"
	Teku       = "teku"
)

var (
	LighthouseBeaconPorts           = []string{"5052:5052"}
	LighthouseValidatorClientPorts  = []string{"5062:5062"}
	LighthouseWeb3SignerAPIEndpoint = "/lighthouse/validators/web3signer"
)

func GetClientBeaconPortsHTTP(clientName string) []string {
	switch clientName {
	case Lighthouse:
		return LighthouseBeaconPorts
	}
	return []string{}
}

func IsConsensusClient(name string) bool {
	switch name {
	case Lighthouse, Prysm, Lodestar, Teku:
		return true
	default:
		return false
	}
}

type ConsensusClientSyncStatus struct {
	Data struct {
		HeadSlot     string `json:"head_slot"`
		SyncDistance string `json:"sync_distance"`
		IsSyncing    bool   `json:"is_syncing"`
	} `json:"data"`
}

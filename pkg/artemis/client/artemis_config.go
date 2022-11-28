package artemis_client

type ArtemisConfig struct {
	Protocol string
	Network  string
}

type ArtemisConfigs []ArtemisConfig

const (
	Mainnet  = "mainnet"
	Goerli   = "goerli"
	Ethereum = "ethereum"
)

var (
	ArtemisEthereumMainnet = NewArtemisConfig(Ethereum, Mainnet)
	ArtemisEthereumGoerli  = NewArtemisConfig(Ethereum, Goerli)
	GlobalArtemisConfigs   = []ArtemisConfig{ArtemisEthereumMainnet, ArtemisEthereumGoerli}
)

func NewArtemisConfig(protocol, network string) ArtemisConfig {
	cfg := ArtemisConfig{
		Protocol: protocol,
		Network:  network,
	}
	return cfg
}

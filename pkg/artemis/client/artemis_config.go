package artemis_client

import "path"

type ArtemisConfig struct {
	Protocol string
	Network  string
}

type ArtemisConfigs []*ArtemisConfig

const (
	Mainnet   = "mainnet"
	Goerli    = "goerli"
	Ephemeral = "ephemeral"
	Ethereum  = "ethereum"
)

var (
	ArtemisEthereumMainnet   = NewArtemisConfig(Ethereum, Mainnet)
	ArtemisEthereumGoerli    = NewArtemisConfig(Ethereum, Goerli)
	ArtemisEthereumEphemeral = NewArtemisConfig(Ethereum, Ephemeral)
	GlobalArtemisConfigs     = []*ArtemisConfig{&ArtemisEthereumMainnet, &ArtemisEthereumGoerli, &ArtemisEthereumEphemeral}
)

func (a *ArtemisConfig) GetV1BetaBaseRoute() string {
	return path.Join("/v1beta", a.Protocol, a.Network)
}

func NewArtemisConfig(protocol, network string) ArtemisConfig {
	cfg := ArtemisConfig{
		Protocol: protocol,
		Network:  network,
	}
	return cfg
}

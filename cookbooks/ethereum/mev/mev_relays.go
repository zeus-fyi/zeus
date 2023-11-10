package ethereum_mev_cookbooks

import "strings"

// mainnet
const (
	flashbotsRelay   = "https://0xac6e77dfe25ecd6110b8e780608cce0dab71fdd5ebea22a16c0205200f2f8e2e3ad3b71d3499c54ad14d6c21b41a37ae@boost-relay.flashbots.net"
	blocknativeRelay = "https://0x9000009807ed12c1f08bf4e81c6da3ba8e3fc3d953898ce0102433094e5f22f21102ec057841fcb81978ed1ea0fa8246@builder-relay-mainnet.blocknative.com"
	edenNetworkRelay = "https://0xb3ee7afcf27f1f1259ac1787876318c6584ee353097a50ed84f51a1f21a323b3736f271a895c7ce918c038e4265918be@relay.edennetwork.io"
	ultraSoundRelay  = "https://0xa1559ace749633b997cb3fdacffb890aeebdb0f5a3b6aaa7eeeaf1a38af0a8fe88b9e4b1f61f236d2e64d95733327a62@relay.ultrasound.money"
)

// goerli
const (
	flashbotsGoerliRelay   = "https://0xafa4c6985aa049fb79dd37010438cfebeb0f2bd42b115b89dd678dab0670c1de38da0c4e9138c9290a398ecd9a0b3110@builder-relay-goerli.flashbots.net"
	blocknativeGoerliRelay = "https://0x8f7b17a74569b7a57e9bdafd2e159380759f5dc3ccbd4bf600414147e8c4e1dc6ebada83c0139ac15850eb6c975e82d0@builder-relay-goerli.blocknative.com"
	edenNetworkGoerliRelay = "https://0xb1d229d9c21298a87846c7022ebeef277dfc321fe674fa45312e20b5b6c400bfde9383f801848d7837ed5fc449083a12@relay-goerli.edennetwork.io"
	ultraSoundGoerliRelay  = "https://0xb1559beef7b5ba3127485bbbb090362d9f497ba64e177ee2c8e7db74746306efad687f2cf8574e38d70067d40ef136dc@relay-stag.ultrasound.money"
)

type RelaysEnabled struct {
	Flashbots   bool `yaml:"flashbots"`
	Blocknative bool `yaml:"blocknative"`
	EdenNetwork bool `yaml:"eden_network"`
	UltraSound  bool `yaml:"ultra_sound"`
}

func (r *RelaysEnabled) GetRelays(network string) []string {
	switch strings.ToLower(network) {
	case "mainnet":
		return r.GetMainnetRelays()
	case "goerli":
		return r.GetGoerliRelays()
	}
	return nil
}

func (r *RelaysEnabled) GetMainnetRelays() []string {
	var relays []string
	if r.Flashbots {
		relays = append(relays, flashbotsRelay)
	}
	if r.Blocknative {
		relays = append(relays, blocknativeRelay)
	}
	if r.EdenNetwork {
		relays = append(relays, edenNetworkRelay)
	}
	if r.UltraSound {
		relays = append(relays, ultraSoundRelay)
	}
	return relays
}

func (r *RelaysEnabled) GetGoerliRelays() []string {
	var relays []string
	if r.Flashbots {
		relays = append(relays, flashbotsGoerliRelay)
	}
	if r.Blocknative {
		relays = append(relays, blocknativeGoerliRelay)
	}
	if r.EdenNetwork {
		relays = append(relays, edenNetworkGoerliRelay)
	}
	if r.UltraSound {
		relays = append(relays, ultraSoundGoerliRelay)
	}
	return relays
}

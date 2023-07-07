package artemis_client

import (
	zeus_ecdsa "github.com/zeus-fyi/zeus/pkg/aegis/crypto/ecdsa"
	resty_base "github.com/zeus-fyi/zeus/zeus/z_client/base"
)

type ArtemisClient struct {
	zeus_ecdsa.Account
	resty_base.Resty
	ArtemisConfigs
}

func NewArtemisClient(baseURL, bearer string) ArtemisClient {
	a := ArtemisClient{}
	a.Resty = resty_base.GetBaseRestyClient(baseURL, bearer)
	a.ArtemisConfigs = GlobalArtemisConfigs
	return a
}

const ArtemisEndpoint = "https://artemis.zeus.fyi"

func NewDefaultArtemisClient(bearer string) ArtemisClient {
	return NewArtemisClient(ArtemisEndpoint, bearer)
}

const ArtemisLocalEndpoint = "http://localhost:9004"

func NewLocalArtemisClient(bearer string) ArtemisClient {
	return NewArtemisClient(ArtemisLocalEndpoint, bearer)
}

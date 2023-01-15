package hestia_client

import (
	zeus_client "github.com/zeus-fyi/zeus/pkg/zeus/client"
	resty_base "github.com/zeus-fyi/zeus/pkg/zeus/client/base"
)

type Hestia struct {
	zeus_client.ZeusClient
}

func NewHestia(baseURL, bearer string) Hestia {
	z := Hestia{}
	z.Resty = resty_base.GetBaseRestyTestClient(baseURL, bearer)
	return z
}

const HestiaEndpoint = "https://hestia.zeus.fyi"

func NewDefaultHestiaClient(bearer string) Hestia {
	return NewHestia(HestiaEndpoint, bearer)
}

const HestiaLocalEndpoint = "http://localhost:9002"

func NewLocalHestiaClient(bearer string) Hestia {
	return NewHestia(HestiaLocalEndpoint, bearer)
}

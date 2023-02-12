package hera_client

import (
	resty_base "github.com/zeus-fyi/zeus/pkg/zeus/client/base"
)

type HeraClient struct {
	resty_base.Resty
}

func NewHeraClient(baseURL, bearer string) HeraClient {
	a := HeraClient{}
	a.Resty = resty_base.GetBaseRestyClient(baseURL, bearer)
	return a
}

const HeraEndpoint = "https://hera.zeus.fyi"

func NewDefaultHeraClient(bearer string) HeraClient {
	return NewHeraClient(HeraEndpoint, bearer)
}

const HeraLocalEndpoint = "http://localhost:9008"

func NewLocalHeraClient(bearer string) HeraClient {
	return NewHeraClient(HeraLocalEndpoint, bearer)
}

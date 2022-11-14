package zeus_client

import (
	resty_base "github.com/zeus-fyi/zeus/pkg/zeus/client/base"
)

type ZeusClient struct {
	resty_base.Resty
}

func NewZeusClient(baseURL, bearer string) ZeusClient {
	z := ZeusClient{}
	z.Resty = resty_base.GetBaseRestyAresTestClient(baseURL, bearer)
	return z
}

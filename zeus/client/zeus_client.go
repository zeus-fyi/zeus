package zeus_client

import resty_base "github.com/zeus-fyi/zeus/zeus/client/base"

type ZeusClient struct {
	resty_base.Resty
}

func NewZeusClient(baseURL, bearer string) ZeusClient {
	z := ZeusClient{}
	z.Resty = resty_base.GetBaseRestyClient(baseURL, bearer)

	return z
}

const ZeusEndpoint = "https://api.zeus.fyi"

func NewDefaultZeusClient(bearer string) ZeusClient {
	return NewZeusClient(ZeusEndpoint, bearer)
}

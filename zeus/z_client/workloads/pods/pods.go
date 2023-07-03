package pods_client

import (
	zeus_client "github.com/zeus-fyi/zeus/zeus/z_client"
	resty_base "github.com/zeus-fyi/zeus/zeus/z_client/base"
)

type PodsClient struct {
	zeus_client.ZeusClient
}

func NewPodsClient(baseURL, bearer string) PodsClient {
	p := PodsClient{}
	p.Resty = resty_base.GetBaseRestyClient(baseURL, bearer)

	return p
}

func NewPodsClientFromZeusClient(z zeus_client.ZeusClient) PodsClient {
	p := PodsClient{
		ZeusClient: z,
	}
	return p
}

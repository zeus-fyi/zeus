package hercules_client

import (
	zeus_client "github.com/zeus-fyi/zeus/pkg/zeus/client"
	resty_base "github.com/zeus-fyi/zeus/pkg/zeus/client/base"
)

type RoutineRequest struct {
	ClientName string `json:"clientName"`
}

type RoutineResp struct {
	Status string `json:"status"`
}

type HerculesClient struct {
	zeus_client.ZeusClient
}

func NewHerculesClient(baseURL, bearer string) HerculesClient {
	z := HerculesClient{}
	z.Resty = resty_base.GetBaseRestyTestClient(baseURL, bearer)
	return z
}

const ZeusEndpoint = "https://api.zeus.fyi"

func NewDefaultHerculesClient(bearer string) HerculesClient {
	return NewHerculesClient(ZeusEndpoint, bearer)
}

const ZeusLocalEndpoint = "http://localhost:9003"

func NewLocalHerculesClient(bearer string) HerculesClient {
	return NewHerculesClient(ZeusLocalEndpoint, bearer)
}

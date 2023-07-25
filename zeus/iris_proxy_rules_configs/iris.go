package iris_proxy_rules_configs

import (
	resty_base "github.com/zeus-fyi/zeus/zeus/z_client/base"
)

const (
	IrisServiceRoute = "https://iris.zeus.fyi"
)

type Iris struct {
	resty_base.Resty
}

func NewIrisClient(bearerToken string) Iris {
	return Iris{
		resty_base.GetBaseRestyClient(IrisServiceRoute, bearerToken),
	}
}

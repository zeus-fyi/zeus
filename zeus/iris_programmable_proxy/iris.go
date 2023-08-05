package iris_programmable_proxy

import (
	resty_base "github.com/zeus-fyi/zeus/zeus/z_client/base"
)

const (
	IrisServiceRoute            = "https://iris.zeus.fyi"
	SelectedRouteResponseHeader = "X-Selected-Route"
	RouteGroupHeader            = "X-Route-Group"
)

type Iris struct {
	resty_base.Resty
}

func NewIrisClient(bearerToken string) Iris {
	return Iris{
		resty_base.GetBaseRestyClient(IrisServiceRoute, bearerToken),
	}
}

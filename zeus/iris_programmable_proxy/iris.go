package iris_programmable_proxy

import (
	resty_base "github.com/zeus-fyi/zeus/zeus/z_client/base"
)

const (
	IrisServiceRoute            = "https://iris.zeus.fyi"
	SelectedRouteResponseHeader = "X-Selected-Route"
	RouteGroupHeader            = "X-Route-Group"

	// AnvilHeader for serverless anvil tx fork/simulate servers, up to 10 minute runtime per invocation
	AnvilHeader = "X-Anvil-Session-Lock-ID"
)

type Iris struct {
	resty_base.Resty
}

func NewIrisClient(bearerToken string) Iris {
	return Iris{
		resty_base.GetBaseRestyClient(IrisServiceRoute, bearerToken),
	}
}

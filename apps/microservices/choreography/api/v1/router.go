package v1_choreography

import (
	"net/http"

	pods_client "github.com/zeus-fyi/zeus/zeus/z_client/workloads/pods"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_common_types"
)

var (
	PodsClient pods_client.PodsClient
	CloudCtxNs zeus_common_types.CloudCtxNs
)

func Routes(e *echo.Echo) *echo.Echo {
	e.GET("/health", Health)
	e.GET("/delete/pods", RestartPods)
	return e
}

func Health(c echo.Context) error {
	return c.String(http.StatusOK, "Healthy")
}

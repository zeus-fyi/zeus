package v1_choreography

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var (
	ZeusClient zeus_client.ZeusClient
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

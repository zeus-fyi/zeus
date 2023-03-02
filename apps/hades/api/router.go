package hades_api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	v1_hades_workloads "github.com/zeus-fyi/hades/api/v1"
	hades_core "github.com/zeus-fyi/zeus/pkg/hades/core"
)

func Routes(e *echo.Echo, hades hades_core.Hades) *echo.Echo {
	e.GET("/health", Health)

	eg := e.Group("/v1")
	eg.Use(middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		AuthScheme: "Bearer",
		Validator: func(token string, c echo.Context) (bool, error) {
			// Insert your own authentication logic here
			return true, nil
		},
	}))
	v1_hades_workloads.InitHadesV1Routes(eg, hades)
	return e
}

func Health(c echo.Context) error {
	return c.String(http.StatusOK, "Healthy")
}

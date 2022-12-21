package hercules_router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	v1_hercules "github.com/zeus-fyi/hercules/api/v1"
)

func Routes(e *echo.Echo) *echo.Echo {
	// Routes
	e.GET("/health", Health)

	eg := e.Group("/v1beta/internal")
	v1_hercules.CommonRoutes(eg)
	return e
}

func Health(c echo.Context) error {
	return c.String(http.StatusOK, "Healthy")
}

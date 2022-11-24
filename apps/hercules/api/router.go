package hercules_router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	v1_hercules "github.com/zeus-fyi/hercules/api/v1"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
)

func Routes(e *echo.Echo, p filepaths.Path) *echo.Echo {
	// Routes
	e.GET("/health", Health)

	eg := e.Group("/v1/internal")
	v1_hercules.CommonRoutes(eg, p)
	return e
}

func Health(c echo.Context) error {
	return c.String(http.StatusOK, "Healthy")
}

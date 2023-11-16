package v1_hera

import (
	"net/http"

	"github.com/labstack/echo/v4"
	ai_codegen "github.com/zeus-fyi/chatgpt/api/v1/codegen"
)

func Routes(e *echo.Echo) *echo.Echo {
	e.GET("/health", Health)
	InitV1Routes(e)
	return e
}

func InitV1Routes(e *echo.Echo) {
	eg := e.Group("/v1")

	ai_codegen.CodeGenRoutes(eg)
}

func Health(c echo.Context) error {
	return c.String(http.StatusOK, "Healthy")
}

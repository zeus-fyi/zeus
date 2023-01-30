package hercules_router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	v1_hercules "github.com/zeus-fyi/hercules/api/v1"
	"github.com/zeus-fyi/hercules/api/v1/common/aegis"
	hercules_ethereum "github.com/zeus-fyi/hercules/api/v1/common/ethereum"
)

func Routes(e *echo.Echo) *echo.Echo {
	// Routes
	e.GET("/health", Health)

	// to match eth keystore syntax styling
	e.POST("/eth/v1/bls/sign/verify", hercules_ethereum.EthereumBLSKeyVerificationHandler)
	e.POST("/eth/v1/keystores", aegis.ImportValidatorsHandler)

	eg := e.Group("/v1beta/internal")
	v1_hercules.CommonRoutes(eg)
	return e
}

func Health(c echo.Context) error {
	return c.String(http.StatusOK, "Healthy")
}

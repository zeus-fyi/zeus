package v1_hercules

import (
	"github.com/labstack/echo/v4"
	v1_common_routes "github.com/zeus-fyi/hercules/api/v1/common"
	hercules_chain_snapshots "github.com/zeus-fyi/hercules/api/v1/common/chain_snapshots"
	hercules_jwt_route "github.com/zeus-fyi/hercules/api/v1/common/jwt"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
)

func CommonRoutes(e *echo.Group, p filepaths.Path) *echo.Group {
	v1_common_routes.CommonManager.Path = p
	e.POST("/jwt/create", hercules_jwt_route.JwtHandler)
	e.POST("/jwt/replace", hercules_jwt_route.JwtReplaceHandler)
	e.POST("/snapshot/download", hercules_chain_snapshots.DownloadChainSnapshotHandler)
	return e
}

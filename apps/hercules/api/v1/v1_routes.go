package v1_hercules

import (
	"github.com/labstack/echo/v4"
	hercules_chain_snapshots "github.com/zeus-fyi/hercules/api/v1/common/chain_snapshots"
	host "github.com/zeus-fyi/hercules/api/v1/common/host_info"
	hercules_jwt_route "github.com/zeus-fyi/hercules/api/v1/common/jwt"
	hercules_routines "github.com/zeus-fyi/hercules/api/v1/common/routines"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
)

func CommonRoutes(e *echo.Group, p filepaths.Path) *echo.Group {
	e.POST("/jwt/create", hercules_jwt_route.JwtHandler)
	e.POST("/jwt/replace", hercules_jwt_route.JwtReplaceHandler)
	e.POST("/snapshot/download", hercules_chain_snapshots.DownloadChainSnapshotHandler)

	e.POST("/routines/suspend", hercules_routines.SuspendRoutineHandler)
	e.POST("/routines/start", hercules_routines.StartAppRoutineHandler)
	e.POST("/routines/resume", hercules_routines.ResumeProcessRoutineHandler)
	e.POST("/routines/kill", hercules_routines.KillProcessRoutineHandler)

	e.POST("/routines/disk/wipe", hercules_routines.WipeDiskHandler)

	e.GET("/host/disk", host.GetDiskStatsHandler)
	e.GET("/host/memory", host.GetMemStatsHandler)
	return e
}

package v1_hercules

import (
	"github.com/labstack/echo/v4"
	hercules_ethereum "github.com/zeus-fyi/hercules/api/v1/common/ethereum"
	host "github.com/zeus-fyi/hercules/api/v1/common/host_info"
	hercules_routines "github.com/zeus-fyi/hercules/api/v1/common/routines"
)

func CommonRoutes(e *echo.Group) *echo.Group {
	e.POST("/routines/suspend", hercules_routines.SuspendRoutineHandler)
	e.POST("/routines/start", hercules_routines.StartAppRoutineHandler)
	e.POST("/routines/resume", hercules_routines.ResumeProcessRoutineHandler)
	e.POST("/routines/kill", hercules_routines.KillProcessRoutineHandler)

	e.GET("/host/disk", host.GetDiskStatsHandler)
	e.GET("/host/memory", host.GetMemStatsHandler)

	e.GET("/ethereum/lighthouse/validator/auth", hercules_ethereum.LighthouseValidatorHandler)
	return e
}

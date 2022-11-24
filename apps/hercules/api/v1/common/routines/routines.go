package hercules_routines

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"github.com/zeus-fyi/hercules/pkg/routines"
)

type RoutineRequest struct {
	ClientName string `json:"clientName"`
}

type RoutineResp struct {
	Status string `json:"status"`
}

func (t *RoutineRequest) Kill(c echo.Context) error {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "func", "Kill")
	appName := routines.GetProcessName(t.ClientName)
	err := routines.KillProcessWithCtx(ctx, appName)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("Kill")
		return c.JSON(http.StatusInternalServerError, err)
	}
	resp := RoutineResp{Status: "killed"}
	return c.JSON(http.StatusOK, resp)
}

func (t *RoutineRequest) Suspend(c echo.Context) error {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "func", "Suspend")
	appName := routines.GetProcessName(t.ClientName)
	err := routines.SuspendProcessWithCtx(ctx, appName)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("SuspendApp")
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, nil)
}

func (t *RoutineRequest) Resume(c echo.Context) error {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "func", "Resume")
	appName := routines.GetProcessName(t.ClientName)
	err := routines.ResumeProcessWithCtx(ctx, appName)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("Resume: ResumeProcessWithCtx")
		return c.JSON(http.StatusInternalServerError, err)
	}
	resp := RoutineResp{Status: "resumed"}
	return c.JSON(http.StatusOK, resp)
}

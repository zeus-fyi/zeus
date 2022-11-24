package hercules_routines

import (
	"context"
	"net/http"
	"os/exec"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

func (t *RoutineRequest) Start(c echo.Context) error {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "func", "Start")

	cmd := exec.Command("sh", "-c", "/scripts/start.sh")
	err := cmd.Start()
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("Start: Start Script")
		return err
	}
	resp := RoutineResp{Status: "started"}
	return c.JSON(http.StatusOK, resp)
}

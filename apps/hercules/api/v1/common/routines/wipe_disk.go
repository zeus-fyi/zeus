package hercules_routines

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	v1_common_routes "github.com/zeus-fyi/hercules/api/v1/common"
	"github.com/zeus-fyi/hercules/pkg/routines"
)

func (t *RoutineRequest) WipeDisk(c echo.Context) error {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "func", "WipeDisk")
	dd := v1_common_routes.CommonManager.DirIn
	err := routines.WipeDisk(ctx, dd)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("WipeDisk")
		return c.JSON(http.StatusInternalServerError, err)
	}
	resp := RoutineResp{Status: fmt.Sprintf("wipeDisk %s", dd)}
	return c.JSON(http.StatusOK, resp)
}

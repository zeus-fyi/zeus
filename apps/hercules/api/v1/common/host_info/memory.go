package host

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"github.com/zeus-fyi/zeus/pkg/utils/host_info"
)

func (h *InfoRequest) GetMemStats(c echo.Context) error {
	ctx := context.Background()
	memInfo, err := host_info.GetVirtualMemoryStats(ctx)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("Start: GetMemStats")
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, memInfo)
}

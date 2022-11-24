package host

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	v1_common_routes "github.com/zeus-fyi/hercules/api/v1/common"
	"github.com/zeus-fyi/hercules/pkg/host_info"
)

type InfoRequest struct {
}

func (h *InfoRequest) GetDiskStats(c echo.Context) error {
	dd := v1_common_routes.CommonManager
	ctx := context.Background()
	diskInfo, err := host_info.GetDiskUsageStats(ctx, dd.DirIn)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("Start: GetDiskStats Script")
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, diskInfo)
}

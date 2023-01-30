package host

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	"github.com/zeus-fyi/zeus/pkg/utils/host_info"
)

type InfoRequest struct {
}

func (h *InfoRequest) GetDiskStats(c echo.Context) error {
	ctx := context.Background()
	p := filepaths.Path{DirIn: "/data"}
	diskInfo, err := host_info.GetDiskUsageStats(ctx, p.DirIn)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("Start: GetDiskStats Script")
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, diskInfo)
}

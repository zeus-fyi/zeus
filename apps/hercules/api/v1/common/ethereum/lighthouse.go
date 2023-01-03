package hercules_ethereum

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
)

var lhAuthTokenPath = filepaths.Path{DirIn: "/data/validators", FnIn: "api-token.txt"}

func LighthouseValidatorHandler(c echo.Context) error {
	request := new(LighthouseValidatorRequest)
	if err := c.Bind(request); err != nil {
		return err
	}
	return request.GetAuthToken(c)
}

type LighthouseValidatorRequest struct {
}

func (l *LighthouseValidatorRequest) GetAuthToken(c echo.Context) error {
	ctx := context.Background()
	a, err := lhAuthTokenPath.ReadFirstFileInPathWithFilter()
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("Start: GetDiskStats Script")
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, string(a))
}

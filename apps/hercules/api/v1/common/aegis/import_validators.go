package aegis

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	aegis_inmemdbs "github.com/zeus-fyi/zeus/pkg/aegis/inmemdbs"
)

type ImportValidatorsRequest struct {
	ValidatorSlice []aegis_inmemdbs.Validator
}

type RoutineResp struct {
	Status string `json:"status"`
}

func (t *ImportValidatorsRequest) ImportValidators(c echo.Context) error {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "func", "ImportValidators")
	aegis_inmemdbs.InsertValidatorsInMemDb(ctx, t.ValidatorSlice)
	resp := RoutineResp{Status: fmt.Sprintf("imported %d validators succesfully", len(t.ValidatorSlice))}
	return c.JSON(http.StatusOK, resp)
}

func ImportValidatorsHandler(c echo.Context) error {
	request := new(ImportValidatorsRequest)
	if err := c.Bind(request); err != nil {
		return err
	}
	return request.ImportValidators(c)
}

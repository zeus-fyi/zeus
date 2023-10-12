package aegis

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	aegis_inmemdbs "github.com/zeus-fyi/zeus/pkg/aegis/inmemdbs"
	signing_automation_ethereum "github.com/zeus-fyi/zeus/pkg/artemis/web3/signing_automation/ethereum"
)

type ImportValidatorsRequest struct {
	ImportKeystores Keystores
}

type Keystores []Keystore

type Keystore struct {
	KeystoreJSON map[string]interface{}
	Password     string
}

type RoutineResp struct {
	Status string `json:"status"`
}

func (t *ImportValidatorsRequest) ImportValidators(c echo.Context) error {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "func", "ImportValidators")
	vs := make([]aegis_inmemdbs.Validator, len(t.ImportKeystores))
	for i, ks := range t.ImportKeystores {
		acc, err := signing_automation_ethereum.DecryptKeystoreCipherIntoEthSignerBLS(ctx, ks.KeystoreJSON, ks.Password)
		if err != nil {
			log.Ctx(ctx).Err(err)
			return c.JSON(http.StatusBadRequest, err)
		}
		vs[i] = aegis_inmemdbs.NewValidator(acc)
	}
	aegis_inmemdbs.InsertValidatorsInMemDb(ctx, vs)
	resp := RoutineResp{Status: fmt.Sprintf("imported %d validators succesfully", len(vs))}
	return c.JSON(http.StatusOK, resp)
}

func ImportValidatorsHandler(c echo.Context) error {
	request := new(ImportValidatorsRequest)
	if err := c.Bind(request); err != nil {
		return err
	}
	return request.ImportValidators(c)
}

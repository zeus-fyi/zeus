package hercules_jwt_route

import (
	"net/http"

	"github.com/labstack/echo/v4"
	v1_common_routes "github.com/zeus-fyi/hercules/api/v1/common"
	init_jwt "github.com/zeus-fyi/zeus/pkg/aegis/jwt"
)

type TokenRequestJWT struct {
	JWT string
}

func (t *TokenRequestJWT) Create(c echo.Context) error {
	err := init_jwt.SetToken(v1_common_routes.CommonManager.Path, t.JWT)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusOK, nil)
}

func (t *TokenRequestJWT) ReplaceJWT(c echo.Context) error {
	err := init_jwt.ReplaceToken(v1_common_routes.CommonManager.Path, t.JWT)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusOK, nil)
}

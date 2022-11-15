package hercules_jwt_route

import "github.com/labstack/echo/v4"

func JwtHandler(c echo.Context) error {
	request := new(TokenRequestJWT)
	if err := c.Bind(request); err != nil {
		return err
	}
	return request.Create(c)
}

package host

import "github.com/labstack/echo/v4"

func GetDiskStatsHandler(c echo.Context) error {
	request := new(InfoRequest)
	if err := c.Bind(request); err != nil {
		return err
	}
	return request.GetDiskStats(c)
}

func GetMemStatsHandler(c echo.Context) error {
	request := new(InfoRequest)
	if err := c.Bind(request); err != nil {
		return err
	}
	return request.GetMemStats(c)
}

package hercules_routines

import "github.com/labstack/echo/v4"

func StartAppRoutineHandler(c echo.Context) error {
	request := new(RoutineRequest)
	if err := c.Bind(request); err != nil {
		return err
	}
	return request.Start(c)
}

func SuspendRoutineHandler(c echo.Context) error {
	request := new(RoutineRequest)
	if err := c.Bind(request); err != nil {
		return err
	}
	return request.Suspend(c)
}

func KillProcessRoutineHandler(c echo.Context) error {
	request := new(RoutineRequest)
	if err := c.Bind(request); err != nil {
		return err
	}
	return request.Kill(c)
}

func ResumeProcessRoutineHandler(c echo.Context) error {
	request := new(RoutineRequest)
	if err := c.Bind(request); err != nil {
		return err
	}
	return request.Resume(c)
}

func WipeDiskHandler(c echo.Context) error {
	request := new(RoutineRequest)
	if err := c.Bind(request); err != nil {
		return err
	}
	return request.WipeDisk(c)
}

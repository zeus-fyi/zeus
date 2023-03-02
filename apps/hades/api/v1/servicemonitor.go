package v1_hades_workloads

import (
	"context"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

func DeployServiceMonitorHandler(c echo.Context) error {
	ctx := context.Background()
	request := new(InternalDeploymentActionRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	if request.ServiceMonitor != nil {
		log.Debug().Interface("kns", request.CloudCtxNs).Msg("DeployServiceMonitorHandler: CreateServiceMonitorIfVersionLabelChangesOrDoesNotExist")
		_, err := Hades.CreateServiceMonitor(ctx, request.CloudCtxNs, request.ServiceMonitor, nil)
		if err != nil {
			log.Err(err).Msg("DeployServiceMonitorHandler")
			return c.JSON(http.StatusInternalServerError, err)
		}
	} else {
		err := errors.New("no servicemonitor workload was supplied")
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, nil)
}

func DestroyDeployServiceMonitorHandler(c echo.Context) error {
	ctx := context.Background()
	request := new(InternalDeploymentActionRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	if request.ServiceMonitor != nil {
		err := Hades.DeleteServiceMonitor(ctx, request.CloudCtxNs, request.ServiceMonitor.Name, nil)
		if err != nil {
			log.Err(err).Msg("DestroyDeployServiceMonitorHandler")
			return c.JSON(http.StatusInternalServerError, err)
		}
	} else {
		err := errors.New("no servicemonitor workload was supplied")
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, nil)
}

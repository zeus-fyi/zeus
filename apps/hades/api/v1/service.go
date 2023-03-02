package v1_hades_workloads

import (
	"context"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

func DeployServiceHandler(c echo.Context) error {
	ctx := context.Background()
	request := new(InternalDeploymentActionRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	if request.Service != nil {
		log.Debug().Interface("kns", request.CloudCtxNs).Msg("DeployServiceHandler: CreateServiceIfVersionLabelChangesOrDoesNotExist")
		_, err := Hades.CreateServiceWithKns(ctx, request.CloudCtxNs, request.Service, nil)
		if err != nil {
			log.Err(err).Msg("DeployServiceHandler")
			return c.JSON(http.StatusInternalServerError, err)
		}
	} else {
		err := errors.New("no service workload was supplied")
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, nil)
}

func DestroyDeployServiceHandler(c echo.Context) error {
	ctx := context.Background()
	request := new(InternalDeploymentActionRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	if request.Service != nil {
		err := Hades.DeleteServiceWithKns(ctx, request.CloudCtxNs, request.Service.Name, nil)
		if err != nil {
			log.Err(err).Msg("DestroyDeployServiceHandler")
			return c.JSON(http.StatusInternalServerError, err)
		}
	} else {
		err := errors.New("no service workload was supplied")
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, nil)
}

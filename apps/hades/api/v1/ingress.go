package v1_hades_workloads

import (
	"context"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

func DeployIngressHandler(c echo.Context) error {
	ctx := context.Background()
	request := new(InternalDeploymentActionRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	if request.Ingress != nil {
		log.Debug().Interface("kns", request.CloudCtxNs).Msg("DeployIngressHandler: CreateIngressIfVersionLabelChangesOrDoesNotExist")
		_, err := Hades.CreateIngressWithKns(ctx, request.CloudCtxNs, request.Ingress, nil)
		if err != nil {
			log.Err(err).Msg("DeployIngressHandler")
			return c.JSON(http.StatusInternalServerError, err)
		}
	} else {
		err := errors.New("no ingress workload was supplied")
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, nil)
}

func DestroyDeployIngressHandler(c echo.Context) error {
	ctx := context.Background()
	request := new(InternalDeploymentActionRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	if request.Ingress != nil {
		err := Hades.DeleteIngressWithKns(ctx, request.CloudCtxNs, request.Ingress.Name, nil)
		if err != nil {
			log.Err(err).Msg("DestroyDeployIngressHandler")
			return c.JSON(http.StatusInternalServerError, err)
		}
	} else {
		err := errors.New("no ingress workload was supplied")
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, nil)
}

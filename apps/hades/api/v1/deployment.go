package v1_hades_workloads

import (
	"context"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

func DeployDeploymentHandler(c echo.Context) error {
	ctx := context.Background()
	request := new(InternalDeploymentActionRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	if request.Deployment != nil {
		log.Debug().Interface("kns", request.CloudCtxNs).Msg("DeployDeploymentHandler: CreateDeploymentIfVersionLabelChangesOrDoesNotExist")
		_, err := Hades.CreateDeployment(ctx, request.CloudCtxNs, request.Deployment, nil)
		if err != nil {
			log.Err(err).Msg("DeployDeploymentHandler")
			return c.JSON(http.StatusInternalServerError, err)
		}
	} else {
		err := errors.New("no deployment workload was supplied")
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, nil)
}

func DestroyDeployDeploymentHandler(c echo.Context) error {
	ctx := context.Background()
	request := new(InternalDeploymentActionRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	if request.Deployment != nil {
		err := Hades.DeleteDeployment(ctx, request.CloudCtxNs, request.Deployment.Name, nil)
		if err != nil {
			log.Err(err).Msg("DestroyDeployDeploymentHandler")
			return c.JSON(http.StatusInternalServerError, err)
		}
	} else {
		err := errors.New("no deployment workload was supplied")
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, nil)
}

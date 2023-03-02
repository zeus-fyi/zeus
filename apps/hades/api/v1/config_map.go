package v1_hades_workloads

import (
	"context"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

func DeployConfigMapHandler(c echo.Context) error {
	ctx := context.Background()
	request := new(InternalDeploymentActionRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	if request.ConfigMap != nil {
		log.Debug().Interface("kns", request.CloudCtxNs).Msg("DeployConfigMapHandler: CreateConfigMapIfVersionLabelChangesOrDoesNotExist")
		_, err := Hades.CreateConfigMapWithKns(ctx, request.CloudCtxNs, request.ConfigMap, nil)
		if err != nil {
			log.Err(err).Msg("DeployConfigMapHandler")
			return c.JSON(http.StatusInternalServerError, err)
		}
	} else {
		err := errors.New("no configmap workload was supplied")
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, nil)
}

func DestroyDeployConfigMapHandler(c echo.Context) error {
	ctx := context.Background()
	request := new(InternalDeploymentActionRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	if request.ConfigMap != nil {
		err := Hades.DeleteConfigMapWithKns(ctx, request.CloudCtxNs, request.ConfigMap.Name, nil)
		if err != nil {
			log.Err(err).Msg("DestroyDeployConfigMapHandler")
			return c.JSON(http.StatusInternalServerError, err)
		}
	} else {
		err := errors.New("no configmap workload was supplied")
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, nil)
}

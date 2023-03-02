package v1_hades_workloads

import (
	"context"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

func DeployStatefulSetHandler(c echo.Context) error {
	ctx := context.Background()
	request := new(InternalDeploymentActionRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	if request.StatefulSet != nil {
		log.Debug().Interface("kns", request.CloudCtxNs).Msg("DeployStatefulSetHandler: CreateStatefulSetIfVersionLabelChangesOrDoesNotExist")
		_, err := Hades.CreateStatefulSet(ctx, request.CloudCtxNs, request.StatefulSet, nil)
		if err != nil {
			log.Err(err).Msg("DeployStatefulSetHandler")
			return c.JSON(http.StatusInternalServerError, err)
		}
	} else {
		err := errors.New("no statefulset workload was supplied")
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, nil)
}

func DestroyDeployStatefulSetHandler(c echo.Context) error {
	ctx := context.Background()
	request := new(InternalDeploymentActionRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	if request.StatefulSet != nil {
		err := Hades.DeleteStatefulSet(ctx, request.CloudCtxNs, request.StatefulSet.Name, nil)
		if err != nil {
			log.Err(err).Msg("DestroyDeployStatefulSetHandler")
			return c.JSON(http.StatusInternalServerError, err)
		}
	} else {
		err := errors.New("no statefulset workload was supplied")
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, nil)
}

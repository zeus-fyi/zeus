package v1_hades_workloads

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

func DeployNamespaceHandler(c echo.Context) error {
	ctx := context.Background()
	request := new(InternalDeploymentActionRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	log.Debug().Interface("kns", request.CloudCtxNs).Msg("DeployNamespaceHandler: CreateNamespaceIfDoesNotExist")
	_, err := Hades.CreateNamespaceIfDoesNotExist(ctx, request.CloudCtxNs)
	if err != nil {
		log.Err(err).Msg("DeployNamespaceHandler")
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, nil)
}
func DestroyDeployNamespaceHandler(c echo.Context) error {
	ctx := context.Background()
	request := new(InternalDeploymentActionRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	err := Hades.DeleteNamespace(ctx, request.CloudCtxNs)
	if err != nil {
		log.Err(err).Msg("DestroyDeployNamespaceHandler")
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, nil)
}

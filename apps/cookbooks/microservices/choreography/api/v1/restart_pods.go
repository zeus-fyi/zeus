package v1_choreography

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

func RestartPods(c echo.Context) error {
	par := zeus_pods_reqs.PodActionRequest{
		TopologyDeployRequest: zeus_req_types.TopologyDeployRequest{
			TopologyID: 0,
			CloudCtxNs: CloudCtxNs,
		},
	}
	resp, err := ZeusClient.DeletePods(context.Background(), par)
	if err != nil {
		log.Err(err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.String(http.StatusOK, string(resp))
}

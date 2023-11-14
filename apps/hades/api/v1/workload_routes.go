package v1_hades_workloads

import (
	"github.com/labstack/echo/v4"
	hades_core "github.com/zeus-fyi/zeus/pkg/hades/core"
	"github.com/zeus-fyi/zeus/zeus/workload_config_drivers/topology_workloads"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_common_types"
)

var Hades = hades_core.Hades{}

type InternalDeploymentActionRequest struct {
	zeus_common_types.CloudCtxNs
	topology_workloads.TopologyBaseInfraWorkload
}

func InitHadesV1Routes(e *echo.Group, hades hades_core.Hades) *echo.Group {
	Hades = hades
	e = HadesV1DeployRoutes(e)
	e = HadesV1DestroyDeployRoutes(e)
	return e
}

func HadesV1DeployRoutes(e *echo.Group) *echo.Group {
	e.POST("/deploy/namespace", DeployNamespaceHandler)
	e.POST("/deploy/deployment", DeployDeploymentHandler)
	e.POST("/deploy/statefulset", DeployStatefulSetHandler)
	e.POST("/deploy/configmap", DeployConfigMapHandler)
	e.POST("/deploy/service", DeployServiceHandler)
	e.POST("/deploy/ingress", DeployIngressHandler)
	e.POST("/deploy/servicemonitor", DeployServiceMonitorHandler)
	return e
}
func HadesV1DestroyDeployRoutes(e *echo.Group) *echo.Group {
	e.POST("/deploy/destroy/namespace", DestroyDeployNamespaceHandler)
	e.POST("/deploy/destroy/deployment", DestroyDeployDeploymentHandler)
	e.POST("/deploy/destroy/statefulset", DestroyDeployStatefulSetHandler)
	e.POST("/deploy/destroy/configmap", DestroyDeployConfigMapHandler)
	e.POST("/deploy/destroy/service", DestroyDeployServiceHandler)
	e.POST("/deploy/destroy/ingress", DestroyDeployIngressHandler)
	e.POST("/deploy/destroy/servicemonitor", DestroyDeployServiceMonitorHandler)
	return e
}

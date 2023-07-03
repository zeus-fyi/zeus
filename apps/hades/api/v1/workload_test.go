package v1_hades_workloads

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
	hades_api_test "github.com/zeus-fyi/hades/test"
	zeus_client "github.com/zeus-fyi/zeus/zeus/client"
	"github.com/zeus-fyi/zeus/zeus/client/zeus_common_types"
	"github.com/zeus-fyi/zeus/zeus/client/zeus_resp_types/topology_workloads"
)

type WorkloadsTestSuite struct {
	hades_api_test.HadesApiBaseTestSuite
}

// You can extend this example to test your own handlers.
func (t *WorkloadsTestSuite) TestNamespaceHandlers() {
	Hades = t.Hades
	t.Eg.POST("/deploy/namespace", DeployNamespaceHandler)
	t.Eg.POST("/deploy/destroy/namespace", DestroyDeployNamespaceHandler)

	start := make(chan struct{}, 1)
	go func() {
		close(start)
		_ = t.E.Start(":8888")
	}()

	client := zeus_client.NewZeusClient("http://localhost:8888", "")

	req := InternalDeploymentActionRequest{
		CloudCtxNs: zeus_common_types.CloudCtxNs{
			CloudProvider: "do",
			Region:        "nyc1",
			Context:       "do-nyc1-do-nyc1-zeus-demo",
			Namespace:     "demo",
			Env:           "test",
		},
		TopologyBaseInfraWorkload: topology_workloads.TopologyBaseInfraWorkload{},
	}
	client.PrintReqJson(req)
	resp, err := client.R().
		SetBody(req).
		Post("/v1/deploy/namespace")
	t.Require().NoError(err)
	t.Require().Equal(http.StatusOK, resp.StatusCode())

	client.PrintReqJson(req)
	resp, err = client.R().
		SetBody(req).
		Post("/v1/deploy/destroy/namespace")
	t.Require().NoError(err)
	t.Require().Equal(http.StatusOK, resp.StatusCode())
}

func TestWorkloadsTestSuite(t *testing.T) {
	suite.Run(t, new(WorkloadsTestSuite))
}

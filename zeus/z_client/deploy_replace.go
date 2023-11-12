package zeus_client

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	zeus_endpoints "github.com/zeus-fyi/zeus/zeus/z_client/endpoints"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_req_types"

	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_resp_types"
)

// DeployReplace will replace the topology at the desired cloud ctx ns only, it won't change the underlying topology
// definition, this is mostly useful for rapid development iteration and quick changes
func (z *ZeusClient) DeployReplace(ctx context.Context, p filepaths.Path, tar zeus_req_types.TopologyDeployRequest) (zeus_resp_types.TopologyDeployStatus, error) {
	respJson := zeus_resp_types.TopologyDeployStatus{}
	err := z.ZipK8sChartToPath(&p)
	if err != nil {
		return respJson, err
	}
	z.PrintReqJson(tar)
	resp, err := z.R().
		SetResult(&respJson).
		SetFormData(map[string]string{
			"topologyID":    fmt.Sprintf("%d", tar.TopologyID),
			"cloudProvider": tar.CloudProvider,
			"region":        tar.Region,
			"context":       tar.Context,
			"namespace":     tar.Namespace,
			"env":           tar.Env,
		}).
		SetFile("chart", p.FileOutPath()).
		Post(zeus_endpoints.ReplaceTopologyV1Path)

	if err != nil || resp.StatusCode() >= 400 {
		if err == nil {
			err = fmt.Errorf("non-OK status code: %d", resp.StatusCode())
		}
		log.Err(err).Msg("ZeusClient: DeployReplace")
		return respJson, err
	}
	z.PrintRespJson(resp.Body())
	return respJson, err
}

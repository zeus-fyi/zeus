package zeus_client

import (
	"context"
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/compression"
	zeus_cluster_config_drivers "github.com/zeus-fyi/zeus/pkg/zeus/cluster_config_drivers"

	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	zeus_endpoints "github.com/zeus-fyi/zeus/pkg/zeus/client/endpoints"
	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_req_types"
	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_resp_types"
)

func (z *ZeusClient) UploadChart(ctx context.Context, p filepaths.Path, tar zeus_req_types.TopologyCreateRequest) (zeus_resp_types.TopologyCreateResponse, error) {
	respJson := zeus_resp_types.TopologyCreateResponse{}
	err := z.ZipK8sChartToPath(&p)
	if err != nil {
		return respJson, err
	}
	z.PrintReqJson(tar)
	resp, err := z.R().
		SetResult(&respJson).
		SetFormData(map[string]string{
			"topologyName":      tar.TopologyName,
			"chartName":         tar.ChartName,
			"chartDescription":  tar.ChartDescription,
			"version":           tar.Version,
			"clusterClassName":  tar.ClusterClassName,
			"componentBaseName": tar.ComponentBaseName,
			"skeletonBaseName":  tar.SkeletonBaseName,
			"tag":               tar.Tag,
		}).
		SetFile("chart", p.FileOutPath()).
		Post(zeus_endpoints.InfraCreateV1Path)

	if err != nil || resp.StatusCode() != http.StatusOK {
		log.Ctx(ctx).Err(err).Msg("ZeusClient: UploadChart")
		if err == nil {
			err = fmt.Errorf("non-OK status code: %d", resp.StatusCode())
		}
		return respJson, err
	}
	z.PrintRespJson(resp.Body())
	return respJson, err
}

func (z *ZeusClient) UploadChartsFromClusterDefinition(ctx context.Context, cdef zeus_cluster_config_drivers.ClusterDefinition, print bool) ([]zeus_resp_types.TopologyCreateResponse, error) {
	sbs, err := cdef.GenerateSkeletonBaseCharts()
	if err != nil {
		log.Ctx(ctx).Err(err)
		return nil, err
	}
	responses := make([]zeus_resp_types.TopologyCreateResponse, len(sbs))
	for i, sb := range sbs {
		resp, rerr := z.UploadChart(ctx, sb.SkeletonBaseNameChartPath, sb.SkeletonBaseChart)
		if rerr != nil {
			log.Ctx(ctx).Err(err)
			return responses, rerr
		}
		if print {
			tar := zeus_req_types.TopologyRequest{TopologyID: resp.TopologyID}
			chartResp, cerr := z.ReadChart(ctx, tar)
			if cerr != nil {
				log.Ctx(ctx).Err(cerr)
			}
			cerr = chartResp.PrintWorkload(sb.SkeletonBaseNameChartPath)
			if cerr != nil {
				log.Ctx(ctx).Err(cerr)
			}
		}
		responses[i] = resp
	}
	return responses, err
}

func (z *ZeusClient) ZipK8sChartToPath(p *filepaths.Path) error {
	comp := compression.NewCompression()
	err := comp.GzipCompressDir(p)
	if err != nil {
		log.Err(err).Interface("path", p).Msg("ZeusClient: ZipK8sChartToPath")
		return err
	}
	return err
}

package api_calls

import (
	"fmt"
	"time"

	"github.com/zeus-fyi/zeus/demos/api_calls/endpoints"
	"github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/compression"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	"github.com/zeus-fyi/zeus/test/configs"
)

var uploadPath = filepaths.Path{}

type TopologyCreateRequest struct {
	TopologyName     string `json:"topologyName"`
	ChartName        string `json:"chartName"`
	ChartDescription string `json:"chartDescription,omitempty"`
	Version          string `json:"version"`
}

type TopologyCreateResponse struct {
	ID int `json:"id"`
}

var comp = compression.NewCompression()

func CreateDemoChartApiCall() error {
	cfg := configs.InitLocalTestConfigs()
	tar := TopologyCreateRequest{
		TopologyName:     "demo topology",
		ChartName:        "demo chart",
		ChartDescription: "demo chart description",
		Version:          fmt.Sprintf("v0.0.%d", time.Now().Unix()),
	}
	PrintReqJson(tar)
	err := comp.GzipCompressDir(&uploadPath)
	if err != nil {
		return err
	}
	client := GetBaseRestyClient()
	resp, err := client.R().
		SetAuthToken(cfg.Bearer).
		SetFormData(map[string]string{
			"topologyName":     tar.TopologyName,
			"chartName":        tar.ChartName,
			"chartDescription": tar.ChartDescription,
			"version":          tar.Version,
		}).
		SetFile("chart", uploadPath.FileOutPath()).
		Post(endpoints.InfraCreateV1Path)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	PrintRespJson(resp.Body())
	return err
}

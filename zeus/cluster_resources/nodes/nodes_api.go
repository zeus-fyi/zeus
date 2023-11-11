package nodes

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	zeus_client "github.com/zeus-fyi/zeus/zeus/z_client"
)

const (
	searchNodesEndpoint = "/v1/search/nodes"
)

type NodeSearchParams struct {
	CloudProviderRegions map[string][]string `json:"cloudProviderRegions"`
	DiskType             string              `json:"diskType,omitempty"`
	ResourceMinMax       ResourceMinMax      `json:"resourceMinMax,omitempty"`
}

type ResourceMinMax struct {
	Max ResourceAggregate `json:"max"`
	Min ResourceAggregate `json:"min"`
}

type ResourceAggregate struct {
	MonthlyPrice float64 `json:"monthlyPrice,omitempty"`
	HourlyPrice  float64 `json:"hourlyPrice,omitempty"`
	MemRequests  string  `json:"memRequests"`
	CpuRequests  string  `json:"cpuRequests"`
}

type NodeSearchRequest struct {
	NodeSearchParams `json:"nodeSearchParams"`
}

func GetNodes(ctx context.Context, z zeus_client.ZeusClient, searchParams NodeSearchParams) (any, error) {
	z.PrintReqJson(searchParams)
	respJson := NodesSlice{}
	req := NodeSearchRequest{
		NodeSearchParams: searchParams,
	}
	resp, err := z.R().
		SetResult(&respJson).
		SetBody(&req).
		Post(searchNodesEndpoint)

	if err != nil || resp.StatusCode() >= 400 {
		if err == nil {
			err = fmt.Errorf("non-OK status code: %d", resp.StatusCode())
		}
		log.Err(err).Msg("ZeusClient: GetNodes")
		return respJson, err
	}
	z.PrintRespJson(resp.Body())
	return respJson, err
}

func DeployNodes() {
	// TODO
}

func ScheduleNodes() {
	// TODO
}

func CreateNodeGroups() {
	// TODO
}

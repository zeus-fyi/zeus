package hercules_client

import (
	"context"
	"errors"
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	hercules_endpoints "github.com/zeus-fyi/zeus/pkg/hercules/client/endpoints"
)

func (a *HerculesClient) GetHostDiskInfo(ctx context.Context) (*disk.UsageStat, error) {
	respJson := &disk.UsageStat{}
	resp, err := a.R().
		SetResult(respJson).
		Get(hercules_endpoints.InternalHostDiskV1Path)
	if err != nil || resp.StatusCode() != http.StatusOK {
		log.Ctx(ctx).Err(err).Msg("HerculesClient: Kill")
		if resp.StatusCode() == http.StatusBadRequest {
			err = errors.New("bad request")
		}
		return respJson, err
	}
	a.PrintRespJson(resp.Body())
	return respJson, err
}

func (a *HerculesClient) GetHostMemInfo(ctx context.Context) (*mem.VirtualMemoryStat, error) {
	respJson := &mem.VirtualMemoryStat{}
	resp, err := a.R().
		SetResult(respJson).
		Get(hercules_endpoints.InternalHostMemV1Path)
	if err != nil || resp.StatusCode() != http.StatusOK {
		log.Ctx(ctx).Err(err).Msg("HerculesClient: GetHostMemInfo")
		if resp.StatusCode() == http.StatusBadRequest {
			err = errors.New("bad request")
		}
		return respJson, err
	}
	a.PrintRespJson(resp.Body())
	return respJson, err
}

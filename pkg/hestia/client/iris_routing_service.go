package hestia_client

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	hestia_endpoints "github.com/zeus-fyi/zeus/pkg/hestia/client/endpoints"
	hestia_req_types "github.com/zeus-fyi/zeus/pkg/hestia/client/req_types"
	hestia_resp_types "github.com/zeus-fyi/zeus/pkg/hestia/client/resp_types"
)

func (h *Hestia) CreateIrisRoutes(ctx context.Context, rr hestia_req_types.IrisOrgGroupRoutesRequest) error {
	h.PrintReqJson(rr)
	resp, err := h.R().
		SetBody(rr).
		Post(hestia_endpoints.IrisCreateRoutesPath)
	if err != nil || resp.StatusCode() >= 400 {
		if err == nil {
			err = fmt.Errorf("non-OK status code: %d", resp.StatusCode())
		}
		log.Ctx(ctx).Err(err).Msg("Hestia: CreateIrisRoutes")
		return err
	}
	h.PrintRespJson(resp.Body())
	return err
}

func (h *Hestia) ReadIrisRoutes(ctx context.Context) (hestia_resp_types.OrgRoutesResponse, error) {
	respJson := hestia_resp_types.OrgRoutesResponse{}
	resp, err := h.R().
		SetResult(&respJson).
		Get(hestia_endpoints.IrisReadRoutesPath)
	if err != nil || resp.StatusCode() >= 400 {
		if err == nil {
			err = fmt.Errorf("non-OK status code: %d", resp.StatusCode())
		}
		log.Ctx(ctx).Err(err).Msg("Hestia: ReadIrisRoutes")
		return respJson, err
	}
	h.PrintRespJson(resp.Body())
	return respJson, err
}

func (h *Hestia) DeleteIrisRoutes(ctx context.Context, rr hestia_req_types.IrisOrgGroupRoutesRequest) (hestia_resp_types.Response, error) {
	h.PrintReqJson(rr)
	respJson := hestia_resp_types.Response{}
	resp, err := h.R().
		SetBody(rr).
		SetResult(&respJson).
		Delete(hestia_endpoints.IrisDeleteRoutesPath)
	if err != nil || resp.StatusCode() >= 400 {
		if err == nil {
			err = fmt.Errorf("non-OK status code: %d", resp.StatusCode())
		}
		log.Ctx(ctx).Err(err).Msg("Hestia: IrisDeleteRoutesPath")
		return respJson, err
	}
	h.PrintRespJson(resp.Body())
	return respJson, err
}

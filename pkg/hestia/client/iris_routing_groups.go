package hestia_client

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	hestia_endpoints "github.com/zeus-fyi/zeus/pkg/hestia/client/endpoints"
	hestia_req_types "github.com/zeus-fyi/zeus/pkg/hestia/client/req_types"
	hestia_resp_types "github.com/zeus-fyi/zeus/pkg/hestia/client/resp_types"
)

func (h *Hestia) CreateIrisGroupRoutes(ctx context.Context, rr hestia_req_types.IrisOrgGroupRoutesRequest) error {
	resp, err := h.R().
		SetBody(rr).
		Post(hestia_endpoints.IrisCreateGroupRoutesPath)
	if err != nil || resp.StatusCode() >= 400 {
		if err == nil {
			err = fmt.Errorf("non-OK status code: %d", resp.StatusCode())
		}
		log.Ctx(ctx).Err(err).Msg("Hestia: IrisCreateRoutesPath")
		return err
	}
	h.PrintRespJson(resp.Body())
	return err
}

func (h *Hestia) ReadIrisGroupRoutes(ctx context.Context, groupName string) (hestia_resp_types.OrgGroupRoutesResponse, error) {
	respJson := hestia_resp_types.OrgGroupRoutesResponse{}
	resp, err := h.R().
		SetResult(&respJson).
		Get(hestia_endpoints.IrisReadGroupRoutesPath)
	if err != nil || resp.StatusCode() >= 400 {
		if err == nil {
			err = fmt.Errorf("non-OK status code: %d", resp.StatusCode())
		}
		log.Ctx(ctx).Err(err).Msg("Hestia: ReadIrisGroupRoutes")
		return respJson, err
	}
	h.PrintRespJson(resp.Body())
	return respJson, err
}

func (h *Hestia) ReadIrisGroupsRoutes(ctx context.Context) (hestia_resp_types.OrgGroupsRoutesResponse, error) {
	respJson := hestia_resp_types.OrgGroupsRoutesResponse{}
	resp, err := h.R().
		SetResult(&respJson).
		Get(hestia_endpoints.IrisReadGroupsRoutesPath)
	if err != nil || resp.StatusCode() >= 400 {
		if err == nil {
			err = fmt.Errorf("non-OK status code: %d", resp.StatusCode())
		}
		log.Ctx(ctx).Err(err).Msg("Hestia: ReadIrisGroupRoutes")
		return respJson, err
	}
	h.PrintRespJson(resp.Body())
	return respJson, err
}

func (h *Hestia) UpdateIrisGroupRoutes(ctx context.Context, rr hestia_req_types.IrisOrgGroupRoutesRequest) (hestia_resp_types.Response, error) {
	h.PrintReqJson(rr)
	respJson := hestia_resp_types.Response{}
	resp, err := h.R().
		SetBody(rr).
		SetResult(&respJson).
		Put(hestia_endpoints.IrisUpdateGroupRoutesPath)
	if err != nil || resp.StatusCode() >= 400 {
		if err == nil {
			err = fmt.Errorf("non-OK status code: %d", resp.StatusCode())
		}
		log.Ctx(ctx).Err(err).Msg("Hestia: UpdateIrisGroupRoutes")
		return respJson, err
	}
	h.PrintRespJson(resp.Body())
	return respJson, err
}

func (h *Hestia) DeleteIrisGroupRoutes(ctx context.Context, rr hestia_req_types.IrisOrgGroupRoutesRequest) (hestia_resp_types.Response, error) {
	h.PrintReqJson(rr)
	respJson := hestia_resp_types.Response{}
	resp, err := h.R().
		SetBody(rr).
		SetResult(&respJson).
		Delete(hestia_endpoints.IrisDeleteGroupRoutesPath)
	if err != nil || resp.StatusCode() >= 400 {
		if err == nil {
			err = fmt.Errorf("non-OK status code: %d", resp.StatusCode())
		}
		log.Ctx(ctx).Err(err).Msg("Hestia: DeleteIrisGroupRoutes")
		return respJson, err
	}
	h.PrintRespJson(resp.Body())
	return respJson, err
}

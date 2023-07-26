package hestia_client

import (
	"context"
	"errors"

	"github.com/rs/zerolog/log"
	hestia_endpoints "github.com/zeus-fyi/zeus/pkg/hestia/client/endpoints"
	hestia_resp_types "github.com/zeus-fyi/zeus/pkg/hestia/client/resp_types"
)

func (h *Hestia) CreateIrisGroupRoutes(ctx context.Context, rr any) (hestia_resp_types.Response, error) {
	h.PrintReqJson(rr)
	respJson := hestia_resp_types.Response{}
	resp, err := h.R().
		SetBody(rr).
		SetResult(&respJson).
		Post(hestia_endpoints.IrisCreateGroupRoutesPath)
	if err != nil || resp.StatusCode() >= 400 {
		if err == nil {
			err = errors.New("bad request")
		}
		log.Ctx(ctx).Err(err).Msg("Hestia: IrisCreateRoutesPath")
		return respJson, err
	}
	h.PrintRespJson(resp.Body())
	return respJson, err
}

func (h *Hestia) ReadIrisGroupRoutes(ctx context.Context, rr any) (hestia_resp_types.Response, error) {
	h.PrintReqJson(rr)
	respJson := hestia_resp_types.Response{}
	resp, err := h.R().
		SetBody(rr).
		SetResult(&respJson).
		Post(hestia_endpoints.IrisReadGroupRoutesPath)
	if err != nil || resp.StatusCode() >= 400 {
		if err == nil {
			err = errors.New("bad request")
		}
		log.Ctx(ctx).Err(err).Msg("Hestia: ReadIrisGroupRoutes")
		return respJson, err
	}
	h.PrintRespJson(resp.Body())
	return respJson, err
}

func (h *Hestia) UpdateIrisGroupRoutes(ctx context.Context, rr any) (hestia_resp_types.Response, error) {
	h.PrintReqJson(rr)
	respJson := hestia_resp_types.Response{}
	resp, err := h.R().
		SetBody(rr).
		SetResult(&respJson).
		Post(hestia_endpoints.IrisUpdateGroupRoutesPath)
	if err != nil || resp.StatusCode() >= 400 {
		if err == nil {
			err = errors.New("bad request")
		}
		log.Ctx(ctx).Err(err).Msg("Hestia: UpdateIrisGroupRoutes")
		return respJson, err

	}
	h.PrintRespJson(resp.Body())
	return respJson, err
}

func (h *Hestia) DeleteIrisGroupRoutes(ctx context.Context, rr any) (hestia_resp_types.Response, error) {
	h.PrintReqJson(rr)
	respJson := hestia_resp_types.Response{}
	resp, err := h.R().
		SetBody(rr).
		SetResult(&respJson).
		Post(hestia_endpoints.IrisDeleteGroupRoutesPath)
	if err != nil || resp.StatusCode() >= 400 {
		if err == nil {
			err = errors.New("bad request")
		}
		log.Ctx(ctx).Err(err).Msg("Hestia: DeleteIrisGroupRoutes")
		return respJson, err
	}
	h.PrintRespJson(resp.Body())
	return respJson, err
}

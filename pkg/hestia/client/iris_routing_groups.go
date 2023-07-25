package hestia_client

import (
	"context"
	"errors"
	"net/http"

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
	if err != nil || resp.StatusCode() != http.StatusAccepted {
		log.Ctx(ctx).Err(err).Msg("Hestia: IrisCreateRoutesPath")
		if resp.StatusCode() == http.StatusBadRequest {
			err = errors.New("bad request")
		}
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
	if err != nil || resp.StatusCode() != http.StatusAccepted {
		log.Ctx(ctx).Err(err).Msg("Hestia: IrisCreateRoutesPath")
		if resp.StatusCode() == http.StatusBadRequest {
			err = errors.New("bad request")
		}
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
	if err != nil || resp.StatusCode() != http.StatusAccepted {
		log.Ctx(ctx).Err(err).Msg("Hestia: UpdateIrisGroupRoutes")
		if resp.StatusCode() == http.StatusBadRequest {
			err = errors.New("bad request")
		}
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
	if err != nil || resp.StatusCode() != http.StatusAccepted {
		log.Ctx(ctx).Err(err).Msg("Hestia: DeleteIrisGroupRoutes")
		if resp.StatusCode() == http.StatusBadRequest {
			err = errors.New("bad request")
		}
		return respJson, err
	}
	h.PrintRespJson(resp.Body())
	return respJson, err
}

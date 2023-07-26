package hestia_client

import (
	"context"
	"fmt"

	hestia_resp_types "github.com/zeus-fyi/zeus/pkg/hestia/client/resp_types"

	"github.com/rs/zerolog/log"
	hestia_endpoints "github.com/zeus-fyi/zeus/pkg/hestia/client/endpoints"
	hestia_req_types "github.com/zeus-fyi/zeus/pkg/hestia/client/req_types"
)

type Response struct {
	Message string `json:"message"`
}

func (h *Hestia) ValidatorsServiceRequest(ctx context.Context, rr hestia_req_types.CreateValidatorServiceRequest) (hestia_resp_types.Response, error) {
	h.PrintReqJson(rr)

	respJson := hestia_resp_types.Response{}
	resp, err := h.R().
		SetBody(rr).
		SetResult(&respJson).
		Post(hestia_endpoints.EthereumValidatorsCreateServiceRequestV1Path)

	if err != nil || resp.StatusCode() >= 400 {
		if err == nil {
			err = fmt.Errorf("non-OK status code: %d", resp.StatusCode())
		}
		log.Ctx(ctx).Err(err).Msg("Hestia: ValidatorsServiceRequest")
		return respJson, err
	}

	h.PrintRespJson(resp.Body())
	return respJson, err
}

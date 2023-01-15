package hestia_client

import (
	"context"
	"errors"
	"net/http"

	"github.com/rs/zerolog/log"
	hestia_endpoints "github.com/zeus-fyi/zeus/pkg/hestia/client/endpoints"
	hestia_req_types "github.com/zeus-fyi/zeus/pkg/hestia/client/req_types"
)

func (h *Hestia) ValidatorsServiceRequest(ctx context.Context, rr hestia_req_types.CreateValidatorServiceRequest) (hestia_req_types.ValidatorServiceOrgGroupSlice, error) {
	h.PrintReqJson(rr)

	respJson := hestia_req_types.ValidatorServiceOrgGroupSlice{}
	resp, err := h.R().
		SetBody(rr).
		SetResult(&respJson).
		Post(hestia_endpoints.EthereumValidatorsCreateServiceRequestV1Path)

	if err != nil || resp.StatusCode() != http.StatusOK {
		log.Ctx(ctx).Err(err).Msg("Hestia: ValidatorBalances")
		if resp.StatusCode() == http.StatusBadRequest {
			err = errors.New("bad request")
		}
		return respJson, err
	}

	h.PrintRespJson(resp.Body())
	return respJson, err
}

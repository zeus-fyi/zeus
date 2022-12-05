package apollo_client

import (
	"context"
	"errors"
	"net/http"

	"github.com/rs/zerolog/log"
	apollo_endpoints "github.com/zeus-fyi/zeus/pkg/apollo/client/endpoints"
	apollo_req_types "github.com/zeus-fyi/zeus/pkg/apollo/client/req_types"
	apollo_resp_types "github.com/zeus-fyi/zeus/pkg/apollo/client/resp_types"
)

func (a *Apollo) ValidatorStatuses(ctx context.Context, rr apollo_req_types.ValidatorsRequest) (apollo_resp_types.Validators, error) {
	a.PrintReqJson(rr)

	respJson := apollo_resp_types.Validators{}
	resp, err := a.R().
		SetBody(rr).
		SetResult(&respJson).
		Post(apollo_endpoints.EthereumValidatorsV1Path)

	if err != nil || resp.StatusCode() != http.StatusOK {
		log.Ctx(ctx).Err(err).Msg("Apollo: ValidatorStatuses")
		if resp.StatusCode() == http.StatusBadRequest {
			err = errors.New("bad request")
		}
		return respJson, err
	}

	a.PrintRespJson(resp.Body())
	return respJson, err
}

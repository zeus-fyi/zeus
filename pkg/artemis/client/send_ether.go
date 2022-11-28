package artemis_client

import (
	"context"
	"errors"
	"net/http"
	"path"

	"github.com/rs/zerolog/log"
	artemis_req_types "github.com/zeus-fyi/zeus/pkg/artemis/client/req_types"
	artemis_resp_types "github.com/zeus-fyi/zeus/pkg/artemis/client/req_types/resp_types"
)

func (a *ArtemisClient) SendEther(ctx context.Context, rr artemis_req_types.SendEtherPayload, networkRoute ArtemisConfig) (artemis_resp_types.Response, error) {
	a.PrintReqJson(rr)
	respJson := artemis_resp_types.Response{}
	resp, err := a.R().
		SetResult(respJson).
		SetBody(rr).
		Post(getSendEtherEndpoint(networkRoute))

	if err != nil || resp.StatusCode() != http.StatusOK {
		log.Ctx(ctx).Err(err).Msg("ArtemisClient: SendEther")
		if resp.StatusCode() == http.StatusBadRequest {
			err = errors.New("bad request")
		}
		return respJson, err
	}
	a.PrintRespJson(resp.Body())
	return respJson, err
}

func getSendEtherEndpoint(networkRoute ArtemisConfig) string {
	return path.Join(networkRoute.GetV1BetaBaseRoute(), "/send")
}

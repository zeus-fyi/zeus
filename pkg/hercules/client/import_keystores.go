package hercules_client

import (
	"context"
	"errors"
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/zeus-fyi/hercules/api/v1/common/aegis"
	hercules_endpoints "github.com/zeus-fyi/zeus/pkg/hercules/client/endpoints"
)

func (a *HerculesClient) ImportKeystores(ctx context.Context, rr aegis.ImportValidatorsRequest) error {
	a.PrintReqJson(rr)
	resp, err := a.R().
		SetBody(rr).
		Post(hercules_endpoints.V1ImportKeystoresEthBLS)

	if err != nil || resp.StatusCode() != http.StatusOK {
		log.Ctx(ctx).Err(err).Msg("HerculesClient: ImportKeystores")
		if resp.StatusCode() == http.StatusBadRequest {
			err = errors.New("bad request")
		}
		return err
	}
	a.PrintRespJson(resp.Body())
	return err
}

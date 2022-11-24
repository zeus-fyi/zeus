package hercules_client

import (
	"context"
	"errors"
	"net/http"

	"github.com/rs/zerolog/log"
	hercules_endpoints "github.com/zeus-fyi/zeus/pkg/hercules/client/endpoints"
)

func (a *HerculesClient) Download(ctx context.Context) error {
	resp, err := a.R().
		Post(hercules_endpoints.InternalDownloadV1Path)

	if err != nil || resp.StatusCode() != http.StatusOK {
		log.Ctx(ctx).Err(err).Msg("HerculesClient: Download")
		if resp.StatusCode() == http.StatusBadRequest {
			err = errors.New("bad request")
		}
		return err
	}
	a.PrintRespJson(resp.Body())
	return err
}

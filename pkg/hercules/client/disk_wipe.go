package hercules_client

import (
	"context"
	"errors"
	"net/http"

	"github.com/rs/zerolog/log"

	hercules_endpoints "github.com/zeus-fyi/zeus/pkg/hercules/client/endpoints"
)

func (a *HerculesClient) DiskWipe(ctx context.Context, rr RoutineRequest) error {
	a.PrintReqJson(rr)
	resp, err := a.R().
		SetBody(rr).
		Post(hercules_endpoints.InternalDiskWipeV1Path)

	if err != nil || resp.StatusCode() != http.StatusOK {
		log.Ctx(ctx).Err(err).Msg("HerculesClient: DiskWipe")
		if resp.StatusCode() == http.StatusBadRequest {
			err = errors.New("bad request")
		}
		return err
	}

	a.PrintRespJson(resp.Body())
	return err
}

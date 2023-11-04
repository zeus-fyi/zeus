package iris_programmable_proxy

import (
	"fmt"

	"github.com/rs/zerolog/log"
	resty_base "github.com/zeus-fyi/zeus/zeus/z_client/base"
)

const (
	IrisServiceRoute            = "https://iris.zeus.fyi"
	SelectedRouteResponseHeader = "X-Selected-Route"
	RouteGroupHeader            = "X-Route-Group"
)

type Iris struct {
	resty_base.Resty
}

func NewIrisClient(bearerToken string) Iris {
	return Iris{
		resty_base.GetBaseRestyClient(IrisServiceRoute, bearerToken),
	}
}

func (i *Iris) EndServerlessEnvironment(sessionID string) error {
	resp, err := i.R().
		Delete(fmt.Sprintf("/v1/serverless/%s", sessionID))
	if err != nil {
		log.Err(err).Msg("EndServerlessSession failed")
		return err
	}
	if resp.StatusCode() >= 400 {
		err = fmt.Errorf("EndServerlessSession: status code: %d", resp.StatusCode())
		log.Err(err).Msg("EndServerlessSession failed")
		return err
	}
	return nil
}

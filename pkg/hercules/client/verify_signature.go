package hercules_client

import (
	"context"
	"errors"
	"net/http"

	"github.com/rs/zerolog/log"
	hercules_ethereum "github.com/zeus-fyi/hercules/api/v1/common/ethereum"
	aegis_inmemdbs "github.com/zeus-fyi/zeus/pkg/aegis/inmemdbs"
	hercules_endpoints "github.com/zeus-fyi/zeus/pkg/hercules/client/endpoints"
)

func (a *HerculesClient) VerifyEthSignatureBLS(ctx context.Context, rr hercules_ethereum.EthereumBLSKeyVerificationRequests) (aegis_inmemdbs.EthereumBLSKeySignatureResponses, error) {
	respJSON := aegis_inmemdbs.EthereumBLSKeySignatureResponses{}
	a.PrintReqJson(rr)
	resp, err := a.R().
		SetBody(rr).
		SetResult(&respJSON).
		Post(hercules_endpoints.V1VerifyEthBLS)

	if err != nil || resp.StatusCode() != http.StatusOK {
		log.Ctx(ctx).Err(err).Msg("HerculesClient: VerifyEthSignatureBLS")
		if resp.StatusCode() == http.StatusBadRequest {
			err = errors.New("bad request")
		}
		return respJSON, err
	}
	a.PrintRespJson(resp.Body())
	return respJSON, err
}

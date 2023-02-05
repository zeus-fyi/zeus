package hercules_ethereum

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	aegis_inmemdbs "github.com/zeus-fyi/zeus/pkg/aegis/inmemdbs"
)

func EthereumBLSKeyVerificationHandler(c echo.Context) error {
	request := new(EthereumBLSKeyVerificationRequests)
	if err := c.Bind(request); err != nil {
		return err
	}
	return request.SignVerificationMessages(c)
}

type EthereumBLSKeyVerificationRequests struct {
	SignatureRequests aegis_inmemdbs.EthereumBLSKeySignatureRequests
}

// SignVerificationMessages requires you to seed your keys into the aegis inmemfs, or replace with another import method
func (b *EthereumBLSKeyVerificationRequests) SignVerificationMessages(c echo.Context) error {
	ctx := context.Background()
	resp, err := aegis_inmemdbs.SignValidatorMessagesFromInMemDb(ctx, b.SignatureRequests)
	if err != nil {
		log.Ctx(ctx).Err(err)
		return c.JSON(http.StatusBadRequest, nil)
	}
	return c.JSON(http.StatusOK, resp)
}

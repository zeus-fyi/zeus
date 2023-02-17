package aegis_aws_secretmanager

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/rs/zerolog/log"
)

func (s *SecretsManagerAuthAWS) UpdateSecretAWS(ctx context.Context, si secretsmanager.UpdateSecretInput) error {
	_, err := s.UpdateSecret(ctx, &si)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("SecretsManagerAuthAWS: error updating secret")
		return err
	}
	return err
}

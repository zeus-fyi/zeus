package aws_secrets

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/rs/zerolog/log"
)

func (s *SecretsManagerAuthAWS) UpdateSecret(ctx context.Context, si secretsmanager.PutSecretValueInput) error {
	_, err := s.PutSecretValue(ctx, &si)
	if err != nil {
		log.Ctx(ctx).Error().Err(err).Msg("SecretsManagerAuthAWS: error updating secret")
		return err
	}
	return err
}

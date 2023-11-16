package aegis_aws_secretmanager

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/rs/zerolog/log"
)

func (s *SecretsManagerAuthAWS) CreateNewSecret(ctx context.Context, si secretsmanager.CreateSecretInput) error {
	_, err := s.CreateSecret(ctx, &si)
	if err != nil {
		log.Err(err).Msg("SecretsManagerAuthAWS: error creating secret")
		return err
	}
	return err
}

package aegis_aws_iam

import (
	"context"
	"github.com/rs/zerolog/log"
)

func (i *IAMClientAWS) CreateLambdaUser(ctx context.Context, upt UserPolicyTemplate) error {
	_, err := i.CreateUser(ctx, upt.UserName)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("CreateLambdaUser: error creating user")
		return err
	}
	return nil
}

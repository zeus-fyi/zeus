package aegis_aws_iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/iam"
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

func (i *IAMClientAWS) DoesUserExist(ctx context.Context, upt UserPolicyTemplate) bool {
	lu := &iam.GetUserInput{UserName: upt.UserName.UserName}
	u, err := i.GetUser(ctx, lu)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("GetLambdaUser: user not found, or other error")
		return false
	}
	if u != nil {
		log.Ctx(ctx).Info().Msg("GetLambdaUser: user found")
		return true
	}
	return false
}

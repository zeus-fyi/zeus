package aegis_aws_iam

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/rs/zerolog/log"
)

func (i *IAMClientAWS) CreateLambdaUserPolicy(ctx context.Context, upt UserPolicyTemplate, fnName string) error {
	createPolicyInput := upt.GetTemplateIAM(ctx, i.GetLambdaResourceARN(fnName))
	createPolicyOutput, err := i.CreatePolicy(ctx, createPolicyInput)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("CreateLambdaUserPolicy: error creating user policy")
		return err
	}
	attachUserPolicyInput := &iam.AttachUserPolicyInput{
		UserName:  upt.UserName.UserName,
		PolicyArn: createPolicyOutput.Policy.Arn,
	}
	_, err = i.AttachUserPolicy(ctx, attachUserPolicyInput)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("CreateLambdaUserPolicy: error attaching user policy")
		return err
	}
	return err
}

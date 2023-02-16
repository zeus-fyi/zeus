package aegis_aws_iam

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/rs/zerolog/log"
)

var (
	LambdaRoleName           = "lambda-role"
	LambdaRolePolicyDocument = `{
        "Version": "2012-10-17",
        "Statement": [
			{
				"Effect": "Allow",
				"Principal": {
					"Service": "lambda.amazonaws.com"
				},
				"Action": "sts:AssumeRole"
			}
        ]
    }`
)

func (i *IAMClientAWS) CreateLambdaRole(ctx context.Context, userName string) (*iam.CreateRoleOutput, error) {
	roleRes, err := i.CreateRole(ctx, &iam.CreateRoleInput{
		RoleName:                 aws.String(LambdaRoleName),
		AssumeRolePolicyDocument: aws.String(LambdaRolePolicyDocument),
	})
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("IAMClientAWS: CreateLambdaRole: error creating lambda role")
		return roleRes, err
	}
	return roleRes, nil
}

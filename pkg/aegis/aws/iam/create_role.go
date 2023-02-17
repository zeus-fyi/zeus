package aegis_aws_iam

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/rs/zerolog/log"
)

var (
	LambdaRoleName           = "lambdaRole"
	LambdaRolePolicyDocument = `{
        "Version": "2012-10-17",
        "Statement": [
			{
				"Effect": "Allow",
				"Principal": {
					"Service": [
						"lambda.amazonaws.com",
                    	"edgelambda.amazonaws.com"
					]
				},
				"Action": "sts:AssumeRole"
			}
        ]
    }`
)

func (i *IAMClientAWS) CreateInternalLambdaRole(ctx context.Context) (*iam.CreateRoleOutput, error) {
	roleRes, err := i.CreateRole(ctx, &iam.CreateRoleInput{
		RoleName:                 aws.String(LambdaRoleName),
		AssumeRolePolicyDocument: aws.String(LambdaRolePolicyDocument),
	})
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("IAMClientAWS: CreateInternalLambdaRole: error creating lambda role")
		return roleRes, err
	}
	return roleRes, nil
}

func (i *IAMClientAWS) GetInternalPolicyARN() string {
	return fmt.Sprintf("arn:aws:iam::%s:policy/%s", i.AccountNumber, internalLambdaPolicyTemplateName)
}
func (i *IAMClientAWS) AddInternalPolicyToLambdaRolePolicies(ctx context.Context) (*iam.AttachRolePolicyOutput, error) {
	roleRes, err := i.Client.AttachRolePolicy(ctx, &iam.AttachRolePolicyInput{
		PolicyArn: aws.String(i.GetInternalPolicyARN()),
		RoleName:  aws.String(LambdaRoleName),
	})
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("IAMClientAWS: CreateInternalLambdaRole: error creating lambda role")
		return roleRes, err
	}
	return roleRes, nil
}

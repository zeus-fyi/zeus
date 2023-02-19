package aegis_aws_iam

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/rs/zerolog/log"
)

var (
	EthereumSignerFunctionName       = "ethereumSignerBLS"
	internalLambdaUserName           = "internalLambdaUser"
	internalLambdaPolicyTemplateName = "internalLambdaPolicy"
	InternalLambdaUserAndPolicy      = UserPolicyTemplate{
		PolicyName: internalLambdaPolicyTemplateName,
		UserName: &iam.CreateUserInput{
			UserName: aws.String(internalLambdaUserName),
		}, Policy: nil,
	}
	ExternalLambdaUserName           = "externalLambdaUser"
	externalLambdaPolicyTemplateName = "externalLambdaPolicy"
	ExternalLambdaUserAndPolicy      = UserPolicyTemplate{
		PolicyName: externalLambdaPolicyTemplateName,
		UserName: &iam.CreateUserInput{
			UserName: aws.String(ExternalLambdaUserName),
		},
		Policy: nil,
	}
)

type UserPolicyTemplate struct {
	PolicyName string
	UserName   *iam.CreateUserInput
	Policy     *iam.CreatePolicyInput
}

// GetLambdaResourceARN uses the value from aws_lambda.EthereumSignerFunctionName
func (i *IAMClientAWS) GetLambdaResourceARN() string {
	return fmt.Sprintf("arn:aws:lambda:%s:%s:function:%s", i.Region, i.AccountNumber, EthereumSignerFunctionName)
}

func (p *UserPolicyTemplate) GetPolicyTemplateIAM(ctx context.Context, resource string) *iam.CreatePolicyInput {
	var iamPolicy string
	var createPolicyInput *iam.CreatePolicyInput
	switch p.PolicyName {
	case internalLambdaPolicyTemplateName:
		iamPolicy = `{
					  "Version": "2012-10-17",
					  "Statement": [
						{
							"Effect": "Allow",
							"Action": [
								"logs:CreateLogGroup",
								"logs:CreateLogStream",
								"logs:PutLogEvents"
							],
							"Resource": "arn:aws:logs:*:*:*"
						},
						{
							"Effect": "Allow",
							"Action": [
								"kms:Decrypt",
								"secretsmanager:GetSecretValue",
								"secretsmanager:DescribeSecret",
								"secretsmanager:ListSecrets"
							],
							"Resource": "*"
						}
					  ]
					}`
		createPolicyInput = &iam.CreatePolicyInput{
			PolicyName:     &internalLambdaPolicyTemplateName,
			PolicyDocument: &iamPolicy,
		}
	case externalLambdaPolicyTemplateName:
	}
	return createPolicyInput
}

func (i *IAMClientAWS) CreateNewLambdaUserPolicy(ctx context.Context, upt UserPolicyTemplate) (*iam.CreatePolicyOutput, error) {
	createPolicyInput := upt.GetPolicyTemplateIAM(ctx, i.GetLambdaResourceARN())
	createPolicyOutput, err := i.CreatePolicy(ctx, createPolicyInput)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("CreateLambdaUserPolicy: error creating user policy")
		return createPolicyOutput, err
	}
	return createPolicyOutput, err
}

func (i *IAMClientAWS) GetExternalPolicyARN() string {
	return "arn:aws:iam::aws:policy/service-role/AWSLambdaRole"
}
func (i *IAMClientAWS) AttachExternalLambdaUserPolicy(ctx context.Context) error {
	policy := &iam.AttachUserPolicyInput{
		PolicyArn: aws.String(i.GetExternalPolicyARN()),
		UserName:  aws.String(ExternalLambdaUserName),
	}
	_, err := i.AttachUserPolicy(ctx, policy)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("CreateLambdaUserPolicy: error creating user policy")
		return err
	}
	return err
}

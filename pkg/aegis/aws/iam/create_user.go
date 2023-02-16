package aegis_aws_iam

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/rs/zerolog/log"
)

var (
	internalLambdaUserTemplateName   = "internalLambdaUser"
	internalLambdaPolicyTemplateName = "internalLambdaPolicy"
	InternalLambdaUserAndPolicy      = UserPolicyTemplate{
		PolicyName: internalLambdaPolicyTemplateName,
		UserName: &iam.CreateUserInput{
			UserName:            aws.String(internalLambdaUserTemplateName),
			Path:                nil,
			PermissionsBoundary: nil,
			Tags:                nil,
		}, Policy: nil,
	}
	externalLambdaUserTemplateName   = "externalLambdaUser"
	externalLambdaPolicyTemplateName = "externalLambdaPolicy"
	ExternalLambdaUserAndPolicy      = UserPolicyTemplate{
		PolicyName: externalLambdaUserTemplateName,
		UserName: &iam.CreateUserInput{
			UserName:            aws.String(externalLambdaPolicyTemplateName),
			Path:                nil,
			PermissionsBoundary: nil,
			Tags:                nil,
		},
		Policy: nil,
	}
)

type UserPolicyTemplate struct {
	PolicyName string
	UserName   *iam.CreateUserInput
	Policy     *iam.CreatePolicyInput
}

func (i *IAMClientAWS) GetLambdaResourceARN(fnName string) string {
	return fmt.Sprintf("arn:aws:lambda:%s:%s:function:%s", i.Region, i.AccountNumber, fnName)
}

func (p *UserPolicyTemplate) GetTemplateIAM(ctx context.Context, resource string) *iam.CreatePolicyInput {
	var iamPolicy string
	var createPolicyInput *iam.CreatePolicyInput
	switch p.PolicyName {
	case internalLambdaPolicyTemplateName:
		iamPolicy = `
	{
	  "Version": "2012-10-17",
	  "Statement": [
       {
            "Effect": "Allow",
            "Principal": {
                "Service": "lambda.amazonaws.com"
            },
            "Action": "sts:AssumeRole"
        },
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
				"lambda:UpdateFunctionCode",
				"lambda:UpdateEventSourceMapping",
				"lambda:InvokeFunctionUrl",
				"lambda:InvokeFunction",
				"lambda:PublishLayerVersion",
				"lambda:PublishVersion",
                "lambda:InvokeFunction"
            ],
            "Resource": "arn:aws:lambda:*:*:*"
        }
	  ]
	}`
		createPolicyInput = &iam.CreatePolicyInput{
			PolicyName:     &internalLambdaPolicyTemplateName,
			PolicyDocument: &iamPolicy,
		}
	case externalLambdaPolicyTemplateName:
		iamPolicy = fmt.Sprintf(`
	{
	  "Version": "2012-10-17",
	  "Statement": [
	    {
	      "Effect": "Allow",
	      "Action": [
			"lambda:InvokeFunction",
	      ],
	      "Resource": "%s"
	    }
	  ]
	}
`, resource)
		createPolicyInput = &iam.CreatePolicyInput{
			PolicyName:     &externalLambdaPolicyTemplateName,
			PolicyDocument: &iamPolicy,
		}

	}
	return createPolicyInput
}

func (i *IAMClientAWS) CreateLambdaUser(ctx context.Context, upt UserPolicyTemplate) error {
	_, err := i.CreateUser(ctx, upt.UserName)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("CreateLambdaUser: error creating user")
		return err
	}
	return nil
}

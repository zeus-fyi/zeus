package aegis_aws_iam

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
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
		iamPolicy = fmt.Sprintf(`
	{
	  "Version": "2012-10-17",
	  "Statement": [
	    {
	      "Effect": "Allow",
	      "Action": [
	        "lambda:UpdateFunctionCode",
			"lambda:UpdateEventSourceMapping",
			"lambda:InvokeFunctionUrl",
			"lambda:InvokeFunction",
			"lambda:PublishLayerVersion",
			"lambda:PublishVersion"
	      ],
	      "Resource": "%s"
	    }
	  ]
	}
`, resource)
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

func (i *IAMClientAWS) CreateLambdaUser(ctx context.Context, upt UserPolicyTemplate, fnName string) error {
	_, err := i.CreateUser(ctx, upt.UserName)
	if err != nil {
		return err
	}
	createPolicyInput := upt.GetTemplateIAM(ctx, i.GetLambdaResourceARN(fnName))
	createPolicyOutput, err := i.CreatePolicy(ctx, createPolicyInput)
	if err != nil {
		return err
	}
	attachUserPolicyInput := &iam.AttachUserPolicyInput{
		UserName:  upt.UserName.UserName,
		PolicyArn: createPolicyOutput.Policy.Arn,
	}
	_, err = i.AttachUserPolicy(ctx, attachUserPolicyInput)
	if err != nil {
		return err
	}
	return nil
}

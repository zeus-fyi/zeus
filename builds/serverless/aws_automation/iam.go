package serverless_aws_automation

import (
	"context"
	"fmt"
	aegis_aws_auth "github.com/zeus-fyi/zeus/pkg/aegis/aws/auth"
	aegis_aws_iam "github.com/zeus-fyi/zeus/pkg/aegis/aws/iam"
)

func InternalUserRolePolicySetupForLambdaDeployment(ctx context.Context, auth aegis_aws_auth.AuthAWS) {
	CreateInternalLambdaUser(ctx, auth)
	CreateInternalLambdaRole(ctx, auth)
	CreateInternalLambdaPolicy(ctx, auth)
	AddInternalLambdaPoliciesToRole(ctx, auth)
}

func ExternalUserRolePolicySetupForLambdaDeployment(ctx context.Context, auth aegis_aws_auth.AuthAWS) {
	CreateExternalLambdaUser(ctx, auth)
	CreateExternalLambdaPolicy(ctx, auth)
}

func CreateExternalLambdaUserAccessKeys(ctx context.Context, auth aegis_aws_auth.AuthAWS) aegis_aws_auth.AuthAWS {
	fmt.Println("INFO: creating access keys for external lambda invocation with username ", aegis_aws_iam.ExternalLambdaUserName)
	iamClient, err := aegis_aws_iam.InitIAMClient(ctx, auth)
	if err != nil {
		panic(err)
	}
	keys, err := iamClient.CreateUserAccessKeys(ctx, aegis_aws_iam.ExternalLambdaUserName)
	if err != nil {
		panic(err)
	}
	return keys
}

func CreateExternalLambdaUser(ctx context.Context, auth aegis_aws_auth.AuthAWS) {
	fmt.Println("INFO: creating iam user for external lambda invocation with username ", aegis_aws_iam.ExternalLambdaUserAndPolicy.UserName)
	iamClient, err := aegis_aws_iam.InitIAMClient(ctx, auth)
	if err != nil {
		panic(err)
	}
	err = iamClient.CreateLambdaUser(ctx, aegis_aws_iam.ExternalLambdaUserAndPolicy)
	if err != nil {
		panic(err)
	}
}

func CreateExternalLambdaPolicy(ctx context.Context, auth aegis_aws_auth.AuthAWS) {
	fmt.Println("INFO: creating policy for external lambda invocation with policy name ", aegis_aws_iam.ExternalLambdaUserAndPolicy.PolicyName)
	iamClient, err := aegis_aws_iam.InitIAMClient(ctx, auth)
	if err != nil {
		panic(err)
	}
	_, err = iamClient.CreateNewLambdaUserPolicy(ctx, aegis_aws_iam.InternalLambdaUserAndPolicy)
	if err != nil {
		panic(err)
	}
}

func CreateInternalLambdaUser(ctx context.Context, auth aegis_aws_auth.AuthAWS) {
	fmt.Println("INFO: creating iam user for lambda deployment with username ", aegis_aws_iam.InternalLambdaUserAndPolicy.UserName)
	iamClient, err := aegis_aws_iam.InitIAMClient(ctx, auth)
	if err != nil {
		panic(err)
	}
	err = iamClient.CreateLambdaUser(ctx, aegis_aws_iam.InternalLambdaUserAndPolicy)
	if err != nil {
		panic(err)
	}
}

func CreateInternalLambdaRole(ctx context.Context, auth aegis_aws_auth.AuthAWS) {
	fmt.Println("INFO: creating role for lambda deployment with role name ", aegis_aws_iam.LambdaRoleName)
	iamClient, err := aegis_aws_iam.InitIAMClient(ctx, auth)
	if err != nil {
		panic(err)
	}
	_, err = iamClient.CreateInternalLambdaRole(ctx)
	if err != nil {
		panic(err)
	}
}

func CreateInternalLambdaPolicy(ctx context.Context, auth aegis_aws_auth.AuthAWS) {
	fmt.Println("INFO: creating policy for lambda deployment with policy name ", aegis_aws_iam.InternalLambdaUserAndPolicy.PolicyName)
	iamClient, err := aegis_aws_iam.InitIAMClient(ctx, auth)
	if err != nil {
		panic(err)
	}
	_, err = iamClient.CreateNewLambdaUserPolicy(ctx, aegis_aws_iam.InternalLambdaUserAndPolicy)
	if err != nil {
		panic(err)
	}
}

func AddInternalLambdaPoliciesToRole(ctx context.Context, auth aegis_aws_auth.AuthAWS) {
	fmt.Println("INFO: adding policy to role for lambda deployment")
	iamClient, err := aegis_aws_iam.InitIAMClient(ctx, auth)
	if err != nil {
		panic(err)
	}
	_, err = iamClient.AddInternalPolicyToLambdaRolePolicies(ctx)
	if err != nil {
		panic(err)
	}
}

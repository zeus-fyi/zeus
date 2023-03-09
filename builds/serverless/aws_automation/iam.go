package serverless_aws_automation

import (
	"context"
	"fmt"
	"strings"

	aegis_aws_auth "github.com/zeus-fyi/zeus/pkg/aegis/aws/auth"
	aegis_aws_iam "github.com/zeus-fyi/zeus/pkg/aegis/aws/iam"
)

func InternalUserRolePolicySetupForLambdaDeployment(ctx context.Context, auth aegis_aws_auth.AuthAWS) error {
	err := CreateInternalLambdaUser(ctx, auth)
	if err != nil {
		return err
	}
	err = CreateInternalLambdaRole(ctx, auth)
	if err != nil {
		return err
	}
	err = CreateInternalLambdaPolicy(ctx, auth)
	if err != nil {
		return err
	}
	err = AddInternalLambdaPoliciesToRole(ctx, auth)
	if err != nil {
		return err
	}
	return err
}

func ExternalUserRolePolicySetupForLambdaDeployment(ctx context.Context, auth aegis_aws_auth.AuthAWS) error {
	err := CreateExternalLambdaUser(ctx, auth)
	if err != nil {
		return err
	}
	err = AddExternalLambdaPolicyToUser(ctx, auth)
	return err
}

func CreateExternalLambdaUserAccessKeys(ctx context.Context, auth aegis_aws_auth.AuthAWS) (aegis_aws_auth.AuthAWS, error) {
	fmt.Println("INFO: creating access keys for external lambda invocation with username ", aegis_aws_iam.ExternalLambdaUserName)
	iamClient, err := aegis_aws_iam.InitIAMClient(ctx, auth)
	if err != nil {
		return aegis_aws_auth.AuthAWS{}, err
	}
	keys, err := iamClient.CreateUserAccessKeys(ctx, aegis_aws_iam.ExternalLambdaUserName)
	if err != nil {
		return aegis_aws_auth.AuthAWS{}, err
	}
	return keys, err
}

func CreateExternalLambdaUser(ctx context.Context, auth aegis_aws_auth.AuthAWS) error {
	fmt.Println("INFO: creating iam user for external lambda invocation with username ", aegis_aws_iam.ExternalLambdaUserAndPolicy.UserName)
	iamClient, err := aegis_aws_iam.InitIAMClient(ctx, auth)
	if err != nil {
		return err
	}
	err = iamClient.CreateLambdaUser(ctx, aegis_aws_iam.ExternalLambdaUserAndPolicy)
	if err != nil {
		if strings.Contains(err.Error(), "EntityAlreadyExists:") {
			fmt.Println("INFO: policy already exists, skipping creation")
			return nil
		}
		return err
	}
	return err
}

func AddExternalLambdaPolicyToUser(ctx context.Context, auth aegis_aws_auth.AuthAWS) error {
	iamClient, err := aegis_aws_iam.InitIAMClient(ctx, auth)
	if err != nil {
		return err
	}
	err = iamClient.AttachExternalLambdaUserPolicy(ctx)
	if err != nil {
		return err
	}
	return err
}

func CreateInternalLambdaUser(ctx context.Context, auth aegis_aws_auth.AuthAWS) error {
	fmt.Println("INFO: creating iam user for lambda deployment with username ", aegis_aws_iam.InternalLambdaUserAndPolicy.UserName)
	iamClient, err := aegis_aws_iam.InitIAMClient(ctx, auth)
	if err != nil {
		return err
	}
	userExists := iamClient.DoesUserExist(ctx, aegis_aws_iam.InternalLambdaUserAndPolicy)
	if userExists {
		fmt.Println("INFO: user already exists, skipping creation")
		return nil
	}
	err = iamClient.CreateLambdaUser(ctx, aegis_aws_iam.InternalLambdaUserAndPolicy)
	if err != nil {
		if strings.Contains(err.Error(), "EntityAlreadyExists:") {
			fmt.Println("INFO: policy already exists, skipping creation")
			return nil
		}
		return err
	}
	return err
}

func CreateInternalLambdaRole(ctx context.Context, auth aegis_aws_auth.AuthAWS) error {
	fmt.Println("INFO: creating role for lambda deployment with role name ", aegis_aws_iam.LambdaRoleName)
	iamClient, err := aegis_aws_iam.InitIAMClient(ctx, auth)
	if err != nil {
		return err
	}
	_, err = iamClient.CreateInternalLambdaRole(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "EntityAlreadyExists:") {
			fmt.Println("INFO: policy already exists, skipping creation")
			return nil
		}
		return err
	}
	return err
}

func CreateInternalLambdaPolicy(ctx context.Context, auth aegis_aws_auth.AuthAWS) error {
	fmt.Println("INFO: creating policy for lambda deployment with policy name ", aegis_aws_iam.InternalLambdaUserAndPolicy.PolicyName)
	iamClient, err := aegis_aws_iam.InitIAMClient(ctx, auth)
	if err != nil {
		return err
	}
	_, err = iamClient.CreateNewLambdaUserPolicy(ctx, aegis_aws_iam.InternalLambdaUserAndPolicy)
	if err != nil {
		if strings.Contains(err.Error(), "EntityAlreadyExists:") {
			fmt.Println("INFO: policy already exists, skipping creation")
			return nil
		}
		return err
	}
	return err
}

func AddInternalLambdaPoliciesToRole(ctx context.Context, auth aegis_aws_auth.AuthAWS) error {
	fmt.Println("INFO: adding policy to role for lambda deployment")
	iamClient, err := aegis_aws_iam.InitIAMClient(ctx, auth)
	if err != nil {
		return err
	}
	_, err = iamClient.AddInternalPolicyToLambdaRolePolicies(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "EntityAlreadyExists:") {
			fmt.Println("INFO: policy already exists, skipping creation")
			return nil
		}
		return err
	}
	return err
}

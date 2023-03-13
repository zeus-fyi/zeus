package serverless_aws_automation

import (
	"context"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/rs/zerolog/log"
	aegis_aws_auth "github.com/zeus-fyi/zeus/pkg/aegis/aws/auth"
	aws_lambda "github.com/zeus-fyi/zeus/pkg/cloud/aws/lambda"
)

func CreateLambdaFunctionSecretsKeyGen(ctx context.Context, auth aegis_aws_auth.AuthAWS) (string, error) {
	fmt.Println("INFO: creating lambda function")
	lm, err := aws_lambda.InitLambdaClient(ctx, auth)
	if err != nil {
		return "", err
	}
	_, err = lm.CreateServerlessBlsSecretsKeyGenLambdaFn(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "Function already exist") {
			log.Ctx(ctx).Info().Msg("INFO: lambda function already exists, skipping creation")
		} else {
			return "", err
		}
	}
	fmt.Println("INFO: creating lambda function url")
	lfUrl, err := lm.MakeLambdaURL(ctx, aws_lambda.EthereumValidatorsSecretsGenFunctionName)
	if err != nil {
		if strings.Contains(err.Error(), " FunctionUrlConfig exists for this Lambda function") {
			log.Ctx(ctx).Info().Msg("INFO: lambda function url already exists, skipping creation")
			lfUrlCfg, lerr := lm.GetExternalLambdaSignerConfigURL(ctx)
			if lerr != nil {
				return "", lerr
			}
			lfUrl = &lambda.CreateFunctionUrlConfigOutput{FunctionUrl: lfUrlCfg.FunctionUrl}
		} else {
			return "", err
		}
	}
	_, err = lm.MakeLambdaFuncAuthIAM(ctx, aws_lambda.EthereumValidatorsSecretsGenFunctionName)
	if err != nil {
		if strings.Contains(err.Error(), "already exists") {
			log.Ctx(ctx).Info().Msg("INFO: lambda function iam auth config already exists, skipping creation")
		} else {
			return "", err
		}
	}
	if lfUrl.FunctionUrl == nil {
		panic("ERROR: lambda function url is nil")
	}
	return *lfUrl.FunctionUrl, err
}

func CreateLambdaFunctionEncryptedKeystoresZip(ctx context.Context, auth aegis_aws_auth.AuthAWS) (string, error) {
	fmt.Println("INFO: creating lambda function")
	lm, err := aws_lambda.InitLambdaClient(ctx, auth)
	if err != nil {
		return "", err
	}
	_, err = lm.CreateServerlessBlsEncryptedKeystoresZipLambdaFn(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "Function already exist") {
			log.Ctx(ctx).Info().Msg("INFO: lambda function already exists, skipping creation")
		} else {
			return "", err
		}
	}
	fmt.Println("INFO: creating lambda function url")
	lfUrl, err := lm.MakeLambdaURL(ctx, aws_lambda.EthereumValidatorsEncryptedSecretsZipGenFunctionName)
	if err != nil {
		if strings.Contains(err.Error(), " FunctionUrlConfig exists for this Lambda function") {
			log.Ctx(ctx).Info().Msg("INFO: lambda function url already exists, skipping creation")
			lfUrlCfg, lerr := lm.GetExternalLambdaSignerConfigURL(ctx)
			if lerr != nil {
				return "", lerr
			}
			lfUrl = &lambda.CreateFunctionUrlConfigOutput{FunctionUrl: lfUrlCfg.FunctionUrl}
		} else {
			return "", err
		}
	}
	_, err = lm.MakeLambdaFuncAuthIAM(ctx, aws_lambda.EthereumValidatorsEncryptedSecretsZipGenFunctionName)
	if err != nil {
		if strings.Contains(err.Error(), "already exists") {
			log.Ctx(ctx).Info().Msg("INFO: lambda function iam auth config already exists, skipping creation")
		} else {
			return "", err
		}
	}
	if lfUrl.FunctionUrl == nil {
		panic("ERROR: lambda function url is nil")
	}
	return *lfUrl.FunctionUrl, err
}

func CreateLambdaFunctionDepositGen(ctx context.Context, auth aegis_aws_auth.AuthAWS) (string, error) {
	fmt.Println("INFO: creating lambda function")
	lm, err := aws_lambda.InitLambdaClient(ctx, auth)
	if err != nil {
		return "", err
	}
	_, err = lm.CreateServerlessValidatorDepositsGenLambdaFn(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "Function already exist") {
			log.Ctx(ctx).Info().Msg("INFO: lambda function already exists, skipping creation")
		} else {
			return "", err
		}
	}
	fmt.Println("INFO: creating lambda function url")
	lfUrl, err := lm.MakeLambdaURL(ctx, aws_lambda.EthereumCreateValidatorsDepositsFunctionName)
	if err != nil {
		if strings.Contains(err.Error(), " FunctionUrlConfig exists for this Lambda function") {
			log.Ctx(ctx).Info().Msg("INFO: lambda function url already exists, skipping creation")
			lfUrlCfg, lerr := lm.GetExternalLambdaSignerConfigURL(ctx)
			if lerr != nil {
				return "", lerr
			}
			lfUrl = &lambda.CreateFunctionUrlConfigOutput{FunctionUrl: lfUrlCfg.FunctionUrl}
		} else {
			return "", err
		}
	}
	_, err = lm.MakeLambdaFuncAuthIAM(ctx, aws_lambda.EthereumCreateValidatorsDepositsFunctionName)
	if err != nil {
		if strings.Contains(err.Error(), "already exists") {
			log.Ctx(ctx).Info().Msg("INFO: lambda function iam auth config already exists, skipping creation")
		} else {
			return "", err
		}
	}
	if lfUrl.FunctionUrl == nil {
		panic("ERROR: lambda function url is nil")
	}
	return *lfUrl.FunctionUrl, err
}

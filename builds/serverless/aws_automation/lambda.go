package serverless_aws_automation

import (
	"context"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/rs/zerolog/log"
	aegis_aws_auth "github.com/zeus-fyi/zeus/pkg/aegis/aws/auth"
	aws_lambda "github.com/zeus-fyi/zeus/pkg/cloud/aws/lambda"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
)

func CreateLambdaFunction(ctx context.Context, auth aegis_aws_auth.AuthAWS, functionName, keystoresLayerName string) (string, error) {
	fmt.Println("INFO: creating lambda function")
	lm, err := aws_lambda.InitLambdaClient(ctx, auth)
	if err != nil {
		return "", err
	}
	_, err = lm.CreateServerlessBLSLambdaFn(ctx, functionName, keystoresLayerName)
	if err != nil {
		if strings.Contains(err.Error(), "Function already exist") {
			log.Ctx(ctx).Info().Msg("INFO: lambda function already exists, skipping creation")
		} else {
			return "", err
		}
	}
	fmt.Println("INFO: creating lambda function url")
	lfUrl, err := lm.MakeLambdaURL(ctx, functionName)
	if err != nil {
		if strings.Contains(err.Error(), " FunctionUrlConfig exists for this Lambda function") {
			log.Ctx(ctx).Info().Msg("INFO: lambda function url already exists, skipping creation")
			lfUrlCfg, lerr := lm.GetLambdaConfigURL(ctx, functionName)
			if lerr != nil {
				return "", lerr
			}
			lfUrl = &lambda.CreateFunctionUrlConfigOutput{FunctionUrl: lfUrlCfg.FunctionUrl}
		} else {
			return "", err
		}
	}
	_, err = lm.MakeLambdaFuncAuthIAM(ctx, functionName)
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

func CreateLambdaFunctionKeystoresLayer(ctx context.Context, auth aegis_aws_auth.AuthAWS, p filepaths.Path, keystoresLayerName string) error {
	fmt.Println("INFO: creating lambda function keystores layer")
	lm, err := aws_lambda.InitLambdaClient(ctx, auth)
	if err != nil {
		return err
	}
	p.FnIn = "keystores.zip"
	keystoresZipBinary := p.ReadFileInPath()
	_, err = lm.CreateServerlessBLSLambdaFnKeystoreLayer(ctx, keystoresLayerName, keystoresZipBinary)
	if err != nil {
		return err
	}
	return err
}

func GetLambdaFunctionUrl(ctx context.Context, auth aegis_aws_auth.AuthAWS, functionName string) (string, error) {
	fmt.Println("INFO: getting lambda function url")
	lm, err := aws_lambda.InitLambdaClient(ctx, auth)
	if err != nil {
		return "", err
	}
	lcfg, err := lm.GetLambdaConfigURL(ctx, functionName)
	if err != nil {
		return "", err
	}
	return *lcfg.FunctionUrl, err
}

func UpdateLambdaFunctionKeystoresLayer(ctx context.Context, auth aegis_aws_auth.AuthAWS, functionName, keystoresLayerName string) error {
	fmt.Println("INFO: updating lambda function keystores layer to latest")
	lm, err := aws_lambda.InitLambdaClient(ctx, auth)
	if err != nil {
		return err
	}
	_, err = lm.UpdateServerlessBLSLambdaFnKeystoreLayer(ctx, functionName, keystoresLayerName)
	if err != nil {
		return err
	}
	return err
}

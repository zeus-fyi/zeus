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

func CreateLambdaFunction(ctx context.Context, auth aegis_aws_auth.AuthAWS) string {
	fmt.Println("INFO: creating lambda function")
	lm, err := aws_lambda.InitLambdaClient(ctx, auth)
	if err != nil {
		panic(err)
	}
	_, err = lm.CreateServerlessBLSLambdaFn(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "Function already exist") {
			log.Ctx(ctx).Info().Msg("INFO: lambda function already exists, skipping creation")
		} else {
			panic(err)
		}
	}
	fmt.Println("INFO: creating lambda function url")
	lfUrl, err := lm.MakeEthereumSignerURL(ctx)
	if err != nil {
		if strings.Contains(err.Error(), " FunctionUrlConfig exists for this Lambda function") {
			log.Ctx(ctx).Info().Msg("INFO: lambda function url already exists, skipping creation")
			lfUrlCfg, lerr := lm.GetExternalLambdaSignerConfigURL(ctx)
			if lerr != nil {
				panic(lerr)
			}
			lfUrl = &lambda.CreateFunctionUrlConfigOutput{FunctionUrl: lfUrlCfg.FunctionUrl}
		} else {
			panic(err)
		}
	}
	_, err = lm.MakeEthereumSignerFuncAuthIAM(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "already exists") {
			log.Ctx(ctx).Info().Msg("INFO: lambda function iam auth config already exists, skipping creation")
		} else {
			panic(err)
		}
	}
	if lfUrl.FunctionUrl == nil {
		panic("ERROR: lambda function url is nil")
	}
	return *lfUrl.FunctionUrl
}

func CreateLambdaFunctionKeystoresLayer(ctx context.Context, auth aegis_aws_auth.AuthAWS) {
	fmt.Println("INFO: creating lambda function keystores layer")
	lm, err := aws_lambda.InitLambdaClient(ctx, auth)
	if err != nil {
		panic(err)
	}
	_, err = lm.CreateServerlessBLSLambdaFnKeystoreLayer(ctx)
	if err != nil {
		panic(err)
	}
}

func GetLambdaFunctionUrl(ctx context.Context, auth aegis_aws_auth.AuthAWS) string {
	fmt.Println("INFO: getting lambda function url")
	lm, err := aws_lambda.InitLambdaClient(ctx, auth)
	if err != nil {
		panic(err)
	}
	lcfg, err := lm.GetExternalLambdaSignerConfigURL(ctx)
	if err != nil {
		panic(err)
	}

	return *lcfg.FunctionUrl
}
func UpdateLambdaFunctionKeystoresLayer(ctx context.Context, auth aegis_aws_auth.AuthAWS, version string) {
	fmt.Println("INFO: updating lambda function keystores layer")
	lm, err := aws_lambda.InitLambdaClient(ctx, auth)
	if err != nil {
		panic(err)
	}
	_, err = lm.UpdateServerlessBLSLambdaFnKeystoreLayer(ctx, version)
	if err != nil {
		panic(err)
	}
}

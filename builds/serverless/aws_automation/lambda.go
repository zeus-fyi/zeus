package serverless_aws_automation

import (
	"context"
	"fmt"
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
		panic(err)
	}
	fmt.Println("INFO: creating lambda function url")
	lfUrl, err := lm.MakeEthereumSignerURL(ctx)
	if err != nil {
		panic(err)
	}
	if lfUrl.FunctionUrl == nil {
		panic("ERROR: lambda function url is nil")
	}
	fmt.Println("INFO: lambda function url: ", lfUrl.FunctionUrl)
	fmt.Println("INFO: making lambda function url public")
	_, err = lm.MakeEthereumSignerFuncPublic(ctx)
	if err != nil {
		panic(err)
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

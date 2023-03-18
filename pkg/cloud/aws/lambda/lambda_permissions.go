package aws_lambda

import (
	"context"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/rs/zerolog/log"
)

const zeusCloudOrigin = "https://cloud.zeus.fyi"

// MakeEthereumSignerFuncPublic uses the EthereumSignerFunctionName to make the function public
func (l *LambdaClientAWS) MakeEthereumSignerFuncPublic(ctx context.Context) (*lambda.AddPermissionOutput, error) {
	input := &lambda.AddPermissionInput{
		Action:              aws.String("lambda:InvokeFunctionUrl"),
		FunctionName:        aws.String(EthereumSignerFunctionName),
		Principal:           aws.String("*"),
		StatementId:         aws.String("FunctionURLAllowPublicAccess"),
		FunctionUrlAuthType: types.FunctionUrlAuthTypeNone,
	}
	resp, err := l.AddPermission(ctx, input)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("LambdaClientAWS: MakeFuncPublic: error making function public")
		return resp, err
	}
	return resp, err
}

func (l *LambdaClientAWS) GetLambdaConfigURL(ctx context.Context, functionName string) (*lambda.GetFunctionUrlConfigOutput, error) {
	input := &lambda.GetFunctionUrlConfigInput{
		FunctionName: aws.String(functionName),
	}
	resp, err := l.GetFunctionUrlConfig(ctx, input)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("LambdaClientAWS: GetLambdaConfigURL: error getting function url config")
		return resp, err
	}
	return resp, err
}

func (l *LambdaClientAWS) MakeLambdaURL(ctx context.Context, lambdaName string) (*lambda.CreateFunctionUrlConfigOutput, error) {
	input := &lambda.CreateFunctionUrlConfigInput{
		AuthType:     types.FunctionUrlAuthTypeAwsIam,
		FunctionName: aws.String(lambdaName),
		Cors: &types.Cors{
			AllowCredentials: aws.Bool(true),
			AllowHeaders:     []string{"*"},
			AllowMethods:     []string{http.MethodPost},
			AllowOrigins:     []string{zeusCloudOrigin},
		},
	}
	resp, err := l.CreateFunctionUrlConfig(ctx, input)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("LambdaClientAWS: MakeLambdaURL: error making function public")
		return resp, err
	}
	return resp, err
}

func (l *LambdaClientAWS) MakeLambdaFuncAuthIAM(ctx context.Context, lambdaName string) (*lambda.AddPermissionOutput, error) {
	input := &lambda.AddPermissionInput{
		Action:              aws.String("lambda:InvokeFunctionUrl"),
		FunctionName:        aws.String(lambdaName),
		Principal:           aws.String("*"),
		StatementId:         aws.String("FunctionURLAllowAuthIAMAccess"),
		FunctionUrlAuthType: types.FunctionUrlAuthTypeAwsIam,
	}
	resp, err := l.AddPermission(ctx, input)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("LambdaClientAWS: MakeLambdaFuncAuthIAM: error making function iam")
		return resp, err
	}
	return resp, err
}

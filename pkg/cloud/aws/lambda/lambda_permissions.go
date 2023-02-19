package aws_lambda

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/rs/zerolog/log"
)

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

// MakeEthereumSignerURL uses the EthereumSignerFunctionName to make the function public
func (l *LambdaClientAWS) MakeEthereumSignerURL(ctx context.Context) (*lambda.CreateFunctionUrlConfigOutput, error) {
	input := &lambda.CreateFunctionUrlConfigInput{
		AuthType:     types.FunctionUrlAuthTypeNone,
		FunctionName: aws.String(EthereumSignerFunctionName),
	}
	resp, err := l.CreateFunctionUrlConfig(ctx, input)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("LambdaClientAWS: MakeFuncPublic: error making function public")
		return resp, err
	}

	return resp, err
}

func (l *LambdaClientAWS) GetExternalLambdaSignerConfigURL(ctx context.Context) (*lambda.GetFunctionUrlConfigOutput, error) {
	input := &lambda.GetFunctionUrlConfigInput{
		FunctionName: aws.String(EthereumSignerFunctionName),
	}
	resp, err := l.GetFunctionUrlConfig(ctx, input)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("LambdaClientAWS: MakeFuncPublic: error making function public")
		return resp, err
	}

	return resp, err
}

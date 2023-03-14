package aws_lambda

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/rs/zerolog/log"
)

func (l *LambdaClientAWS) GetExternalLambdaFuncInfo(ctx context.Context) (*lambda.GetFunctionOutput, error) {
	fnInfo, err := l.GetFunction(ctx, &lambda.GetFunctionInput{
		FunctionName: aws.String(EthereumSignerFunctionName),
		Qualifier:    nil,
	})
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("GetExternalLambdaFuncInfo: error getting lambda function info")
		return nil, err
	}
	return fnInfo, err
}

func (l *LambdaClientAWS) GetLambdaFuncInfo(ctx context.Context, functionName string) (*lambda.GetFunctionOutput, error) {
	fnInfo, err := l.GetFunction(ctx, &lambda.GetFunctionInput{
		FunctionName: aws.String(functionName),
		Qualifier:    nil,
	})
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("GetLambdaFuncInfo: error getting lambda function info")
		return nil, err
	}
	return fnInfo, err
}

func (l *LambdaClientAWS) GetKeystoreLayerInfo(ctx context.Context) (*lambda.ListLayerVersionsOutput, error) {
	fnInfo, err := l.ListLayerVersions(ctx, &lambda.ListLayerVersionsInput{
		LayerName: aws.String(KeystoresLayerName),
	})
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("GetExternalLambdaFuncInfo: error getting lambda layer info")
		return nil, err
	}
	return fnInfo, err
}

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

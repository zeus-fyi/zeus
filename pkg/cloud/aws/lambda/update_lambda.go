package aws_lambda

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/rs/zerolog/log"
	"github.com/zeus-fyi/zeus/builds"
)

func (l *LambdaClientAWS) UpdateServerlessBLSLambdaFnBinary(ctx context.Context, functionName string) (*lambda.UpdateFunctionCodeOutput, error) {
	builds.ChangeToBuildsDir()
	update := &lambda.UpdateFunctionCodeInput{
		FunctionName: aws.String(functionName),
		ZipFile:      blsMainZipFilePath.ReadFileInPath(),
	}
	lf, err := l.UpdateFunctionCode(ctx, update)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("CreateNewLambdaFn: error creating lambda function")
		return nil, err
	}
	return lf, err
}

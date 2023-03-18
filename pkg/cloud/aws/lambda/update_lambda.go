package aws_lambda

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/rs/zerolog/log"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
)

func (l *LambdaClientAWS) UpdateServerlessBLSLambdaFnBinary(ctx context.Context, functionName string, p filepaths.Path) (*lambda.UpdateFunctionCodeOutput, error) {
	update := &lambda.UpdateFunctionCodeInput{
		FunctionName: aws.String(functionName),
		ZipFile:      p.ReadFileInPath(),
	}
	lf, err := l.UpdateFunctionCode(ctx, update)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("CreateNewLambdaFn: error creating lambda function")
		return nil, err
	}
	return lf, err
}

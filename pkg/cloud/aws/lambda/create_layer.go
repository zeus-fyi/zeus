package aws_lambda

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/rs/zerolog/log"
)

func (l *LambdaClientAWS) CreateServerlessBLSLambdaFnKeystoreLayer(ctx context.Context, keystoresLayerName string, blsKeystoresZipFileBinary []byte) (*lambda.PublishLayerVersionOutput, error) {
	input := &lambda.PublishLayerVersionInput{
		Content: &types.LayerVersionContentInput{
			ZipFile: blsKeystoresZipFileBinary,
		},
		LayerName:               aws.String(keystoresLayerName),
		CompatibleArchitectures: []types.Architecture{types.ArchitectureX8664},
		CompatibleRuntimes:      []types.Runtime{types.RuntimeGo1x},
		Description:             aws.String("-"),
		LicenseInfo:             aws.String("-"),
	}
	ly, err := l.PublishLayerVersion(ctx, input)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("LambdaClientAWS: CreateServerlessBLSLambdaFnKeystoreLayer: error creating lambda layer")
		return nil, err
	}
	return ly, err
}

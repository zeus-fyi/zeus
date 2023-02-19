package aws_lambda

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/rs/zerolog/log"
	"github.com/zeus-fyi/zeus/builds"
)

func (l *LambdaClientAWS) UpdateServerlessBLSLambdaFnKeystoreLayer(ctx context.Context) (*lambda.UpdateFunctionConfigurationOutput, error) {
	builds.ChangeToBuildsDir()

	layerVersion, err := l.GetKeystoreLayerInfo(ctx)
	if err != nil || layerVersion == nil {
		log.Ctx(ctx).Err(err).Msg("CreateNewLambdaFn: error getting lambda function keystore layer info")
		return nil, err
	}

	input := &lambda.UpdateFunctionConfigurationInput{
		FunctionName: aws.String(EthereumSignerFunctionName),
		Layers:       []string{l.GetLambdaExtensionARN(), *layerVersion.LayerVersions[0].LayerVersionArn},
	}
	ly, err := l.UpdateFunctionConfiguration(ctx, input)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("LambdaClientAWS: UpdateServerlessBLSLambdaFnKeystoreLayer: error updating to fn to use new lambda layer")
		return nil, err
	}
	return ly, err
}

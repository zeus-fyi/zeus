package aws_lambda

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/rs/zerolog/log"
	"github.com/zeus-fyi/zeus/builds"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
)

var (
	KeystoresLayerName      = "blsKeystores"
	blsKeystoresZipFilePath = filepaths.Path{DirIn: "./serverless/bls_signatures", FnIn: "keystores.zip"}
)

func (l *LambdaClientAWS) CreateServerlessBLSLambdaFnKeystoreLayer(ctx context.Context) (*lambda.PublishLayerVersionOutput, error) {
	builds.ChangeToBuildsDir()

	b := blsKeystoresZipFilePath.ReadFileInPath()
	input := &lambda.PublishLayerVersionInput{
		Content: &types.LayerVersionContentInput{
			ZipFile: b,
		},
		LayerName:               aws.String(KeystoresLayerName),
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

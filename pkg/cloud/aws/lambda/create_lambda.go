package aws_lambda

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/rs/zerolog/log"
	"github.com/zeus-fyi/zeus/builds"
	aegis_aws_iam "github.com/zeus-fyi/zeus/pkg/aegis/aws/iam"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
)

const region = "us-west-1"

var (
	EthereumSignerFunctionName = "ethereumSignerBLS"
	blsMainZipFilePath         = filepaths.Path{DirIn: "./serverless/bls_signatures", FnIn: "main.zip"}
	blsFnParams                = &lambda.CreateFunctionInput{
		Code: &types.FunctionCode{
			ZipFile: nil,
		},
		FunctionName:      aws.String(EthereumSignerFunctionName),
		Role:              nil,
		Architectures:     []types.Architecture{types.ArchitectureX8664},
		Description:       aws.String("BLS Ethereum Validator Signer Lambda Function"),
		FileSystemConfigs: nil,
		Handler:           aws.String("main"),
		Layers:            []string{},
		MemorySize:        nil,
		PackageType:       types.PackageTypeZip,
		Publish:           false,
		Runtime:           types.RuntimeGo1x,
		Tags:              make(map[string]string),
		Timeout:           aws.Int32(3),
		TracingConfig:     nil,
	}
)

/*
Creates a Lambda function. To create a function, you need a deployment package
(https://docs.aws.amazon.com/lambda/latest/dg/gettingstarted-package.html) and
 an execution role
(https://docs.aws.amazon.com/lambda/latest/dg/intro-permission-model.html#lambda-intro-execution-role).
*/

// aegis_aws_iam

// GetLambdaRole references a role created in aegis_aws_iam
func (l *LambdaClientAWS) GetLambdaRole() string {
	return fmt.Sprintf("arn:aws:iam::%s:role/%s", l.AccountNumber, aegis_aws_iam.LambdaRoleName)
}

// GetLambdaExtensionARN uses the us-west-1 specific number
// more info: https://docs.aws.amazon.com/systems-manager/latest/userguide/ps-integration-lambda-extensions.html
func (l *LambdaClientAWS) GetLambdaExtensionARN() string {
	return fmt.Sprintf("arn:aws:lambda:%s:997803712105:layer:AWS-Parameters-and-Secrets-Lambda-Extension:4", l.Region)
}

// GetLambdaKeystoreLayerARN uses version 1, you'll need to update if you add new versions to this layer
func (l *LambdaClientAWS) GetLambdaKeystoreLayerARN(version string) string {
	return fmt.Sprintf("arn:aws:lambda:%s:%s:layer:%s:%s", l.Region, l.AccountNumber, KeystoresLayerName, version)
}

func (l *LambdaClientAWS) CreateServerlessBLSLambdaFn(ctx context.Context) (*lambda.CreateFunctionOutput, error) {
	builds.ChangeToBuildsDir()
	blsFnParams.Role = aws.String(l.GetLambdaRole())
	blsFnParams.Code.ZipFile = blsMainZipFilePath.ReadFileInPath()
	layerVersion, err := l.GetKeystoreLayerInfo(ctx)
	if err != nil || layerVersion == nil {
		log.Ctx(ctx).Err(err).Msg("CreateNewLambdaFn: error getting lambda function keystore layer info")
		return nil, err
	}
	blsFnParams.Layers = []string{l.GetLambdaExtensionARN(), *layerVersion.LayerVersions[0].LayerVersionArn}
	lf, err := l.CreateFunction(ctx, blsFnParams)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("CreateNewLambdaFn: error creating lambda function")
		return nil, err
	}
	return lf, err
}

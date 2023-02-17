package aws_lambda

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/rs/zerolog/log"
	aws_aegis_auth "github.com/zeus-fyi/zeus/pkg/aegis/aws/auth"
)

type LambdaClientAWS struct {
	AccountNumber string
	Region        string
	*lambda.Client
}

func InitLambdaClient(ctx context.Context, auth aws_aegis_auth.AuthAWS) (LambdaClientAWS, error) {
	creds := credentials.NewStaticCredentialsProvider(auth.AccessKey, auth.SecretKey, "")
	cfg, err := config.LoadDefaultConfig(ctx, config.WithCredentialsProvider(creds), config.WithRegion(auth.Region))
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("InitLambdaClient: error loading config")
		return LambdaClientAWS{}, err
	}
	cfg.Region = auth.Region
	log.Ctx(ctx).Info().Interface("region", auth.Region).Msg("InitLambdaClient")
	lambdaClient := lambda.NewFromConfig(cfg)
	return LambdaClientAWS{AccountNumber: auth.AccountNumber, Client: lambdaClient, Region: auth.Region}, err
}

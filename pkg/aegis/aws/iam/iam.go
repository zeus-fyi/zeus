package aegis_aws_iam

import (
	"context"
	aws_aegis_auth "github.com/zeus-fyi/zeus/pkg/aegis/aws/auth"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/rs/zerolog/log"
)

type IAMClientAWS struct {
	AccountNumber string `json:"accountNumber"`
	Region        string `json:"region"`
	*iam.Client
}

func InitIAMClient(ctx context.Context, auth aws_aegis_auth.AuthAWS) (IAMClientAWS, error) {
	creds := credentials.NewStaticCredentialsProvider(auth.AccessKey, auth.SecretKey, "")
	cfg, err := config.LoadDefaultConfig(ctx, config.WithCredentialsProvider(creds), config.WithRegion(auth.Region))
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("InitLambdaClient: error loading config")
		return IAMClientAWS{}, err
	}
	cfg.Region = auth.Region
	iamClient := iam.NewFromConfig(cfg)
	log.Ctx(ctx).Info().Interface("region", auth.Region).Msg("InitIAMClient")
	return IAMClientAWS{AccountNumber: auth.AccountNumber, Region: auth.Region, Client: iamClient}, err
}
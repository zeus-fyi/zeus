package aegis_aws_iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/rs/zerolog/log"
	aws_aegis_auth "github.com/zeus-fyi/zeus/pkg/aegis/aws/auth"
)

type IAMClientAWS struct {
	AccountNumber string `json:"accountNumber"`
	Region        string `json:"region"`
	*iam.Client
}

func InitIAMClient(ctx context.Context, auth aws_aegis_auth.AuthAWS) (IAMClientAWS, error) {
	cfg, err := auth.CreateConfig(ctx)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("InitLambdaClient: error loading config")
		return IAMClientAWS{}, err
	}
	iamClient := iam.NewFromConfig(cfg)
	log.Ctx(ctx).Info().Interface("region", auth.Region).Msg("InitIAMClient")
	return IAMClientAWS{AccountNumber: auth.AccountNumber, Region: auth.Region, Client: iamClient}, err
}

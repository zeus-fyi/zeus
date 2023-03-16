package aegis_aws_iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/rs/zerolog/log"
	aws_aegis_auth "github.com/zeus-fyi/zeus/pkg/aegis/aws/auth"
)

func (i *IAMClientAWS) CreateUserAccessKeys(ctx context.Context, userName string) (aws_aegis_auth.AuthAWS, error) {
	auth := aws_aegis_auth.AuthAWS{
		Region:    i.Region,
		AccessKey: "",
		SecretKey: "",
	}
	createAccessKeyInput := &iam.CreateAccessKeyInput{
		UserName: &userName,
	}
	createAccessKeyOutput, err := i.CreateAccessKey(ctx, createAccessKeyInput)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("CreateUserAccessKeys: error creating access key")
		return auth, nil
	}
	auth.AccessKey = *createAccessKeyOutput.AccessKey.AccessKeyId
	auth.SecretKey = *createAccessKeyOutput.AccessKey.SecretAccessKey
	return auth, nil
}

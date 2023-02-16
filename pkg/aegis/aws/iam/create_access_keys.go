package aegis_aws_iam

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/rs/zerolog/log"
	aegis_aws_secretmanager "github.com/zeus-fyi/zeus/pkg/aegis/aws/secretmanager"
)

func (i *IAMClientAWS) CreateUserAccessKeys(ctx context.Context, userName string) (aegis_aws_secretmanager.AuthAWS, error) {
	auth := aegis_aws_secretmanager.AuthAWS{
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
	println("Access Key:", auth.AccessKey)
	println("Secret Key:", auth.SecretKey)

	return auth, nil
}

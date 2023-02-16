package aegis_aws_iam

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/rs/zerolog/log"
)

func (i *IAMClientAWS) CreateUserAccessKeys(ctx context.Context, userName string) error {
	createAccessKeyInput := &iam.CreateAccessKeyInput{
		UserName: &userName,
	}
	createAccessKeyOutput, err := i.CreateAccessKey(ctx, createAccessKeyInput)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("CreateUserAccessKeys: error creating access key")
		return err
	}
	accessKey := createAccessKeyOutput.AccessKey.AccessKeyId
	secretKey := createAccessKeyOutput.AccessKey.SecretAccessKey
	println("Access Key:", accessKey)
	println("Secret Key:", secretKey)
	return nil
}

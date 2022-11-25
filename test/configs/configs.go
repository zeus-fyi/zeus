package configs

import (
	"github.com/spf13/viper"
)

var testCont TestContainer

type TestContainer struct {
	Env                string
	Bearer             string
	PresignedBucketURL string
}

func InitLocalTestConfigs() TestContainer {
	InitEnvFromConfig(ForceDirToConfigLocation())
	testCont.Env = viper.GetString("ENV")
	testCont.Bearer = viper.GetString("BEARER")
	testCont.PresignedBucketURL = viper.GetString("PRESIGNED_BUCKET_URL")

	return testCont
}

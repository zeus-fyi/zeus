package configs

import (
	"github.com/spf13/viper"
)

var testCont TestContainer

type TestContainer struct {
	Env                string
	Bearer             string
	PresignedBucketURL string

	LocalEcsdaTestPkey  string
	LocalEcsdaTestPkey2 string
}

func InitLocalTestConfigs() TestContainer {
	InitEnvFromConfig(ForceDirToConfigLocation())
	testCont.Env = viper.GetString("ENV")
	testCont.Bearer = viper.GetString("BEARER")

	// snapshot testing
	testCont.PresignedBucketURL = viper.GetString("PRESIGNED_BUCKET_URL")

	// artemis testing
	testCont.LocalEcsdaTestPkey = viper.GetString("LOCAL_TESTING_ECDSA_PKEY")
	testCont.LocalEcsdaTestPkey2 = viper.GetString("LOCAL_TESTING_ECDSA_PKEY_2")
	return testCont
}

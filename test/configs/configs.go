package configs

import (
	"github.com/spf13/viper"
)

var testCont TestContainer

type TestContainer struct {
	Env    string
	Bearer string
}

func InitLocalTestConfigs() TestContainer {
	InitEnvFromConfig(ForceDirToConfigLocation())
	testCont.Env = viper.GetString("ENV")
	testCont.Bearer = viper.GetString("BEARER_TOKEN")

	return testCont
}

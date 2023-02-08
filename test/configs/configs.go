package configs

import (
	"github.com/spf13/viper"
)

var testCont TestContainer

type TestContainer struct {
	Env                string
	Bearer             string
	PresignedBucketURL string
	NodeURL            string
	GoerliNodeURL      string
	EphemeralNodeURL   string

	LocalMnemonic24Words string
	LocalBLSTestPkey     string
	LocalWdBLSTestPkey   string

	LocalEcsdaTestPkey  string
	LocalEcsdaTestPkey2 string

	HDWalletPassword string

	Web3SignerDockerImage string
	Web3SignerAuthURL     string

	AccessKeyAWS string
	SecretKeyAWS string

	AgePubKey  string
	AgePrivKey string

	EncryptedKeystoresPath   string
	UnencryptedKeystoresPath string

	ServerlessSignerFuncBLS        string
	ServerlessSignerFuncSecretName string
}

func InitLocalTestConfigs() TestContainer {
	InitEnvFromConfig(ForceDirToConfigLocation())
	testCont.Env = viper.GetString("ENV")
	testCont.Bearer = viper.GetString("BEARER")

	// snapshot testing
	testCont.PresignedBucketURL = viper.GetString("PRESIGNED_BUCKET_URL")

	// artemis testing
	testCont.NodeURL = viper.GetString("NODE_URL")
	testCont.GoerliNodeURL = viper.GetString("GOERLI_NODE_URL")
	testCont.EphemeralNodeURL = viper.GetString("EPHEMERAL_NODE_URL")
	testCont.LocalEcsdaTestPkey = viper.GetString("LOCAL_TESTING_ECDSA_PKEY")
	testCont.LocalEcsdaTestPkey2 = viper.GetString("LOCAL_TESTING_ECDSA_PKEY_2")
	testCont.LocalBLSTestPkey = viper.GetString("LOCAL_TESTING_BLS_SECRET_KEY")
	testCont.LocalWdBLSTestPkey = viper.GetString("LOCAL_TESTING_BLS_WD_SECRET_KEY")
	testCont.LocalMnemonic24Words = viper.GetString("MNEMONIC_24_WORDS")
	testCont.HDWalletPassword = viper.GetString("HD_WALLET_PASSWORD")
	testCont.Web3SignerDockerImage = viper.GetString("WEB3SIGNER_DOCKER_IMG")
	testCont.Web3SignerAuthURL = viper.GetString("WEB3SIGNER_AUTH_URL")

	testCont.AccessKeyAWS = viper.GetString("AWS_ACCESS_KEY")
	testCont.SecretKeyAWS = viper.GetString("AWS_SECRET_KEY")

	testCont.AgePubKey = viper.GetString("AGE_PUBKEY")
	testCont.AgePrivKey = viper.GetString("AGE_PRIV_KEY")

	testCont.EncryptedKeystoresPath = viper.GetString("ENCRYPTED_KEYSTORES_PATH")
	testCont.UnencryptedKeystoresPath = viper.GetString("RAW_UNENCRYPTED_KEYSTORES_PATH")
	testCont.ServerlessSignerFuncBLS = viper.GetString("BLS_SERVERLESS_LAMBA_FUNC_ADDR")
	testCont.ServerlessSignerFuncSecretName = viper.GetString("BLS_SERVERLESS_LAMBA_FUNC_SECRET_NAME")
	return testCont
}

package main

import (
	"context"
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
	"time"

	"github.com/gochain/gochain/v4/core/types"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zeus-fyi/gochain/web3/accounts"
	"github.com/zeus-fyi/gochain/web3/web3_actions"
	"github.com/zeus-fyi/zeus/builds"
	serverless_aws_automation "github.com/zeus-fyi/zeus/builds/serverless/aws_automation"
	ethereum_automation_cookbook "github.com/zeus-fyi/zeus/cookbooks/ethereum/automation"
	aws_aegis_auth "github.com/zeus-fyi/zeus/pkg/aegis/aws/auth"
	signing_automation_ethereum "github.com/zeus-fyi/zeus/pkg/artemis/signing_automation/ethereum"
	age_encryption "github.com/zeus-fyi/zeus/pkg/crypto/age"
	bls_signer "github.com/zeus-fyi/zeus/pkg/crypto/bls"
	aegis_random "github.com/zeus-fyi/zeus/pkg/crypto/random"
	hestia_req_types "github.com/zeus-fyi/zeus/pkg/hestia/client/req_types"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	strings_filter "github.com/zeus-fyi/zeus/pkg/utils/strings"
	"k8s.io/apimachinery/pkg/util/rand"
)

func main() {
	if err := Cmd.Execute(); err != nil {
		log.Err(err)
	}
}

var (
	nodeURL                string
	network                string
	agePubKey              string
	agePrivKey             string
	mnemonic               string
	hdWalletPassword       string
	numKeysToGen, hdOffset int
	keystoresPath          = filepaths.Path{
		DirIn:  "",
		DirOut: "",
	}
	keyGenSecrets        bool
	genValidatorDeposits bool
	sendDeposits         bool
	eth1AddrPrivKey      string

	automateSetupOnAWS            bool
	mnemonicAndHDWalletSecretName = "mnemonicAndHDWallet"
	ageEncryptionSecretName       = "ageEncryptionKey"
	externalZeusLambdaAccessKeys  = "externalZeusLambdaAccessKeys"
	awsRegion                     = "us-west-1"
	awsAuth                       = aws_aegis_auth.AuthAWS{
		AccountNumber: "",
		Region:        awsRegion,
		AccessKey:     "",
		SecretKey:     "",
	}
	feeRecipient              string
	bearerToken               string
	keyGroupName              string
	submitValidatorServiceReq bool
	externalAwsAuth           = hestia_req_types.AuthLamdbaAWS{
		ServiceURL: "",
		SecretName: "",
		AccessKey:  "",
		SecretKey:  "",
	}

	automationSteps = ""
)

func init() {
	dir := ForceDirToRootDir()
	viper.AddConfigPath(dir)
	_ = viper.ReadInConfig()
	err := bls_signer.InitEthBLS()
	if err != nil {
		panic(err)
	}
	/*
		########################################
			AWS settings (internal access)
		########################################
	*/
	// aws automation settings for lambda set up
	// your aws account number should not have dashes here, even thought it is displayed with dashes in the UI
	Cmd.Flags().StringVar(&awsAuth.AccountNumber, "aws-account-number", "", "AWS_ACCOUNT_NUMBER: aws account number")
	Cmd.Flags().StringVar(&awsAuth.AccessKey, "aws-access-key", "", "AWS_ACCESS_KEY: your private aws access key, which needs permissions to create iam users, roles, policies, secrets, and lambda functions and layers")
	Cmd.Flags().StringVar(&awsAuth.SecretKey, "aws-secret-key", "", "AWS_SECRET_KEY: your private aws secret key")
	if viper.GetString("AWS_ACCESS_KEY") != "" && viper.GetString("AWS_SECRET_KEY") != "" && viper.GetString("AWS_ACCOUNT_NUMBER") != "" {
		awsAuth.AccountNumber = viper.GetString("AWS_ACCOUNT_NUMBER")
		awsAuth.AccessKey = viper.GetString("AWS_ACCESS_KEY")
		awsAuth.SecretKey = viper.GetString("AWS_SECRET_KEY")
	}
	// actions
	Cmd.Flags().BoolVar(&automateSetupOnAWS, "aws-automation-on", viper.GetBool("AWS_AUTOMATION"), "AWS_AUTOMATION: automate the entire setup process on aws, requires you provide aws credentials")
	/*
		########################################
			AWS settings (external access)
		########################################
	*/
	// aws service
	Cmd.Flags().StringVar(&externalAwsAuth.ServiceURL, "ext-aws-lambda-url", viper.GetString("AWS_LAMBDA_FUNC_URL"), "AWS_LAMBDA_FUNC_URL: your lambda func url for validator service on zeus")
	Cmd.Flags().StringVar(&externalAwsAuth.AccessKey, "ext-aws-access-key", "", "AWS_EXTERNAL_ACCESS_KEY: external access token for validator service on zeus")
	Cmd.Flags().StringVar(&externalAwsAuth.SecretKey, "ext-aws-secret-key", "", "AWS_EXTERNAL_SECRET_KEY: external secret token for validator service on zeus")
	Cmd.Flags().StringVar(&externalAwsAuth.SecretName, "ext-aws-age-secret-name", viper.GetString("AWS_AGE_DECRYPTION_SECRET_NAME"), "AWS_AGE_DECRYPTION_SECRET_NAME: the name of the secret that holds your age decryption keys for validator service on zeus")
	if viper.GetString("AWS_EXTERNAL_ACCESS_KEY") != "" && viper.GetString("AWS_EXTERNAL_SECRET_KEY") != "" {
		externalAwsAuth.AccessKey = viper.GetString("AWS_EXTERNAL_ACCESS_KEY")
		externalAwsAuth.SecretKey = viper.GetString("AWS_EXTERNAL_SECRET_KEY")
	}

	// secret key generation for serverless
	Cmd.Flags().StringVar(&agePrivKey, "age-private-key", "", "AGE_PRIVKEY: age private key")
	Cmd.Flags().StringVar(&agePubKey, "age-public-key", "", "AGE_PUBKEY: age public key")
	if viper.GetString("AGE_PRIVKEY") != "" && viper.GetString("AGE_PUBKEY") != "" {
		agePrivKey = viper.GetString("AGE_PRIVKEY")
		agePubKey = viper.GetString("AGE_PUBKEY")
	}
	// actions
	Cmd.Flags().BoolVar(&keyGenSecrets, "keygen", viper.GetBool("KEYGEN_SECRETS"), "KEYGEN_SECRETS: generates secrets for validator encryption and generation")
	/*
		###################################################
			Validator network, keygen, and deposit settings
		###################################################
	*/
	// validator key generation for deposits settings
	if viper.GetString("NETWORK") == "" || viper.GetString("NODE_URL") == "" {
		log.Info().Msg("no network or node url provided, using ephemery network and eth.ephemeral.zeus.fyi node url")
		viper.Set("NETWORK", "ephemery")
		viper.Set("NODE_URL", "https://eth.ephemeral.zeus.fyi")
	}
	Cmd.Flags().StringVar(&nodeURL, "node-url", viper.GetString("NODE_URL"), "NODE_URL: beacon for getting network data for validator deposit generation & submitting deposits")
	Cmd.Flags().StringVar(&network, "network", viper.GetString("NETWORK"), "NETWORK: network to run on mainnet, goerli, ephemery, etc")
	Cmd.Flags().StringVar(&eth1AddrPrivKey, "eth1-addr-priv-key", "", "ETH1_PRIVATE_KEY: eth1 address private key for submitting deposits")
	if viper.GetString("ETH1_PRIVATE_KEY") != "" {
		eth1AddrPrivKey = viper.GetString("ETH1_PRIVATE_KEY")
	}

	// validator secret key generation
	Cmd.Flags().StringVar(&mnemonic, "mnemonic", "", "MNEMONIC_24_WORDS: twenty four word mnemonic to generate keystores")
	Cmd.Flags().StringVar(&hdWalletPassword, "hd-wallet-pw", "", "HD_WALLET_PASSWORD: hd wallet password")

	if viper.GetString("MNEMONIC_24_WORDS") != "" && viper.GetString("HD_WALLET_PASSWORD") != "" {
		mnemonic = viper.GetString("MNEMONIC_24_WORDS")
		hdWalletPassword = viper.GetString("HD_WALLET_PASSWORD")
	}

	Cmd.Flags().IntVar(&numKeysToGen, "validator-count", viper.GetInt("VALIDATORS_COUNT"), "VALIDATORS_COUNT: number of keys to generate")
	Cmd.Flags().IntVar(&hdOffset, "hd-offset", viper.GetInt("HD_OFFSET_VALIDATORS"), "HD_OFFSET_VALIDATORS: offset to start generating keys from hd wallet")
	// validator key generation paths
	Cmd.Flags().StringVar(&keystoresPath.DirIn, "keystores-dir-in", viper.GetString("KEYSTORE_DIR_IN"), "KEYSTORE_DIR_IN: keystores directory in location (relative to builds dir)")
	Cmd.Flags().StringVar(&keystoresPath.DirOut, "keystores-dir-out", viper.GetString("KEYSTORE_DIR_OUT"), "KEYSTORE_DIR_OUT: keystores directory out location (relative to builds dir)")
	if keystoresPath.DirIn == "" || keystoresPath.DirOut == "" {
		log.Info().Msg("no keystore path provided, using ./serverless/keystores as default, this is relative to the builds directory")
		keystoresPath.DirIn = "./serverless/keystores"
		keystoresPath.DirOut = "./serverless/keystores"
	}
	// actions
	Cmd.Flags().BoolVar(&genValidatorDeposits, "keygen-validators", viper.GetBool("KEYGEN_VALIDATORS"), "KEYGEN_VALIDATORS: generates validator deposits, with additional encrypted age keystore")
	Cmd.Flags().BoolVar(&sendDeposits, "submit-deposits", viper.GetBool("SUBMIT_DEPOSITS"), "SUBMIT_DEPOSITS: submits validator deposits in keystore directory to the network for activation")
	/*
		###################################################
			Validator service values for Zeus
		###################################################
	*/
	// service on zeus
	Cmd.Flags().StringVar(&bearerToken, "bearer", viper.GetString("BEARER"), "BEARER: bearer token for validator service on zeus")
	Cmd.Flags().StringVar(&keyGroupName, "key-group-name", viper.GetString("KEY_GROUP_NAME"), "KEY_GROUP_NAME: name for validator service group on zeus")
	Cmd.Flags().StringVar(&feeRecipient, "fee-recipient", viper.GetString("FEE_RECIPIENT_ADDR"), "FEE_RECIPIENT_ADDR: fee recipient address for validators service on zeus")
	Cmd.Flags().BoolVar(&submitValidatorServiceReq, "submit-validator-service-req", viper.GetBool("SUBMIT_SERVICE_REQUEST"), "SUBMIT_SERVICE_REQUEST: sends a request to zeus to setup a validator service")
	/*
		###################################################
			Automation Steps
		###################################################
	*/
	if viper.GetString("AUTOMATION_STEPS") == "" {
		viper.Set("AUTOMATION_STEPS", "all")
	}
	Cmd.Flags().StringVar(&automationSteps, "automation-steps", viper.GetString("AUTOMATION_STEPS"), "AUTOMATION_STEPS: select which steps to automate and which order, using a comma separated list of numbers. default is all steps in order")
}

// Cmd represents the base command when called without any subcommands
var Cmd = &cobra.Command{
	Use:   "Validator Key Generation and AWS Lambda Serverless Setup Automation",
	Short: "Automates the entire setup process for validator keys and serverless setup on AWS",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()

		if automateSetupOnAWS {
			if awsAuth.AccountNumber == "" || awsAuth.AccessKey == "" || awsAuth.SecretKey == "" {
				panic("ERROR: aws credentials and/or account number missing")
			}
		}

		w3Client := signing_automation_ethereum.Web3SignerClient{
			Web3Actions: web3_actions.Web3Actions{
				NodeURL: nodeURL,
				Network: network,
			},
		}

		if automationSteps == "all" {
			automationSteps = "1,2,3,4,5,6,7,8,9"
			automateSetupOnAWS = true
		}
		if automationSteps == "serverless" {
			automationSteps = "1,2,3,4,5,6,7"
			automateSetupOnAWS = true
		}

		for _, automationStep := range strings.Split(automationSteps, ",") {
			switch automationStep {
			case "getMnemonicHDWalletPasswordSecret":
				s, err := serverless_aws_automation.GetSecret(ctx, awsAuth, mnemonicAndHDWalletSecretName)
				if err != nil {
					panic(err)
				}
				fmt.Println(s)
			case "getAgeEncryptionKeySecret":
				s, err := serverless_aws_automation.GetSecret(ctx, awsAuth, ageEncryptionSecretName)
				if err != nil {
					panic(err)
				}
				fmt.Println(s)
			case "getExternalLambdaAccessKeys":
				s, err := serverless_aws_automation.GetSecret(ctx, awsAuth, externalZeusLambdaAccessKeys)
				if err != nil {
					panic(err)
				}
				fmt.Println(s)
			case "updateLambdaKeystoresLayerToLatest":
				serverless_aws_automation.UpdateLambdaFunctionKeystoresLayer(ctx, awsAuth)
			case "1", "createSecretsAndStoreInAWS":
				if agePubKey == "" || agePrivKey == "" {
					fmt.Println("INFO: no credentials provided, generating new age keypair")
					agePubKey, agePrivKey = age_encryption.GenerateNewKeyPair()
					fmt.Println("agePubKey: ", agePubKey)
					fmt.Println("agePrivKey: ", agePrivKey)
				}
				if hdWalletPassword == "" {
					fmt.Println("INFO: no credentials provided, generating random string password")
					hdWalletPassword = rand.String(32)
					fmt.Println("hdWalletPassword: ", hdWalletPassword)
				}
				if mnemonic == "" {
					fmt.Println("INFO: no mnemonic provided, generating new mnemonic")
					mnemonic24Words, err := aegis_random.GenerateMnemonic()
					if err != nil {
						panic(err)
					}
					mnemonic = mnemonic24Words
					fmt.Println("mnemonic: ", mnemonic)
				}
				if automateSetupOnAWS {
					serverless_aws_automation.AddMnemonicHDWalletSecretInAWSSecretManager(ctx, awsAuth, mnemonicAndHDWalletSecretName, hdWalletPassword, mnemonic)
					serverless_aws_automation.AddAgeEncryptionKeyInAWSSecretManager(ctx, awsAuth, ageEncryptionSecretName, agePubKey, agePrivKey)
				}
			case "2", "createInternalLambdaUser":
				fmt.Println("INFO: creating internal iam user, role, policies for serverless deployment")
				serverless_aws_automation.InternalUserRolePolicySetupForLambdaDeployment(ctx, awsAuth)
			case "3", "generateValidatorDeposits":
				fmt.Println("INFO: generating keystores, deposit data, and encypting keystores with age encryption")
				s, err := serverless_aws_automation.GetSecret(ctx, awsAuth, ageEncryptionSecretName)
				if err != nil {
					panic(err)
				}
				mns, err := serverless_aws_automation.GetSecret(ctx, awsAuth, mnemonicAndHDWalletSecretName)
				if err != nil {
					panic(err)
				}
				mnemonic = mns["mnemonic"]
				hdWalletPassword = mns["hdWalletPassword"]

				for pubkey, privkey := range s {
					agePubKey = pubkey
					agePrivKey = privkey
				}
				vdg := signing_automation_ethereum.ValidatorDepositGenerationParams{
					Fp:                   keystoresPath,
					Mnemonic:             mnemonic,
					Pw:                   hdWalletPassword,
					ValidatorIndexOffset: hdOffset,
					NumValidators:        numKeysToGen,
					Network:              network,
				}
				enc := age_encryption.NewAge(agePrivKey, agePubKey)
				builds.ChangeToBuildsDir()
				err = ethereum_automation_cookbook.GenerateValidatorDepositsAndCreateAgeEncryptedKeystores(ctx, w3Client, vdg, enc, hdWalletPassword)
				if err != nil {
					panic(err)
				}
			case "4", "createLambdaFunctionKeystoresLayer":
				fmt.Println("INFO: creating lambda keystore layer using your encrypted keystores from step 3")
				serverless_aws_automation.CreateLambdaFunctionKeystoresLayer(ctx, awsAuth)
			case "5", "createLambdaFunction":
				fmt.Println("INFO: creating lambda function")
				lambdaFnUrl := serverless_aws_automation.CreateLambdaFunction(ctx, awsAuth)
				externalAwsAuth.ServiceURL = lambdaFnUrl
				fmt.Println("lambdaFnUrl: ", externalAwsAuth.ServiceURL)
			case "6", "createExternalLambdaUser":
				fmt.Println("INFO: creating external iam user, role, policies for us to send validator messages to your lambda function")
				serverless_aws_automation.ExternalUserRolePolicySetupForLambdaDeployment(ctx, awsAuth)
				if externalAwsAuth.AccessKey == "" || externalAwsAuth.SecretKey == "" {
					s, err := serverless_aws_automation.GetExternalAccessKeySecretIfExists(ctx, awsAuth, externalZeusLambdaAccessKeys)
					if err != nil {
						panic(err)
					}
					externalAwsAuth.AccessKey = s.AccessKey
					externalAwsAuth.SecretKey = s.SecretKey
					if externalAwsAuth.AccessKey == "" || externalAwsAuth.SecretKey == "" {
						externalAccessKeys := serverless_aws_automation.CreateExternalLambdaUserAccessKeys(ctx, awsAuth)
						serverless_aws_automation.AddExternalAccessKeysInAWSSecretManager(ctx, awsAuth, externalZeusLambdaAccessKeys, externalAccessKeys)
						externalAwsAuth.AccessKey = externalAccessKeys.AccessKey
						externalAwsAuth.SecretKey = externalAccessKeys.SecretKey
					}
				}
			case "7", "verifyLambdaFunction":
				if externalAwsAuth.ServiceURL == "" {
					fmt.Println("INFO: no lambda fn url, looking up url if exists")
					externalAwsAuth.ServiceURL = serverless_aws_automation.GetLambdaFunctionUrl(ctx, awsAuth)
				}
				if externalAwsAuth.ServiceURL == "" {
					panic("ERROR: lambda function url not provided or configured")
				}
				if externalAwsAuth.AccessKey == "" || externalAwsAuth.SecretKey == "" {
					fmt.Println("INFO: checking if external access keys are stored in aws secret manager")
					s, err := serverless_aws_automation.GetExternalAccessKeySecretIfExists(ctx, awsAuth, externalZeusLambdaAccessKeys)
					if err != nil {
						panic(err)
					}
					externalAwsAuth.AccessKey = s.AccessKey
					externalAwsAuth.SecretKey = s.SecretKey
					if externalAwsAuth.AccessKey == "" || externalAwsAuth.SecretKey == "" {
						fmt.Println("INFO: no external access keys found in aws secret manager, generating and storing new access keys")
						externalAccessKeys := serverless_aws_automation.CreateExternalLambdaUserAccessKeys(ctx, awsAuth)
						externalAwsAuth.AccessKey = externalAccessKeys.AccessKey
						externalAwsAuth.SecretKey = externalAccessKeys.SecretKey
						serverless_aws_automation.AddExternalAccessKeysInAWSSecretManager(ctx, awsAuth, externalZeusLambdaAccessKeys, externalAccessKeys)
					}
				}
				lambdaAccessAuth := aws_aegis_auth.AuthAWS{
					Region:    "us-west-1",
					AccessKey: externalAwsAuth.AccessKey,
					SecretKey: externalAwsAuth.SecretKey,
				}
				fmt.Println("INFO: verifying we can send validator messages to your lambda function")
				serverless_aws_automation.VerifyLambdaSigner(ctx, lambdaAccessAuth, keystoresPath, externalAwsAuth.ServiceURL, ageEncryptionSecretName)
			case "8", "createValidatorServiceRequestOnZeus":
				fmt.Println("INFO: creating zeus validator service request")
				if bearerToken == "" {
					panic("ERROR: bearer token not provided")
				}
				if feeRecipient == "" {
					panic("ERROR: fee recipient not provided")
				}
				err := web3_actions.ValidateToAddress(ctx, feeRecipient)
				if err != nil {
					panic("ERROR: fee recipient is not a valid ethereum address")
				}
				feeRecipient = strings_filter.AddHexPrefix(feeRecipient)
				serverless_aws_automation.ExternalUserRolePolicySetupForLambdaDeployment(ctx, awsAuth)
				if externalAwsAuth.AccessKey == "" || externalAwsAuth.SecretKey == "" {
					fmt.Println("INFO: no credentials provided, generating new aws access key pair")
					keys, err := serverless_aws_automation.GetExternalAccessKeySecret(ctx, awsAuth, externalZeusLambdaAccessKeys)
					if err != nil {
						panic("ERROR: failed to get external access key secret")
					}
					externalAwsAuth.AccessKey = keys.AccessKey
					externalAwsAuth.SecretKey = keys.SecretKey
				}
				if keyGroupName == "" {
					fmt.Println("INFO: no key group name provided, generating a key group name")
					keyGroupName = fmt.Sprintf("keyGroup-%d", time.Now().Unix())
					fmt.Println("INFO: generated key group name: ", keyGroupName)
				}
				sr := hestia_req_types.ServiceRequestWrapper{
					GroupName:         keyGroupName,
					ProtocolNetworkID: hestia_req_types.ProtocolNetworkStringToID(network),
					Enabled:           true,
					ServiceAuth: hestia_req_types.ServiceAuthConfig{
						AuthLamdbaAWS: &hestia_req_types.AuthLamdbaAWS{
							ServiceURL: externalAwsAuth.ServiceURL,
							SecretName: ageEncryptionSecretName,
							AccessKey:  externalAwsAuth.AccessKey,
							SecretKey:  externalAwsAuth.SecretKey,
						},
					}}
				err = sr.ServiceAuth.Validate()
				if err != nil {
					panic(err)
				}
				serverless_aws_automation.CreateHestiaValidatorsServiceRequest(ctx, keystoresPath, sr, bearerToken, feeRecipient)
			case "9", "sendValidatorDeposits":
				fmt.Println("INFO: submitting validator deposits to the network")
				if eth1AddrPrivKey == "" {
					panic("eth1 address private key is required for submitting deposits, you'll also need 32 Eth per validator + gas fees")
				}
				fmt.Println("INFO: depositing validators, using the dir in path relative to the build dir: default is builds/serverless/keystores")
				fmt.Println("INFO: your keystore directory search path is: ", keystoresPath.DirIn)
				acc, err := accounts.ParsePrivateKey(eth1AddrPrivKey)
				if err != nil {
					panic(err)
				}
				w3Client.Account = acc
				builds.ChangeToBuildsDir()
				filter := &strings_filter.FilterOpts{StartsWith: "deposit_data", DoesNotInclude: []string{"keystores.tar.gz.age", ".DS_Store"}}
				keystoresPath.FilterFiles = filter
				dpSlice, err := signing_automation_ethereum.ParseValidatorDepositSliceJSON(ctx, keystoresPath)
				if err != nil {
					panic(err)
				}
				if len(dpSlice) <= 0 {
					panic("no deposit data found in the dir, please check the path and make sure you have generated validator deposits")
				}
				txToBroadcast := make([]*types.Transaction, len(dpSlice))
				for i, d := range dpSlice {
					signedTx, serr := w3Client.SignValidatorDepositTxToBroadcastFromJSON(ctx, d)
					if serr != nil {
						panic(serr)
					}
					txToBroadcast[i] = signedTx
					rx, serr := w3Client.SubmitSignedTxAndReturnTxData(ctx, signedTx)
					if serr != nil {
						panic(serr)
					}
					fmt.Println("tx receipt: ", rx.Hash.String())
				}
			}
		}

	}}

func ForceDirToRootDir() string {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "")
	err := os.Chdir(dir)
	if err != nil {
		panic(err.Error())
	}
	return dir
}

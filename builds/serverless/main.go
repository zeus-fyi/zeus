package main

import (
	"context"
	"errors"
	"fmt"
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
	awsRegion                     = "us-west-1"
	awsAuth                       = aws_aegis_auth.AuthAWS{
		AccountNumber: "",
		Region:        awsRegion,
		AccessKey:     "",
		SecretKey:     "",
	}
	feeRecipient              string
	lambdaFnUrl               string
	bearerToken               string
	keyGroupName              string
	submitValidatorServiceReq bool
	externalAwsAuth           = hestia_req_types.AuthLamdbaAWS{
		ServiceURL: lambdaFnUrl,
		SecretName: "",
		AccessKey:  "",
		SecretKey:  "",
	}
)

func init() {
	viper.AutomaticEnv()
	err := bls_signer.InitEthBLS()
	if err != nil {
		panic(err)
	}

	/*
		########################################
			AWS settings (internal access)
		########################################
	*/
	// aws automation settings for lambda setup
	Cmd.Flags().StringVar(&awsAuth.AccountNumber, "aws-account-number", "", "AWS_ACCOUNT_NUMBER: aws account number")
	Cmd.Flags().StringVar(&awsAuth.AccessKey, "aws-access-key", "", "AWS_ACCESS_KEY: aws access key, which needs permissions to create iam users, roles, policies, secrets, and lambda functions and layers")
	Cmd.Flags().StringVar(&awsAuth.SecretKey, "aws-secret-key", "", "AWS_SECRET_KEY: aws secret key")
	// actions
	Cmd.Flags().BoolVar(&automateSetupOnAWS, "aws-automation-on", false, "AWS_AUTOMATION: automate the entire setup process on aws, requires you provide aws credentials")
	/*
		########################################
			AWS settings (external access)
		########################################
	*/
	// aws service
	Cmd.Flags().StringVar(&externalAwsAuth.ServiceURL, "ext-aws-lambda-url", "", "AWS_LAMBDA_FUNC_URL: bearer token for validator service on zeus")
	Cmd.Flags().StringVar(&externalAwsAuth.AccessKey, "ext-aws-access-key", "", "AWS_EXTERNAL_ACCESS_KEY: bearer token for validator service on zeus")
	Cmd.Flags().StringVar(&externalAwsAuth.SecretKey, "ext-aws-secret-key", "", "AWS_EXTERNAL_SECRET_KEY: bearer token for validator service on zeus")
	Cmd.Flags().StringVar(&externalAwsAuth.SecretName, "ext-aws-age-secret-name", "", "AWS_AGE_DECRYPTION_SECRET_NAME: bearer token for validator service on zeus")
	// secret key generation for serverless
	Cmd.Flags().StringVar(&agePrivKey, "age-private-key", "", "AGE_PRIVKEY: age private key")
	Cmd.Flags().StringVar(&agePubKey, "age-public-key", "", "AGE_PUBKEY: age public key")
	// actions
	Cmd.Flags().BoolVar(&keyGenSecrets, "keygen", true, "KEYGEN_SECRETS: generates secrets for validator encryption and generation")
	/*
		###################################################
			Validator network, keygen, and deposit settings
		###################################################
	*/
	// validator key generation for deposits settings
	Cmd.Flags().StringVar(&nodeURL, "node-url", "https://eth.ephemeral.zeus.fyi", "NODE_URL: beacon for getting network data for validator deposit generation & submitting deposits")
	Cmd.Flags().StringVar(&network, "network", "ephemery", "NETWORK: network to run on (mainnet, goerli, ephemery, etc")
	Cmd.Flags().StringVar(&eth1AddrPrivKey, "eth1-addr-priv-key", "", "ETH1_PRIVATE_KEY: eth1 address private key for submitting deposits")
	// validator secret key generation
	Cmd.Flags().StringVar(&mnemonic, "mnemonic", "", "MNEMONIC_24_WORDS: twenty four word mnemonic to generate keystores")
	Cmd.Flags().StringVar(&hdWalletPassword, "hd-wallet-pw", "", "HD_WALLET_PASSWORD: hd wallet password")
	Cmd.Flags().IntVar(&numKeysToGen, "validator-count", 3, "VALIDATORS_COUNT: number of keys to generate")
	Cmd.Flags().IntVar(&hdOffset, "hd-offset", 0, "HD_OFFSET_VALIDATORS: offset to start generating keys from hd wallet")
	// validator key generation paths
	Cmd.Flags().StringVar(&keystoresPath.DirIn, "keystores-dir-in", "./serverless/keystores", "KEYSTORE_DIR_IN: keystores directory in location (relative to builds dir)")
	Cmd.Flags().StringVar(&keystoresPath.DirOut, "keystores-dir-out", "./serverless/keystores", "KEYSTORE_DIR_OUT: keystores directory out location (relative to builds dir)")
	// actions
	Cmd.Flags().BoolVar(&genValidatorDeposits, "keygen-validators", true, "KEYGEN_VALIDATORS: generates validator deposits, with additional encrypted age keystore")
	Cmd.Flags().BoolVar(&sendDeposits, "submit-deposits", false, "SUBMIT_DEPOSITS: submits validator deposits in keystore directory to the network for activation")
	/*
		###################################################
			Validator service values for Zeus
		###################################################
	*/
	// service on zeus
	Cmd.Flags().StringVar(&bearerToken, "bearer", "", "BEARER: bearer token for validator service on zeus")
	Cmd.Flags().StringVar(&keyGroupName, "key-group-name", "", "KEY_GROUP_NAME: name for validator service group on zeus")
	Cmd.Flags().StringVar(&feeRecipient, "fee-recipient", "", "FEE_RECIPIENT_ADDR: fee recipient address for validators service on zeus")
	Cmd.Flags().BoolVar(&submitValidatorServiceReq, "submit-validator-service-req", false, "SUBMIT_SERVICE_REQUEST: sends a request to zeus to setup a validator service")
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
		if keyGenSecrets {
			if agePubKey == "" || agePrivKey == "" {
				fmt.Println("INFO: no credentials provided, generating new age keypair")
				agePubKey, agePrivKey = age_encryption.GenerateNewKeyPair()
				fmt.Println("agePubKey: ", agePubKey)
				fmt.Println("agePrivKey: ", agePrivKey)
			}
			if hdWalletPassword == "" {
				fmt.Println("INFO: no credentials provided, using default password")
				hdWalletPassword = "password"
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
		}

		if automateSetupOnAWS {
			fmt.Println("INFO: creating internal iam user, role, policies for serverless deployment")
			serverless_aws_automation.InternalUserRolePolicySetupForLambdaDeployment(ctx, awsAuth)
		}
		if genValidatorDeposits {
			fmt.Println("INFO: generating keystores, deposit data, and encypting keystores with age encryption")
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
			err := ethereum_automation_cookbook.GenerateValidatorDepositsAndCreateAgeEncryptedKeystores(ctx, w3Client, vdg, enc, hdWalletPassword)
			if err != nil {
				panic(err)
			}
			if automateSetupOnAWS {
				serverless_aws_automation.AddMnemonicHDWalletSecretInAWSSecretManager(ctx, awsAuth, mnemonicAndHDWalletSecretName, hdWalletPassword, mnemonic)
				serverless_aws_automation.AddAgeEncryptionKeyInAWSSecretManager(ctx, awsAuth, ageEncryptionSecretName, agePubKey, agePrivKey)
			}
		}

		if automateSetupOnAWS {
			serverless_aws_automation.CreateLambdaFunctionKeystoresLayer(ctx, awsAuth)
			lambdaFnUrl = serverless_aws_automation.CreateLambdaFunction(ctx, awsAuth)
		}
		if automateSetupOnAWS {
			if lambdaFnUrl == "" {
				panic("ERROR: lambda function url not provided")
			}
			serverless_aws_automation.VerifyLambdaSigner(ctx, keystoresPath, lambdaFnUrl, ageEncryptionSecretName)
		}

		// Creates service request
		if automateSetupOnAWS {
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
			keys := serverless_aws_automation.CreateExternalLambdaUserAccessKeys(ctx, awsAuth)
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
						ServiceURL: lambdaFnUrl,
						SecretName: ageEncryptionSecretName,
						AccessKey:  keys.AccessKey,
						SecretKey:  keys.SecretKey,
					},
				}}
			err = sr.ServiceAuth.Validate()
			if err != nil {
				panic(err)
			}
			serverless_aws_automation.CreateHestiaValidatorsServiceRequest(ctx, keystoresPath, sr, bearerToken, feeRecipient)
		}

		if sendDeposits {
			if w3Client.Account == nil {
				panic(errors.New("eth1 acount is required for submitting deposits, you'll also need 32 Eth per validator + gas fees"))
			}
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
			filter := &strings_filter.FilterOpts{StartsWith: "deposit_data"}
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
	}}

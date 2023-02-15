package main

import (
	"context"
	"fmt"
	"github.com/gochain/gochain/v4/core/types"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zeus-fyi/gochain/web3/accounts"
	"github.com/zeus-fyi/gochain/web3/web3_actions"
	"github.com/zeus-fyi/zeus/builds"
	ethereum_automation_cookbook "github.com/zeus-fyi/zeus/cookbooks/ethereum/automation"
	signing_automation_ethereum "github.com/zeus-fyi/zeus/pkg/artemis/signing_automation/ethereum"
	age_encryption "github.com/zeus-fyi/zeus/pkg/crypto/age"
	bls_signer "github.com/zeus-fyi/zeus/pkg/crypto/bls"
	aegis_random "github.com/zeus-fyi/zeus/pkg/crypto/random"
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
	keyGen          bool
	sendDeposits    bool
	eth1AddrPrivKey string
)

func init() {
	viper.AutomaticEnv()
	err := bls_signer.InitEthBLS()
	if err != nil {
		panic(err)
	}
	Cmd.Flags().StringVar(&nodeURL, "node-url", "https://eth.ephemeral.zeus.fyi", "beacon for getting network data for validator deposit generation & submitting deposits")
	Cmd.Flags().StringVar(&network, "network", "ephemery", "network to run on (mainnet, goerli, ephemery, etc")
	Cmd.Flags().StringVar(&agePrivKey, "age-private-key", "", "age private key")
	Cmd.Flags().StringVar(&agePubKey, "age-public-key", "", "age public key")
	Cmd.Flags().StringVar(&mnemonic, "mnemonic", "", "twenty four word mnemonic to generate keystores")
	Cmd.Flags().StringVar(&hdWalletPassword, "hd-wallet-pw", "", "hd wallet password")
	Cmd.Flags().IntVar(&numKeysToGen, "num-keys", 3, "number of keys to generate")
	Cmd.Flags().IntVar(&hdOffset, "hd-offset", 0, "offset to start generating keys from hd wallet")
	Cmd.Flags().StringVar(&keystoresPath.DirIn, "keystores-dir-in", "./serverless/keystores", "keystores directory in location (relative to builds dir)")
	Cmd.Flags().StringVar(&keystoresPath.DirOut, "keystores-dir-out", "./serverless/keystores", "keystores directory out location (relative to builds dir)")
	Cmd.Flags().BoolVar(&keyGen, "keygen", true, "generates full keygen procedure")
	Cmd.Flags().BoolVar(&sendDeposits, "submit-deposits", false, "submits validator deposits in keystore directory to the network for activation")
	Cmd.Flags().StringVar(&eth1AddrPrivKey, "eth1-addr-priv-key", "", "eth1 address private key for submitting deposits")
}

// Cmd represents the base command when called without any subcommands
var Cmd = &cobra.Command{
	Use:   "Web3 Middleware",
	Short: "A web3 infra middleware manager for apps on Olympus",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()

		w3Client := signing_automation_ethereum.Web3SignerClient{
			Web3Actions: web3_actions.Web3Actions{
				NodeURL: nodeURL,
				Network: network,
			},
		}
		if keyGen {
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
		}

		if sendDeposits {
			if eth1AddrPrivKey == "" {
				panic("eth1 address private key is required for submitting deposits, you'll also need 32 Eth per validator + gas fees")
			}
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

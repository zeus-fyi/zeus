package ethereum_automation_cookbook

import (
	"context"
	"fmt"

	"github.com/gochain/gochain/v4/core/types"
	"github.com/wealdtech/go-ed25519hd"
	util "github.com/wealdtech/go-eth2-util"
	signing_automation_ethereum "github.com/zeus-fyi/zeus/pkg/artemis/signing_automation/ethereum"
	bls_signer "github.com/zeus-fyi/zeus/pkg/crypto/bls"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	strings_filter "github.com/zeus-fyi/zeus/pkg/utils/strings"
)

var ctx = context.Background()

// you can place your own deposit data values here to check your gen vs your own cli deposit values
var depositDataPath = filepaths.Path{
	PackageName: "",
	DirIn:       "./ethereum/automation/validator_keys/ephemery",
	DirOut:      "./ethereum/automation/validator_keys/ephemery",
	FnIn:        "",
	FnOut:       "",
	Env:         "",
	FilterFiles: strings_filter.FilterOpts{},
}

func (t *EthereumAutomationCookbookTestSuite) TestFullDepositAutomation() {
	offset := 0
	numKeys := 3

	vdg := signing_automation_ethereum.ValidatorDepositGenerationParams{
		Fp:                   depositDataPath,
		Mnemonic:             t.Tc.LocalMnemonic24Words,
		Pw:                   t.Tc.HDWalletPassword,
		ValidatorIndexOffset: offset,
		NumValidators:        numKeys,
	}
	err := vdg.GenerateAndEncryptValidatorKeysFromSeedAndPath(ctx, "ephemery")
	t.Require().Nil(err)

	dpSlice, err := t.Web3SignerClientTestClient.GenerateEphemeryDepositDataWithDefaultWd(ctx, vdg)
	t.Require().Nil(err)

	signing_automation_ethereum.PrintJSONSlice(depositDataPath, dpSlice)

	txToBroadcast := make([]*types.Transaction, len(dpSlice))
	for i, d := range dpSlice {
		signedTx, serr := t.Web3SignerClientTestClient.SignValidatorDepositTxToBroadcast(ctx, d)
		t.Require().Nil(serr)
		txToBroadcast[i] = signedTx

		rx, serr := t.Web3SignerClientTestClient.SubmitSignedTxAndReturnTxData(ctx, signedTx)
		t.Require().Nil(serr)
		t.Assert().NotEmpty(rx)
	}
}

func (t *EthereumAutomationCookbookTestSuite) TestDepositGenVsEthStakingCliDeposit() {
	seed, err := ed25519hd.SeedFromMnemonic(t.Tc.LocalMnemonic24Words, t.Tc.HDWalletPassword)
	t.Require().Nil(err)
	pk, err := util.PrivateKeyFromSeedAndPath(seed, "m/12381/3600/0/0/0")
	t.Require().Nil(err)
	pubkeyFromMnemonicSeed := bls_signer.ConvertBytesToString(pk.PublicKey().Marshal())
	fmt.Println(pubkeyFromMnemonicSeed)

	ea := bls_signer.NewEthSignerBLSFromExistingKey(t.Tc.LocalBLSTestPkey)
	expPubKey := "8a7addbf2857a72736205d861169c643545283a74a1ccb71c95dd2c9652acb89de226ca26d60248c4ef9591d7e010288"
	t.Assert().Equal(expPubKey, ea.PublicKeyString())
	t.Assert().Equal(expPubKey, pubkeyFromMnemonicSeed)

	pkWd, err := util.PrivateKeyFromSeedAndPath(seed, "m/12381/3600/0/0")
	t.Require().Nil(err)

	expWdPubKey := "85bf7eaa189475664975408b4ce6190f59ca6c3f0d97b4eca3b3ace556813a45e6bdcab2d55dd4ec38c43ff94ae1f6ed"
	derivedWdKeyBytes := pkWd.PublicKey().Marshal()
	derivedWdKeyStr := bls_signer.ConvertBytesToString(derivedWdKeyBytes)
	t.Assert().Equal(expWdPubKey, derivedWdKeyStr)

	wd, err := signing_automation_ethereum.ValidateAndReturnBLSPubkeyBytes(derivedWdKeyStr)
	t.Require().Nil(err)

	expPaddedBlsWd := "00a3f653b8d4e23c82652a29c4c235effe4a46e2dc4893d0704e6776b991480d"
	t.Assert().Equal(expPaddedBlsWd, bls_signer.ConvertBytesToString(wd))

	vdg := signing_automation_ethereum.ValidatorDepositGenerationParams{
		Fp:                   depositDataPath,
		Mnemonic:             t.Tc.LocalMnemonic24Words,
		Pw:                   t.Tc.HDWalletPassword,
		ValidatorIndexOffset: 0,
		NumValidators:        1,
	}
	dp, err := t.Web3SignerClientTestClient.GenerateEphemeryDepositDataWithDefaultWd(ctx, vdg)
	t.Require().Nil(err)

	t.Require().Len(dp, 1)
	sv := dp[0].FormatJSON()
	t.Assert().Equal(expPubKey, sv.Pubkey)

	expSig := "8a8cfd56f8750776a880c0e00bddc9ad9a6e225a4f363c47370e2059606fb604c9f3e27723b68a8cf4d5acfd254651660fced47790a35a7b0be7b51e56cf8dd953a21da1ee990590522b5346f04ddbef7443cc7a1bdbfc544a2d5f191715ac75"
	t.Assert().Equal(expSig, sv.Signature)
	t.Assert().Equal(expPaddedBlsWd, sv.WithdrawalCredentials)

	expDepositRoot := "fefcc4c487d9892e61843907098ff35bdb9330dd6c5f62fdfd0e73c4532a763d"
	t.Assert().Equal(expDepositRoot, sv.DepositDataRoot)
}

func (t *EthereumAutomationCookbookTestSuite) getFirstDepositValue() signing_automation_ethereum.Keystore {
	exp := depositDataPath
	exp.DirIn = "./ethereum/automation/validator_keys"
	ks, err := signing_automation_ethereum.ParseKeystoreJSON(ctx, exp)
	t.Require().Nil(err)
	for _, k := range ks {
		return k
	}
	return signing_automation_ethereum.Keystore{}
}

/*

you can use these to print out values in the eth cli py code directly, replace with your own values

cs = BaseChainSetting(NETWORK_NAME="ephemery", GENESIS_FORK_VERSION=bytes.fromhex('1000101b'))
c = Credential(mnemonic="", mnemonic_password="", index=0, amount=32000000000, chain_setting=cs, hex_eth1_withdrawal_address=None)
print("signingkey")
print(c.signing_key_path)
print("signingkeysecret")

print("pubkey")
print(c.deposit_datum_dict['pubkey'].hex())
print("withdrawal_credentials")
print(c.deposit_datum_dict['withdrawal_credentials'].hex())
print("deposit_data_root")
print(c.deposit_datum_dict['deposit_data_root'].hex())
print("signature")
print(c.deposit_datum_dict['signature'].hex())

print("deposit_datum_dict")
print(c.deposit_datum_dict)
*/

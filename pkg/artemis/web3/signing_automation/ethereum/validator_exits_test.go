package signing_automation_ethereum

import (
	"time"

	consensusclient "github.com/attestantio/go-eth2-client"
	"github.com/wealdtech/ethdo/util"
	"github.com/zeus-fyi/zeus/test/configs"
)

func (t *Web3SignerClientTestSuite) TestVoluntaryExits() {
	configs.ForceDirToConfigLocation()
	offset := 0
	numKeys := 3

	vdg := ValidatorDepositGenerationParams{
		Fp:                   depositDataPath,
		Mnemonic:             t.TestMnemonic,
		Pw:                   t.TestHDWalletPassword,
		ValidatorIndexOffset: offset,
		NumValidators:        numKeys,
		Network:              "Goerli",
	}
	t.NodeURL = "http://localhost:5052"

	eth2Client, err := util.ConnectToBeaconNode(ctx, &util.ConnectOpts{
		Address:       t.NodeURL,
		Timeout:       time.Minute,
		AllowInsecure: true,
	})
	genesis, err := eth2Client.(consensusclient.GenesisProvider).Genesis(ctx)
	t.Require().Nil(err)

	for i := 0; i <= 400; i++ {
		signer, err := vdg.GenerateDerivedKeySigner(ctx, i)
		t.Require().Nil(err)
		pubkey := signer.ZeroXPrefixedPublicKeyString()
		vi, err := GetValidatorIndexFromPubkey(ctx, t.NodeURL, pubkey)
		t.Require().Nil(err)
		t.Require().NotEmpty(vi)
		fv, err := GetCurrentForkVersion(ctx, t.NodeURL)
		t.Require().Nil(err)

		exitMsg, err := t.Web3SignerClientTestClient.GenerateVoluntaryExit(ctx, signer, fv, genesis.GenesisValidatorsRoot, vi)
		t.Require().Nil(err)
		t.Require().NotNil(exitMsg)
		err = SubmitVoluntaryExit(ctx, t.NodeURL, exitMsg)
		t.Require().Nil(err)
	}
}

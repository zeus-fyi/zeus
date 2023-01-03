package wallet

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/wealdtech/go-ed25519hd"
	util "github.com/wealdtech/go-eth2-util"
	e2wallet "github.com/wealdtech/go-eth2-wallet"
	keystorev4 "github.com/wealdtech/go-eth2-wallet-encryptor-keystorev4"
	scratch "github.com/wealdtech/go-eth2-wallet-store-scratch"
	bls_signer "github.com/zeus-fyi/zeus/pkg/crypto/bls"
	aegis_random "github.com/zeus-fyi/zeus/pkg/crypto/random"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	strings_filter "github.com/zeus-fyi/zeus/pkg/utils/strings"
	"github.com/zeus-fyi/zeus/test/configs"
)

type WalletTestSuite struct {
	suite.Suite
}

var depositDataPath = filepaths.Path{
	PackageName: "",
	DirOut:      "../mocks/validator_keys",
	FnOut:       "",
	Env:         "",
	FilterFiles: &strings_filter.FilterOpts{},
}

func (s *WalletTestSuite) TestHDWalletCreation() {
	ctx := context.Background()
	m, err := aegis_random.GenerateMnemonic()
	s.Require().Nil(err)
	s.Assert().Len(strings.Fields(m), 24)
	password := "ssdfsdfasdfgdasfrd"

	seed, err := ed25519hd.SeedFromMnemonic(m, password)
	s.Require().Nil(err)
	s.Assert().Len(seed, 64)

	// for a real application you can use this style store to replace the test item: scratch.New()
	// store := filesystem.New(filesystem.WithPassphrase([]byte(password)), filesystem.WithLocation(p.DirOut))
	store := scratch.New()

	w := CreateHDWalletFromMnemonic("testWallet", password, m, store)
	s.Assert().NotEmpty(w)

	s.Assert().Equal("hierarchical deterministic", w.Type())
	s.Assert().Equal("testWallet", w.Name())

	for wallet := range e2wallet.Wallets() {
		fmt.Printf("Found wallet %s\n", wallet.Name())
		for account := range wallet.Accounts(ctx) {

			fmt.Printf("Wallet %s has account %s\n", wallet.Name(), account.Name())
		}
	}
	err = bls_signer.InitEthBLS()
	s.Require().Nil(err)
	path := "m/12381/3600/0/0/0"
	sk, err := util.PrivateKeyFromSeedAndPath(seed, path)
	s.Require().Nil(err)
	fmt.Println(bls_signer.ConvertBytesToString(sk.Marshal()))

	path = "m/12381/3600/1/0/0"
	sk, err = util.PrivateKeyFromSeedAndPath(seed, path)
	s.Require().Nil(err)
	fmt.Println(bls_signer.ConvertBytesToString(sk.Marshal()))

	ks := keystorev4.New()

	enc, err := ks.Encrypt(sk.Marshal(), password)
	s.Require().Nil(err)
	fmt.Println(enc)

	b, err := json.Marshal(enc)
	s.Require().Nil(err)

	configs.ForceDirToConfigLocation()

	slashSplit := strings.Split(path, "/")
	underScoreStr := strings.Join(slashSplit, "_")

	depositDataPath.FnOut = fmt.Sprintf("keystore-ephemeral_%s", underScoreStr)
	err = depositDataPath.WriteToFileOutPath(b)
	s.Require().Nil(err)
}

func TestWalletTestSuite(t *testing.T) {
	suite.Run(t, new(WalletTestSuite))
}

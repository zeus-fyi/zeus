package ethereum_web3signer_actions

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	ethereum_cookbook_test_suite "github.com/zeus-fyi/zeus/cookbooks/ethereum/test"
	validator_cookbooks "github.com/zeus-fyi/zeus/cookbooks/ethereum/validators"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	strings_filter "github.com/zeus-fyi/zeus/pkg/utils/strings"
)

var keystorePath = filepaths.Path{
	PackageName: "",
	DirIn:       "./ethereum/automation/validator_keys/ephemery",
	DirOut:      "./ethereum/automation/validator_keys/ephemery",
	FnIn:        "keystore-ephemery-m_12381_3600_0_0_0.json",
	FnOut:       "",
	Env:         "",
	FilterFiles: strings_filter.FilterOpts{},
}

var ctx = context.Background()

func (t *EthereumWeb3SignerCookbookTestSuite) TestImportWeb3SignerKeysViaKeystoreAPI() {
	kns := validator_cookbooks.ValidatorCloudCtxNs
	w3 := Web3SignerActionsClient{t.ZeusTestClient}
	resp, err := w3.ImportKeystores(ctx, kns, keystorePath, t.Tc.HDWalletPassword)
	t.Require().Nil(err)
	t.Assert().NotEmpty(resp)
}

func (t *EthereumWeb3SignerCookbookTestSuite) TestReadKeystoreFile() {
	k := Web3SignerKeystores{}
	k.ReadKeystoreDirAndAppendPw(ctx, keystorePath, t.Tc.HDWalletPassword)
	t.Assert().NotEmpty(k.Keystores)
}

func (t *EthereumWeb3SignerCookbookTestSuite) TestGetWeb3SignerKeysViaKeystoreAPI() {
	kns := validator_cookbooks.ValidatorCloudCtxNs
	w3 := Web3SignerActionsClient{t.ZeusTestClient}
	resp, err := w3.GetKeystores(ctx, kns)
	t.Require().Nil(err)
	t.Assert().NotEmpty(resp)
}

type EthereumWeb3SignerCookbookTestSuite struct {
	ethereum_cookbook_test_suite.EthereumCookbookTestSuite
}

func TestEthereumWeb3SignerCookbookTestSuite(t *testing.T) {
	suite.Run(t, new(EthereumWeb3SignerCookbookTestSuite))
}

package web3signer_cmds_ai_generated

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type Web3SignerCmdTestSuite struct {
	suite.Suite
}

func (s *Web3SignerCmdTestSuite) TestWeb3SignerCmdStr() {
	cliCmdArgs, err := Web3SignerStd.CreateFieldsForCLI("eth2")
	s.Require().Nil(err)
	fmt.Println(cliCmdArgs)
}

func (s *Web3SignerCmdTestSuite) TestWeb3SignerAPICmdStr() {
	cliCmdArgs, err := Web3SignerAPICmd.CreateFieldsForCLI("eth2")
	s.Require().Nil(err)
	fmt.Println(cliCmdArgs)

}
func TestCWeb3SignerCmdTestSuite(t *testing.T) {
	suite.Run(t, new(Web3SignerCmdTestSuite))
}

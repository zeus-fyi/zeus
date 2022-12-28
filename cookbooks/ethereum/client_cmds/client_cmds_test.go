package client_cmds

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ClientCmdTestSuite struct {
	suite.Suite
}

func (s *ClientCmdTestSuite) TestGethCmdStr() {
	cliCmdArgs := GethEphemeralConfigTemplate.BuildCliCmd()
	fmt.Println(cliCmdArgs)
}

func TestClientCmdTestSuite(t *testing.T) {
	suite.Run(t, new(ClientCmdTestSuite))
}

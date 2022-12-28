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
	cliCmdStr := GethEphemeralConfigTemplate.BuildCliCmd()
	fmt.Println(cliCmdStr)
}

func TestClientCmdTestSuite(t *testing.T) {
	suite.Run(t, new(ClientCmdTestSuite))
}

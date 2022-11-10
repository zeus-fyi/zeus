package test_suites

import (
	"testing"

	"github.com/stretchr/testify/suite"
	test_base "github.com/zeus-fyi/zeus/test"
	"github.com/zeus-fyi/zeus/test/configs"
)

type BaseTestSuite struct {
	suite.Suite
}

func (s *BaseTestSuite) ChangeToTestDir() {
	test_base.ForceDirToTestDirLocation()
}

func (s *BaseTestSuite) TestConfigReader() {
	tc := configs.InitLocalTestConfigs()
	s.Assert().Equal("local", tc.Env)
}

func TestBaseTestSuite(t *testing.T) {
	suite.Run(t, new(BaseTestSuite))
}

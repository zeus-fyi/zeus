package test_suites

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/zeus-fyi/zeus/test/configs"
)

type BaseTestSuiteTester struct {
	suite.Suite
}

func (s *BaseTestSuiteTester) TestConfigReader() {
	tc := configs.InitLocalTestConfigs()
	s.Assert().Equal("local", tc.Env)
}

func TestBaseTestSuiteTester(t *testing.T) {
	suite.Run(t, new(BaseTestSuiteTester))
}

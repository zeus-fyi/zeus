package test_suites

import (
	"testing"

	"github.com/stretchr/testify/suite"
	test_base "github.com/zeus-fyi/zeus/test"
)

type BaseTestSuite struct {
	suite.Suite
}

func (s *BaseTestSuite) ChangeToTestDir() {
	test_base.ForceDirToTestDirLocation()
}

func TestBaseTestSuite(t *testing.T) {
	suite.Run(t, new(BaseTestSuite))
}

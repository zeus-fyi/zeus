package test_suites

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type BaseTestSuite struct {
	suite.Suite
}

func TestBaseTestSuite(t *testing.T) {
	suite.Run(t, new(BaseTestSuite))
}

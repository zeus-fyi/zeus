package mb_json_schemas

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/zeus-fyi/zeus/test/test_suites"
)

type MbJsonSchemasTestSuite struct {
	test_suites.BaseTestSuite
}

func TestMbJsonSchemasTestSuite(t *testing.T) {
	suite.Run(t, new(MbJsonSchemasTestSuite))
}

package hera_client

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
	test_base "github.com/zeus-fyi/zeus/test"
	"github.com/zeus-fyi/zeus/test/configs"
	"github.com/zeus-fyi/zeus/test/test_suites"
)

type HeraClientTestSuite struct {
	test_suites.BaseTestSuite
	HeraTestClient HeraClient
}

var ctx = context.Background()

func (t *HeraClientTestSuite) TestTokenCountApproximate() {
	bytes, err := os.ReadFile("./mocks/hera/tokenizer_example/example.txt")
	t.Require().Nil(err)
	tokenCount := t.HeraTestClient.GetTokenApproximate(string(bytes))
	t.Assert().Equal(61, tokenCount)
	// NOTE open gpt-3 https://beta.openai.com/tokenizer returns 64 tokens as the count
	// there's no opensource transformer for this, so use this + some margin when sending requests
	// 2048 is the max token count for most models, the max size - prompt size, is your limitation on completion
	// tokens
}

func (t *HeraClientTestSuite) SetupTest() {
	// points dir to test/configs
	tc := configs.InitLocalTestConfigs()

	// uses the bearer token from test/configs/config.yaml
	//t.HeraTestClient = NewLocalHeraClient(tc.Bearer)
	t.HeraTestClient = NewDefaultHeraClient(tc.Bearer)
	// points working dir to inside /test
	test_base.ForceDirToTestDirLocation()

	// generates outputs to /test/outputs dir
	// uses inputs from /test/mocks dir
}

func TestHeraClientTestSuite(t *testing.T) {
	suite.Run(t, new(HeraClientTestSuite))
}

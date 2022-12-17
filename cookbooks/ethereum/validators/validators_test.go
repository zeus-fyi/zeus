package validator_cookbooks

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/zeus-fyi/zeus/cookbooks"
	zeus_client "github.com/zeus-fyi/zeus/pkg/zeus/client"
	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_req_types"
	"github.com/zeus-fyi/zeus/test/configs"
	"github.com/zeus-fyi/zeus/test/test_suites"
)

func (t *ValidatorCookbookTestSuite) TestCreateClusterValidatorBase() {
	ctx := context.Background()
	basesInsert := []string{consensusValidatorClientComponentBaseName}
	cc := zeus_req_types.TopologyCreateOrAddBasesToClassesRequest{
		ClassName:      ValidatorClusterClassName,
		ClassBaseNames: basesInsert,
	}
	_, err := t.ZeusTestClient.AddBasesToClass(ctx, cc)
	t.Require().Nil(err)
}

func (t *ValidatorCookbookTestSuite) TestCreateClusterValidatorSkeletonBase() {
	ctx := context.Background()
	cc := zeus_req_types.TopologyCreateOrAddBasesToClassesRequest{
		ClassName:      consensusValidatorClientComponentBaseName,
		ClassBaseNames: []string{validatorSkeletonBaseName},
	}
	_, err := t.ZeusTestClient.AddSkeletonBasesToClass(ctx, cc)
	t.Require().Nil(err)
}

type ValidatorCookbookTestSuite struct {
	test_suites.BaseTestSuite
	ZeusTestClient zeus_client.ZeusClient
}

func (t *ValidatorCookbookTestSuite) SetupTest() {
	// points dir to test/configs
	tc := configs.InitLocalTestConfigs()

	// uses the bearer token from test/configs/config.yaml
	t.ZeusTestClient = zeus_client.NewDefaultZeusClient(tc.Bearer)
	//t.ZeusTestClient = zeus_client.NewZeusClient("http://localhost:9001", tc.Bearer)
	cookbooks.ChangeToCookbookDir()
}

func TestValidatorCookbookTestSuite(t *testing.T) {
	suite.Run(t, new(ValidatorCookbookTestSuite))
}

package topology_workloads

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/zeus-fyi/zeus/test/configs"
	"github.com/zeus-fyi/zeus/test/test_suites"
	v1 "k8s.io/api/core/v1"
)

type WorkloadValidationTestSuite struct {
	test_suites.BaseTestSuite
}

func (t *WorkloadValidationTestSuite) SetupTest() {
	// points dir to test/configs
	tc := configs.InitLocalTestConfigs()
	t.Tc = tc
	// points dir to cookbooks
}
func (t *WorkloadValidationTestSuite) TestValidation() {
	tw := TopologyBaseInfraWorkload{}

	tw.Service = &v1.Service{}

	err := tw.ValidateWorkloads()
	t.Require().Nil(err)
}
func TestWorkloadValidationTestSuite(t *testing.T) {
	suite.Run(t, new(WorkloadValidationTestSuite))
}

package ethereum_automation_cookbook

import (
	"testing"

	"github.com/stretchr/testify/suite"
	ethereum_cookbook_test_suite "github.com/zeus-fyi/zeus/cookbooks/ethereum/test"
)

type EthereumAutomationCookbookTestSuite struct {
	//
	ethereum_cookbook_test_suite.EthereumCookbookTestSuite
}

func TestEthereumAutomationCookbookTestSuite(t *testing.T) {
	suite.Run(t, new(EthereumAutomationCookbookTestSuite))
}

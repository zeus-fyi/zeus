package smart_contracts

import (
	"testing"

	"github.com/stretchr/testify/suite"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	strings_filter "github.com/zeus-fyi/zeus/pkg/utils/strings"
	"github.com/zeus-fyi/zeus/test/test_suites"
)

type SmartContractABIParserTestSuite struct {
	test_suites.BaseTestSuite
}

func (s *SmartContractABIParserTestSuite) SetupTest() {
	s.ChangeToTestDir()
}
func (s *SmartContractABIParserTestSuite) TestABIContractParser() {
	p := filepaths.Path{
		PackageName: "",
		DirIn:       "./mocks/contract_abis",
		DirOut:      "./",
		FnIn:        "eth_deposit_contract_abi.json",
		Env:         "",
		FilterFiles: strings_filter.FilterOpts{},
	}

	b := p.ReadFileInPath()

	sca, err := NewSmartContractABI(b)
	s.Require().Nil(err)
	s.Assert().NotEmpty(sca)

	v, ok := sca.Functions["deposit"]
	s.Require().True(ok)

	s.Assert().True(v.Payable)
	s.Assert().Len(v.Inputs, 4)

	for i, inputs := range v.Inputs {
		switch i {
		case 0:
			s.Assert().Equal("pubkey", inputs.Name)
			s.Assert().Equal("bytes", inputs.Type.String())
		case 1:
			s.Assert().Equal("withdrawal_credentials", inputs.Name)
			s.Assert().Equal("bytes", inputs.Type.String())
		case 2:
			s.Assert().Equal("signature", inputs.Name)
			s.Assert().Equal("bytes", inputs.Type.String())
		case 3:
			s.Assert().Equal("deposit_data_root", inputs.Name)
			s.Assert().Equal("bytes32", inputs.Type.String())
		}
	}
}

func TestSmartContractABIParserTestSuite(t *testing.T) {
	suite.Run(t, new(SmartContractABIParserTestSuite))
}

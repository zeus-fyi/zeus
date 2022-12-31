package aegis_random

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/wealdtech/go-ed25519hd"
)

type MnemonicTestSuite struct {
	suite.Suite
}

func (s *MnemonicTestSuite) TestMnemonicGen() {
	m, err := GenerateMnemonic()
	s.Require().Nil(err)
	fmt.Println(m)
	s.Assert().Len(strings.Fields(m), 24)

	seed, err := ed25519hd.SeedFromMnemonic(m, "alkjdkl35klksmgkolds")
	s.Require().Nil(err)
	s.Assert().Len(seed, 64)
}

func TestMnemonicTestSuite(t *testing.T) {
	suite.Run(t, new(MnemonicTestSuite))
}

package aegis_jwt

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type JWTTestSuite struct {
	suite.Suite
}

func (s *JWTTestSuite) TestJWT() {
	j := NewAegisJWT()
	j.GenerateJwtTokenString()
	s.Assert().NotEmpty(j.Raw)
	fmt.Println(j.Raw)
}

func TestJWTTestSuite(t *testing.T) {
	suite.Run(t, new(JWTTestSuite))
}

package aegis_jwt

import (
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
}

func TestJWTTestSuite(t *testing.T) {
	suite.Run(t, new(JWTTestSuite))
}

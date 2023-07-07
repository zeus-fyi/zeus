package aegis_jwt

import (
	"github.com/golang-jwt/jwt"
	aegis_random "github.com/zeus-fyi/zeus/pkg/aegis/crypto/random"
)

type AegisJWT struct {
	jwt.Token
}

func NewAegisJWT() AegisJWT {
	return AegisJWT{
		Token: jwt.Token{},
	}
}

func (j *AegisJWT) GenerateJwtTokenString() {
	h := aegis_random.Hex(32)
	j.Raw = h
}

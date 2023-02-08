package bls_serverless_signatures

import (
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/suite"
	aegis_inmemdbs "github.com/zeus-fyi/zeus/pkg/aegis/inmemdbs"
	"github.com/zeus-fyi/zeus/test/test_suites"
	"testing"
)

type ServerlessInMemDBsTestSuite struct {
	test_suites.BaseTestSuite
}

func (s *ServerlessInMemDBsTestSuite) TestServerlessSigningFunc() {
	r := resty.New()
	r.SetBaseURL(s.Tc.ServerlessSignerFuncBLS)
	respMsgMap := make(map[string]aegis_inmemdbs.EthereumBLSKeySignatureResponse)
	signedEventResponse := aegis_inmemdbs.EthereumBLSKeySignatureResponses{
		Map: respMsgMap,
	}
	sr := SignatureRequests{
		SecretName:        s.Tc.ServerlessSignerFuncSecretName,
		SignatureRequests: aegis_inmemdbs.EthereumBLSKeySignatureRequests{Map: make(map[string]aegis_inmemdbs.EthereumBLSKeySignatureRequest)},
	}
	key := "0x8a7addbf2857a72736205d861169c643545283a74a1ccb71c95dd2c9652acb89de226ca26d60248c4ef9591d7e010288"
	signMsg := aegis_inmemdbs.EthereumBLSKeySignatureRequest{Message: "This is my first request"}
	sr.SignatureRequests.Map[key] = signMsg
	resp, err := r.R().
		SetResult(&signedEventResponse).
		SetBody(sr).Post("/")
	s.Require().Nil(err)
	s.Require().Equal(200, resp.StatusCode())
	s.Assert().NotEmpty(respMsgMap)
	expSig := "0xa28328b6fb687caddd889160cdebfb1f667e367cb1b62eb58ab4d6dd34b74c1951fe396fdc4b056a5e5a5d3cbcb8f5b1061bb47d3bba5f639953277a395ffbe50978a023b07405986c934a341c49cf61762d2be355fbbdae911028a50754b090"
	s.Assert().Equal(expSig, signedEventResponse.Map[key].Signature)
}

func TestInMemDBsTestSuite(t *testing.T) {
	suite.Run(t, new(ServerlessInMemDBsTestSuite))
}

package bls_serverless_signatures

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/suite"
	bls_serverless_signing "github.com/zeus-fyi/zeus/pkg/aegis/aws/serverless_signing"
	aegis_inmemdbs "github.com/zeus-fyi/zeus/pkg/aegis/inmemdbs"
	bls_signer "github.com/zeus-fyi/zeus/pkg/crypto/bls"
	"github.com/zeus-fyi/zeus/test/test_suites"
	"testing"
)

type ServerlessInMemFSTestSuite struct {
	test_suites.BaseTestSuite
}

var ctx context.Context

func (s *ServerlessInMemFSTestSuite) TestServerlessSigningFunc() {
	r := resty.New()
	r.SetBaseURL(s.Tc.ServerlessSignerFuncBLS)
	respMsgMap := make(map[string]aegis_inmemdbs.EthereumBLSKeySignatureResponse)
	signedEventResponse := aegis_inmemdbs.EthereumBLSKeySignatureResponses{
		Map: respMsgMap,
	}
	sr := bls_serverless_signing.SignatureRequests{
		SecretName:        s.Tc.ServerlessSignerFuncSecretName,
		SignatureRequests: aegis_inmemdbs.EthereumBLSKeySignatureRequests{Map: make(map[string]aegis_inmemdbs.EthereumBLSKeySignatureRequest)},
	}
	key := "0x8a7addbf2857a72736205d861169c643545283a74a1ccb71c95dd2c9652acb89de226ca26d60248c4ef9591d7e010288"

	hexMessage, err := aegis_inmemdbs.RandomHex(10)
	s.Require().Nil(err)
	signMsg := aegis_inmemdbs.EthereumBLSKeySignatureRequest{Message: hexMessage}
	sr.SignatureRequests.Map[key] = signMsg
	resp, err := r.R().
		SetResult(&signedEventResponse).
		SetBody(sr).Post("/")
	s.Require().Nil(err)
	s.Require().Equal(200, resp.StatusCode())
	s.Assert().NotEmpty(respMsgMap)

	err = bls_signer.InitEthBLS()
	s.Require().Nil(err)
	verified, err := signedEventResponse.VerifySignatures(ctx, sr.SignatureRequests, true)
	s.Require().Nil(err)

	s.Assert().Len(verified, 1)
	s.Require().Equal(key, verified[0])
}

func TestServerlessInMemFSTestSuite(t *testing.T) {
	suite.Run(t, new(ServerlessInMemFSTestSuite))
}

package hercules_client

import (
	hercules_ethereum "github.com/zeus-fyi/hercules/api/v1/common/ethereum"
	aegis_inmemdbs "github.com/zeus-fyi/zeus/pkg/aegis/inmemdbs"
)

func (t *HerculesClientTestSuite) TestVerifySignatureBLS() {
	t.TestImportKeystores()

	pubKeyExp := "8a7addbf2857a72736205d861169c643545283a74a1ccb71c95dd2c9652acb89de226ca26d60248c4ef9591d7e010288"
	msg, err := aegis_inmemdbs.RandomHex(10)
	t.Require().Nil(err)
	rr := hercules_ethereum.EthereumBLSKeyVerificationRequests{SignatureRequests: aegis_inmemdbs.EthereumBLSKeySignatureRequests{Map: make(map[string]aegis_inmemdbs.EthereumBLSKeySignatureRequest)}}
	tmp := rr.SignatureRequests.Map[pubKeyExp]
	tmp.Message = msg
	rr.SignatureRequests.Map[pubKeyExp] = tmp
	resp, err := t.HerculesTestClient.VerifyEthSignatureBLS(ctx, rr)
	t.Require().Nil(err)
	t.Require().NotEmpty(resp)

	sigVerify, err := resp.VerifySignatures(ctx, rr.SignatureRequests, true)
	t.Require().Nil(err)
	t.Require().Len(sigVerify, 1)
	t.Assert().Equal("0x"+pubKeyExp, sigVerify[0])
}

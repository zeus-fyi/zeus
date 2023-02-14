package aegis_inmemdbs

import (
	"context"
	"encoding/hex"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
	types "github.com/wealdtech/go-eth2-types/v2"
	"github.com/zeus-fyi/gochain/web3/accounts"
	bls_signer "github.com/zeus-fyi/zeus/pkg/crypto/bls"
	"github.com/zeus-fyi/zeus/pkg/crypto/ecdsa"
	strings_filter "github.com/zeus-fyi/zeus/pkg/utils/strings"
)

type InMemDBsTestSuite struct {
	suite.Suite
}

var ctx = context.Background()

func (s *InMemDBsTestSuite) SetupTest() {
	InitEcdsaAccountsDB()
	InitValidatorDB()
}
func (s *InMemDBsTestSuite) TestEcdsaInMemDB() {
	i := 0
	numAccounts := 3
	insertAccountSlice := make([]ecdsa.Account, numAccounts)
	for i < numAccounts {
		acc, err := accounts.CreateAccount()
		s.Require().Nil(err)
		s.Assert().NotEmpty(acc)
		insertAccountSlice[i] = ecdsa.Account{Account: acc}
		i++
	}
	InsertEcdsaAccounts(ctx, insertAccountSlice)
	signer := insertAccountSlice[0]
	fetchedSigner := ReadOnlyEcdsaAccountFromInMemDb(ctx, signer)
	s.Assert().Equal(signer, fetchedSigner)
}

func (s *InMemDBsTestSuite) TestBatchValidatorsInMemDB() {
	i := 0
	numAccounts := 3
	insertAccountSlice := make([]Validator, numAccounts)
	signReqs := make(map[string]EthereumBLSKeySignatureRequest)
	batchSignReqs := EthereumBLSKeySignatureRequests{Map: signReqs}

	for i < numAccounts {
		acc := bls_signer.NewEthBLSAccount()
		insertAccountSlice[i] = NewValidator(acc)
		hexStr, err := RandomHex(10)
		s.Require().Nil(err)
		batchSignReqs.Map[acc.PublicKeyString()] = EthereumBLSKeySignatureRequest{hexStr}
		i++
	}
	InsertValidatorsInMemDb(ctx, insertAccountSlice)

	resp, err := SignValidatorMessagesFromInMemDb(ctx, batchSignReqs)
	s.Require().Nil(err)
	s.Assert().NotEmpty(resp.Map)

	s.Require().Len(resp.Map, numAccounts)

	for k, v := range resp.Map {
		signReqMessage := batchSignReqs.Map[k].Message
		signedResp := v.Signature

		signedResp = strings.TrimLeft(signedResp, "0x")
		data, derr := hex.DecodeString(strings_filter.Trim0xPrefix(signedResp))
		s.Require().Nil(derr)

		sig, serr := types.BLSSignatureFromBytes(data)
		s.Require().Nil(serr)

		pubkeyHexStr, herr := hex.DecodeString(strings_filter.Trim0xPrefix(k))
		s.Require().Nil(herr)

		pubkey, perr := types.BLSPublicKeyFromBytes(pubkeyHexStr)
		s.Require().Nil(perr)
		s.Require().True(sig.Verify([]byte(signReqMessage), pubkey))
	}

	verifiedKeys, err := resp.VerifySignatures(ctx, batchSignReqs)
	s.Require().Nil(err)
	s.Require().Len(batchSignReqs.Map, len(verifiedKeys))

	fmt.Println(verifiedKeys)
}

func (s *InMemDBsTestSuite) TestValidatorsInMemDB() {
	i := 0
	numAccounts := 3
	insertAccountSlice := make([]Validator, numAccounts)
	for i < numAccounts {
		acc := bls_signer.NewEthBLSAccount()
		insertAccountSlice[i] = NewValidator(acc)
		i++
	}
	InsertValidatorsInMemDb(ctx, insertAccountSlice)
	for i < numAccounts {
		v := insertAccountSlice[i]
		inMemDBVal := ReadOnlyValidatorFromInMemDb(ctx, v.PublicKeyString())
		s.Assert().Equal(v, inMemDBVal)
	}
}

func TestInMemDBsTestSuite(t *testing.T) {
	suite.Run(t, new(InMemDBsTestSuite))
}

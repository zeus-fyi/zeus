package aegis_inmemdbs

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/zeus-fyi/gochain/web3/accounts"
	bls_signer "github.com/zeus-fyi/zeus/pkg/crypto/bls"
	"github.com/zeus-fyi/zeus/pkg/crypto/ecdsa"
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

func (s *InMemDBsTestSuite) TestValidatorsInMemDB() {
	i := 0
	numAccounts := 3
	insertAccountSlice := make([]Validator, numAccounts)
	for i < numAccounts {
		key := bls_signer.NewKeyBLS()
		insertAccountSlice[i] = NewValidator(i, key)
		i++
	}

	InsertValidatorsInMemDb(ctx, insertAccountSlice)
	v := insertAccountSlice[0]
	fetchedValidator := ReadOnlyValidatorFromInMemDb(ctx, v)
	s.Require().NotEmpty(fetchedValidator)
	s.Assert().Equal(v.Index, fetchedValidator.Index)

	msg := []byte("hello foo")
	sig := v.Sign(msg)
	s.Assert().True(v.Verify(*sig, msg))
	msgUnauthorized := []byte("hello bar")
	s.Assert().False(v.Verify(*sig, msgUnauthorized))
	// tests that the fetched validator matches
	s.Assert().True(fetchedValidator.Verify(*sig, msg))
	s.Assert().False(fetchedValidator.Verify(*sig, msgUnauthorized))
}
func TestInMemDBsTestSuite(t *testing.T) {
	suite.Run(t, new(InMemDBsTestSuite))
}

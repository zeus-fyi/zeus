package age_encryption

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/zeus-fyi/zeus/test/test_suites"
)

type AgeEncryptionTestSuite struct {
	test_suites.BaseTestSuite
	Age Age
}

func (s *AgeEncryptionTestSuite) SetupTest() {
}

func (s *AgeEncryptionTestSuite) TestGenerateNewAge() {
	a := GenerateNewAgeCredentials()
	s.Assert().NotEmpty(a)
}

func (s *AgeEncryptionTestSuite) TestEncryption() {
	GenerateNewAgeCredentials()
}

// use age-keygen -o private_key.txt to create a pubkey/private key pair for here
func (s *AgeEncryptionTestSuite) TestDecryption() {
}

func (s *AgeEncryptionTestSuite) TestDecryptionToInMemFs() {
}

func TestAgeEncryptionTestSuite(t *testing.T) {
	suite.Run(t, new(AgeEncryptionTestSuite))
}

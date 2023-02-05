package serverless_inmemdb

import (
	"context"
	"github.com/stretchr/testify/suite"
	aegis_inmemdbs "github.com/zeus-fyi/zeus/pkg/aegis/inmemdbs"
	age_encryption "github.com/zeus-fyi/zeus/pkg/crypto/age"
	"github.com/zeus-fyi/zeus/test/test_suites"
	"testing"
)

type ServerlessInMemDBsTestSuite struct {
	test_suites.BaseTestSuite
}

var ctx = context.Background()

// TestInMemDBsImportPrep is how you would probably want to use a lighter weight keystore method for better performance
// it's much more efficient to just encrypt the batch of keystores, vs having each one having decryption overhead
// use the age encryption package to encrypt the batch of keystores
func (s *ServerlessInMemDBsTestSuite) TestInMemDBsImportPrep() {
	// change to actual path
	KeystorePath.DirIn = s.Tc.EncryptedKeystoresPath
	vs := aegis_inmemdbs.DecryptedValidators{HDPassword: s.Tc.HDWalletPassword, DecryptPath: s.Tc.UnencryptedKeystoresPath}
	err := KeystorePath.WalkAndApplyFuncToFileType(".json", vs.ReadValidatorFromKeystoreAndGenerateRawKeyfiles)
	s.Require().Nil(err)
}

func (s *ServerlessInMemDBsTestSuite) TestLightweightGzipAgeEncrypt() {
	KeystorePath.FnIn = "keystores"
	KeystorePath.DirIn = s.Tc.UnencryptedKeystoresPath
	KeystorePath.DirOut = s.Tc.EncryptedKeystoresPath

	enc := age_encryption.NewAge(s.Tc.AgePrivKey, s.Tc.AgePubKey)
	err := enc.GzipAndEncrypt(&KeystorePath)
	s.Require().Nil(err)

	// creates a keystores.tar.gz.age file
}

// TODO
func (s *ServerlessInMemDBsTestSuite) TestLightweightGzipAgeDecrypt() {

}

func TestInMemDBsTestSuite(t *testing.T) {
	suite.Run(t, new(ServerlessInMemDBsTestSuite))
}

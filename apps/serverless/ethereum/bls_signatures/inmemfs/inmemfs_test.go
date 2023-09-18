package serverless_inmemdb

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	aegis_inmemdbs "github.com/zeus-fyi/zeus/pkg/aegis/inmemdbs"
	age_encryption "github.com/zeus-fyi/zeus/pkg/crypto/age"
	bls_signer "github.com/zeus-fyi/zeus/pkg/crypto/bls"
	"github.com/zeus-fyi/zeus/test/test_suites"
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

func (s *ServerlessInMemDBsTestSuite) TestDecAndInsertIntoInMemDB() {
	KeystorePath.FnIn = "keystores.tar.gz.age"
	KeystorePath.DirIn = s.Tc.EncryptedKeystoresPath

	enc := age_encryption.NewAge(s.Tc.AgePrivKey, s.Tc.AgePubKey)
	err := ImportIntoInMemFs(ctx, enc)
	s.Require().Nil(err)
	pubkey := "0x8a0a0acdd7e062d2a99548e578790d03fd73ea011775866fd51d966f3245d61eb38f9b4a3229beada55c2df03197f714"

	KeystorePath.DirIn = "keystores"
	KeystorePath.FnIn = pubkey
	b, err := InMemFs.ReadFile(KeystorePath.FileInPath())
	s.Assert().NotEmpty(b)

	acc := bls_signer.NewEthSignerBLSFromExistingKey(string(b))
	s.Require().Equal(pubkey, acc.ZeroXPrefixedPublicKeyString())
}

func (s *ServerlessInMemDBsTestSuite) TestDecryptionToInMemFS() {
	KeystorePath.FnIn = "keystores.tar.gz.age"
	KeystorePath.DirIn = s.Tc.EncryptedKeystoresPath

	enc := age_encryption.NewAge(s.Tc.AgePrivKey, s.Tc.AgePubKey)

	err := enc.DecryptAndUnGzipToInMemFs(&KeystorePath, InMemFs, "./keystores")
	s.Require().Nil(err)
	// creates a keystores.tar.gz.age file
}

func (s *ServerlessInMemDBsTestSuite) TestLightweightGzipAgeEncrypt() {
	KeystorePath.FnIn = "keystores"
	KeystorePath.DirIn = s.Tc.UnencryptedKeystoresPath
	KeystorePath.DirOut = s.Tc.EncryptedKeystoresPath

	enc := age_encryption.NewAge(s.Tc.AgePrivKey, s.Tc.AgePubKey)
	err := enc.GzipAndEncrypt(&KeystorePath)
	s.Require().Nil(err)
}

func (s *ServerlessInMemDBsTestSuite) TestLightweightGzipAgeDecrypt() {
	enc := age_encryption.NewAge(s.Tc.AgePrivKey, s.Tc.AgePubKey)
	KeystorePath.FnIn = "keystores"
	KeystorePath.DirIn = s.Tc.EncryptedKeystoresPath
	KeystorePath.DirOut = s.Tc.UnencryptedKeystoresPath
	err := enc.UnGzipAndDecrypt(&KeystorePath)
	s.Require().Nil(err)
}

func TestInMemDBsTestSuite(t *testing.T) {
	suite.Run(t, new(ServerlessInMemDBsTestSuite))
}

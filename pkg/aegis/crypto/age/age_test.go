package age_encryption

import (
	"context"
	"encoding/json"
	"os"
	"path"
	"runtime"
	"testing"

	"github.com/stretchr/testify/suite"
	aws_aegis_auth "github.com/zeus-fyi/zeus/pkg/aegis/aws/auth"
	aegis_aws_secretmanager "github.com/zeus-fyi/zeus/pkg/aegis/aws/secretmanager"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	"github.com/zeus-fyi/zeus/test/configs"
	"github.com/zeus-fyi/zeus/test/test_suites"
)

type AgeEncryptionTestSuite struct {
	test_suites.BaseTestSuite
	Age Age
}

var ctx = context.Background()

func (s *AgeEncryptionTestSuite) SetupTest() {
	s.Tc = configs.InitLocalTestConfigs()

	region := "us-west-1"
	a := aws_aegis_auth.AuthAWS{
		AccessKey: s.Tc.AccessKeyAWS,
		SecretKey: s.Tc.SecretKeyAWS,
		Region:    region,
	}
	sm, err := aegis_aws_secretmanager.InitSecretsManager(ctx, a)
	s.Require().Nil(err)
	s.Require().NotNil(sm)

	secretInfo := aegis_aws_secretmanager.SecretInfo{
		Region: region,
		Name:   "ageEncryptionKeyEphemery",
	}
	b, err := sm.GetSecretBinary(ctx, secretInfo)
	s.Require().Nil(err)

	m := make(map[string]any)
	err = json.Unmarshal(b, &m)
	s.Require().Nil(err)

	for pubkey, privkey := range m {
		s.Age = NewAge(privkey.(string), pubkey)
	}
}

func (s *AgeEncryptionTestSuite) TestGenerateNewAge() {
	a := GenerateNewAgeCredentials()
	s.Assert().NotEmpty(a)
}

func (s *AgeEncryptionTestSuite) TestEncryption() {
	forceDirToAgeTestSuite()
	p := filepaths.Path{DirIn: "./", FnIn: "tmp.txt", DirOut: "./"}
	err := s.Age.Encrypt(&p)
	s.Require().Nil(err)
}

// use age-keygen -o private_key.txt to create a pubkey/private key pair for here
func (s *AgeEncryptionTestSuite) TestDecryption() {
	forceDirToAgeTestSuite()
	p := filepaths.Path{DirIn: "./", FnIn: "keystores.tar.gz.age", DirOut: "./", FnOut: "keystores.tar.gz"}

	err := s.Age.Decrypt(&p)
	s.Require().Nil(err)
}

func (s *AgeEncryptionTestSuite) TestDecryptionToInMemFs() {
}

func TestAgeEncryptionTestSuite(t *testing.T) {
	suite.Run(t, new(AgeEncryptionTestSuite))
}

func forceDirToAgeTestSuite() string {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "")
	err := os.Chdir(dir)
	if err != nil {
		panic(err.Error())
	}
	return dir
}

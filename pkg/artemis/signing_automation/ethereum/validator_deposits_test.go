package signing_automation_ethereum

import (
	"context"

	spec "github.com/attestantio/go-eth2-client/spec/phase0"
)

var ctx = context.Background()

func (t *Web3SignerClientTestSuite) TestValidatorABI() {
	ForceDirToEthSigningDirLocation()
	f, err := ABIOpenFile(ctx, validatorAbiFileLocation)
	t.Require().Nil(err)
	t.Require().NotEmpty(f)

	depositExists := false
	for _, mn := range f.Methods {
		if mn.Name == validatorDepositMethodName {
			depositExists = true
		}
	}
	t.Require().True(depositExists)
}

func (t *Web3SignerClientTestSuite) TestFetchEphemeralForkVersion() {
	versionByteArr, err := GetEphemeralForkVersion(ctx)
	t.Require().Nil(err)
	t.Require().NotEmpty(versionByteArr)

	var expectedVersion spec.Version
	copy(expectedVersion[:], []byte{0x10, 0x00, 0x10, 0x1b})
	t.Assert().Equal(expectedVersion, *versionByteArr)
}

func (t *Web3SignerClientTestSuite) TestFetchGoerliForkVersion() {
	versionByteArr, err := GetForkVersion(ctx, t.NodeURL)
	t.Require().Nil(err)
	t.Require().NotEmpty(versionByteArr)
	t.Require().Nil(err)

	var expectedVersion spec.Version
	copy(expectedVersion[:], []byte{0x00, 0x00, 0x10, 0x20})
	t.Assert().Equal(expectedVersion, *versionByteArr)
}

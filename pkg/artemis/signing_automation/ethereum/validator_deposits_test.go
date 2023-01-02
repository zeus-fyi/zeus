package signing_automation_ethereum

import (
	"context"
	"encoding/hex"
	"strings"
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
	forkVersion, err := hex.DecodeString(strings.TrimPrefix("0x1000101b", "0x"))
	t.Require().Nil(err)
	t.Assert().Equal(forkVersion, []byte{0x10, 0x00, 0x10, 0x1b})
	t.Assert().Equal(versionByteArr, []byte{0x10, 0x00, 0x10, 0x1b})
}

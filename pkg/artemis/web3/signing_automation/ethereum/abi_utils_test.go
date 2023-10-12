package signing_automation_ethereum

func (t *Web3SignerClientTestSuite) TestReaderABI() {
	ForceDirToEthSigningDirLocation()
	f, err := ABIOpenFile(ctx, validatorAbiFileLocation)
	t.Require().Nil(err)
	t.Require().NotEmpty(f)
}

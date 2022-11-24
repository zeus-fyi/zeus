package hercules_client

func (t *HerculesClientTestSuite) TestHostDiskInfo() {
	resp, err := t.HerculesTestClient.GetHostDiskInfo(ctx)
	t.Assert().Nil(err)
	t.Assert().NotNil(resp)
}

func (t *HerculesClientTestSuite) TestHostMemInfo() {
	resp, err := t.HerculesTestClient.GetHostMemInfo(ctx)
	t.Assert().Nil(err)
	t.Assert().NotNil(resp)
}

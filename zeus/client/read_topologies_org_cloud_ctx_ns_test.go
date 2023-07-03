package zeus_client

func (t *ZeusClientTestSuite) TestReadTopologiesOrgCloudCtxNs() {
	resp, err := t.ZeusTestClient.ReadTopologiesOrgCloudCtxNs(ctx)
	t.Require().Nil(err)
	t.Assert().NotEmpty(resp)
}

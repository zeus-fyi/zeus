package zeus_client

func (t *ZeusClientTestSuite) TestReadTopologies() {
	resp, err := t.ZeusTestClient.ReadTopologies(ctx)
	t.Require().Nil(err)
	t.Assert().NotEmpty(resp)
}

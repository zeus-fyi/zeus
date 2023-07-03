package zeus_client

func (t *ZeusClientTestSuite) TestDestroyDeploy() {
	resp, err := t.ZeusTestClient.DestroyDeploy(ctx, deployKnsReq)
	t.Require().Nil(err)
	t.Assert().NotEmpty(resp)
}

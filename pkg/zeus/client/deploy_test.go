package zeus_client

// TestChartUpload will return a topology id associated with this workload
func (t *ZeusClientTestSuite) TestDeploy() {
	resp, err := t.ZeusTestClient.Deploy(ctx, deployKnsReq)
	t.Require().Nil(err)
	t.Assert().NotEmpty(resp)
}

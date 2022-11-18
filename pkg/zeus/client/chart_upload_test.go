package zeus_client

// TestChartUpload will return a topology id associated with this workload
func (t *ZeusClientTestSuite) TestChartUpload() int {
	resp, err := t.ZeusTestClient.UploadChart(ctx, demoChartPath, uploadChart)
	t.Require().Nil(err)
	t.Assert().NotZero(resp.ID)

	return resp.ID
}

package zeus_client

func (t *ZeusClientTestSuite) TestChartUploadAndRead() {
	topologyID := t.TestChartUpload()
	deployKnsReq.TopologyID = topologyID
	t.TestReadDemoChart()
}

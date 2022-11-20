package zeus_client

// TestDeployReplace will replace the components at this location, but does not change the underlying topology
// definitions. In other words, this is a localized change.
func (t *ZeusClientTestSuite) TestDeployReplace() {
	demoChartPath.DirIn += "/alt_config"
	resp, err := t.ZeusTestClient.DeployReplace(ctx, demoChartPath, deployKnsReq)
	t.Require().Nil(err)
	t.Assert().NotEmpty(resp)
}

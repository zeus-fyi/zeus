package zeus_client

import "github.com/zeus-fyi/zeus/pkg/zeus/client/req_types"

func (t *ZeusClientTestSuite) TestReadDemoChart() {
	tar := req_types.TopologyRequest{TopologyID: deployKnsReq.TopologyID}
	resp, err := t.ZeusTestClient.ReadChart(ctx, tar)
	t.Require().Nil(err)
	t.Assert().NotEmpty(resp)

	// prints the chart output for inspection
	err = resp.PrintWorkload(demoChartPath)
	t.Require().Nil(err)
}

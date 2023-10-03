package sui_actions

import (
	"fmt"
	"strings"

	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_common_types"
)

func (t *SuiActionsCookbookTestSuite) TestGetMetrics() {
	cloudCtxNs := zeus_common_types.CloudCtxNs{
		CloudProvider: "do",
		Region:        "nyc1",
		Context:       "do-nyc1-do-nyc1-zeus-demo",
		Namespace:     "sui-testnet-do-5ab4e8aa",
	}
	t.su.PrintResp = false
	resp, err := t.su.GetMetrics(ctx, cloudCtxNs)
	t.NoError(err)
	for _, v := range resp.ReplyBodies {
		rs := strings.NewReader(string(v))
		fmt.Println(rs)
	}
}

package sui_actions

import (
	"fmt"
	"strings"

	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_common_types"
)

func (t *SuiActionsCookbookTestSuite) TestGetMetrics() {
	cloudCtxNs := zeus_common_types.CloudCtxNs{
		CloudProvider: "aws",
		Region:        "us-west-1",
		Context:       "zeus-us-west-1",
		Namespace:     "sui-3f454d91",
	}
	t.su.PrintResp = false
	resp, err := t.su.GetMetrics(ctx, cloudCtxNs)
	t.NoError(err)
	for _, v := range resp.ReplyBodies {
		rs := strings.NewReader(string(v))
		fmt.Println(rs)
	}
}

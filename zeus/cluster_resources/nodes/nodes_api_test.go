package nodes

import "github.com/zeus-fyi/zeus/zeus/cluster_resources/on_demand_resources"

func (t *NodesConfigTestSuite) TestNodeSearch() {
	searchParams := NodeSearchParams{
		CloudProviderRegions: on_demand_resources.CloudProviderRegions,
		ResourceMinMax: ResourceMinMax{
			Max: ResourceAggregate{
				MonthlyPrice: 500,
				MemRequests:  "10Gi",
				CpuRequests:  "20",
			},
			Min: ResourceAggregate{
				MonthlyPrice: 100,
				MemRequests:  "20Gi",
				CpuRequests:  "10",
			},
		},
	}
	resp, err := GetNodes(ctx, t.ZeusLocalTestClient, searchParams)
	t.NoError(err)
	t.NotNil(resp)
}

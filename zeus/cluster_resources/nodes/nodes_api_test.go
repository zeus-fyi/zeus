package nodes

func (t *NodesConfigTestSuite) TestNodeSearch() {
	searchParams := NodeSearchParams{
		ResourceMinMax: ResourceMinMax{
			Max: ResourceAggregate{
				Price:       500,
				MemRequests: "10Gi",
				CpuRequests: "20",
			},
			Min: ResourceAggregate{
				Price:       100,
				MemRequests: "20Gi",
				CpuRequests: "10",
			},
		},
	}
	resp, err := GetNodes(ctx, t.ZeusLocalTestClient, searchParams)
	t.NoError(err)
	t.NotNil(resp)
}

package hercules_client

var clientName = "lighthouse"

func (t *HerculesClientTestSuite) TestResume() {
	rr := RoutineRequest{ClientName: clientName}
	err := t.HerculesTestClient.Resume(ctx, rr)
	t.Assert().Nil(err)
}

func (t *HerculesClientTestSuite) TestSuspend() {
	rr := RoutineRequest{ClientName: clientName}
	err := t.HerculesTestClient.Suspend(ctx, rr)
	t.Assert().Nil(err)
}

func (t *HerculesClientTestSuite) TestKill() {
	rr := RoutineRequest{ClientName: clientName}
	err := t.HerculesTestClient.Kill(ctx, rr)
	t.Assert().Nil(err)
}

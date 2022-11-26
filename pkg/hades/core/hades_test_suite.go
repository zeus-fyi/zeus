package hades_core

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type HadesBaseTestSuite struct {
	suite.Suite
	H Hades
}

func (h *HadesBaseTestSuite) SetupTest() {
	h.H = Hades{}
	h.H.ConnectToK8s()
}

func TestHadesBaseTestSuite(t *testing.T) {
	suite.Run(t, new(HadesBaseTestSuite))
}

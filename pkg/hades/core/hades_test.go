package hades_core

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type HadesTestSuite struct {
	HadesBaseTestSuite
}

func (h *HadesBaseTestSuite) TestHades() {
	kctx, err := h.H.GetContexts()
	h.Nil(err)
	h.Greater(len(kctx), 0)
	fmt.Println(kctx)
}

func TestHadesTestSuite(t *testing.T) {
	suite.Run(t, new(HadesTestSuite))
}

package hades_api_test

import (
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
	hades_core "github.com/zeus-fyi/zeus/pkg/hades/core"
)

type HadesApiBaseTestSuite struct {
	suite.Suite
	E  *echo.Echo
	Eg *echo.Group
	hades_core.Hades
}

func (t *HadesApiBaseTestSuite) SetupTest() {
	t.Hades.ConnectToK8s()
	t.E = echo.New()
	t.Eg = t.E.Group("/v1")
}

func TestHadesApiBaseTestSuite(t *testing.T) {
	suite.Run(t, new(HadesApiBaseTestSuite))
}

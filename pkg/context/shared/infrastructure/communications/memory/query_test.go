package memory_test

import (
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/roles"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/suite"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/communications"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/communications/memory"
)

type QueryBusTestSuite struct {
	communications.QueryBusSuite
}

func (s *QueryBusTestSuite) SetupSuite() {
	s.QueryBusSuite.Handler = new(communications.QueryHandlerMock)

	s.QueryBusSuite.SUT = &memory.QueryBus{
		Handlers: make(map[string]roles.QueryHandler),
	}
}

func TestIntegrationQueryBusSuite(t *testing.T) {
	suite.Run(t, new(QueryBusTestSuite))
}

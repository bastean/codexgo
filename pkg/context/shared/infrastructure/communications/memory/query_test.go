package memory_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/roles"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/communications"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/communications/memory"
)

type QueryBusTestSuite struct {
	communications.QueryBusSuite
}

func (s *QueryBusTestSuite) SetupTest() {
	s.QueryBusSuite.Handler = new(communications.QueryHandlerMock)

	s.QueryBusSuite.SUT = &memory.QueryBus{
		Handlers: make(map[messages.Key]roles.QueryHandler),
	}
}

func TestIntegrationQueryBusSuite(t *testing.T) {
	suite.Run(t, new(QueryBusTestSuite))
}

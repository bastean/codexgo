package memory_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/roles"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/communications"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/communications/memory"
)

type CommandBusTestSuite struct {
	communications.CommandBusSuite
}

func (s *CommandBusTestSuite) SetupTest() {
	s.CommandBusSuite.Handler = new(communications.CommandHandlerMock)

	s.CommandBusSuite.SUT = &memory.CommandBus{
		Handlers: make(map[messages.Key]roles.CommandHandler),
	}
}

func TestIntegrationCommandBusSuite(t *testing.T) {
	suite.Run(t, new(CommandBusTestSuite))
}

package memory_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/roles"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/communications"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/communications/memory"
)

type EventBusTestSuite struct {
	communications.EventBusSuite
}

func (s *EventBusTestSuite) SetupTest() {
	s.EventBusSuite.Event = messages.Random()

	s.EventBusSuite.Consumer = new(communications.EventConsumerMock)

	s.EventBusSuite.SUT = &memory.EventBus{
		Consumers: make(map[messages.Key][]roles.EventConsumer),
	}
}

func TestIntegrationEventBusSuite(t *testing.T) {
	suite.Run(t, new(EventBusTestSuite))
}

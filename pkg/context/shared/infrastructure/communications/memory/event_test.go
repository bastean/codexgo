package memory_test

import (
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/events"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/communications"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/communications/memory"
	"github.com/stretchr/testify/suite"
)

type EventBusTestSuite struct {
	suite.Suite
	sut      events.Bus
	consumer *communications.EventConsumerMock
}

func (s *EventBusTestSuite) SetupTest() {
	s.consumer = new(communications.EventConsumerMock)

	s.sut = &memory.EventBus{
		Consumers: make(map[events.Key][]events.Consumer),
	}
}

func (s *EventBusTestSuite) TestSubscribe() {
	s.NoError(s.sut.Subscribe(messages.Random[events.Event]().Key, s.consumer))
}

func (s *EventBusTestSuite) TestPublish() {
	event := messages.Random[events.Event]()

	s.NoError(s.sut.Subscribe(event.Key, s.consumer))

	s.consumer.Mock.On("On", event)

	s.NoError(s.sut.Publish(event))

	s.consumer.AssertExpectations(s.T())
}

func (s *EventBusTestSuite) TestPublishErrMissingConsumer() {
	event := messages.Random[events.Event]()

	err := s.sut.Publish(event)

	var actual *errors.Internal

	s.ErrorAs(err, &actual)

	expected := &errors.Internal{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "Publish",
		What:  "Failure to execute a Event without a Consumer",
		Why: errors.Meta{
			"Event": event.Key,
		},
	}}

	s.EqualError(expected, actual.Error())
}

func TestIntegrationEventBusSuite(t *testing.T) {
	suite.Run(t, new(EventBusTestSuite))
}

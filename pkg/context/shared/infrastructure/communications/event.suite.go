package communications

import (
	"time"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/roles"
)

type EventBusSuite struct {
	suite.Suite
	SUT      roles.EventBus
	Consumer *EventConsumerMock
	Event    *messages.Message
}

func (s *EventBusSuite) TestSubscribe() {
	s.NoError(s.SUT.Subscribe(s.Event.Key, s.Consumer))
}

func (s *EventBusSuite) TestPublish() {
	s.Consumer.Mock.On("On", s.Event)

	s.NoError(s.SUT.Subscribe(s.Event.Key, s.Consumer))

	s.NoError(s.SUT.Publish(s.Event))

	s.Eventually(func() bool {
		return s.Consumer.Mock.AssertExpectations(s.T())
	}, 10*time.Second, 30*time.Millisecond)
}

func (s *EventBusSuite) TestPublishErrMissingConsumer() {
	event := messages.Random()

	err := s.SUT.Publish(event)

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

	s.Equal(expected, actual)
}

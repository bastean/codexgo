package communications

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/roles"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/suite"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/time"
)

type EventBusSuite struct {
	suite.Default
	SUT      roles.EventBus
	Consumer *EventConsumerMock
	Event    *messages.Message
}

func (s *EventBusSuite) SetupTest() {
	s.Event = messages.Mother().MessageValid()
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
	s.Event = messages.Mother().MessageValid()

	err := s.SUT.Publish(s.Event)

	var actual *errors.Internal

	s.ErrorAs(err, &actual)

	s.Contains(actual.Where, "Publish")

	expected := &errors.Internal{Bubble: &errors.Bubble{
		ID:    actual.ID,
		When:  actual.When,
		Where: actual.Where,
		What:  "Failure to execute a Event without a Consumer",
		Why: errors.Meta{
			"ID":  s.Event.ID.Value(),
			"Key": s.Event.Key.Value(),
		},
	}}

	s.Equal(expected, actual)
}

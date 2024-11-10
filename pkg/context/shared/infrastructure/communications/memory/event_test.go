package memory_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/events"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/communications"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/communications/memory"
)

type EventBusTestSuite struct {
	suite.Suite
	sut      events.Bus
	consumer *communications.EventConsumerMock
}

func (suite *EventBusTestSuite) SetupTest() {
	suite.consumer = new(communications.EventConsumerMock)

	suite.sut = &memory.EventBus{
		Consumers: make(map[events.Key][]events.Consumer),
	}
}

func (suite *EventBusTestSuite) TestSubscribe() {
	suite.NoError(suite.sut.Subscribe(messages.Random[events.Event]().Key, suite.consumer))
}

func (suite *EventBusTestSuite) TestPublish() {
	event := messages.Random[events.Event]()

	suite.NoError(suite.sut.Subscribe(event.Key, suite.consumer))

	suite.consumer.Mock.On("On", event)

	suite.NoError(suite.sut.Publish(event))

	suite.consumer.AssertExpectations(suite.T())
}

func (suite *EventBusTestSuite) TestPublishErrMissingConsumer() {
	event := messages.Random[events.Event]()

	err := suite.sut.Publish(event)

	var actual *errors.Internal

	suite.ErrorAs(err, &actual)

	expected := &errors.Internal{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "Publish",
		What:  "Failure to execute a Event without a Consumer",
		Why: errors.Meta{
			"Event": event.Key,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func TestIntegrationEventBusSuite(t *testing.T) {
	suite.Run(t, new(EventBusTestSuite))
}

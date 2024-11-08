package memory_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/events"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/communications"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/communications/memory"
)

type EventBusTestSuite struct {
	suite.Suite
	sut      events.Bus
	event    *events.Event
	consumer *communications.EventConsumerMock
}

func (suite *EventBusTestSuite) SetupTest() {
	suite.event = &events.Event{
		Key: "event.testing",
	}

	suite.consumer = new(communications.EventConsumerMock)

	suite.sut = &memory.EventBus{
		Consumers: make(map[events.Key][]events.Consumer),
	}
}

func (suite *EventBusTestSuite) TestSubscribe() {
	suite.NoError(suite.sut.Subscribe(suite.event.Key, suite.consumer))
}

func (suite *EventBusTestSuite) TestPublish() {
	suite.NoError(suite.sut.Subscribe(suite.event.Key, suite.consumer))

	suite.consumer.Mock.On("On", suite.event)

	suite.NoError(suite.sut.Publish(suite.event))

	suite.consumer.AssertExpectations(suite.T())
}

func (suite *EventBusTestSuite) TestPublishErrMissingConsumer() {
	err := suite.sut.Publish(suite.event)

	var actual *errors.Internal

	suite.ErrorAs(err, &actual)

	expected := &errors.Internal{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "Publish",
		What:  "Failure to execute a Event without a Consumer",
		Why: errors.Meta{
			"Event": suite.event.Key,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func TestIntegrationEventBusSuite(t *testing.T) {
	suite.Run(t, new(EventBusTestSuite))
}

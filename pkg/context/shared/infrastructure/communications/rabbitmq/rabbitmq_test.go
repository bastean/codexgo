package rabbitmq_test

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/events"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/communications"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/communications/rabbitmq"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/records/log"
)

type RabbitMQTestSuite struct {
	suite.Suite
	sut        events.Bus
	logger     *log.Log
	consumer   *communications.EventConsumerMock
	queue      events.Recipient
	routingKey events.Key
}

func (s *RabbitMQTestSuite) SetupTest() {
	var err error

	s.routingKey = messages.NewKey(&messages.KeyComponents{
		Service: "publisher",
		Version: "1",
		Type:    messages.Type.Event,
		Entity:  "publisher",
		Event:   "test",
		Status:  messages.Status.Succeeded,
	})

	s.queue = messages.NewRecipient(&messages.RecipientComponents{
		Service: "queue",
		Entity:  "queue",
		Action:  "assert",
		Event:   "test",
		Status:  "succeeded",
	})

	queues := rabbitmq.Queues{
		s.routingKey: &rabbitmq.Recipient{
			Name:       s.queue,
			BindingKey: events.Key("#.event.#.test.succeeded"),
		},
	}

	s.consumer = new(communications.EventConsumerMock)

	s.logger = log.New()

	s.sut, err = rabbitmq.Open(
		os.Getenv("CODEXGO_BROKER_RABBITMQ_URI"),
		os.Getenv("CODEXGO_BROKER_RABBITMQ_NAME"),
		queues,
		s.logger,
	)

	if err != nil {
		errors.Panic(err.Error(), "SetupTest")
	}
}

func (s *RabbitMQTestSuite) TestPublish() {
	event := messages.RandomWithKey[events.Event](s.routingKey)

	s.consumer.Mock.On("On", event)

	go func() {
		s.NoError(s.sut.Subscribe(event.Key, s.consumer))
	}()

	s.NoError(s.sut.Publish(event))

	s.Eventually(func() bool {
		return s.consumer.AssertExpectations(s.T())
	}, 3*time.Second, 30*time.Millisecond)
}

func TestIntegrationRabbitMQSuite(t *testing.T) {
	suite.Run(t, new(RabbitMQTestSuite))
}

package rabbitmq_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/communications"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/communications/rabbitmq"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/records/log"
)

type RabbitMQTestSuite struct {
	communications.EventBusSuite
}

func (s *RabbitMQTestSuite) SetupTest() {
	var err error

	routingKey := messages.NewKey(&messages.KeyComponents{
		Service: "publisher",
		Version: "1",
		Type:    messages.Type.Event,
		Entity:  "publisher",
		Event:   "test",
		Status:  messages.Status.Succeeded,
	})

	queue := messages.NewRecipient(&messages.RecipientComponents{
		Service: "queue",
		Entity:  "queue",
		Action:  "assert",
		Event:   "test",
		Status:  "succeeded",
	})

	queues := rabbitmq.QueueMapper{
		routingKey: &rabbitmq.Recipient{
			Name:       queue,
			BindingKey: messages.Key("#.event.#.test.succeeded"),
		},
	}

	logger := log.New()

	consumeCycle, _ := context.WithTimeout(context.Background(), 5*time.Second)

	s.EventBusSuite.Event = messages.RandomWithKey(routingKey)

	s.EventBusSuite.Consumer = new(communications.EventConsumerMock)

	s.EventBusSuite.SUT, err = rabbitmq.Open(
		os.Getenv("CODEXGO_BROKER_RABBITMQ_URI"),
		os.Getenv("CODEXGO_BROKER_RABBITMQ_NAME"),
		logger,
		consumeCycle,
	)

	if err != nil {
		errors.Panic(err.Error(), "SetupTest")
	}

	err = rabbitmq.AddQueueMapper(s.EventBusSuite.SUT.(*rabbitmq.RabbitMQ), queues)

	if err != nil {
		errors.Panic(err.Error(), "SetupTest")
	}
}

func TestIntegrationRabbitMQSuite(t *testing.T) {
	suite.Run(t, new(RabbitMQTestSuite))
}

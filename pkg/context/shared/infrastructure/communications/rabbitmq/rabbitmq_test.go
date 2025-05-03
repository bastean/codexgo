package rabbitmq_test

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/time"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/communications"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/communications/rabbitmq"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/records/log"
)

type RabbitMQTestSuite struct {
	communications.EventBusSuite
}

func (s *RabbitMQTestSuite) SetupSuite() {
	var err error

	routingKey, _ := values.New[*messages.Key](messages.ParseKey(&messages.KeyComponents{
		Service: "publisher",
		Version: "1",
		Type:    messages.Type.Event,
		Entity:  "publisher",
		Action:  "test",
		Status:  messages.Status.Succeeded,
	}))

	queue, _ := values.New[*messages.Recipient](messages.ParseRecipient(&messages.RecipientComponents{
		Service: "queue",
		Entity:  "queue",
		Trigger: "assert",
		Action:  "test",
		Status:  "succeeded",
	}))

	queues := rabbitmq.Mapper{
		routingKey: []*rabbitmq.Queue{
			{
				Name:       queue,
				BindingKey: "#.event.#.test.succeeded",
			},
		},
	}

	logger := log.New()

	consumeCycle, _ := context.WithTimeout(context.Background(), 5*time.Second) //nolint:govet

	s.EventBusSuite.Event = messages.Mother.MessageValidWithKey(routingKey)

	s.EventBusSuite.Consumer = new(communications.EventConsumerMock)

	s.EventBusSuite.SUT, err = rabbitmq.Open(
		os.Getenv("CODEXGO_BROKER_RABBITMQ_URI"),
		os.Getenv("CODEXGO_BROKER_RABBITMQ_NAME"),
		logger,
		consumeCycle,
	)

	if err != nil {
		errors.Panic(err)
	}

	err = rabbitmq.AddQueueMapper(s.EventBusSuite.SUT.(*rabbitmq.RabbitMQ), queues)

	if err != nil {
		errors.Panic(err)
	}
}

func TestIntegrationRabbitMQSuite(t *testing.T) {
	suite.Run(t, new(RabbitMQTestSuite))
}

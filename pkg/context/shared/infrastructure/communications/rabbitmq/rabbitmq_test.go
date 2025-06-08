package rabbitmq_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/suite"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/time"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/communications"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/communications/rabbitmq"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/records/log"
)

type RabbitMQTestSuite struct {
	communications.EventBusSuite
}

func (s *RabbitMQTestSuite) SetupSuite() {
	var err error

	logger := log.New()

	consumeCycle, _ := context.WithTimeout(context.Background(), 5*time.Second) //nolint:govet

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
}

func (s *RabbitMQTestSuite) SetupTest() {
	s.EventBusSuite.SetupTest()

	var routingKeyComponents *messages.KeyComponents

	s.NotPanics(func() { routingKeyComponents = messages.ParseKey(s.EventBusSuite.Event.Key.Value()) })

	err := rabbitmq.AddQueueMapper(s.EventBusSuite.SUT.(*rabbitmq.RabbitMQ), rabbitmq.Mapper{
		s.EventBusSuite.Event.Key.Value(): []*rabbitmq.Queue{
			{
				Name: messages.Mother().RecipientValid(),
				BindingKey: fmt.Sprintf("#.%s.#.%s.%s",
					routingKeyComponents.Type,
					routingKeyComponents.Action,
					routingKeyComponents.Status),
			},
		},
	})

	if err != nil {
		errors.Panic(err)
	}
}

func TestIntegrationRabbitMQSuite(t *testing.T) {
	suite.Run(t, new(RabbitMQTestSuite))
}

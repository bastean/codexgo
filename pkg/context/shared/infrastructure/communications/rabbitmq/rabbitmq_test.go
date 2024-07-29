package rabbitmq_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/communications"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/communications/rabbitmq"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/records"
	"github.com/stretchr/testify/suite"
)

type RabbitMQTestSuite struct {
	suite.Suite
	sut      messages.Broker
	logger   *records.LoggerMock
	router   *messages.Router
	queue    *messages.Queue
	consumer *communications.ConsumerMock
	messages []*messages.Message
}

func (suite *RabbitMQTestSuite) SetupTest() {
	suite.logger = new(records.LoggerMock)

	suite.sut, _ = rabbitmq.Open(
		os.Getenv("BROKER_RABBITMQ_URI"),
		suite.logger,
	)

	suite.router = &messages.Router{
		Name: os.Getenv("BROKER_RABBITMQ_NAME"),
	}

	suite.queue = &messages.Queue{
		Name: messages.NewRecipientName(&messages.RecipientNameComponents{
			Service: "queue",
			Entity:  "queue",
			Action:  "assert",
			Event:   "test",
			Status:  "succeeded",
		}),
	}

	suite.consumer = new(communications.ConsumerMock)

	message := messages.New(
		messages.NewRoutingKey(&messages.RoutingKeyComponents{
			Service: "publisher",
			Version: "1",
			Type:    messages.Type.Event,
			Entity:  "publisher",
			Event:   "test",
			Status:  messages.Status.Succeeded,
		}),
		messages.Attributes{},
		messages.Meta{},
	)

	message.Id = "0"

	message.OccurredOn = "0"

	suite.messages = append(suite.messages, message)
}

func (suite *RabbitMQTestSuite) TestBroker() {
	suite.NoError(suite.sut.AddRouter(suite.router))

	suite.NoError(suite.sut.AddQueue(suite.queue))

	bindingKeys := []string{"#.event.#.test.succeeded"}

	bindingSucceeded := fmt.Sprintf("binding queue [%s] to exchange [%s] with binding key [%s]", suite.queue.Name, suite.router.Name, bindingKeys[0])

	suite.logger.Mock.On("Info", bindingSucceeded)

	suite.NoError(suite.sut.AddQueueMessageBind(suite.queue, bindingKeys))

	suite.consumer.Mock.On("SubscribedTo").Return([]*messages.Queue{suite.queue})

	suite.NoError(suite.sut.AddQueueConsumer(suite.consumer))

	// TODO?(goroutine): suite.consumer.Mock.On("On", suite.messages[0])

	suite.NoError(suite.sut.PublishMessages(suite.messages))

	suite.logger.AssertExpectations(suite.T())

	suite.consumer.AssertExpectations(suite.T())
}

func TestIntegrationRabbitMQSuite(t *testing.T) {
	suite.Run(t, new(RabbitMQTestSuite))
}

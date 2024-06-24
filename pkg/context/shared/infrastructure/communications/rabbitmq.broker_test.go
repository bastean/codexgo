package communications_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/bastean/codexgo/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/pkg/context/shared/infrastructure/communications"
	"github.com/bastean/codexgo/pkg/context/shared/infrastructure/loggers"
	"github.com/stretchr/testify/suite"
)

type RabbitMQBrokerTestSuite struct {
	suite.Suite
	sut      messages.Broker
	logger   *loggers.LoggerMock
	router   *messages.Router
	queue    *messages.Queue
	consumer *communications.ConsumerMock
	messages []*messages.Message
}

func (suite *RabbitMQBrokerTestSuite) SetupTest() {
	suite.logger = new(loggers.LoggerMock)

	uri := os.Getenv("BROKER_URI")

	suite.sut, _ = communications.NewRabbitMQ(uri, suite.logger)

	suite.router = &messages.Router{Name: "test"}

	queueName := messages.NewRecipientName(&messages.RecipientNameComponents{
		Service: "queue",
		Entity:  "queue",
		Action:  "assert",
		Event:   "test",
		Status:  "succeeded",
	})

	suite.queue = &messages.Queue{Name: queueName}

	suite.consumer = new(communications.ConsumerMock)

	messageRoutingKey := messages.NewRoutingKey(&messages.RoutingKeyComponents{
		Service: "publisher",
		Version: "1",
		Type:    messages.Type.Event,
		Entity:  "publisher",
		Event:   "test",
		Status:  messages.Status.Succeeded,
	})

	messageAttributes := messages.Attributes{}

	messageMeta := messages.Meta{}

	message := messages.NewMessage(messageRoutingKey, messageAttributes, messageMeta)

	message.Id = "0"

	message.OccurredOn = "0"

	suite.messages = append(suite.messages, message)
}

func (suite *RabbitMQBrokerTestSuite) TestBroker() {
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

func TestIntegrationRabbitMQBrokerSuite(t *testing.T) {
	suite.Run(t, new(RabbitMQBrokerTestSuite))
}

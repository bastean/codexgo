package communications_test

import (
	"os"
	"testing"

	"github.com/bastean/codexgo/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/pkg/context/shared/domain/models"
	"github.com/bastean/codexgo/pkg/context/shared/domain/queues"
	"github.com/bastean/codexgo/pkg/context/shared/domain/routers"
	"github.com/bastean/codexgo/pkg/context/shared/infrastructure/communications"
	"github.com/bastean/codexgo/pkg/context/shared/infrastructure/loggers"
	"github.com/stretchr/testify/suite"
)

type RabbitMQBrokerTestSuite struct {
	suite.Suite
	sut      models.Broker
	logger   *loggers.LoggerMock
	router   *routers.Router
	queue    *queues.Queue
	consumer *communications.ConsumerMock
	messages []*messages.Message
}

func (suite *RabbitMQBrokerTestSuite) SetupTest() {
	suite.logger = new(loggers.LoggerMock)

	uri := os.Getenv("BROKER_URI")

	suite.sut, _ = communications.NewRabbitMQ(uri, suite.logger)

	suite.router = &routers.Router{Name: "test"}

	queueName := queues.NewQueueName(&queues.QueueName{
		Module: "queue",
		Action: "assert",
		Event:  "test.succeeded",
	})

	suite.queue = &queues.Queue{Name: queueName}

	suite.consumer = new(communications.ConsumerMock)

	messageRoutingKey := messages.NewRoutingKey(&messages.MessageRoutingKey{
		Module:    "publisher",
		Version:   "1",
		Type:      messages.Type.Event,
		Aggregate: "publisher",
		Event:     "test",
		Status:    messages.Status.Succeeded,
	})

	messageAttributes := []byte{}

	messageMeta := []byte{}

	message := messages.NewMessage(messageRoutingKey, messageAttributes, messageMeta)

	message.Id = "0"

	message.OccurredOn = "0"

	suite.messages = append(suite.messages, message)
}

func (suite *RabbitMQBrokerTestSuite) TestBroker() {
	suite.NoError(suite.sut.AddRouter(suite.router))

	suite.NoError(suite.sut.AddQueue(suite.queue))

	suite.NoError(suite.sut.AddQueueMessageBind(suite.queue, []string{"#.event.#.test.succeeded"}))

	suite.consumer.Mock.On("SubscribedTo").Return([]*queues.Queue{suite.queue})

	suite.NoError(suite.sut.AddQueueConsumer(suite.consumer))

	// TODO?(goroutine): suite.consumer.Mock.On("On", suite.messages[0])

	suite.NoError(suite.sut.PublishMessages(suite.messages))

	suite.consumer.AssertExpectations(suite.T())
}

func TestIntegrationRabbitMQBrokerSuite(t *testing.T) {
	suite.Run(t, new(RabbitMQBrokerTestSuite))
}

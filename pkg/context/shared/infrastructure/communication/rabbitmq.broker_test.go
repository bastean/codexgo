package communication_test

import (
	"os"
	"testing"

	"github.com/bastean/codexgo/pkg/context/shared/domain/message"
	"github.com/bastean/codexgo/pkg/context/shared/domain/model"
	"github.com/bastean/codexgo/pkg/context/shared/domain/queue"
	"github.com/bastean/codexgo/pkg/context/shared/domain/router"
	"github.com/bastean/codexgo/pkg/context/shared/infrastructure/communication"
	communicationMock "github.com/bastean/codexgo/pkg/context/shared/infrastructure/communication/mock"
	loggerMock "github.com/bastean/codexgo/pkg/context/shared/infrastructure/logger/mock"
	"github.com/stretchr/testify/suite"
)

type RabbitMQBrokerTestSuite struct {
	suite.Suite
	sut      model.Broker
	logger   *loggerMock.LoggerMock
	router   *router.Router
	queue    *queue.Queue
	consumer *communicationMock.ConsumerMock
	messages []*message.Message
}

func (suite *RabbitMQBrokerTestSuite) SetupTest() {
	suite.logger = new(loggerMock.LoggerMock)

	uri := os.Getenv("BROKER_URI")
	suite.sut, _ = communication.NewRabbitMQ(uri, suite.logger)

	suite.router = &router.Router{Name: "test"}

	queueName := queue.NewQueueName(&queue.QueueName{
		Module: "queue",
		Action: "assert",
		Event:  "test.succeeded",
	})
	suite.queue = &queue.Queue{Name: queueName}

	suite.consumer = new(communicationMock.ConsumerMock)

	messageRoutingKey := message.NewRoutingKey(&message.MessageRoutingKey{
		Module:    "publisher",
		Version:   "1",
		Type:      message.Type.Event,
		Aggregate: "publisher",
		Event:     "test",
		Status:    message.Status.Succeeded,
	})
	messageAttributes := []byte{}
	messageMeta := []byte{}
	message := message.NewMessage(messageRoutingKey, messageAttributes, messageMeta)
	message.Id = "0"
	message.OccurredOn = "0"
	suite.messages = append(suite.messages, message)
}

func (suite *RabbitMQBrokerTestSuite) TestBroker() {
	suite.NoError(suite.sut.AddRouter(suite.router))

	suite.NoError(suite.sut.AddQueue(suite.queue))

	suite.NoError(suite.sut.AddQueueMessageBind(suite.queue, []string{"#.event.#.test.succeeded"}))

	suite.consumer.Mock.On("SubscribedTo").Return([]*queue.Queue{suite.queue})

	suite.NoError(suite.sut.AddQueueConsumer(suite.consumer))

	// TODO?(goroutine): suite.consumer.Mock.On("On", suite.messages[0])

	suite.NoError(suite.sut.PublishMessages(suite.messages))

	suite.consumer.AssertExpectations(suite.T())
}

func TestIntegrationRabbitMQBrokerSuite(t *testing.T) {
	suite.Run(t, new(RabbitMQBrokerTestSuite))
}

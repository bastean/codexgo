package communication_test

import (
	"testing"

	"github.com/bastean/codexgo/pkg/context/shared/domain/exchange"
	"github.com/bastean/codexgo/pkg/context/shared/domain/message"
	"github.com/bastean/codexgo/pkg/context/shared/domain/model"
	"github.com/bastean/codexgo/pkg/context/shared/domain/queue"
	"github.com/bastean/codexgo/pkg/context/shared/infrastructure/communication"
	communicationMock "github.com/bastean/codexgo/pkg/context/shared/infrastructure/communication/mock"
	"github.com/stretchr/testify/suite"
)

type RabbitMQBrokerTestSuite struct {
	suite.Suite
	sut      model.Broker
	exchange *exchange.Exchange
	queue    *queue.Queue
	consumer *communicationMock.ConsumerMock
	messages []*message.Message
}

func (suite *RabbitMQBrokerTestSuite) SetupTest() {
	suite.sut = communication.NewRabbitMQ()

	suite.exchange = exchange.NewExchange("test")

	queueName := queue.NewQueueName(&queue.QueueName{Module: "queue", Action: "assert", Event: "test.succeeded"})
	suite.queue = queue.NewQueue(queueName)

	suite.consumer = communicationMock.NewConsumerMock()

	messageRoutingKey := message.NewMessageRoutingKey(&message.MessageRoutingKey{Module: "publisher", Version: "1", Type: message.Event, Aggregate: "publisher", Event: "test", Status: message.Succeeded})
	messageAttributes := []byte{}
	messageMeta := []byte{}
	message := message.NewMessage(messageRoutingKey, messageAttributes, messageMeta)
	message.Id = "0"
	message.OccurredOn = "0"
	suite.messages = append(suite.messages, message)

	suite.sut.AddExchange(suite.exchange)

	suite.sut.AddQueue(suite.queue)

	suite.sut.AddQueueMessageBind(suite.queue, []string{"#.event.#.test.succeeded"})
}

func (suite *RabbitMQBrokerTestSuite) TestBroker() {
	suite.consumer.Mock.On("SubscribedTo").Return([]*queue.Queue{suite.queue})

	suite.sut.AddQueueConsumer(suite.consumer)

	// TODO?: suite.consumer.Mock.On("On", suite.messages[0])

	suite.NotPanics(func() { suite.sut.PublishMessages(suite.messages) })

	suite.consumer.AssertExpectations(suite.T())
}

func TestRabbitMQBrokerSuite(t *testing.T) {
	suite.Run(t, new(RabbitMQBrokerTestSuite))
}

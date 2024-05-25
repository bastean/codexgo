package scommunication_test

import (
	"os"
	"testing"

	"github.com/bastean/codexgo/pkg/context/shared/domain/smessage"
	"github.com/bastean/codexgo/pkg/context/shared/domain/smodel"
	"github.com/bastean/codexgo/pkg/context/shared/domain/squeue"
	"github.com/bastean/codexgo/pkg/context/shared/domain/srouter"
	"github.com/bastean/codexgo/pkg/context/shared/infrastructure/scommunication"
	"github.com/bastean/codexgo/pkg/context/shared/infrastructure/slogger"
	"github.com/stretchr/testify/suite"
)

type RabbitMQBrokerTestSuite struct {
	suite.Suite
	sut      smodel.Broker
	logger   *slogger.LoggerMock
	router   *srouter.Router
	queue    *squeue.Queue
	consumer *scommunication.ConsumerMock
	messages []*smessage.Message
}

func (suite *RabbitMQBrokerTestSuite) SetupTest() {
	suite.logger = new(slogger.LoggerMock)

	uri := os.Getenv("BROKER_URI")

	suite.sut, _ = scommunication.NewRabbitMQ(uri, suite.logger)

	suite.router = &srouter.Router{Name: "test"}

	queueName := squeue.NewQueueName(&squeue.QueueName{
		Module: "queue",
		Action: "assert",
		Event:  "test.succeeded",
	})

	suite.queue = &squeue.Queue{Name: queueName}

	suite.consumer = new(scommunication.ConsumerMock)

	messageRoutingKey := smessage.NewRoutingKey(&smessage.MessageRoutingKey{
		Module:    "publisher",
		Version:   "1",
		Type:      smessage.Type.Event,
		Aggregate: "publisher",
		Event:     "test",
		Status:    smessage.Status.Succeeded,
	})

	messageAttributes := []byte{}

	messageMeta := []byte{}

	message := smessage.NewMessage(messageRoutingKey, messageAttributes, messageMeta)

	message.Id = "0"

	message.OccurredOn = "0"

	suite.messages = append(suite.messages, message)
}

func (suite *RabbitMQBrokerTestSuite) TestBroker() {
	suite.NoError(suite.sut.AddRouter(suite.router))

	suite.NoError(suite.sut.AddQueue(suite.queue))

	suite.NoError(suite.sut.AddQueueMessageBind(suite.queue, []string{"#.event.#.test.succeeded"}))

	suite.consumer.Mock.On("SubscribedTo").Return([]*squeue.Queue{suite.queue})

	suite.NoError(suite.sut.AddQueueConsumer(suite.consumer))

	// TODO?(goroutine): suite.consumer.Mock.On("On", suite.messages[0])

	suite.NoError(suite.sut.PublishMessages(suite.messages))

	suite.consumer.AssertExpectations(suite.T())
}

func TestIntegrationRabbitMQBrokerSuite(t *testing.T) {
	suite.Run(t, new(RabbitMQBrokerTestSuite))
}

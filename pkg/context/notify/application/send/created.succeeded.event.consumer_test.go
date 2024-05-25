package send_test

import (
	"encoding/json"
	"testing"

	"github.com/bastean/codexgo/pkg/context/notify/application/send"
	"github.com/bastean/codexgo/pkg/context/notify/infrastructure/communication"
	"github.com/bastean/codexgo/pkg/context/shared/domain/smodel"
	"github.com/bastean/codexgo/pkg/context/shared/domain/squeue"
	"github.com/bastean/codexgo/pkg/context/shared/domain/stype"
	"github.com/stretchr/testify/suite"
)

type CreatedSucceededEventConsumerTestSuite struct {
	suite.Suite
	sut       smodel.Consumer
	useCase   smodel.UseCase[any, *stype.Empty]
	transport *communication.TransportMock
	queues    []*squeue.Queue
}

func (suite *CreatedSucceededEventConsumerTestSuite) SetupTest() {
	queueName := squeue.NewQueueName(&squeue.QueueName{
		Module: "queue",
		Action: "assert",
		Event:  "test.succeeded",
	})

	suite.queues = append(suite.queues, &squeue.Queue{
		Name: queueName,
	})

	suite.transport = new(communication.TransportMock)

	suite.useCase = &send.Send{
		Transport: suite.transport,
	}

	suite.sut = &send.CreatedSucceededEventConsumer{
		UseCase: suite.useCase,
		Queues:  suite.queues,
	}
}

func (suite *CreatedSucceededEventConsumerTestSuite) TestEventConsumer() {
	message := send.RandomEvent()

	attributes := new(send.CreatedSucceededEventAttributes)

	json.Unmarshal(message.Attributes, attributes)

	suite.transport.On("Submit", attributes)

	suite.NoError(suite.sut.On(message))

	suite.transport.AssertExpectations(suite.T())
}

func TestUnitCreatedSucceededEventConsumerSuite(t *testing.T) {
	suite.Run(t, new(CreatedSucceededEventConsumerTestSuite))
}

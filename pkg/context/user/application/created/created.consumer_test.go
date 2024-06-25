package created_test

import (
	"encoding/json"
	"testing"

	"github.com/bastean/codexgo/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/pkg/context/shared/domain/models"
	"github.com/bastean/codexgo/pkg/context/shared/domain/types"
	"github.com/bastean/codexgo/pkg/context/user/application/created"
	"github.com/bastean/codexgo/pkg/context/user/domain/event"
	"github.com/bastean/codexgo/pkg/context/user/infrastructure/communication"
	"github.com/stretchr/testify/suite"
)

type CreatedConsumerTestSuite struct {
	suite.Suite
	sut       messages.Consumer
	usecase   models.UseCase[*event.CreatedSucceeded, types.Empty]
	transport *communication.TransportMock
	queues    []*messages.Queue
}

func (suite *CreatedConsumerTestSuite) SetupTest() {
	queueName := messages.NewRecipientName(&messages.RecipientNameComponents{
		Service: "queue",
		Entity:  "queue",
		Action:  "assert",
		Event:   "test",
		Status:  "succeeded",
	})

	suite.queues = append(suite.queues, &messages.Queue{
		Name: queueName,
	})

	suite.transport = new(communication.TransportMock)

	suite.usecase = &created.Created{
		Transport: suite.transport,
	}

	suite.sut = &created.Consumer{
		UseCase: suite.usecase,
		Queues:  suite.queues,
	}
}

func (suite *CreatedConsumerTestSuite) TestCreatedSucceeded() {
	message := event.RandomCreatedSucceeded()

	user := new(event.CreatedSucceeded)

	user.Attributes = new(event.CreatedSucceededAttributes)

	suite.NoError(json.Unmarshal(message.Attributes, user.Attributes))

	suite.transport.On("Submit", user.Attributes)

	suite.NoError(suite.sut.On(message))

	suite.transport.AssertExpectations(suite.T())
}

func TestUnitCreatedConsumerSuite(t *testing.T) {
	suite.Run(t, new(CreatedConsumerTestSuite))
}

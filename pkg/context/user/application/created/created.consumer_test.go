package created_test

import (
	"encoding/json"
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/transports"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/created"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/usecase"
	"github.com/stretchr/testify/suite"
)

type CreatedConsumerTestSuite struct {
	suite.Suite
	sut      messages.Consumer
	created  usecase.Created
	transfer *transports.TransferMock
	queues   []*messages.Queue
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

	suite.transfer = new(transports.TransferMock)

	suite.created = &created.Created{
		Transfer: suite.transfer,
	}

	suite.sut = &created.Consumer{
		Created: suite.created,
		Queues:  suite.queues,
	}
}

func (suite *CreatedConsumerTestSuite) TestCreatedSucceeded() {
	message := user.RandomCreatedSucceeded()

	event := new(user.CreatedSucceeded)

	event.Attributes = new(user.CreatedSucceededAttributes)

	suite.NoError(json.Unmarshal(message.Attributes, event.Attributes))

	suite.transfer.On("Submit", event.Attributes)

	suite.NoError(suite.sut.On(message))

	suite.transfer.AssertExpectations(suite.T())
}

func TestUnitCreatedConsumerSuite(t *testing.T) {
	suite.Run(t, new(CreatedConsumerTestSuite))
}

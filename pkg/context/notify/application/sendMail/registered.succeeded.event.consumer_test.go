package sendMail_test

import (
	"encoding/json"
	"testing"

	"github.com/bastean/codexgo/pkg/context/notify/application/sendMail"
	eventMother "github.com/bastean/codexgo/pkg/context/notify/application/sendMail/mother"
	"github.com/bastean/codexgo/pkg/context/notify/domain/model"
	"github.com/bastean/codexgo/pkg/context/notify/domain/template"
	communicationMock "github.com/bastean/codexgo/pkg/context/notify/infrastructure/communication/mock"
	sharedModel "github.com/bastean/codexgo/pkg/context/shared/domain/model"
	"github.com/bastean/codexgo/pkg/context/shared/domain/queue"
	"github.com/bastean/codexgo/pkg/context/shared/domain/types"
	"github.com/stretchr/testify/suite"
)

type RegisteredSucceededEventConsumerTestSuite struct {
	suite.Suite
	sut     sharedModel.Consumer
	useCase sharedModel.UseCase[model.MailTemplate, *types.Empty]
	mail    *communicationMock.MailMock
	queues  []*queue.Queue
}

func (suite *RegisteredSucceededEventConsumerTestSuite) SetupTest() {
	queueName := queue.NewQueueName(&queue.QueueName{
		Module: "notify",
		Action: "assert",
		Event:  "test.succeeded",
	})
	suite.queues = append(suite.queues, &queue.Queue{
		Name: queueName,
	})
	suite.mail = new(communicationMock.MailMock)
	suite.useCase = &sendMail.SendMail{
		Mail: suite.mail,
	}
	suite.sut = &sendMail.RegisteredSucceededEventConsumer{
		UseCase: suite.useCase,
		Queues:  suite.queues,
	}
}

func (suite *RegisteredSucceededEventConsumerTestSuite) TestEventConsumer() {
	message := eventMother.Random()

	attributes := new(sendMail.RegisteredSucceededEventAttributes)

	json.Unmarshal(message.Attributes, attributes)

	accountConfirmationTemplate := &template.AccountConfirmationMail{
		Mail: &template.Mail{
			To: []string{attributes.Email},
		},
		Username:         attributes.Username,
		ConfirmationLink: attributes.Id,
	}

	suite.mail.On("Send", accountConfirmationTemplate)

	suite.NoError(suite.sut.On(message))

	suite.mail.AssertExpectations(suite.T())
}

func TestUnitRegisteredSucceededEventConsumerSuite(t *testing.T) {
	suite.Run(t, new(RegisteredSucceededEventConsumerTestSuite))
}
